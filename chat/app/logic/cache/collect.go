package cache

import (
	"context"
	"fmt"

	"chat/app/logic/model"
	"chat/pkg/cache"
)


// CollectCache 用户收藏
type CollectCache struct {
	cache cache.Driver
	//localCache cache.Driver
}

// NewCollectCache new一个收藏cache
func NewCollectCache() *CollectCache {
	return &CollectCache{
		cache: newCache(nil),
	}
}

func (u *CollectCache) GetCacheListKey(id uint32) string {
	return fmt.Sprintf(CollectListCacheKey, id)
}

func (u *CollectCache) SetCacheList(ctx context.Context, id uint32, field string, list []*model.CollectModel) error {
	if len(list) == 0 {
		return u.cache.HSetCacheWithNotFound(u.GetCacheListKey(id), field)
	}
	return u.cache.HSet(u.GetCacheListKey(id), field, list, defaultExpireTime)
}

func (u *CollectCache) GetCacheList(ctx context.Context, id uint32, field string) (list []*model.CollectModel, err error) {
	err = u.cache.HGet(u.GetCacheListKey(id), field, &list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// DelCacheList 删除列表缓存
func (u *CollectCache) DelCacheList(ctx context.Context, id uint32) error {
	return u.cache.Del(u.GetCacheListKey(id))
}