package cache

import (
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

// NewMemoryCache 实例化一个内存cache
func NewMemoryCache(keyPrefix string, encoding Encoding) (Driver, error) {
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
func (m *memoryCache) Set(key string, val interface{}, expiration time.Duration) error {
	buf, err := Marshal(m.encoding, val)
	if err != nil {
		return errors.Wrapf(err, "marshal data err, value is %+v", val)
	}
	m.Store.SetWithTTL(BuildCacheKey(m.KeyPrefix, key), buf, 1, expiration)
	return nil
}

// Get data
func (m *memoryCache) Get(key string, val interface{}) error {
	cacheKey := BuildCacheKey(m.KeyPrefix, key)
	data, ok := m.Store.Get(cacheKey)
	if !ok {
		return nil
	}
	err := Unmarshal(m.encoding, data.([]byte), val)
	if err != nil {
		return errors.Wrapf(err, "unmarshal data error, key=%s, cacheKey=%s type=%v, json is %+v ",
			key, cacheKey, reflect.TypeOf(val), string(data.([]byte)))
	}
	return nil
}

// Del 删除
func (m *memoryCache) Del(keys ...string) error {
	if len(keys) == 0 {
		return nil
	}

	m.Store.Del(BuildCacheKey(m.KeyPrefix, keys[0]))
	return nil
}

func (m *memoryCache) HSet(key string, field string, val interface{}, expiration time.Duration) error {
	return nil
}

func (m *memoryCache) HGet(key string, field string, val interface{}) error {
	return nil
}

// MultiSet 批量set
func (m *memoryCache) MultiSet(valMap map[string]interface{}, expiration time.Duration) error {
	// TODO
	return nil
}

// MultiGet 批量获取
func (m *memoryCache) MultiGet(keys []string, val interface{}) error {
	// TODO
	return nil
}

// Incr 自增
func (m *memoryCache) Incr(key string, step int64) (int64, error) {
	// TODO
	return 0, nil
}

// Decr 自减
func (m *memoryCache) Decr(key string, step int64) (int64, error) {
	// TODO
	return 0, nil
}

func (m *memoryCache) SetCacheWithNotFound(key string) error {
	if m.Store.Set(key, NotFoundPlaceholder, int64(DefaultNotFoundExpireTime)) {
		return nil
	}
	return ErrSetMemoryWithNotFound
}

func (m *memoryCache) HSetCacheWithNotFound(key, field string) error {
	return nil
}