package service

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"

	"mall/app/idl"
	"mall/app/model"
	"mall/pkg/log"
	"mall/pkg/utils"
)

var (
	//ErrGoodsEmpty 商品为空
	ErrGoodsEmpty = errors.New("goods empty")
	//ErrCouponNotUse 不可使用
	ErrCouponNotUse = errors.New("coupon not use")
	//ErrOrderNotFound 订单不存在
	ErrOrderNotFound = errors.New("order not found")
)

//IOrder 订单服务
type IOrder interface {
	SubmitOrder(ctx context.Context, userID int, cartIDs []string, addressID, couponID int, remark string) (int, error)
	SubmitOrderGoods(ctx context.Context, userID, goodsID, skuID, num, addressID, couponID int, remark string) (int, error)
	OrderDetail(ctx context.Context, id, userID int) (*model.Order, error)
	OrderCancel(ctx context.Context, userID int, orderNo string) error
	MyOrderList(ctx context.Context, userID, status, offset, limit int) ([]*model.OrderList, error)
	OrderPayNotify(ctx context.Context, userID, amount int, pType int8, orderNo, tradeNo, transHash string) error
	OrderRefund(ctx context.Context, userID int, orderNo, content string) error
	OrderConfirmReceipt(ctx context.Context, userID int, orderNo string) error
	OrderComment(ctx context.Context, userID int, rate int8, orderNo, content string, goodsIDs []int) error
}

//SubmitOrder 购物车提交订单
func (s *Service) SubmitOrder(ctx context.Context, userID int, cartIDs []string, addressID, couponID int, remark string) (int, error) {
	carts, err := s.repo.GetCartsByIds(ctx, userID, cartIDs)
	if err != nil {
		return 0, errors.Wrapf(err, "[service.order] get carts by uid: %v ids: %v", userID, cartIDs)
	}
	if len(carts) == 0 {
		return 0, ErrGoodsEmpty
	}

	// 订单商品
	var total int
	var goods []*model.OrderGoodsModel
	for _, cart := range carts {
		goods = append(goods, &model.OrderGoodsModel{
			GID:        model.GID{GoodsID: cart.GoodsID},
			GoodsName:  cart.GoodsName,
			GoodsCover: cart.Cover,
			GoodsPrice: cart.Price,
			BuyCount:   cart.Num,
			Price:      cart.Price * cart.Num,
			Attrs:      cart.SkuName,
		})
		total += cart.Price * cart.Num
	}

	orderID, err := s.saveOrder(ctx, userID, couponID, addressID, total, remark, goods)
	if err != nil {
		return 0, err
	}

	//删除购物车
	err = s.repo.DelCart(ctx, userID, cartIDs)
	if err != nil {
		log.Warnf("[service.order] del cart by uid:%v ids: %v err: %v", userID, cartIDs, err)
	}

	return orderID, nil
}

//SubmitOrderGoods 商品直接提交订单
func (s *Service) SubmitOrderGoods(ctx context.Context, userID, goodsID, skuID, num, addressID, couponID int, remark string) (int, error) {
	goods, err := s.repo.GoodsDetail(ctx, goodsID)
	if err != nil {
		return 0, errors.Wrapf(err, "[service.cart] detail by goods_id: %v", goodsID)
	}
	if goods.ID == 0 {
		return 0, ErrGoodsNotFound
	}
	price := goods.Price
	skuName := ""
	if skuID > 0 {
		sku, err := s.repo.GetSkuByID(ctx, skuID)
		if err != nil {
			return 0, errors.Wrapf(err, "[service.cart] sku by id: %v", skuID)
		}
		if sku.ID == 0 {
			return 0, ErrGoodsSkuNotFound
		}
		price = sku.Price
		skuName = sku.ValueNames
	}

	total := goods.Price * num
	oGoods := &model.OrderGoodsModel{
		GID:        model.GID{GoodsID: goodsID},
		GoodsName:  goods.Title,
		GoodsCover: goods.Cover,
		GoodsPrice: price,
		BuyCount:   num,
		Price:      total,
		Attrs:      skuName,
	}

	return s.saveOrder(ctx, userID, couponID, addressID, total, remark, []*model.OrderGoodsModel{oGoods})
}

//OrderDetail 订单详情
func (s *Service) OrderDetail(ctx context.Context, id, userID int) (*model.Order, error) {
	order, err := s.repo.GetOrderByID(ctx, id, userID)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.order] find by id: %v", id)
	}
	if order.ID == 0 {
		return nil, ErrOrderNotFound
	}

	return idl.TransferOrder(order), nil
}

