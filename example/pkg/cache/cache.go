package cache

import (
	"context"
	"errors"
	"strings"
	"time"
)

const (
	// DefaultExpireTime 默认过期时间
	DefaultExpireTime = time.Hour * 24
	// DefaultNotFoundExpireTime 结果为空时的过期时间 1分钟, 常用于数据为空时的缓存时间(缓存穿透)
	DefaultNotFoundExpireTime = time.Minute
	// NotFoundPlaceholder .
	NotFoundPlaceholder = "*"
)

var (
	// DefaultClient 生成一个缓存客户端，其中keyPrefix 一般为业务前缀
	DefaultClient Cache

	//ErrPlaceholder 空数据标识
	ErrPlaceholder = errors.New("cache: placeholder")
	//ErrSetMemoryWithNotFound 设置缓存失败
	ErrSetMemoryWithNotFound = errors.New("cache: set memory cache err for not found")
)

// Cache 定义cache驱动接口
type Cache interface {
	Set(ctx context.Context, key string, val interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string, val interface{}) error
	MultiSet(ctx context.Context, valMap map[string]interface{}, expiration time.Duration) error
	MultiGet(ctx context.Context, keys []string, valueMap interface{}, newObject func() interface{}) error
	Del(ctx context.Context, keys ...string) error
	SetCacheWithNotFound(ctx context.Context, key string) error
}

// Set 数据
func Set(ctx context.Context, key string, val interface{}, expiration time.Duration) error {
	return DefaultClient.Set(ctx, key, val, expiration)
}

// Get 数据
func Get(ctx context.Context, key string, val interface{}) error {
	return DefaultClient.Get(ctx, key, val)
}

// MultiSet 批量set
func MultiSet(ctx context.Context, valMap map[string]interface{}, expiration time.Duration) error {
	return DefaultClient.MultiSet(ctx, valMap, expiration)
}

// MultiGet 批量获取
func MultiGet(ctx context.Context, keys []string, valueMap interface{}, newObject func() interface{}) error {
	return DefaultClient.MultiGet(ctx, keys, valueMap, newObject)
}

// Del 批量删除
func Del(ctx context.Context, keys ...string) error {
	return DefaultClient.Del(ctx, keys...)
}

// SetCacheWithNotFound 设置空
func SetCacheWithNotFound(ctx context.Context, key string) error {
	return DefaultClient.SetCacheWithNotFound(ctx, key)
}

// BuildCacheKey 构建一个带有前缀的缓存key
func BuildCacheKey(keyPrefix string, key string) (cacheKey string, err error) {
	if key == "" {
		return "", errors.New("[cache] key should not be empty")
	}

	cacheKey = key
	if keyPrefix != "" {
		cacheKey, err = strings.Join([]string{keyPrefix, key}, ":"), nil
	}

	return
}