package cache

import (
	"context"
	"reflect"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"

	"lib/pkg/encoding"
	"lib/pkg/logger"
)

// redisCache redis cache结构体
type redisCache struct {
	client            *redis.Client
	KeyPrefix         string
	encoding          encoding.Encoding
	DefaultExpireTime time.Duration
}

// NewRedisCache new一个cache, client 参数是可传入的，方便进行单元测试
func NewRedisCache(client *redis.Client, keyPrefix string, encoding encoding.Encoding) Cache {
	return &redisCache{
		client:    client,
		KeyPrefix: keyPrefix,
		encoding:  encoding,
	}
}

func (c *redisCache) Set(ctx context.Context, key string, val interface{}, expiration time.Duration) error {
	buf, err := encoding.Marshal(c.encoding, val)
	if err != nil {
		return errors.Wrapf(err, "marshal data err, value is %+v", val)
	}

	cacheKey, err := BuildCacheKey(c.KeyPrefix, key)
	if err != nil {
		return errors.Wrapf(err, "build cache key err, key is %+v", key)
	}
	if expiration == 0 {
		expiration = DefaultExpireTime
	}
	err = c.client.Set(ctx, cacheKey, buf, expiration).Err()
	if err != nil {
		return errors.Wrapf(err, "redis set err: %+v", err)
	}
	return nil
}

func (c *redisCache) Get(ctx context.Context, key string, val interface{}) error {
	cacheKey, err := BuildCacheKey(c.KeyPrefix, key)
	if err != nil {
		return errors.Wrapf(err, "build cache key err, key is %+v", key)
	}

	bytes, err := c.client.Get(ctx, cacheKey).Bytes()
	if err == redis.Nil {
		return nil
	} else if err != nil {
		return errors.Wrapf(err, "get data error from redis, key is %+v", cacheKey)
	}

	// 防止data为空时，Unmarshal报错
	if string(bytes) == "" {
		return nil
	}
	if string(bytes) == NotFoundPlaceholder {
		return ErrPlaceholder
	}
	err = encoding.Unmarshal(c.encoding, bytes, val)
	if err != nil {
		return errors.Wrapf(err, "unmarshal data error, key=%s, cacheKey=%s type=%v, json is %+v ",
			key, cacheKey, reflect.TypeOf(val), string(bytes))
	}
	return nil
}

func (c *redisCache) MultiSet(ctx context.Context, valueMap map[string]interface{}, expiration time.Duration) error {
	if len(valueMap) == 0 {
		return nil
	}
	if expiration == 0 {
		expiration = DefaultExpireTime
	}
	// key-value是成对的，所以这里的容量是map的2倍
	paris := make([]interface{}, 0, 2*len(valueMap))
	for key, value := range valueMap {
		buf, err := encoding.Marshal(c.encoding, value)
		if err != nil {
			continue
		}
		cacheKey, err := BuildCacheKey(c.KeyPrefix, key)
		if err != nil {
			continue
		}
		paris = append(paris, []byte(cacheKey))
		paris = append(paris, buf)
	}
	pipeline := c.client.Pipeline()
	err := pipeline.MSet(ctx, paris...).Err()
	if err != nil {
		return errors.Wrapf(err, "redis multi set error")
	}
	for i := 0; i < len(paris); i = i + 2 {
		switch paris[i].(type) {
		case []byte:
			pipeline.Expire(ctx, string(paris[i].([]byte)), expiration)
		default:
			logger.Warnf("redis expire is unsupported key type: %+v", reflect.TypeOf(paris[i]))
		}
	}
	_, err = pipeline.Exec(ctx)
	if err != nil {
		return errors.Wrapf(err, "redis multi set pipeline exec error")
	}
	return nil
}

func (c *redisCache) MultiGet(ctx context.Context, keys []string, value interface{}, newObject func() interface{}) error {
	if len(keys) == 0 {
		return nil
	}
	cacheKeys := make([]string, len(keys))
	for index, key := range keys {
		cacheKey, err := BuildCacheKey(c.KeyPrefix, key)
		if err != nil {
			return errors.Wrapf(err, "build cache key err, key is %+v", key)
		}
		cacheKeys[index] = cacheKey
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
		if newObject == nil {//单纯字符串，无需解析，直接填充
			if v, ok := val.(string); ok {
				valueMap.SetMapIndex(reflect.ValueOf(keys[i]), reflect.ValueOf(v))
			}
			continue
		}
		object := newObject()
		if val.(string) == NotFoundPlaceholder {
			valueMap.SetMapIndex(reflect.ValueOf(keys[i]), reflect.ValueOf(object))
			continue
		}

		err = encoding.Unmarshal(c.encoding, []byte(val.(string)), object)
		if err != nil {
			logger.Warnf("unmarshal data error: %+v, key=%s, cacheKey=%s type=%v", err,
				keys[i], cacheKeys[i], reflect.TypeOf(val))
			continue
		}
		valueMap.SetMapIndex(reflect.ValueOf(cacheKeys[i]), reflect.ValueOf(object))
	}
	return nil
}

func (c *redisCache) Del(ctx context.Context, keys ...string) error {
	if len(keys) == 0 {
		return nil
	}

	// 批量构建cacheKey
	cacheKeys := make([]string, len(keys))
	for index, key := range keys {
		cacheKey, err := BuildCacheKey(c.KeyPrefix, key)
		if err != nil {
			continue
		}
		cacheKeys[index] = cacheKey
	}
	err := c.client.Del(ctx, cacheKeys...).Err()
	if err != nil {
		return errors.Wrapf(err, "redis delete error, keys is %+v", keys)
	}
	return nil
}

func (c *redisCache) SetCacheWithNotFound(ctx context.Context, key string) error {
	cacheKey, err := BuildCacheKey(c.KeyPrefix, key)
	if err != nil {
		return errors.Wrapf(err, "build cache key err, key is %+v", key)
	}
	return c.client.Set(ctx, cacheKey, NotFoundPlaceholder, DefaultNotFoundExpireTime).Err()
}
