package repo

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"mall/app/model"
)

//IOrder 订单接口
type IOrder interface {
	CreateOrder(ctx context.Context, tx *gorm.DB, order *model.OrderModel) (int, error)
	OrderSave(ctx context.Context, tx *gorm.DB, order *model.OrderModel) error
	OrderDelete(ctx context.Context, order *model.OrderModel) error
	GetOrderByID(ctx context.Context, id, userID int) (*model.OrderModel, error)
	GetOrderByNo(ctx context.Context, userID int, orderNo string) (*model.OrderModel, error)
	GetOrderList(ctx context.Context, userID, status, offset, limit int) (list []*model.OrderModel, err error)
	CreateOrderGoods(ctx context.Context, tx *gorm.DB, goods []*model.OrderGoodsModel) error
	CreateOrderAddress(ctx context.Context, tx *gorm.DB, address *model.OrderAddressModel) error
	CreateOrderRefund(ctx context.Context, tx *gorm.DB, refund *model.OrderRefundModel) error
}

//CreateOrder 创建订单
func (r *Repo) CreateOrder(ctx context.Context, tx *gorm.DB, order *model.OrderModel) (int, error) {
	err := tx.WithContext(ctx).Create(order).Error
	if err != nil {
		return 0, errors.Wrapf(err, "[repo.order] create")
	}
	return order.ID, nil
}

//GetOrderByID 获取订单详情
func (r *Repo) GetOrderByID(ctx context.Context, id, userID int) (*model.OrderModel, error) {
	order := new(model.OrderModel)
	err := r.db.WithContext(ctx).Where("id=? and user_id=?", id, userID).
		Preload("Goods").Preload("Address").First(order).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(err, "[repo.order] detail id: %v", id)
	}
	return order, nil
}

//GetOrderByNo 订单号获取订单信息
func (r *Repo) GetOrderByNo(ctx context.Context, userID int, orderNo string) (*model.OrderModel, error) {
	order := new(model.OrderModel)
	err := r.db.WithContext(ctx).Where("order_no=? and user_id=?", orderNo, userID).First(order).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(err, "[repo.order] detail no: %v", orderNo)
	}
	return order, nil
}

//GetOrderList 用户订单列表
func (r *Repo) GetOrderList(ctx context.Context, userID, status, offset, limit int) (list []*model.OrderModel, err error) {
	err = r.db.WithContext(ctx).Where("user_id=?", userID).Order(model.DefaultOrder).
		Scopes(orderScopesStatus(status),model.OffsetPage(offset, limit)).Preload("Goods").Find(&list).Error
	if err != nil {
		return nil, errors.Wrapf(err, "[repo.order] find")
	}
	return
}

//OrderSave 订单保存
func (r *Repo) OrderSave(ctx context.Context, tx *gorm.DB, order *model.OrderModel) (err error) {
	if tx == nil {
		err = r.db.WithContext(ctx).Save(order).Error
	} else {
		err = tx.WithContext(ctx).Save(order).Error
	}
	if err != nil {
		return errors.Wrapf(err, "[repo.order] save")
	}
	return nil
}

//OrderDelete 订单删除
func (r *Repo) OrderDelete(ctx context.Context, order *model.OrderModel) error {
	err := r.db.WithContext(ctx).Delete(order).Error
	if err != nil {
		return errors.Wrapf(err, "[repo.order] delete")
	}

	return nil
}

//orderScopesStatus 状态筛选
func orderScopesStatus(status int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if status == 0 {
			return db
		}
		return db.Where("status=?", status)
	}
}