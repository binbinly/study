package util

import (
	"common/constvar"
	"context"
	"pkg/cache"
	"pkg/encoding"
	"pkg/redis"
)

const (
	// _prefixKey 缓存前缀
	_prefixKey = "cache"
)

//Cache 缓存结构
type Cache struct {
	Redis cache.Cache
	//localCache cache.Driver

	key string
}

// NewCache 实例化cache
func NewCache() *Cache {
	encoding := encoding.JSONEncoding{}
	return &Cache{
		Redis: cache.NewRedisCache(redis.Client, _prefixKey, encoding),
	}
}

// SetCache 写入缓存
func (c *Cache) SetCache(ctx context.Context, key string, data interface{}) error {
	return c.Redis.Set(ctx, key, data, constvar.CacheExpireTime)
}

// GetCache 获取缓存
func (c *Cache) GetCache(ctx context.Context, key string, data interface{}) (err error) {
	err = c.Redis.Get(ctx, key, data)
	if err != nil {
		return err
	}
	return nil
}

// MultiGetCache 批量获取用户cache
func (c *Cache) MultiGetCache(ctx context.Context, keys []string, valueMap interface{}, obj func() interface{}) (err error) {
	err = c.Redis.MultiGet(ctx, keys, valueMap, obj)
	if err != nil {
		return err
	}
	return nil
}

// DelCache 删除缓存
func (c *Cache) DelCache(ctx context.Context, key string) error {
	return c.Redis.Del(ctx, key)
}

// SetCacheWithNotFound 设置空,防缓存穿透
func (c *Cache) SetCacheWithNotFound(ctx context.Context, key string) error {
	return c.Redis.SetCacheWithNotFound(ctx, key)
}
