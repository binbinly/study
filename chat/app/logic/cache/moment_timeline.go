package cache

import (
	"context"
	"fmt"

	"chat/pkg/cache"
)

// TimelineCache 朋友圈时间线
type TimelineCache struct {
	cache cache.Driver
	//localCache cache.Driver
}

// NewTimelineCache new一个朋友圈时间线cache
func NewTimelineCache() *TimelineCache {
	return &TimelineCache{
		cache: newCache(nil),
	}
}

//GetCacheKey 获取缓存键
func (u *TimelineCache) GetCacheKey(uid, mid uint32) string {
	return fmt.Sprintf(momentTimelineCacheKey, uid, mid)
}

//SetCache 设置缓存
func (u *TimelineCache) SetCache(ctx context.Context, uid, mid uint32, c int64) error {
	if c == 0 {
		return u.cache.SetCacheWithNotFound(ctx, u.GetCacheKey(uid, mid))
	}
	return u.cache.Set(ctx, u.GetCacheKey(uid, mid), c, defaultExpireTime)
}

// GetCache 获取缓存
func (u *TimelineCache) GetCache(ctx context.Context, uid, mid uint32) (c int64, err error) {
	err = u.cache.Get(ctx, u.GetCacheKey(uid, mid), &c)
	if err != nil {
		return 0, err
	}
	return c, nil
}