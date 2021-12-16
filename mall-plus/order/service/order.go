package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/pkg/errors"
	"go-micro.dev/v4/logger"

	"common/errno"
	"common/message"
	"common/orm"
	pb "common/proto/order"
	ware "common/proto/warehouse"
	"common/util"
	"order/idl"
	"order/model"
)

//IOrder 订单接口
type IOrder interface {
	OrderSubmit(ctx context.Context, memberID, addressID, couponID int64, skuIds []int64, note string) (int64, error)
	SubmitSkuOrder(ctx context.Context, memberID, skuID, addressID, couponID int64, num int, note string) (int64, error)
	SubmitSeckillOrder(ctx context.Context, memberID, skuID, addressID int64, price, num int, orderNo string) error
	OrderDetail(ctx context.Context, memberID, id int64) (*pb.OrderInfo, error)
	OrderCancel(ctx context.Context, memberID, id int64) error
	MyOrderList(ctx context.Context, memberID int64, status, offset, limit int) ([]*pb.OrderList, error)
	OrderPayNotify(ctx context.Context, memberID int64, amount int, pType int8, orderNo, tradeNo, transHash string) error
	OrderConfirmReceipt(ctx context.Context, memberID, orderID int64) error
	OrderRefund(ctx context.Context, memberID, orderID int64, content string) error
	OrderComment(ctx context.Context, memberID, orderID int64, skuIds []int64, star int8, content, resources string) error
	OrderInfo(ctx context.Context, orderID int64) (*pb.OrderInfo, error)
}

//OrderConfirm 确认订单
func (s *Service) OrderConfirm(ctx context.Context) string {
	script := `if redis.call('get', KEYS[1]) == ARGV[1] then return redis.call('del',KEYS[1]) else return 0 end`
	return script
}

//OrderSubmit 提交订单
func (s *Service) OrderSubmit(ctx context.Context, memberID, addressID, couponID int64, skuIds []int64, note string) (int64, error) {
	//购物车获取需要购买的商品信息
	carts, err := s.batchGetCarts(ctx, memberID, skuIds)
	if err != nil {
		return 0, err
	}
	if len(carts) == 0 {
		return 0, errno.ErrOrderSkuEmpty
	}
	var total int
	//1，订单信息
	//2，商品spu信息（暂不处理）
	//3，商品sku信息
	//4，优惠信息（暂不处理）
	//5，积分信息,与价格同步 1元 = 1积分，1成长值
	items := make([]*model.OrderItemModel, 0, len(carts))
	for _, cart := range carts {
		amount := util.FormatAmount(cart.Price) * int(cart.Num)
		items = append(items, &model.OrderItemModel{
			Sku:             orm.Sku{SkuID: cart.SkuId},
			SkuName:         cart.Title,
			SkuImg:          cart.Cover,
			SkuPrice:        util.FormatAmount(cart.Price),
			SkuAttrs:        cart.SkuAttr,
			Num:             int(cart.Num),
			RealAmount:      amount,
			GiveIntegration: int(cart.Price) * int(cart.Num),
			GiveGrowth:      int(cart.Price) * int(cart.Num),
		})
		total += amount
	}

	orderID, err := s.saveOrder(ctx, memberID, couponID, addressID, total, note, "", items)
	if err != nil {
		return 0, err
	}

	//删除购物车
	if err = s.delCart(ctx, memberID, skuIds); err != nil {
		logger.Warnf("[service.order] del cart by uid:%v ids: %v err: %v", memberID, skuIds, err)
	}

	return orderID, nil
}

