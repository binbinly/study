package cache

import (
	"context"
	"fmt"

	"chat/app/logic/model"
	"chat/pkg/cache"
)

// GroupCache 群成员
type GroupCache struct {
	cache cache.Driver
	//localCache cache.Driver
}

// NewGroupCache new一个群cache
func NewGroupCache() *GroupCache {
	return &GroupCache{
		cache: newCache(nil),
	}
}

func (u *GroupCache) GetCacheKey(id uint32) string {
	return fmt.Sprintf(GroupCacheKey, id)
}

func (u *GroupCache) GetCacheAllKey(id uint32) string {
	return fmt.Sprintf(GroupAllCacheKey, id)
}

func (u *GroupCache) SetCache(ctx context.Context, id uint32, data *model.GroupModel) error {
	if data == nil || data.ID == 0 {
		return nil
	}
	return u.cache.Set(u.GetCacheKey(id), data, cache.DefaultExpireTime)
}

func (u *GroupCache) SetCacheAll(ctx context.Context, uid uint32, all []*model.GroupList) error {
	if len(all) == 0 {
		return u.cache.SetCacheWithNotFound(u.GetCacheAllKey(uid))
	}
	return u.cache.Set(u.GetCacheAllKey(uid), all, cache.DefaultExpireTime)
}

func (u *GroupCache) GetCache(ctx context.Context, id uint32) (data *model.GroupModel, err error) {
	err = u.cache.Get(u.GetCacheKey(id), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (u *GroupCache) GetCacheAll(ctx context.Context, uid uint32) (all []*model.GroupList, err error) {
	err = u.cache.Get(u.GetCacheAllKey(uid), &all)
	if err != nil {
		return nil, err
	}
	return all, nil
}

func (u *GroupCache) DelCache(ctx context.Context, id uint32) error {
	return u.cache.Del(u.GetCacheKey(id))
}

func (u *GroupCache) DelCacheAll(ctx context.Context, id uint32) error {
	return u.cache.Del(u.GetCacheAllKey(id))
}

// SetCacheWithNotFound 设置空
func (u *GroupCache) SetCacheWithNotFound(ctx context.Context, id uint32) error {
	return u.cache.SetCacheWithNotFound(u.GetCacheKey(id))
}