package cache

import (
	"github.com/pkg/errors"
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
	ErrPlaceholder           = errors.New("cache: placeholder")
	ErrSetMemoryWithNotFound = errors.New("cache: set memory cache err for not found")
)

// Client 生成一个缓存客户端，其中keyPrefix 一般为业务前缀
var Client Driver

// Driver 定义cache驱动接口
type Driver interface {
	Set(key string, val interface{}, expiration time.Duration) error
	Get(key string, val interface{}) error
	HSet(key string, field string, val interface{}, expiration time.Duration) error
	HGet(key string, field string, val interface{}) error
	MultiSet(valMap map[string]interface{}, expiration time.Duration) error
	MultiGet(keys []string, valueMap interface{}) error
	Del(keys ...string) error
	Incr(key string, step int64) (int64, error)
	Decr(key string, step int64) (int64, error)
	SetCacheWithNotFound(key string) error
	HSetCacheWithNotFound(key, field string) error
}

// Set 设置缓存
func Set(key string, val interface{}, expiration time.Duration) error {
	return Client.Set(key, val, expiration)
}

// Get 获取缓存
func Get(key string, val interface{}) error {
	return Client.Get(key, val)
}

// MultiSet 批量设置缓存
func MultiSet(valMap map[string]interface{}, expiration time.Duration) error {
	return Client.MultiSet(valMap, expiration)
}

// MultiGet 批量获取缓存
func MultiGet(keys []string, valueMap interface{}) error {
	return Client.MultiGet(keys, valueMap)
}

// Del 删除缓存
func Del(keys ...string) error {
	return Client.Del(keys...)
}

// Incr 自增
func Incr(key string, step int64) (int64, error) {
	return Client.Incr(key, step)
}

// Decr 自减
func Decr(key string, step int64) (int64, error) {
	return Client.Decr(key, step)
}

// SetCacheWithNotFound 设置空
func SetCacheWithNotFound(key string) error {
	return Client.SetCacheWithNotFound(key)
}