//MyOrderList 订单列表
func (s *Service) MyOrderList(ctx context.Context, userID, status, offset, limit int) ([]*model.OrderList, error) {
	list, err := s.repo.GetOrderList(ctx, userID, status, offset, limit)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.order] list uid: %v", userID)
	}

	return idl.TransferOrderList(list), nil
}

//OrderPayNotify 支付成功回调处理
func (s *Service) OrderPayNotify(ctx context.Context, userID, amount int, pType int8, orderNo, tradeNo, transHash string) error {
	order, err := s.repo.GetOrderByNo(ctx, userID, orderNo)
	if err != nil {
		return errors.Wrapf(err, "[service.order] get by no: %v", orderNo)
	}
	if order.ID == 0 || order.Status != model.OrderStatusInit {
		return ErrOrderNotFound
	}

	//检查是否真实支付
	var pays []*model.ConfigPayList
	err = s.repo.GetConfigByName(ctx, model.ConfigKeyPayList, &pays)
	if err != nil {
		return errors.Wrapf(err, "[service.comm] get home cat")
	}
	var address string
	for _, pay := range pays {
		if pay.ID == pType {
			address = pay.Address
		}
	}
	if address == "" {
		return ErrOrderNotFound
	}
	//连接合约
	err = s.contract.Connect(pType, address)
	if err != nil {
		return errors.Wrapf(err, "[service.order] connect contract address %v", address)
	}
	//调用合约
	check, err := s.contract.CheckPay(pType, orderNo)
	if err != nil {
		return errors.Wrapf(err, "[service.order] contract call checkPay")
	}
	if !check {//未支付
		return ErrOrderNotFound
	}

	order.Status = model.OrderStatusDelivered
	order.PayAt = time.Now().Unix()
	order.PayType = pType
	order.PayAmount = amount
	order.PayStatus = model.OrderPayStatue
	order.TradeNo = tradeNo
	order.TransHash = transHash
	err = s.repo.OrderSave(ctx, nil, order)
	if err != nil {
		return errors.Wrapf(err, "[service.order] save")
	}

	return nil
}

//OrderConfirmReceipt 确认收货
func (s *Service) OrderConfirmReceipt(ctx context.Context, userID int, orderNo string) error {
	order, err := s.repo.GetOrderByNo(ctx, userID, orderNo)
	if err != nil {
		return errors.Wrapf(err, "[service.order] get by no: %v", orderNo)
	}
	if order.ID == 0 || (order.Status != model.OrderStatusDelivered && order.Status != model.OrderStatusShipped) {
		return ErrOrderNotFound
	}

	order.Status = model.OrderStatusReceived
	err = s.repo.OrderSave(ctx, nil, order)
	if err != nil {
		return errors.Wrapf(err, "[service.order] save")
	}

	return nil
}

//OrderRefund 退款
func (s *Service) OrderRefund(ctx context.Context, userID int, orderNo, content string) error {
	order, err := s.repo.GetOrderByNo(ctx, userID, orderNo)
	if err != nil {
		return errors.Wrapf(err, "[service.order] get by no: %v", orderNo)
	}
	if order.ID == 0 || (order.Status != model.OrderStatusReceived && order.Status != model.OrderStatusDelivered) {
		return ErrOrderNotFound
	}

	// 开启事务
	db := model.GetDB()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	order.Status = model.OrderStatusPendingRefund
	err = s.repo.OrderSave(ctx, tx, order)
	if err != nil {
		tx.Rollback()
		return errors.Wrapf(err, "[service.order] save")
	}

	err = s.repo.CreateOrderRefund(ctx, tx, &model.OrderRefundModel{
		OrderID: order.ID,
		Amount:  order.PayAmount,
		Content: content,
	})
	if err != nil {
		tx.Rollback()
		return errors.Wrapf(err, "[service.order] create refund")
	}

	//提交事务
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "[service.order] tx commit refund")
	}

	return nil
}

func (s *Service) OrderComment(ctx context.Context, userID int, rate int8, orderNo, content string, goodsIDs []int) error {
	order, err := s.repo.GetOrderByNo(ctx, userID, orderNo)
	if err != nil {
		return errors.Wrapf(err, "[service.order] get by no: %v", orderNo)
	}
	if order.ID == 0 || order.Status != model.OrderStatusReceived {
		return ErrOrderNotFound
	}

	// 开启事务
	db := model.GetDB()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	order.Status = model.OrderStatusFinish
	order.FinishAt = time.Now().Unix()
	err = s.repo.OrderSave(ctx, tx, order)
	if err != nil {
		tx.Rollback()
		return errors.Wrapf(err, "[service.order] save")
	}

	var comments []*model.GoodsCommentModel
	for _, d := range goodsIDs {
		comments = append(comments, &model.GoodsCommentModel{
			UID:        model.UID{UserID: userID},
			GID:        model.GID{GoodsID: d},
			OrderNo:    order.OrderNo,
			Rate:       rate,
			Content:    content,
		})
	}

	err = s.repo.CreateGoodsComment(ctx, tx, comments)
	if err != nil {
		tx.Rollback()
		return errors.Wrapf(err, "[service.order] create comment")
	}

	//提交事务
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "[service.order] tx commit refund")
	}

	return nil
}

