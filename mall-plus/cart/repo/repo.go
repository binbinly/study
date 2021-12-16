package repo

import (
	"context"

	"github.com/go-redis/redis/v8"

	"cart/model"
)

var _ IRepo = (*Repo)(nil)

//IRepo 数据仓库接口
type IRepo interface {
	AddCart(ctx context.Context, userID int64, cart *model.CartModel) error
	EditCart(ctx context.Context, userID int64, oldID int64, cart *model.CartModel) error
	GetCartByID(ctx context.Context, userID int64, skuID int64) (*model.CartModel, error)
	GetCartsByIds(ctx context.Context, userID int64, ids []int64) ([]*model.CartModel, error)
	DelCart(ctx context.Context, userID int64, ids []int64) error
	EmptyCart(ctx context.Context, userID int64) error
	CartList(ctx context.Context, userID int64) ([]*model.CartModel, error)
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
