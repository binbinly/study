package cache

import (
	"context"
	"time"

	"github.com/pkg/errors"
)

const (
	// DefaultExpireTime 默认过期时间
	DefaultExpireTime = time.Hour * 24
	// DefaultNotFoundExpireTime 结果为空时的过期时间 5分钟, 常用于数据为空时的缓存时间(缓存穿透)
	DefaultNotFoundExpireTime = time.Minute * 5
	// NotFoundPlaceholder .
	NotFoundPlaceholder = "*"
)

var (
	//ErrPlaceholder 空数据标识
	ErrPlaceholder = errors.New("cache: placeholder")
	//ErrSetMemoryWithNotFound 设置缓存失败
	ErrSetMemoryWithNotFound = errors.New("cache: set memory cache err for not found")
)

// Client 生成一个缓存客户端，其中keyPrefix 一般为业务前缀
var Client Driver

// Driver 定义cache驱动接口
type Driver interface {
	Set(ctx context.Context, key string, val interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string, val interface{}) error
	HSet(ctx context.Context, key string, field string, val interface{}, expiration time.Duration) error
	HGet(ctx context.Context, key string, field string, val interface{}) error
	MultiSet(ctx context.Context, valMap map[string]interface{}, expiration time.Duration) error
	MultiGet(ctx context.Context, keys []string, valueMap interface{}) error
	Del(ctx context.Context, keys ...string) error
	Incr(ctx context.Context, key string, step int64) (int64, error)
	Decr(ctx context.Context, key string, step int64) (int64, error)
	SetCacheWithNotFound(ctx context.Context, key string) error
	HSetCacheWithNotFound(ctx context.Context, key, field string) error
}

// Set 设置缓存
func Set(ctx context.Context, key string, val interface{}, expiration time.Duration) error {
	return Client.Set(ctx, key, val, expiration)
}

// Get 获取缓存
func Get(ctx context.Context, key string, val interface{}) error {
	return Client.Get(ctx, key, val)
}

// MultiSet 批量设置缓存
func MultiSet(ctx context.Context, valMap map[string]interface{}, expiration time.Duration) error {
	return Client.MultiSet(ctx, valMap, expiration)
}

// MultiGet 批量获取缓存
func MultiGet(ctx context.Context, keys []string, valueMap interface{}) error {
	return Client.MultiGet(ctx, keys, valueMap)
}

// Del 删除缓存
func Del(ctx context.Context, keys ...string) error {
	return Client.Del(ctx, keys...)
}

// Incr 自增
func Incr(ctx context.Context, key string, step int64) (int64, error) {
	return Client.Incr(ctx, key, step)
}

// Decr 自减
func Decr(ctx context.Context, key string, step int64) (int64, error) {
	return Client.Decr(ctx, key, step)
}

// SetCacheWithNotFound 设置空
func SetCacheWithNotFound(ctx context.Context, key string) error {
	return Client.SetCacheWithNotFound(ctx, key)
}
