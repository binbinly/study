package cache

import (
	"context"

	"chat/app/chat/model"
	"chat/app/constvar"
	"chat/internal"
)

const (
	_groupCacheKey = "group:%d"
)

// GroupCache 群成员
type GroupCache struct {
	*internal.Cache
}

// NewGroupCache new一个群cache
func NewGroupCache() *GroupCache {
	return &GroupCache{internal.NewCache(_groupCacheKey, nil)}
}

//SetCache 设置缓存
func (u *GroupCache) SetCache(ctx context.Context, id uint32, data *model.GroupModel) error {
	if data == nil || data.ID == 0 {
		return nil
	}
	return u.Redis.Set(ctx, u.BuildCacheKey([]interface{}{id}), data, constvar.CacheExpireTime)
}

//GetCache 获取缓存
func (u *GroupCache) GetCache(ctx context.Context, id uint32) (data *model.GroupModel, err error) {
	err = u.Redis.Get(ctx, u.BuildCacheKey([]interface{}{id}), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
