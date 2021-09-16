package cache

import (
	"context"
	"fmt"
	"time"

	"mall/pkg/cache"
	"mall/pkg/redis"
)

const (
	// _prefixKey 缓存前缀
	_prefixKey = "cache"
	// _cacheExpireTime 缓存过期时间
	_cacheExpireTime = time.Hour * 24
)

//Cache 缓存结构
type Cache struct {
	Redis cache.Driver
	//localCache cache.Driver

	key string
}

// NewCache 实例化cache
func NewCache(key string, obj func() interface{}) *Cache {
	encoding := cache.JSONEncoding{}
	return &Cache{
		Redis: cache.NewRedisCache(redis.Client, _prefixKey, encoding, obj),
		key:   key,
	}
}

// BuildCacheKey 构建缓存键
func (c *Cache) BuildCacheKey(ids []interface{}) string {
	return fmt.Sprintf(c.key, ids...)
}

// DelCache 删除缓存
func (c *Cache) DelCache(ctx context.Context, id int) error {
	return c.Redis.Del(ctx, c.BuildCacheKey([]interface{}{id}))
}

// DelCache2 删除缓存
func (c *Cache) DelCache2(ctx context.Context, a, b int) error {
	return c.Redis.Del(ctx, c.BuildCacheKey([]interface{}{a, b}))
}

// DelCacheByKey 删除缓存
func (c *Cache) DelCacheByKey(ctx context.Context, key string) error {
	return c.Redis.Del(ctx, key)
}

// SetCacheWithNotFound 设置空
func (c *Cache) SetCacheWithNotFound(ctx context.Context, id int) error {
	return c.Redis.SetCacheWithNotFound(ctx, c.BuildCacheKey([]interface{}{id}))
}

// SetCacheWithNotFound 设置空
func (c *Cache) SetCacheWithNotFound2(ctx context.Context, a, b int) error {
	return c.Redis.SetCacheWithNotFound(ctx, c.BuildCacheKey([]interface{}{a, b}))
}
