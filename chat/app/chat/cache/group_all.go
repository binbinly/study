package cache

import (
	"context"

	"chat/app/chat/model"
	"chat/app/constvar"
	"chat/internal"
)

const (
	_groupAllCacheKey = "group:all:%d"
)

// GroupCache 群成员
type GroupAllCache struct {
	*internal.Cache
}

// NewGroupAllCache new一个群cache
func NewGroupAllCache() *GroupAllCache {
	return &GroupAllCache{internal.NewCache(_groupAllCacheKey, nil)}
}

//SetCache 设置缓存
func (u *GroupAllCache) SetCache(ctx context.Context, uid uint32, all []*model.GroupList) error {
	if len(all) == 0 {
		return u.Redis.SetCacheWithNotFound(ctx, u.BuildCacheKey([]interface{}{uid}))
	}
	return u.Redis.Set(ctx, u.BuildCacheKey([]interface{}{uid}), all, constvar.CacheExpireTime)
}

//GetCache 获取缓存
func (u *GroupAllCache) GetCache(ctx context.Context, uid uint32) (all []*model.GroupList, err error) {
	err = u.Redis.Get(ctx, u.BuildCacheKey([]interface{}{uid}), &all)
	if err != nil {
		return nil, err
	}
	return all, nil
}
