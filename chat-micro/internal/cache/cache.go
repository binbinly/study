package cache

import (
	"context"
	"time"

	"chat-micro/pkg/cache"
	"chat-micro/pkg/encoding"
	"chat-micro/pkg/redis"
)

//Cache 缓存结构
type Cache struct {
	Redis cache.Cache
	//localCache cache.Driver

	opts Options
}

// NewCache 实例化cache
func NewCache(opts ...Option) *Cache {
	o := NewOptions(opts...)
	return &Cache{
		Redis: cache.NewRedisCache(redis.Client, o.prefix, encoding.JSONEncoding{}),
		opts:  o,
	}
}

// Set 写入缓存
func (c *Cache) Set(ctx context.Context, key string, data interface{}) error {
	return c.Redis.Set(ctx, key, data, c.opts.expire)
}

// SetEX 写入缓存
func (c *Cache) SetEX(ctx context.Context, key string, data interface{}, d time.Duration) error {
	return c.Redis.Set(ctx, key, data, d)
}

// Get 获取缓存
func (c *Cache) Get(ctx context.Context, key string, data interface{}) (err error) {
	return c.Redis.Get(ctx, key, data)
}

//HSet 设置缓存
func (c *Cache) HSet(ctx context.Context, key, field string, data interface{}) error {
	return c.Redis.HSet(ctx, key, field, data, c.opts.expire)
}

//HGet 获取缓存
func (c *Cache) HGet(ctx context.Context, key, field string, data interface{}) (err error) {
	return c.Redis.HGet(ctx, key, field, data)
}

// MultiGet 批量获取用户cache
func (c *Cache) MultiGet(ctx context.Context, keys []string, valueMap interface{}, obj func() interface{}) (err error) {
	return c.Redis.MultiGet(ctx, keys, valueMap, obj)
}

// Del 删除缓存
func (c *Cache) Del(ctx context.Context, key string) error {
	return c.Redis.Del(ctx, key)
}

// SetNotFound 设置空,防缓存穿透
func (c *Cache) SetNotFound(ctx context.Context, key string) error {
	return c.Redis.SetCacheWithNotFound(ctx, key)
}

// HSetNotFound 设置空,防缓存穿透
func (c *Cache) HSetNotFound(ctx context.Context, key, field string) error {
	return c.Redis.HSetCacheWithNotFound(ctx, key, field)
}

// IsNotFound 是否存在
func (c *Cache) IsNotFound(err error) bool {
	return err == cache.ErrPlaceholder
}
