package cache

import (
	"context"
	"reflect"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"

	"mall/pkg/log"
)

// redisCache redis cache结构体
type redisCache struct {
	client            *redis.Client
	KeyPrefix         string
	encoding          Encoding
	DefaultExpireTime time.Duration
	newObject         func() interface{}
}

// NewRedisCache new一个cache cache, redis client 参数是可传入的，这样方便进行单元测试
func NewRedisCache(client *redis.Client, keyPrefix string, encoding Encoding, newObject func() interface{}) Driver {
	return &redisCache{
		client:    client,
		KeyPrefix: keyPrefix,
		encoding:  encoding,
		newObject: newObject,
	}
}

// Set 设置缓存
func (c *redisCache) Set(ctx context.Context, key string, val interface{}, expiration time.Duration) error {
	buf, err := Marshal(c.encoding, val)
	if err != nil {
		return errors.Wrapf(err, "marshal data err, value is %+v", val)
	}

	if expiration == 0 {
		expiration = DefaultExpireTime
	}
	err = c.client.Set(ctx, BuildCacheKey(c.KeyPrefix, key), buf, expiration).Err()
	if err != nil {
		return errors.Wrapf(err, "redis set error")
	}
	return nil
}

// Get 获取缓存
func (c *redisCache) Get(ctx context.Context, key string, val interface{}) error {
	cacheKey := BuildCacheKey(c.KeyPrefix, key)
	data, err := c.client.Get(ctx, cacheKey).Bytes()
	if err != nil {
		if err != redis.Nil {
			return errors.Wrapf(err, "get data error from redis, key is %+v", cacheKey)
		}
	}

	// 防止data为空时，Unmarshal报错
	if string(data) == "" {
		return nil
	}
	if string(data) == NotFoundPlaceholder {
		return ErrPlaceholder
	}
	err = Unmarshal(c.encoding, data, val)
	if err != nil {
		return errors.Wrapf(err, "unmarshal data error, key=%s, cacheKey=%s type=%v, json is %+v ",
			key, cacheKey, reflect.TypeOf(val), string(data))
	}
	return nil
}

func (c *redisCache) HSet(ctx context.Context, key string, field string, val interface{}, expiration time.Duration) error {
	buf, err := Marshal(c.encoding, val)
	if err != nil {
		return errors.Wrapf(err, "marshal data err, value is %+v", val)
	}

	if expiration == 0 {
		expiration = DefaultExpireTime
	}
	cacheKey := BuildCacheKey(c.KeyPrefix, key)
	err = c.client.HSet(ctx, cacheKey, field, buf).Err()
	if err != nil {
		return errors.Wrapf(err, "redis hSet err")
	}
	err = c.client.Expire(ctx, cacheKey, expiration).Err()
	if err != nil {
		return errors.Wrapf(err, "redis expire err")
	}
	return nil
}

func (c *redisCache) HGet(ctx context.Context, key string, field string, val interface{}) error {
	cacheKey := BuildCacheKey(c.KeyPrefix, key)
	data, err := c.client.HGet(ctx, cacheKey, field).Bytes()
	if err != nil {
		if err != redis.Nil {
			return errors.Wrapf(err, "hGet data error from redis, key is %+v", cacheKey)
		}
	}

	// 防止data为空时，Unmarshal报错
	if string(data) == "" {
		return nil
	}
	if string(data) == NotFoundPlaceholder {
		return ErrPlaceholder
	}
	err = Unmarshal(c.encoding, data, val)
	if err != nil {
		return errors.Wrapf(err, "unmarshal data error, key=%s, cacheKey=%s type=%v, json is %+v ",
			key, cacheKey, reflect.TypeOf(val), string(data))
	}
	return nil
}