//OrderCancel 取消订单
func (s *Service) OrderCancel(ctx context.Context, userID int, orderNo string) error {
	order, err := s.repo.GetOrderByNo(ctx, userID, orderNo)
	if err != nil {
		return errors.Wrapf(err, "[service.order] get by no: %v", orderNo)
	}
	if order.ID == 0 || order.Status != model.OrderStatusInit {
		return ErrOrderNotFound
	}

	err = s.repo.OrderDelete(ctx, order)
	if err != nil {
		return errors.Wrapf(err, "[service.order] delete")
	}

	return nil
}

func (s *Service) saveOrder(ctx context.Context, userID, couponID, addressID, total int, remark string, goods []*model.OrderGoodsModel) (int, error) {
	address, err := s.repo.GetUserAddressByID(ctx, addressID, userID)
	if err != nil {
		return 0, errors.Wrapf(err, "[service.order] get address by id: %v uid: %v", addressID, userID)
	}
	if address.ID == 0 {
		return 0, ErrUserAddressNotFound
	}

	//订单金额
	amount := total
	var coupon *model.CouponModel
	if couponID > 0 { //选择了优惠券
		cUser, err := s.repo.GetCouponUserByID(ctx, couponID)
		if err != nil {
			return 0, errors.Wrapf(err, "[service.order] get coupon user by id: %v", couponID)
		}
		if cUser.ID == 0 || cUser.UserID != userID || cUser.IsUsed == 1 {
			return 0, ErrCouponNotFound
		}
		coupon, err = s.repo.GetCouponByID(ctx, cUser.CouponID)
		if err != nil {
			return 0, errors.Wrapf(err, "[service.order] get coupon by id: %v", couponID)
		}
		now := int(time.Now().Unix())
		if coupon.ID == 0 || now > coupon.EndAt || now < coupon.StartAt {
			return 0, ErrCouponNotFound
		}

		if total < coupon.MinPrice {
			return 0, ErrCouponNotUse
		}
		amount = total - coupon.Value
		if coupon.Type == model.CouponTypeDiscount {
			amount = total * coupon.Value / 1000
		}
	}

	orderNo, err := utils.GenShortID()
	if err != nil {
		return 0, errors.Wrapf(err, "[service.order] gen orderNo")
	}

	// 开启事务
	db := model.GetDB()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	//创建订单
	order := &model.OrderModel{
		UID:          model.UID{UserID: userID},
		OrderNo:      orderNo,
		UserNote:     remark,
		TotalPrice:   total,
		Amount:       amount,
		CouponAmount: total - amount,
		Status:       model.OrderStatusInit,
	}
	orderID, err := s.repo.CreateOrder(ctx, tx, order)
	if err != nil {
		tx.Rollback()
		return 0, errors.Wrapf(err, "[service.order] create")
	}

	if couponID > 0 { // 设置优惠券已使用
		err = s.repo.SetCouponUserUsed(ctx, tx, couponID, userID)
		if err != nil {
			tx.Rollback()
			return 0, errors.Wrapf(err, "[service.order] set coupon used")
		}
	}

	// 创建收货地址
	oAddr := &model.OrderAddressModel{
		UID:       model.UID{UserID: userID},
		OrderID:   orderID,
		AddressID: addressID,
		Name:      address.Name,
		Phone:     address.Phone,
		Area:      fmt.Sprintf("%s %s %s", address.Province, address.City, address.County),
		Detail:    address.Detail,
	}
	err = s.repo.CreateOrderAddress(ctx, tx, oAddr)
	if err != nil {
		tx.Rollback()
		return 0, errors.Wrapf(err, "[service.order] create address")
	}

	// 订单商品设置订单号
	for _, good := range goods {
		good.OrderID = orderID
	}
	err = s.repo.CreateOrderGoods(ctx, tx, goods)
	if err != nil {
		tx.Rollback()
		return 0, errors.Wrapf(err, "[service.order] create goods")
	}

	//提交事务
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return 0, errors.Wrap(err, "[service.order] tx commit")
	}

	return orderID, err
}
