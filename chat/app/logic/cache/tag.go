package cache

import (
	"context"
	"fmt"

	"chat/app/logic/model"
	"chat/pkg/cache"
)

// TagCache 用户标签
type TagCache struct {
	cache cache.Driver
	//localCache cache.Driver
}

// NewTagCache new一个收藏cache
func NewTagCache() *TagCache {
	return &TagCache{
		cache: newCache(nil),
	}
}

//GetCacheAllKey 获取缓存键
func (u *TagCache) GetCacheAllKey(uid uint32) string {
	return fmt.Sprintf(tagAllCacheKey, uid)
}

//SetCacheAll 设置缓存
func (u *TagCache) SetCacheAll(ctx context.Context, uid uint32, list []*model.UserTag) error {
	if len(list) == 0 {
		return u.cache.SetCacheWithNotFound(ctx, u.GetCacheAllKey(uid))
	}
	return u.cache.Set(ctx, u.GetCacheAllKey(uid), list, defaultExpireTime)
}

// GetCacheAll 获取缓存
func (u *TagCache) GetCacheAll(ctx context.Context, uid uint32) (list []*model.UserTag, err error) {
	err = u.cache.Get(ctx, u.GetCacheAllKey(uid), &list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// DelCacheAll 删除列表缓存
func (u *TagCache) DelCacheAll(ctx context.Context, uid uint32) error {
	return u.cache.Del(ctx, u.GetCacheAllKey(uid))
}