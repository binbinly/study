package repo

import (
	"context"

	"gorm.io/gorm"

	"common/util"
	"warehouse/model"
)

var _ IRepo = (*Repo)(nil)

//IRepo 数据仓库接口
type IRepo interface {
	GetWareSkuStock(ctx context.Context, skuID int64) (*model.WareSkuStock, error)
	BatchGetWareSkuStocks(ctx context.Context, spuID int64, skuIds []int64) ([]*model.WareSkuStock, error)
	BatchGetWareSkus(ctx context.Context, skuIds []int64) ([]*model.WareSkuModel, error)
	WareSkuSave(ctx context.Context, tx *gorm.DB, ware *model.WareSkuModel) error
	CreateWareTask(ctx context.Context, tx *gorm.DB, task *model.WareTaskModel) error
	UpdateWareTaskStatus(ctx context.Context, tx *gorm.DB, orderID int64, status int) error
	BatchCreateWareTaskDetail(ctx context.Context, tx *gorm.DB, items []*model.WareTaskDetailModel) error
	GetTaskByOrderID(ctx context.Context, orderID int64) (task *model.WareTaskModel, err error)
	GetTaskDetailByID(ctx context.Context, taskID int64) (list []*model.WareTaskDetailModel, err error)

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
