package pkg

import "chat/pkg/redis"

// IDAlloc define struct
type IDAlloc struct {
	idGenerator *redis.IDAlloc
}

// NewIDAlloc create a id alloc
func NewIdAlloc() *IDAlloc {
	return &IDAlloc{
		idGenerator: redis.NewIDAlloc(redis.Client),
	}
}

// GetUserID generate user id from redis
func (i *IDAlloc) GetUserId() (int64, error) {
	return i.idGenerator.GetNewID("user_id", 1)
}