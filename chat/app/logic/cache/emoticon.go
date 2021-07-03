package cache

import (
	"context"
	"fmt"

	"chat/app/logic/model"
	"chat/pkg/cache"
)

// EmoticonCache 表情包
type EmoticonCache struct {
	cache cache.Driver
	//localCache cache.Driver
}

// NewEmoticonCache new一个收藏cache
func NewEmoticonCache() *EmoticonCache {
	return &EmoticonCache{
		cache: newCache(nil),
	}
}

//GetCacheKey 获取缓存键
func (u *EmoticonCache) GetCacheKey(cat string) string {
	return fmt.Sprintf(emoticonListCacheKey, cat)
}

//SetCache 设置缓存
func (u *EmoticonCache) SetCache(ctx context.Context, cat string, list []*model.Emoticon) error {
	if len(list) == 0 { // 空列表设置占位符
		return u.cache.SetCacheWithNotFound(ctx, u.GetCacheKey(cat))
	}
	return u.cache.Set(ctx, u.GetCacheKey(cat), list, defaultExpireTime)
}

//GetCache 获取缓存
func (u *EmoticonCache) GetCache(ctx context.Context, cat string) (list []*model.Emoticon, err error) {
	err = u.cache.Get(ctx, u.GetCacheKey(cat), &list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

//SetCatCache 设置缓存
func (u *EmoticonCache) SetCatCache(ctx context.Context, list []*model.Emoticon) error {
	return u.cache.Set(ctx, emoticonCatAllCacheKey, list, defaultExpireTime)
}

//GetCatCache 获取缓存
func (u *EmoticonCache) GetCatCache(ctx context.Context) (list []*model.Emoticon, err error) {
	err = u.cache.Get(ctx, emoticonCatAllCacheKey, &list)
	if err != nil {
		return nil, err
	}
	return list, nil
}
