package app

import (
	"context"

	"mall/pkg/redis"
)

// IDAlloc define struct ID生成器
type IDAlloc struct {
	idGenerator *redis.IDAlloc
}

// NewIDAlloc create a id alloc
func NewIDAlloc() *IDAlloc {
	return &IDAlloc{
		idGenerator: redis.NewIDAlloc(redis.Client),
	}
}

// GetUserID generate user id from redis
func (i *IDAlloc) GetUserID() (int64, error) {
	return i.idGenerator.GetNewID(context.Background(), "user_id", 1)
}
