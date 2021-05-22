package cache

import (
	"context"
	"fmt"

	"chat/app/logic/model"
	"chat/pkg/cache"
)

// userCache 用户缓存
type UserCache struct {
	cache cache.Driver
	//localCache cache.Driver
}

// NewUserCache new一个用户cache
func NewUserCache() *UserCache {
	return &UserCache{cache: newCache(func() interface{} {
		return &model.UserModel{}
	})}
}

// GetUserCacheKey 获取缓存键
func (u *UserCache) GetUserCacheKey(userId uint32) string {
	return fmt.Sprintf(UserCacheKey, userId)
}

// SetUserCache 写入用户缓存
func (u *UserCache) SetUserCache(ctx context.Context, userId uint32, user *model.UserModel) error {
	if user == nil || user.ID == 0 {
		return nil
	}
	return u.cache.Set(u.GetUserCacheKey(userId), user, defaultExpireTime)
}

// GetUserCache 获取用户缓存
func (u *UserCache) GetUserCache(ctx context.Context, userId uint32) (user *model.UserModel, err error) {
	err = u.cache.Get(u.GetUserCacheKey(userId), &user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// MultiGetUserCache 批量获取用户cache
func (u *UserCache) MultiGetUserCache(ctx context.Context, userIds []uint32) (map[string]*model.UserModel, error) {
	var keys []string
	for _, v := range userIds {
		keys = append(keys, u.GetUserCacheKey(v))
	}
	// 需要在这里make实例化，如果在返回参数里直接定义会报 nil map
	userMap := make(map[string]*model.UserModel)
	err := u.cache.MultiGet(keys, userMap)
	if err != nil {
		return nil, err
	}
	return userMap, nil
}

// DelUserCache 删除用户缓存
func (u *UserCache) DelUserCache(ctx context.Context, userId uint32) error {
	return u.cache.Del(u.GetUserCacheKey(userId))
}

// SetCacheWithNotFound 设置空
func (u *UserCache) SetCacheWithNotFound(ctx context.Context, userId uint32) error {
	return u.cache.SetCacheWithNotFound(u.GetUserCacheKey(userId))
}
