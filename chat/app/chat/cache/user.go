package cache

import (
	"context"

	"chat/app/constvar"
	"chat/internal"
	"chat/proto/base"
)

const _userCacheKey = "user:%d"

// UserCache 用户缓存
type UserCache struct {
	*internal.Cache
}

// NewUserCache new一个用户cache
func NewUserCache() *UserCache {
	return &UserCache{internal.NewCache(_userCacheKey, func() interface{} {
		return &base.UserInfo{}
	})}
}

// SetCache 写入用户缓存
func (u *UserCache) SetCache(ctx context.Context, id uint32, user *base.UserInfo) error {
	if user == nil || user.Id == 0 {
		return nil
	}
	return u.Redis.Set(ctx, u.BuildCacheKey([]interface{}{id}), user, constvar.CacheExpireTime)
}

// GetCache 获取用户缓存
func (u *UserCache) GetCache(ctx context.Context, id uint32) (user *base.UserInfo, err error) {
	err = u.Redis.Get(ctx, u.BuildCacheKey([]interface{}{id}), &user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// MultiGetCache 批量获取用户cache
func (u *UserCache) MultiGetCache(ctx context.Context, userIds []uint32) (map[string]*base.UserInfo, error) {
	var keys []string
	for _, v := range userIds {
		keys = append(keys, u.BuildCacheKey([]interface{}{v}))
	}
	// 需要在这里make实例化，如果在返回参数里直接定义会报 nil map
	userMap := make(map[string]*base.UserInfo)
	err := u.Redis.MultiGet(ctx, keys, userMap)
	if err != nil {
		return nil, err
	}
	return userMap, nil
}
