package repo

import (
	"chat-micro/app/logic/model"
	"context"
	"reflect"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"

	"chat-micro/internal/cache"
	"chat-micro/internal/orm"
	"chat-micro/pkg/logger"
	"chat-micro/pkg/redis"
)

var _ IRepo = (*Repo)(nil)

//IRepo 数据仓库接口
type IRepo interface {
	IUser
	IApply
	ICollect
	IFriend
	IGroup
	IGroupUser
	IMessage
	IMoment
	IMomentComment
	IMomentTimeline
	IMomentLike
	IReport
	IUserTag
	IReport
	IEmoticon

	CreateMessage(ctx context.Context, message model.MessageModel) (id uint32, err error)
	Close() error
}

// Repo mysql struct
type Repo struct {
	db    *gorm.DB
	cache *cache.Cache
}

// New new a Dao and return
func New(db *gorm.DB, cache *cache.Cache) IRepo {
	return &Repo{
		db:    db,
		cache: cache,
	}
}

// Close release mysql connection
func (r *Repo) Close() error {
	return orm.CloseDB()
}

// queryCache 查询启用缓存
// 缓存的更新策略使用 Cache Aside Pattern
// see: https://coolshell.cn/articles/17416.html
func (r *Repo) queryCache(ctx context.Context, doKey string, data interface{}, query func(interface{}) error) (err error) {
	start := time.Now()
	defer func() {
		logger.Infof("[repo] queryCache key: %v cost: %d μs", doKey, time.Since(start).Microseconds())
	}()

	// 从cache获取
	err = r.cache.Get(ctx, doKey, data)
	if r.cache.IsNotFound(err) {
		//实例化
		if reflect.TypeOf(data).Elem().Kind() == reflect.Ptr {
			reflect.ValueOf(data).Elem().Set(reflect.New(reflect.TypeOf(data).Elem().Elem()))
		}
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
			if err = r.cache.SetNotFound(ctx, doKey); err != nil {
				logger.Warnf("[repo] SetCacheWithNotFound err, key: %s", doKey)
			}
			return data, nil
		} else if err != nil {
			return nil, err
		}
		//数组类型，并且长度为0，需要设置空
		if reflect.TypeOf(data).Elem().Kind() == reflect.Slice &&
			reflect.ValueOf(data).Elem().Len() == 0 {
			if err = r.cache.SetNotFound(ctx, doKey); err != nil {
				logger.Warnf("[repo] SetCacheWithNotFound by key: %s", doKey)
			}
			return data, nil
		}

		//set cache
		if err = r.cache.Set(ctx, doKey, data); err != nil {
			return nil, errors.Wrapf(err, "[repo] set data to cache key: %s", doKey)
		}
		return data, nil
	}

	g := singleflight.Group{}

	_, err, _ = g.Do(doKey, getDataFn)
	if err != nil {
		return errors.Wrapf(err, "[repo] get err via single flight do key: %s", doKey)
	}

	return nil
}

// queryListCache 查询分页列表启用缓存
// 缓存的更新策略使用 Cache Aside Pattern
// see: https://coolshell.cn/articles/17416.html
func (r *Repo) queryListCache(ctx context.Context, doKey string, offset int, data interface{}, query func(interface{}) error) (err error) {
	start := time.Now()
	defer func() {
		logger.Infof("[repo] queryListCache key: %v offset: %v cost: %d μs", doKey, offset, time.Since(start).Microseconds())
	}()

	field := strconv.Itoa(offset)
	// 从cache获取
	err = r.cache.HGet(ctx, doKey, field, data)
	if r.cache.IsNotFound(err) {
		return nil
	} else if err != nil && err != redis.Nil {
		return errors.Wrapf(err, "[repo] hget cache by key: %s", doKey)
	}

	if !reflect.ValueOf(data).Elem().IsNil() {
		logger.Infof("[repo] hget from cache, key: %v", doKey)
		return
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
			if err = r.cache.HSetNotFound(ctx, doKey, field); err != nil {
				logger.Warnf("[repo] HSetNotFound err, key: %s", doKey)
			}
			return data, nil
		} else if err != nil {
			return nil, err
		}

		//长度为0，需要设置空
		if reflect.ValueOf(data).Elem().Len() == 0 {
			if err = r.cache.HSetNotFound(ctx, doKey, field); err != nil {
				logger.Warnf("[repo] HSetNotFound by key: %s", doKey)
			}
			return data, nil
		}

		//set cache
		if err = r.cache.HSet(ctx, doKey, field, data); err != nil {
			return nil, errors.Wrapf(err, "[repo] hset data to cache key: %s", doKey)
		}
		return data, nil
	}

	g := singleflight.Group{}

	_, err, _ = g.Do(doKey, getDataFn)
	if err != nil {
		return errors.Wrapf(err, "[repo] get err via single flight do key: %s", doKey)
	}

	return nil
}