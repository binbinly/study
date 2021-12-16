package repo

import (
	"context"

	"gorm.io/gorm"

	"common/util"
	"order/model"
)

var _ IRepo = (*Repo)(nil)

//IRepo 数据仓库接口
type IRepo interface {
	CreateOrder(ctx context.Context, tx *gorm.DB, order *model.OrderModel) (int64, error)
	GetOrderDetail(ctx context.Context, id, memberID int64) (*model.OrderModel, error)
	GetOrderByID(ctx context.Context, id int64) (*model.OrderModel, error)
	GetOrderByNo(ctx context.Context, orderNo string) (*model.OrderModel, error)
	GetOrderList(ctx context.Context, memberID int64, status, offset, limit int) (list []*model.OrderModel, err error)
	OrderSave(ctx context.Context, tx *gorm.DB, order *model.OrderModel) (err error)
	OrderDelete(ctx context.Context, order *model.OrderModel) error
	CreateOrderRefund(ctx context.Context, tx *gorm.DB, refund *model.OrderRefundModel) error
	BatchCreateOrderItem(ctx context.Context, tx *gorm.DB, items []*model.OrderItemModel) error
	CreateOrderBill(ctx context.Context, tx *gorm.DB, bill *model.OrderBillModel) error
	Close() error
}

// Repo mysql struct
type Repo struct {
	util.Repo
}

// New new a Dao and return
func New(db *gorm.DB, cache *util.Cache) IRepo {
	return &Repo{util.Repo{
		DB:    db,
		Cache: cache,
	}}
}
