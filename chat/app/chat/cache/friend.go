package cache

import (
	"context"

	"chat/app/chat/model"
	"chat/app/constvar"
	"chat/internal"
)

const (
	_friendCacheKey = "friend:%d_%d"
)

// FriendCache 好友
type FriendCache struct {
	*internal.Cache
}

// NewFriendCache new一个收藏cache
func NewFriendCache() *FriendCache {
	return &FriendCache{internal.NewCache(_friendCacheKey, nil)}
}

//SetCache 设置缓存
func (u *FriendCache) SetCache(ctx context.Context, uid, fid uint32, friend *model.FriendModel) error {
	if friend == nil || friend.ID == 0 {
		return nil
	}
	return u.Redis.Set(ctx, u.BuildCacheKey([]interface{}{uid, fid}), friend, constvar.CacheExpireTime)
}

//GetCache 获取缓存
func (u *FriendCache) GetCache(ctx context.Context, uid, fid uint32) (friend *model.FriendModel, err error) {
	err = u.Redis.Get(ctx, u.BuildCacheKey([]interface{}{uid, fid}), &friend)
	if err != nil {
		return nil, err
	}
	return friend, nil
}