//SubmitSkuOrder 商品不通过购物车，直接提交订单
func (s *Service) SubmitSkuOrder(ctx context.Context, memberID, skuID, addressID, couponID int64, num int, note string) (int64, error) {
	//获取sku详情
	sku, err := s.getSkuByID(ctx, skuID)
	if err != nil {
		return 0, err
	}

	amount := int(sku.Price)
	items := []*model.OrderItemModel{
		{
			Sku:             orm.Sku{SkuID: skuID},
			SkuName:         sku.Title,
			SkuImg:          sku.Cover,
			SkuPrice:        amount,
			SkuAttrs:        sku.AttrValue,
			Num:             num,
			RealAmount:      amount,
			GiveIntegration: int(sku.Price) * num,
			GiveGrowth:      int(sku.Price) * num,
		},
	}

	return s.saveOrder(ctx, memberID, couponID, addressID, amount, note, "", items)
}

//SubmitSeckillOrder 秒杀订单提交
func (s *Service) SubmitSeckillOrder(ctx context.Context, memberID, skuID, addressID int64, price, num int, orderNo string) error {
	//获取sku详情
	sku, err := s.getSkuByID(ctx, skuID)
	if err != nil {
		return err
	}
	items := []*model.OrderItemModel{
		{
			Sku:             orm.Sku{SkuID: skuID},
			SkuName:         sku.Title,
			SkuImg:          sku.Cover,
			SkuPrice:        price,
			SkuAttrs:        sku.AttrValue,
			Num:             num,
			RealAmount:      price,
			GiveIntegration: int(sku.Price) * num,
			GiveGrowth:      int(sku.Price) * num,
		},
	}
	_, err = s.saveOrder(ctx, memberID, 0, addressID, price, "", orderNo, items)
	return err
}

//OrderDetail 订单详情
func (s *Service) OrderDetail(ctx context.Context, memberID, id int64) (*pb.OrderInfo, error) {
	order, err := s.repo.GetOrderDetail(ctx, id, memberID)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.order] find by id: %v", id)
	}
	if order.ID == 0 {
		return nil, errno.ErrOrderNotFound
	}

	return idl.TransferOrder(order), nil
}

//OrderCancel 取消订单
func (s *Service) OrderCancel(ctx context.Context, memberID, id int64) error {
	order, err := s.repo.GetOrderByID(ctx, id)
	if err != nil {
		return errors.Wrapf(err, "[service.order] find by id: %v", id)
	}
	if order.MemberID != memberID || order.Status != model.OrderStatusInit {
		return errno.ErrOrderNotFound
	}

	err = s.repo.OrderDelete(ctx, order)
	if err != nil {
		return errors.Wrapf(err, "[service.order] delete")
	}
	//订单取消，解锁库存
	if err = s.wareEvent.Publish(ctx, &ware.Event{
		OrderId: order.ID,
		Finish:  false,
	}); err != nil {
		logger.Warnf("[service.order] event ware stock by order_id: %v", order.ID)
	}

	return nil
}

//MyOrderList 订单列表
func (s *Service) MyOrderList(ctx context.Context, memberID int64, status, offset, limit int) ([]*pb.OrderList, error) {
	list, err := s.repo.GetOrderList(ctx, memberID, status, offset, limit)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.order] list uid: %v", memberID)
	}

	return idl.TransferOrderList(list), nil
}

//OrderPayNotify 支付成功回调处理
func (s *Service) OrderPayNotify(ctx context.Context, memberID int64, amount int, pType int8, orderNo, tradeNo, transHash string) error {
	order, err := s.repo.GetOrderByNo(ctx, orderNo)
	if err != nil {
		return errors.Wrapf(err, "[service.order] get by no: %v", orderNo)
	}
	if order.MemberID != memberID || order.Status != model.OrderStatusInit {
		return errno.ErrOrderNotFound
	}

	//检查是否已支付
	//获取支付方式配置
	pays, err := s.getPayConfig(ctx)
	if err != nil {
		return nil
	}
	//以太坊合约地址
	var address string
	for _, pay := range pays {
		if pay.Id == int64(pType) {
			address = pay.Address
		}
	}
	if address == "" {
		return errno.ErrPayActionInvalid
	}
	if err = s.checkEthPay(ctx, int64(pType), address, orderNo); err != nil {
		return err
	}

	order.Status = model.OrderStatusDelivered
	order.PayAt = time.Now().Unix()
	order.PayType = pType
	order.PayAmount = amount
	order.TradeNo = tradeNo
	order.TransHash = transHash
	if err = s.repo.OrderSave(ctx, nil, order); err != nil {
		return errors.Wrap(err, "[service.order] save")
	}

	return nil
}

