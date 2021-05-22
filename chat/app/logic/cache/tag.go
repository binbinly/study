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

func (u *TagCache) GetCacheAllKey(uid uint32) string {
	return fmt.Sprintf(TagAllCacheKey, uid)
}

func (u *TagCache) SetCacheAll(ctx context.Context, uid uint32, list []*model.UserTag) error {
	if len(list) == 0 {
		return u.cache.SetCacheWithNotFound(u.GetCacheAllKey(uid))
	}
	return u.cache.Set(u.GetCacheAllKey(uid), list, defaultExpireTime)
}

func (u *TagCache) GetCacheAll(ctx context.Context, uid uint32) (list []*model.UserTag, err error) {
	err = u.cache.Get(u.GetCacheAllKey(uid), &list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// DelCacheList 删除列表缓存
func (u *TagCache) DelCacheAll(ctx context.Context, uid uint32) error {
	return u.cache.Del(u.GetCacheAllKey(uid))
}