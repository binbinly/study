package cache

import (
	"context"
	"fmt"

	"chat/app/logic/model"
	"chat/pkg/cache"
)


// FriendCache 用户收藏
type FriendCache struct {
	cache cache.Driver
	//localCache cache.Driver
}

// NewFriendCache new一个收藏cache
func NewFriendCache() *FriendCache {
	return &FriendCache{
		cache: newCache(nil),
	}
}

func (u *FriendCache) GetCacheKey(uid, fid uint32) string {
	return fmt.Sprintf(FriendCacheKey, uid, fid)
}

func (u *FriendCache) GetCacheAllKey(uid uint32) string {
	return fmt.Sprintf(FriendAllCacheKey, uid)
}

func (u *FriendCache) SetCache(ctx context.Context, uid, fid uint32, friend *model.FriendModel) error {
	if friend == nil || friend.ID == 0 {
		return nil
	}
	return u.cache.Set(u.GetCacheKey(uid, fid), friend, defaultExpireTime)
}

func (u *FriendCache) SetCacheAll(ctx context.Context, uid uint32, all []*model.FriendModel) error {
	if len(all) == 0 { // 空列表设置占位符
		return u.cache.SetCacheWithNotFound(u.GetCacheAllKey(uid))
	}
	return u.cache.Set(u.GetCacheAllKey(uid), all, defaultExpireTime)
}

func (u *FriendCache) GetCache(ctx context.Context, uid, fid uint32) (friend *model.FriendModel, err error) {
	err = u.cache.Get(u.GetCacheKey(uid, fid), &friend)
	if err != nil {
		return nil, err
	}
	return friend, nil
}

func (u *FriendCache) GetCacheAll(ctx context.Context, uid uint32) (all []*model.FriendModel, err error) {
	err = u.cache.Get(u.GetCacheAllKey(uid), &all)
	if err != nil {
		return nil, err
	}
	return all, nil
}

func (u *FriendCache) DelCache(ctx context.Context, uid, fid uint32) error {
	return u.cache.Del(u.GetCacheKey(uid, fid))
}

func (u *FriendCache) DelCacheAll(ctx context.Context, uid uint32) error {
	return u.cache.Del(u.GetCacheAllKey(uid))
}

// SetCacheWithNotFound 设置空
func (u *FriendCache) SetCacheWithNotFound(ctx context.Context, uid, fid uint32) error {
	return u.cache.SetCacheWithNotFound(u.GetCacheKey(uid, fid))
}