package cache

import (
	"context"

	"chat/app/chat/model"
	"chat/app/constvar"
	"chat/internal"
)

const (
	_emoticonListCacheKey   = "emoticon:%s"
	_emoticonCatAllCacheKey = "emoticon:cat:all"
)

// EmoticonCache 表情包
type EmoticonCache struct {
	*internal.Cache
}

// NewEmoticonCache new一个收藏cache
func NewEmoticonCache() *EmoticonCache {
	return &EmoticonCache{internal.NewCache(_emoticonListCacheKey, nil)}
}

//SetCache 设置缓存
func (u *EmoticonCache) SetCache(ctx context.Context, cat string, list []*model.Emoticon) error {
	if len(list) == 0 { // 空列表设置占位符
		return u.Redis.SetCacheWithNotFound(ctx, u.BuildCacheKey([]interface{}{cat}))
	}
	return u.Redis.Set(ctx, u.BuildCacheKey([]interface{}{cat}), list, constvar.CacheExpireTime)
}

//GetCache 获取缓存
func (u *EmoticonCache) GetCache(ctx context.Context, cat string) (list []*model.Emoticon, err error) {
	err = u.Redis.Get(ctx, u.BuildCacheKey([]interface{}{cat}), &list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

//SetCatCache 设置表情分类缓存
func (u *EmoticonCache) SetCatCache(ctx context.Context, list []*model.Emoticon) error {
	return u.Redis.Set(ctx, _emoticonCatAllCacheKey, list, constvar.CacheExpireTime)
}

//GetCatCache 获取表情分类
func (u *EmoticonCache) GetCatCache(ctx context.Context) (list []*model.Emoticon, err error) {
	err = u.Redis.Get(ctx, _emoticonCatAllCacheKey, &list)
	if err != nil {
		return nil, err
	}
	return list, nil
}