//OrderConfirmReceipt 确认收货
func (s *Service) OrderConfirmReceipt(ctx context.Context, memberID, orderID int64) error {
	order, err := s.repo.GetOrderByID(ctx, orderID)
	if err != nil {
		return errors.Wrapf(err, "[service.order] find by id: %v", orderID)
	}
	//已支付，已发货，才可以确认收货
	if order.MemberID != memberID || (order.Status != model.OrderStatusDelivered && order.Status != model.OrderStatusShipped) {
		return errno.ErrOrderNotFound
	}

	order.Status = model.OrderStatusReceived
	order.ReceiveAt = time.Now().Unix()
	order.IsConfirm = 1
	err = s.repo.OrderSave(ctx, nil, order)
	if err != nil {
		return errors.Wrapf(err, "[service.order] save")
	}

	return nil
}

//OrderRefund 退款
func (s *Service) OrderRefund(ctx context.Context, memberID, orderID int64, content string) error {
	order, err := s.repo.GetOrderByID(ctx, orderID)
	if err != nil {
		return errors.Wrapf(err, "[service.order] find by id: %v", orderID)
	}
	//已支付，已收货状态下方可退款
	if order.MemberID != memberID || (order.Status != model.OrderStatusReceived && order.Status != model.OrderStatusDelivered) {
		return errno.ErrOrderNotFound
	}

	// 开启事务
	db := orm.GetDB()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	order.Status = model.OrderStatusPendingRefund
	if err = s.repo.OrderSave(ctx, tx, order); err != nil {
		tx.Rollback()
		return errors.Wrapf(err, "[service.order] save")
	}

	if err = s.repo.CreateOrderRefund(ctx, tx, &model.OrderRefundModel{
		OID:     model.OID{OrderID: order.ID},
		Amount:  order.PayAmount,
		Content: content,
	}); err != nil {
		tx.Rollback()
		return errors.Wrapf(err, "[service.order] create refund")
	}

	//提交事务
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return errors.Wrap(err, "[service.order] tx commit refund")
	}

	return nil
}

//OrderComment 评价
func (s *Service) OrderComment(ctx context.Context, memberID, orderID int64, skuIds []int64, star int8, content, resources string) error {
	order, err := s.repo.GetOrderByID(ctx, orderID)
	if err != nil {
		return errors.Wrapf(err, "[service.order] find by id: %v", orderID)
	}
	//已收货方可评价
	if order.MemberID != memberID || order.Status != model.OrderStatusReceived {
		return errno.ErrOrderNotFound
	}

	// 开启事务
	db := orm.GetDB()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	order.Status = model.OrderStatusFinish
	order.CommentAt = time.Now().Unix()
	if err = s.repo.OrderSave(ctx, tx, order); err != nil {
		tx.Rollback()
		return errors.Wrapf(err, "[service.order] save")
	}

	if err = s.spuComment(ctx, skuIds, memberID, orderID, int32(star), content, resources); err != nil {
		tx.Rollback()
		return err
	}

	//提交事务
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return errors.Wrap(err, "[service.order] tx commit refund")
	}

	//订单完成，扣减库存
	if err = s.wareEvent.Publish(ctx, &ware.Event{
		OrderId: order.ID,
		Finish:  true,
	}); err != nil {
		logger.Warnf("[service.order] event ware stock by order_id: %v", order.ID)
	}

	return nil
}

