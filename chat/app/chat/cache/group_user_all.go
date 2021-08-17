package cache

import (
	"context"

	"chat/app/chat/model"
	"chat/app/constvar"
	"chat/internal"
)

const (
	_groupUserAllCacheKey = "group_user:all:%d"
)

// GroupUserAllCache 群成员
type GroupUserAllCache struct {
	*internal.Cache
}

// NewGroupUserAllCache new一个群成员cache
func NewGroupUserAllCache() *GroupUserAllCache {
	return &GroupUserAllCache{internal.NewCache(_groupUserAllCacheKey, nil)}
}

//SetCache 设置缓存
func (u *GroupUserAllCache) SetCache(ctx context.Context, uid uint32, all []*model.GroupUserModel) error {
	if len(all) == 0 {
		return u.Redis.SetCacheWithNotFound(ctx, u.BuildCacheKey([]interface{}{uid}))
	}
	return u.Redis.Set(ctx, u.BuildCacheKey([]interface{}{uid}), all, constvar.CacheExpireTime)
}

//GetCache 获取缓存
func (u *GroupUserAllCache) GetCache(ctx context.Context, uid uint32) (all []*model.GroupUserModel, err error) {
	err = u.Redis.Get(ctx, u.BuildCacheKey([]interface{}{uid}), &all)
	if err != nil {
		return nil, err
	}
	return all, nil
}
