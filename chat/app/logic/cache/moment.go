package cache

import (
	"context"
	"fmt"

	"chat/app/logic/model"
	"chat/pkg/cache"
)

// MomentCache 朋友圈缓存
type MomentCache struct {
	cache cache.Driver
	//localCache cache.Driver
}

// NewMomentCache new一个朋友圈cache
func NewMomentCache() *MomentCache {
	return &MomentCache{cache: newCache(func() interface{} {
		return &model.MomentModel{}
	})}
}

// GetCacheKey 获取缓存键
func (u *MomentCache) GetCacheKey(id uint32) string {
	return fmt.Sprintf(momentCacheKey, id)
}

// SetCache 写入用户缓存
func (u *MomentCache) SetCache(ctx context.Context, id uint32, user *model.MomentModel) error {
	if user == nil || user.ID == 0 {
		return nil
	}
	return u.cache.Set(ctx, u.GetCacheKey(id), user, defaultExpireTime)
}

// GetCache 获取用户缓存
func (u *MomentCache) GetCache(ctx context.Context, id uint32) (user *model.MomentModel, err error) {
	err = u.cache.Get(ctx, u.GetCacheKey(id), &user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// MultiGetCache 批量获取用户cache
func (u *MomentCache) MultiGetCache(ctx context.Context, ids []uint32) (map[string]*model.MomentModel, error) {
	var keys []string
	for _, v := range ids {
		keys = append(keys, u.GetCacheKey(v))
	}
	// 需要在这里make实例化，如果在返回参数里直接定义会报 nil map
	userMap := make(map[string]*model.MomentModel)
	err := u.cache.MultiGet(ctx, keys, userMap)
	if err != nil {
		return nil, err
	}
	return userMap, nil
}

// DelCache 删除用户缓存
func (u *MomentCache) DelCache(ctx context.Context, id uint32) error {
	return u.cache.Del(ctx, u.GetCacheKey(id))
}

// SetCacheWithNotFound 设置空
func (u *MomentCache) SetCacheWithNotFound(ctx context.Context, id uint32) error {
	return u.cache.SetCacheWithNotFound(ctx, u.GetCacheKey(id))
}
