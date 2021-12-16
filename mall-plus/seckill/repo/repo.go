package repo

import (
	"context"

	"github.com/go-redis/redis/v8"

	"seckill/model"
)

var _ IRepo = (*Repo)(nil)

//IRepo 数据仓库接口
type IRepo interface {
	GetSessionAll(ctx context.Context) ([]*model.SessionModel, error)
	GetSkuByID(ctx context.Context, skuID int64) (*model.SkuModel, error)
	GetSkusBySessionID(ctx context.Context, sessionID int64) ([]*model.SkuModel, error)
}

// Repo mysql struct
type Repo struct {
	redis *redis.Client
}

// New new a Dao and return
func New(redis *redis.Client) IRepo {
	return &Repo{
		redis: redis,
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
