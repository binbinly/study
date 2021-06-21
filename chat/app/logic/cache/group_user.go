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

//GetCacheKey 获取缓存键
func (u *GroupUserCache) GetCacheKey(uid, gid uint32) string {
	return fmt.Sprintf(groupUserCacheKey, uid, gid)
}

//GetCacheAllKey 获取缓存键
func (u *GroupUserCache) GetCacheAllKey(uid uint32) string {
	return fmt.Sprintf(groupUserAllCacheKey, uid)
}

//SetCache 设置缓存
func (u *GroupUserCache) SetCache(ctx context.Context, uid, gid uint32, data *model.GroupUserModel) error {
	if data == nil || data.ID == 0 {
		return nil
	}
	return u.cache.Set(ctx, u.GetCacheKey(uid, gid), data, cache.DefaultExpireTime)
}

//SetCacheAll 设置缓存
func (u *GroupUserCache) SetCacheAll(ctx context.Context, uid uint32, all []*model.GroupUserModel) error {
	if len(all) == 0 {
		return u.cache.SetCacheWithNotFound(ctx, u.GetCacheAllKey(uid))
	}
	return u.cache.Set(ctx, u.GetCacheAllKey(uid), all, cache.DefaultExpireTime)
}

//GetCache 获取缓存
func (u *GroupUserCache) GetCache(ctx context.Context, uid, gid uint32) (data *model.GroupUserModel, err error) {
	err = u.cache.Get(ctx, u.GetCacheKey(uid, gid), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//GetCacheAll 获取缓存
func (u *GroupUserCache) GetCacheAll(ctx context.Context, uid uint32) (all []*model.GroupUserModel, err error) {
	err = u.cache.Get(ctx, u.GetCacheAllKey(uid), &all)
	if err != nil {
		return nil, err
	}
	return all, nil
}

//DelCache 删除缓存
func (u *GroupUserCache) DelCache(ctx context.Context, uid, gid uint32) error {
	return u.cache.Del(ctx, u.GetCacheKey(uid, gid))
}

//DelCacheAll 删除缓存
func (u *GroupUserCache) DelCacheAll(ctx context.Context, uid uint32) error {
	return u.cache.Del(ctx, u.GetCacheAllKey(uid))
}

// SetCacheWithNotFound 设置空
func (u *GroupUserCache) SetCacheWithNotFound(ctx context.Context, uid, gid uint32) error {
	return u.cache.SetCacheWithNotFound(ctx, u.GetCacheKey(uid, gid))
}
