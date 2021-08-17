package cache

import (
	"context"

	"chat/app/chat/model"
	"chat/app/constvar"
	"chat/internal"
)

const _momentCacheKey = "moment:%d"

// MomentCache 朋友圈缓存
type MomentCache struct {
	*internal.Cache
}

// NewMomentCache new一个朋友圈cache
func NewMomentCache() *MomentCache {
	return &MomentCache{internal.NewCache(_momentCacheKey, func() interface{} {
		return &model.MomentModel{}
	})}
}

// SetCache 写入用户缓存
func (u *MomentCache) SetCache(ctx context.Context, id uint32, user *model.MomentModel) error {
	if user == nil || user.ID == 0 {
		return nil
	}
	return u.Redis.Set(ctx, u.BuildCacheKey([]interface{}{id}), user, constvar.CacheExpireTime)
}

// GetCache 获取用户缓存
func (u *MomentCache) GetCache(ctx context.Context, id uint32) (user *model.MomentModel, err error) {
	err = u.Redis.Get(ctx, u.BuildCacheKey([]interface{}{id}), &user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// MultiGetCache 批量获取用户cache
func (u *MomentCache) MultiGetCache(ctx context.Context, ids []uint32) (map[string]*model.MomentModel, error) {
	var keys []string
	for _, v := range ids {
		keys = append(keys, u.BuildCacheKey([]interface{}{v}))
	}
	// 需要在这里make实例化，如果在返回参数里直接定义会报 nil map
	userMap := make(map[string]*model.MomentModel)
	err := u.Redis.MultiGet(ctx, keys, userMap)
	if err != nil {
		return nil, err
	}
	return userMap, nil
}
