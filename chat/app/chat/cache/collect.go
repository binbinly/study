package cache

import (
	"context"

	"chat/app/chat/model"
	"chat/app/constvar"
	"chat/internal"
)

const _collectListCacheKey = "collect:list:%d"

// CollectCache 用户收藏
type CollectCache struct {
	*internal.Cache
}

// NewCollectCache new一个收藏cache
func NewCollectCache() *CollectCache {
	return &CollectCache{internal.NewCache(_collectListCacheKey, nil)}
}

//SetCacheList 设置缓存
func (u *CollectCache) SetCacheList(ctx context.Context, id uint32, field string, list []*model.CollectModel) error {
	if len(list) == 0 {
		return u.Redis.HSetCacheWithNotFound(ctx, u.BuildCacheKey([]interface{}{id}), field)
	}
	return u.Redis.HSet(ctx, u.BuildCacheKey([]interface{}{id}), field, list, constvar.CacheExpireTime)
}

//GetCacheList 获取缓存
func (u *CollectCache) GetCacheList(ctx context.Context, id uint32, field string) (list []*model.CollectModel, err error) {
	err = u.Redis.HGet(ctx, u.BuildCacheKey([]interface{}{id}), field, &list)
	if err != nil {
		return nil, err
	}
	return list, nil
}
