package cache

import (
	"context"

	"mall/app/model"
)

const _userCacheKey = "user:%d"

// UserCache 用户缓存
type UserCache struct {
	*Cache
}

// NewUserCache new一个用户cache
func NewUserCache() *UserCache {
	return &UserCache{NewCache(_userCacheKey, func() interface{} {
		return &model.UserModel{}
	})}
}

// SetCache 写入用户缓存
func (u *UserCache) SetCache(ctx context.Context, id int, user *model.UserModel) error {
	if user == nil || user.ID == 0 {
		return nil
	}
	return u.Redis.Set(ctx, u.BuildCacheKey([]interface{}{id}), user, _cacheExpireTime)
}

// GetCache 获取用户缓存
func (u *UserCache) GetCache(ctx context.Context, id int) (user *model.UserModel, err error) {
	err = u.Redis.Get(ctx, u.BuildCacheKey([]interface{}{id}), &user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
