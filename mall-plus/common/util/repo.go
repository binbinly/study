package util

import (
	"context"
	"pkg/lock"
	"reflect"
	"time"

	"github.com/pkg/errors"
	"go-micro.dev/v4/logger"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"

	"pkg/cache"
	"pkg/redis"
)

//ErrNotRecordUpdate 没有记录更新
var ErrNotRecordUpdate = errors.New("Not Record Update")

// Repo mysql struct
type Repo struct {
	DB    *gorm.DB
	Cache *Cache
	Lock  bool
}

// Ping ping mysql
func (r *Repo) Ping(c context.Context) error {
	return nil
}

// Close release mysql connection
func (r *Repo) Close() error {
	db, err := r.DB.DB()
	if err != nil {
		return nil
	}
	return db.Close()
}

// QueryCache 查询启用缓存
// 缓存的更新策略使用 Cache Aside Pattern
// see: https://coolshell.cn/articles/17416.html
func (r *Repo) QueryCache(ctx context.Context, doKey string, data interface{}, query func(interface{}) error) (err error) {
	start := time.Now()
	defer func() {
		logger.Infof("[repo] queryCache key: %v cost: %d μs", doKey, time.Since(start).Microseconds())
	}()

	// 从cache获取
	err = r.Cache.GetCache(ctx, doKey, data)
	if err == cache.ErrPlaceholder {
		return nil
	} else if err != nil && err != redis.Nil {
		return errors.Wrapf(err, "[repo] get cache by key: %s", doKey)
	}

	elem := reflect.ValueOf(data).Elem()
	switch elem.Kind() {
	case reflect.String:
		if elem.String() != "" {
			logger.Infof("[repo] get from string cache, key: %v", doKey)
			return
		}
	default:
		if !elem.IsNil() {
			logger.Infof("[repo] get from cache, key: %v", doKey)
			return
		}
	}

	// use sync/singleflight mode to get data
	// why not use redis lock? see this topic: https://redis.io/topics/distlock
	// demo see: https://github.com/go-demo/singleflight-demo/blob/master/main.go
	// https://juejin.cn/post/6844904084445593613
	getDataFn := func() (interface{}, error) {
		// 从数据库中获取
		err = query(data)
		// if data is empty, set not found cache to prevent cache penetration(缓存穿透)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err = r.Cache.SetCacheWithNotFound(ctx, doKey); err != nil {
				logger.Warnf("[repo] SetCacheWithNotFound err, key: %s", doKey)
			}
			return data, nil
		} else if err != nil {
			return nil, err
		}
		//数组类型，并且长度为0，需要设置空
		if reflect.TypeOf(data).Elem().Kind() == reflect.Slice &&
			reflect.ValueOf(data).Elem().Len() == 0 {
			if err = r.Cache.SetCacheWithNotFound(ctx, doKey); err != nil {
				logger.Warnf("[repo] SetCacheWithNotFound by key: %s", doKey)
			}
			return data, nil
		}

		//set cache
		if err = r.Cache.SetCache(ctx, doKey, data); err != nil {
			return nil, errors.Wrapf(err, "[repo] set data to cache key: %s", doKey)
		}
		return data, nil
	}

	// 开启分布式锁
	if r.Lock {
		lk := lock.NewRedisLock(redis.Client, doKey)
		lk.Lock(ctx, time.Second*15)
		_, err = getDataFn()
		lk.Unlock(ctx)
		if err != nil {
			return errors.Wrapf(err, "[repo] get err via redis lock do key: %s", doKey)
		}
		return nil
	}

	g := singleflight.Group{}

	_, err, _ = g.Do(doKey, getDataFn)
	if err != nil {
		return errors.Wrapf(err, "[repo] get err via single flight do key: %s", doKey)
	}

	return nil
}
