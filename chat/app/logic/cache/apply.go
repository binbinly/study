package cache

import (
	"context"
	"fmt"

	"chat/app/logic/model"
	"chat/pkg/cache"
)


// ApplyCache 好友申请
type ApplyCache struct {
	cache cache.Driver
	//localCache cache.Driver
}

// NewApplyCache new一个好友申请cache
func NewApplyCache() *ApplyCache {
	return &ApplyCache{
		cache: newCache(nil),
	}
}

// GetCacheListKey 申请列表key
func (u *ApplyCache) GetCacheListKey(id uint32) string {
	return fmt.Sprintf(applyListCacheKey, id)
}

//SetCacheList 设置缓存
func (u *ApplyCache) SetCacheList(ctx context.Context, id uint32, field string, list []*model.ApplyModel) error {
	if len(list) == 0 {
		return u.cache.HSetCacheWithNotFound(ctx, u.GetCacheListKey(id), field)
	}
	return u.cache.HSet(ctx, u.GetCacheListKey(id), field, list, defaultExpireTime)
}

//GetCacheList 获取缓存
func (u *ApplyCache) GetCacheList(ctx context.Context, id uint32, field string) (list []*model.ApplyModel, err error) {
	err = u.cache.HGet(ctx, u.GetCacheListKey(id), field, &list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// DelCacheList 删除列表缓存
func (u *ApplyCache) DelCacheList(ctx context.Context, id uint32) error {
	return u.cache.Del(ctx, u.GetCacheListKey(id))
}
