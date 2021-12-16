package repo

import (
	"context"

	"gorm.io/gorm"

	"common/util"
	"member/model"
)

var _ IRepo = (*Repo)(nil)

//IRepo 数据仓库接口
type IRepo interface {
	IMember
	IMemberAddress
	CreateLoginLog(ctx context.Context, log *model.LoginLogModel) error

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
