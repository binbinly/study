package cache

import (
	"context"

	"chat/app/constvar"
	"chat/internal"
)

const _momentLikeCacheKey = "moment:like:%d"

// LikeCache 好友申请
type LikeCache struct {
	*internal.Cache
}

// NewLikeCache new一个朋友圈点赞cache
func NewLikeCache() *LikeCache {
	return &LikeCache{internal.NewCache(_momentLikeCacheKey, func() interface{} {
		return &[]uint32{}
	})}
}

//SetCache 设置缓存
func (u *LikeCache) SetCache(ctx context.Context, id uint32, userIds []uint32) error {
	if len(userIds) == 0 {
		return u.Redis.SetCacheWithNotFound(ctx, u.BuildCacheKey([]interface{}{id}))
	}
	return u.Redis.Set(ctx, u.BuildCacheKey([]interface{}{id}), userIds, constvar.CacheExpireTime)
}

//GetCache 获取缓存
func (u *LikeCache) GetCache(ctx context.Context, id uint32) (userIds *[]uint32, err error) {
	err = u.Redis.Get(ctx, u.BuildCacheKey([]interface{}{id}), &userIds)
	if err != nil {
		return nil, err
	}
	return userIds, nil
}

// MultiGetCache 批量获取缓存
func (u *LikeCache) MultiGetCache(ctx context.Context, ids []uint32) (map[string]*[]uint32, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := u.BuildCacheKey([]interface{}{v})
		keys = append(keys, cacheKey)
	}

	// 需要在这里make实例化，如果在返回参数里直接定义会报 nil map
	likeMap := make(map[string]*[]uint32)
	err := u.Redis.MultiGet(ctx, keys, likeMap)
	if err != nil {
		return nil, err
	}
	return likeMap, nil
}
