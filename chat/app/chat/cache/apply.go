package cache

import (
	"context"

	"chat/app/chat/model"
	"chat/app/constvar"
	"chat/internal"
)

const _applyListCacheKey = "apply:list:%d"

// ApplyCache 好友申请
type ApplyCache struct {
	*internal.Cache
}

// NewApplyCache new一个好友申请cache
func NewApplyCache() *ApplyCache {
	return &ApplyCache{internal.NewCache(_applyListCacheKey, nil)}
}

//SetCacheList 设置缓存
func (u *ApplyCache) SetCacheList(ctx context.Context, id uint32, field string, list []*model.ApplyModel) error {
	if len(list) == 0 {
		return u.Redis.HSetCacheWithNotFound(ctx, u.BuildCacheKey([]interface{}{id}), field)
	}
	return u.Redis.HSet(ctx, u.BuildCacheKey([]interface{}{id}), field, list, constvar.CacheExpireTime)
}

//GetCacheList 获取缓存
func (u *ApplyCache) GetCacheList(ctx context.Context, id uint32, field string) (list []*model.ApplyModel, err error) {
	err = u.Redis.HGet(ctx, u.BuildCacheKey([]interface{}{id}), field, &list)
	if err != nil {
		return nil, err
	}
	return list, nil
}
