package repo

import (
	"context"

	"gorm.io/gorm"

	"common/util"
	"market/model"
)

var _ IRepo = (*Repo)(nil)

//IRepo 数据仓库接口
type IRepo interface {
	ICoupon

	GetConfigByName(ctx context.Context, name string, v interface{}) (err error)
	AppPageData(ctx context.Context, page int) (list []*model.AppSettingModel, err error)
	AppHomePageData(ctx context.Context, catID int) (list []*model.AppSettingModel, err error)
	GetNoticeList(ctx context.Context, offset, limit int) (list []*model.AppNoticeModel, err error)

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
