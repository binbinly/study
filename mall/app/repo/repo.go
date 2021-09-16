package repo

import (
	"context"
	"mall/app/model"

	"gorm.io/gorm"

	"mall/app/cache"
)

var _ IRepo = (*Repo)(nil)

//IRepo 数据仓库接口
type IRepo interface {
	IUser
	ICart
	ICoupon
	IUserAddress
	IOrder
	IGoods

	GetAreaAll(ctx context.Context) (list []*model.Area, err error)
	GetConfigByName(ctx context.Context, name string, v interface{}) (err error)
	AppPageData(ctx context.Context, page, catID int) (list []*model.AppSetting, err error)
	GetNoticeList(ctx context.Context, offset, limit int) (list []*model.AppNoticeModel, err error)

	GetAttrsValByIds(ctx context.Context, ids []int) (list []*model.SkuAttrValModel, err error)
	Close() error
}

// Repo mysql struct
type Repo struct {
	db *gorm.DB

	userCache *cache.UserCache
}

// New new a Dao and return
func New(db *gorm.DB) IRepo {
	return &Repo{
		db:        db,
		userCache: cache.NewUserCache(),
	}
}

// Ping ping mysql
func (r *Repo) Ping(c context.Context) error {
	return nil
}

// Close release mysql connection
func (r *Repo) Close() error {
	return nil
}
