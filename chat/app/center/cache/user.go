package cache

import (
	"context"

	"chat/app/center/model"
	"chat/app/constvar"
	"chat/internal"
)

const _userCacheKey = "user:%d"

// UserCache 用户缓存
type UserCache struct {
	*internal.Cache
}

// NewUserCache new一个用户cache
func NewUserCache() *UserCache {
	return &UserCache{internal.NewCache(_userCacheKey, func() interface{} {
		return &model.UserModel{}
	})}
}

// SetCache 写入用户缓存
func (u *UserCache) SetCache(ctx context.Context, id uint32, user *model.UserModel) error {
	if user == nil || user.ID == 0 {
		return nil
	}
	return u.Redis.Set(ctx, u.BuildCacheKey([]interface{}{id}), user, constvar.CacheExpireTime)
}

// GetCache 获取用户缓存
func (u *UserCache) GetCache(ctx context.Context, id uint32) (user *model.UserModel, err error) {
	err = u.Redis.Get(ctx, u.BuildCacheKey([]interface{}{id}), &user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