// MultiSet 批量设置缓存
func (c *redisCache) MultiSet(ctx context.Context, valueMap map[string]interface{}, expiration time.Duration) error {
	if len(valueMap) == 0 {
		return nil
	}
	// key-value是成对的，所以这里的容量是map的2倍
	paris := make([]interface{}, 0, 2*len(valueMap))
	for key, value := range valueMap {
		buf, err := Marshal(c.encoding, value)
		if err != nil {
			continue
		}
		cacheKey := BuildCacheKey(c.KeyPrefix, key)
		if expiration == 0 {
			expiration = DefaultExpireTime
		}
		paris = append(paris, []byte(cacheKey))
		paris = append(paris, buf)
	}
	if expiration == 0 {
		expiration = DefaultExpireTime
	}
	err := c.client.MSet(ctx, paris...).Err()
	if err != nil {
		return errors.Wrapf(err, "redis multi set error")
	}
	for i := 0; i < len(paris); i = i + 2 {
		switch paris[i].(type) {
		case []byte:
			c.client.Expire(ctx, string(paris[i].([]byte)), expiration)
		}
	}
	return nil
}

// MultiGet 批量获取缓存
func (c *redisCache) MultiGet(ctx context.Context, keys []string, value interface{}) error {
	if len(keys) == 0 {
		return nil
	}
	cacheKeys := make([]string, len(keys))
	for index, key := range keys {
		cacheKeys[index] = BuildCacheKey(c.KeyPrefix, key)
	}
	values, err := c.client.MGet(ctx, cacheKeys...).Result()
	if err != nil {
		return errors.Wrapf(err, "redis MGet error, keys is %+v", keys)
	}

	// 通过反射注入到map
	valueMap := reflect.ValueOf(value)
	for i, val := range values {
		if val == nil {
			continue
		}
		object := c.newObject()
		if val.(string) == NotFoundPlaceholder {
			valueMap.SetMapIndex(reflect.ValueOf(keys[i]), reflect.ValueOf(object))
			continue
		}

		err = Unmarshal(c.encoding, []byte(val.(string)), &object)
		if err != nil {
			log.Warnf("unmarshal data error: %+v, key=%s, type=%v val=%v", err,
				keys[i], reflect.TypeOf(val), val)
			continue
		}
		valueMap.SetMapIndex(reflect.ValueOf(keys[i]), reflect.ValueOf(object))
	}
	return nil
}

// Del 删除缓存
func (c *redisCache) Del(ctx context.Context, keys ...string) error {
	if len(keys) == 0 {
		return nil
	}

	// 批量构建cacheKey
	cacheKeys := make([]string, len(keys))
	for index, key := range keys {
		cacheKeys[index] = BuildCacheKey(c.KeyPrefix, key)
	}
	err := c.client.Del(ctx, cacheKeys...).Err()
	if err != nil {
		return errors.Wrapf(err, "redis delete error, keys is %+v", keys)
	}
	return nil
}

// Incr 原子自增
func (c *redisCache) Incr(ctx context.Context, key string, step int64) (int64, error) {
	cacheKey := BuildCacheKey(c.KeyPrefix, key)
	affectRow, err := c.client.IncrBy(ctx, cacheKey, step).Result()
	if err != nil {
		return 0, errors.Wrapf(err, "redis incr, keys is %+v", key)
	}
	return affectRow, nil
}

// Decr 原子自减
func (c *redisCache) Decr(ctx context.Context, key string, step int64) (int64, error) {
	cacheKey := BuildCacheKey(c.KeyPrefix, key)
	affectRow, err := c.client.DecrBy(ctx, cacheKey, step).Result()
	if err != nil {
		return 0, errors.Wrapf(err, "redis incr, keys is %+v", key)
	}
	return affectRow, nil
}

// SetCacheWithNotFound 设置空值
func (c *redisCache) SetCacheWithNotFound(ctx context.Context, key string) error {
	return c.client.Set(ctx, BuildCacheKey(c.KeyPrefix, key), NotFoundPlaceholder, DefaultNotFoundExpireTime).Err()
}

// HSetCacheWithNotFound 设置空值
func (c *redisCache) HSetCacheWithNotFound(ctx context.Context, key, field string) error {
	cacheKey := BuildCacheKey(c.KeyPrefix, key)
	err := c.client.HSet(ctx, cacheKey, field, NotFoundPlaceholder).Err()
	if err != nil {
		return errors.Wrapf(err, "redis hset empty err")
	}
	return c.client.Expire(ctx, cacheKey, DefaultNotFoundExpireTime).Err()
}
