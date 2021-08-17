package cache

import (
	"context"

	"chat/app/constvar"
	"chat/internal"
)

const _momentTimelineCacheKey = "moment:timeline:count:%d_%d"

// TimelineCache 朋友圈时间线
type TimelineCache struct {
	*internal.Cache
}

// NewTimelineCache new一个朋友圈时间线cache
func NewTimelineCache() *TimelineCache {
	return &TimelineCache{internal.NewCache(_momentTimelineCacheKey, nil)}
}

//SetCache 设置缓存
func (u *TimelineCache) SetCache(ctx context.Context, uid, mid uint32, c int64) error {
	if c == 0 {
		return u.Redis.SetCacheWithNotFound(ctx, u.BuildCacheKey([]interface{}{uid, mid}))
	}
	return u.Redis.Set(ctx, u.BuildCacheKey([]interface{}{uid, mid}), c, constvar.CacheExpireTime)
}

// GetCache 获取缓存
func (u *TimelineCache) GetCache(ctx context.Context, uid, mid uint32) (c int64, err error) {
	err = u.Redis.Get(ctx, u.BuildCacheKey([]interface{}{uid, mid}), &c)
	if err != nil {
		return 0, err
	}
	return c, nil
}
