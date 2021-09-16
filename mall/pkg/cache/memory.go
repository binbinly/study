package cache

import (
	"context"
	"reflect"
	"time"

	"github.com/dgraph-io/ristretto"
	"github.com/pkg/errors"
)

type memoryCache struct {
	Store     *ristretto.Cache
	KeyPrefix string
	encoding  Encoding
}

// NewMemoryCache create a memory cache
func NewMemoryCache(keyPrefix string, encoding Encoding) (Driver, error) {
	// see: https://dgraph.io/blog/post/introducing-ristretto-high-perf-go-cache/
	//		https://www.start.io/blog/we-chose-ristretto-cache-for-go-heres-why/
	config := &ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	}
	store, err := ristretto.NewCache(config)
	if err != nil {
		return nil, err
	}
	return &memoryCache{
		Store:     store,
		KeyPrefix: keyPrefix,
		encoding:  encoding,
	}, nil
}

// Set add cache
func (m *memoryCache) Set(ctx context.Context, key string, val interface{}, expiration time.Duration) error {
	buf, err := Marshal(m.encoding, val)
	if err != nil {
		return errors.Wrapf(err, "marshal data err, value is %+v", val)
	}
	cacheKey := BuildCacheKey(m.KeyPrefix, key)
	m.Store.SetWithTTL(cacheKey, buf, 0, expiration)
	return nil
}

// Get data
func (m *memoryCache) Get(ctx context.Context, key string, val interface{}) error {
	cacheKey := BuildCacheKey(m.KeyPrefix, key)
	data, ok := m.Store.Get(cacheKey)
	if !ok {
		return nil
	}
	if data == NotFoundPlaceholder {
		return ErrPlaceholder
	}
	err := Unmarshal(m.encoding, data.([]byte), val)
	if err != nil {
		return errors.Wrapf(err, "unmarshal data error, key=%s, cacheKey=%s type=%v, json is %+v ",
			key, cacheKey, reflect.TypeOf(val), string(data.([]byte)))
	}
	return nil
}

// Del 删除
func (m *memoryCache) Del(ctx context.Context, keys ...string) error {
	if len(keys) == 0 {
		return nil
	}

	m.Store.Del(BuildCacheKey(m.KeyPrefix, keys[0]))
	return nil
}

func (m *memoryCache) HSet(ctx context.Context, key string, field string, val interface{}, expiration time.Duration) error {
	return nil
}

func (m *memoryCache) HGet(ctx context.Context, key string, field string, val interface{}) error {
	return nil
}

// MultiSet 批量set
func (m *memoryCache) MultiSet(ctx context.Context, valMap map[string]interface{}, expiration time.Duration) error {
	// TODO
	return nil
}

// MultiGet 批量获取
func (m *memoryCache) MultiGet(ctx context.Context, keys []string, val interface{}) error {
	// TODO
	return nil
}

// Incr 自增
func (m *memoryCache) Incr(ctx context.Context, key string, step int64) (int64, error) {
	// TODO
	return 0, nil
}

// Decr 自减
func (m *memoryCache) Decr(ctx context.Context, key string, step int64) (int64, error) {
	// TODO
	return 0, nil
}

func (m *memoryCache) SetCacheWithNotFound(ctx context.Context, key string) error {
	if m.Store.Set(key, NotFoundPlaceholder, int64(DefaultNotFoundExpireTime)) {
		return nil
	}
	return ErrSetMemoryWithNotFound
}

func (m *memoryCache) HSetCacheWithNotFound(ctx context.Context, key, field string) error {
	return nil
}