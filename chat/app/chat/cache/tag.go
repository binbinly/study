package cache

import (
	"context"

	"chat/app/chat/model"
	"chat/app/constvar"
	"chat/internal"
)

const _tagAllCacheKey = "tag:all:%d"

// TagCache 用户标签
type TagCache struct {
	*internal.Cache
}

// NewTagCache new一个收藏cache
func NewTagCache() *TagCache {
	return &TagCache{internal.NewCache(_tagAllCacheKey, nil)}
}

//SetCache 设置缓存
func (u *TagCache) SetCache(ctx context.Context, uid uint32, list []*model.UserTag) error {
	if len(list) == 0 {
		return u.Redis.SetCacheWithNotFound(ctx, u.BuildCacheKey([]interface{}{uid}))
	}
	return u.Redis.Set(ctx, u.BuildCacheKey([]interface{}{uid}), list, constvar.CacheExpireTime)
}

// GetCache 获取缓存
func (u *TagCache) GetCache(ctx context.Context, uid uint32) (list []*model.UserTag, err error) {
	err = u.Redis.Get(ctx, u.BuildCacheKey([]interface{}{uid}), &list)
	if err != nil {
		return nil, err
	}
	return list, nil
}
