package cache

import (
	"context"
	"fmt"

	"chat/pkg/cache"
)

// LikeCache 好友申请
type LikeCache struct {
	cache cache.Driver
	//localCache cache.Driver
}

// NewLikeCache new一个朋友圈点赞cache
func NewLikeCache() *LikeCache {
	return &LikeCache{
		cache: newCache(func() interface{} {
			return &[]uint32{}
		}),
	}
}

//GetCacheKey 获取缓存键
func (u *LikeCache) GetCacheKey(id uint32) string {
	return fmt.Sprintf(momentLikeCacheKey, id)
}

//SetCache 设置缓存
func (u *LikeCache) SetCache(ctx context.Context, id uint32, userIds []uint32) error {
	if len(userIds) == 0 {
		return u.cache.SetCacheWithNotFound(ctx, u.GetCacheKey(id))
	}
	return u.cache.Set(ctx, u.GetCacheKey(id), userIds, defaultExpireTime)
}

//GetCache 获取缓存
func (u *LikeCache) GetCache(ctx context.Context, id uint32) (userIds *[]uint32, err error) {
	err = u.cache.Get(ctx, u.GetCacheKey(id), &userIds)
	if err != nil {
		return nil, err
	}
	return userIds, nil
}

// MultiGetCache 批量获取缓存
func (u *LikeCache) MultiGetCache(ctx context.Context, ids []uint32) (map[string]*[]uint32, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := u.GetCacheKey(v)
		keys = append(keys, cacheKey)
	}

	// 需要在这里make实例化，如果在返回参数里直接定义会报 nil map
	likeMap := make(map[string]*[]uint32)
	err := u.cache.MultiGet(ctx, keys, likeMap)
	if err != nil {
		return nil, err
	}
	return likeMap, nil
}

// DelCache 删除列表缓存
func (u *LikeCache) DelCache(ctx context.Context, id uint32) error {
	return u.cache.Del(ctx, u.GetCacheKey(id))
}