package cache

import (
	"context"

	"chat/app/chat/model"
	"chat/app/constvar"
	"chat/internal"
)

const (
	_friendAllCacheKey = "friend:all:%d"
)

// FriendAllCache 全部好友
type FriendAllCache struct {
	*internal.Cache
}

// NewFriendAllCache 实例化
func NewFriendAllCache() *FriendAllCache {
	return &FriendAllCache{internal.NewCache(_friendAllCacheKey, nil)}
}

//SetCache 设置缓存
func (u *FriendAllCache) SetCache(ctx context.Context, uid uint32, all []*model.FriendModel) error {
	if len(all) == 0 { // 空列表设置占位符
		return u.Redis.SetCacheWithNotFound(ctx, u.BuildCacheKey([]interface{}{uid}))
	}
	return u.Redis.Set(ctx, u.BuildCacheKey([]interface{}{uid}), all, constvar.CacheExpireTime)
}

//GetCache 获取缓存
func (u *FriendAllCache) GetCache(ctx context.Context, uid uint32) (all []*model.FriendModel, err error) {
	err = u.Redis.Get(ctx, u.BuildCacheKey([]interface{}{uid}), &all)
	if err != nil {
		return nil, err
	}
	return all, nil
}
