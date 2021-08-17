package cache

import (
	"context"

	"chat/app/chat/model"
	"chat/app/constvar"
	"chat/internal"
)

const (
	_groupUserCacheKey = "group_user:%d_%d"
)

// GroupUserCache 群成员
type GroupUserCache struct {
	*internal.Cache
}

// NewGroupUserCache new一个群成员cache
func NewGroupUserCache() *GroupUserCache {
	return &GroupUserCache{internal.NewCache(_groupUserCacheKey, nil)}
}

//SetCache 设置缓存
func (u *GroupUserCache) SetCache(ctx context.Context, uid, gid uint32, data *model.GroupUserModel) error {
	if data == nil || data.ID == 0 {
		return nil
	}
	return u.Redis.Set(ctx, u.BuildCacheKey([]interface{}{uid, gid}), data, constvar.CacheExpireTime)
}

//GetCache 获取缓存
func (u *GroupUserCache) GetCache(ctx context.Context, uid, gid uint32) (data *model.GroupUserModel, err error) {
	err = u.Redis.Get(ctx, u.BuildCacheKey([]interface{}{uid, gid}), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
