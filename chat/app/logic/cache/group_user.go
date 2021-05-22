package cache

import (
	"context"
	"fmt"

	"chat/app/logic/model"
	"chat/pkg/cache"
	"chat/pkg/redis"
)


// GroupUserCache 群成员
type GroupUserCache struct {
	cache cache.Driver
	//localCache cache.Driver
}

// NewGroupUserCache new一个群成员cache
func NewGroupUserCache() *GroupUserCache {
	encoding := cache.JSONEncoding{}
	return &GroupUserCache{
		cache: cache.NewRedisCache(redis.Client, prefixKey, encoding, nil),
	}
}

func (u *GroupUserCache) GetCacheKey(uid, gid uint32) string {
	return fmt.Sprintf(GroupUserCacheKey, uid, gid)
}

func (u *GroupUserCache) GetCacheAllKey(uid uint32) string {
	return fmt.Sprintf(GroupUserAllCacheKey, uid)
}

func (u *GroupUserCache) SetCache(ctx context.Context, uid, gid uint32, data *model.GroupUserModel) error {
	if data == nil || data.ID == 0 {
		return nil
	}
	return u.cache.Set(u.GetCacheKey(uid, gid), data, cache.DefaultExpireTime)
}

func (u *GroupUserCache) SetCacheAll(ctx context.Context, uid uint32, all []*model.GroupUserModel) error {
	if len(all) == 0 {
		return u.cache.SetCacheWithNotFound(u.GetCacheAllKey(uid))
	}
	return u.cache.Set(u.GetCacheAllKey(uid), all, cache.DefaultExpireTime)
}

func (u *GroupUserCache) GetCache(ctx context.Context, uid, gid uint32) (data *model.GroupUserModel, err error) {
	err = u.cache.Get(u.GetCacheKey(uid, gid), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (u *GroupUserCache) GetCacheAll(ctx context.Context, uid uint32) (all []*model.GroupUserModel, err error) {
	err = u.cache.Get(u.GetCacheAllKey(uid), &all)
	if err != nil {
		return nil, err
	}
	return all, nil
}

func (u *GroupUserCache) DelCache(ctx context.Context, uid, gid uint32) error {
	return u.cache.Del(u.GetCacheKey(uid, gid))
}

func (u *GroupUserCache) DelCacheAll(ctx context.Context, uid uint32) error {
	return u.cache.Del(u.GetCacheAllKey(uid))
}

// SetCacheWithNotFound 设置空
func (u *GroupUserCache) SetCacheWithNotFound(ctx context.Context, uid, gid uint32) error {
	return u.cache.SetCacheWithNotFound(u.GetCacheKey(uid, gid))
}