//OrderInfo 订单详情
func (s *Service) OrderInfo(ctx context.Context, orderID int64) (*pb.OrderInfo, error) {
	order, err := s.repo.GetOrderByID(ctx, orderID)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.order] find by id: %v", orderID)
	}
	//已收货方可评价
	if order.ID == 0 {
		return nil, errno.ErrOrderNotFound
	}
	return idl.TransferOrder(order), nil
}

//saveOrder 保存订单
//分布式事务 - 柔性事务-可靠消息 + 最终一致性方案->rabbitmq延迟队列
func (s *Service) saveOrder(ctx context.Context, memberID, couponID, addressID int64, total int, note, orderNo string, items []*model.OrderItemModel) (int64, error) {
	address, err := s.getAddressInfo(ctx, addressID, memberID)
	if err != nil {
		return 0, err
	}
	//订单金额
	amount := total
	var couponAmount int
	if couponID > 0 {
		coupon, err := s.getCouponInfo(ctx, couponID, memberID)
		if err != nil {
			return 0, err
		}
		couponAmount = util.FormatAmount(coupon.Amount)
		amount -= couponAmount
	}
	//生成订单号
	if orderNo == "" {
		orderNo = util.BuildOrderNo()
	}
	order := &model.OrderModel{
		MID:               orm.MID{MemberID: memberID},
		OrderNo:           orderNo,
		CouponID:          couponID,
		Username:          "",
		TotalAmount:       total,
		PayAmount:         0,
		FreightAmount:     0,
		PromotionAmount:   0,
		IntegrationAmount: 0,
		CouponAmount:      couponAmount,
		DiscountAmount:    0,
		Amount:            amount,
		PayType:           0,
		SourceType:        model.OrderSourceTypeApp,
		Status:            model.OrderStatusInit,
		AutoConfirmDay:    15,
		Integration:       total / 100,
		Growth:            total / 100,
		AddressName:       address.Name,
		AddressPhone:      address.Phone,
		AddressProvince:   address.Province,
		AddressCity:       address.City,
		AddressCounty:     address.County,
		AddressDetail:     address.Detail,
		Note:              note,
		UseIntegration:    0,
	}

	// 开启事务
	db := orm.GetDB()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	//创建订单
	orderID, err := s.repo.CreateOrder(ctx, tx, order)
	if err != nil {
		tx.Rollback()
		return 0, errors.Wrapf(err, "[service.order] create")
	}

	skuNums := make(map[int64]int32, len(items))
	// 子订单批量创建
	for _, item := range items {
		item.OrderID = orderID
		item.OrderNo = orderNo
		skuNums[item.SkuID] = int32(item.Num)
	}

	if err = s.repo.BatchCreateOrderItem(ctx, tx, items); err != nil {
		tx.Rollback()
		return 0, errors.Wrapf(err, "[service.order] create items")
	}

	if couponID > 0 { // 设置优惠券已使用
		if err = s.setCouponUsed(ctx, couponID, memberID, orderID); err != nil {
			return 0, err
		}
	}

	//锁定库存
	if err = s.batchLockSkuStock(ctx, order, skuNums); err != nil {
		return 0, err
	}

	//提交事务
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		//提交订单失败了，库存已锁定成功，需要发送解锁库存消息
		if err = s.wareEvent.Publish(ctx, &ware.Event{
			OrderId: orderID,
			Finish:  false,
		}); err != nil {
			logger.Warnf("[service.order] event ware stock by order_id: %v", orderID)
		}
		return 0, errors.Wrap(err, "[service.order] tx commit")
	}

	//创建订单成功，发送延迟队列，自动取消订单
	msg, _ := json.Marshal(&message.OrderMessage{OrderID: orderID, MemberID: memberID})
	if err = s.event.Publish(msg); err != nil {
		logger.Warnf("[service.order] event order create by order_id: %v", orderID)
	}

	return orderID, nil
}
