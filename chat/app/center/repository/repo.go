package repository

import (
	"context"

	"gorm.io/gorm"

	"chat/app/center/cache"
)

var _ IRepo = (*Repo)(nil)

//IRepo 数据仓库接口
type IRepo interface {
	IUser

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
	db, err := r.db.DB()
	if err != nil {
		return nil
	}
	return db.Close()
}
