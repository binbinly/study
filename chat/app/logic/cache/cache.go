package cache

import (
	"context"
	"time"

	"chat/pkg/cache"
	"chat/pkg/redis"
)

const (
	// defaultExpireTime 默认过期时间
	defaultExpireTime = time.Hour * 24
	// prefixKey 缓存前缀
	prefixKey = "cache:"

	UserCacheKey = "user:%d"
	ApplyListCacheKey  = "apply:list:%d"
	TagAllCacheKey  = "tag:all:%d"
	MomentTimelineCacheKey  = "moment:timeline:count:%d_%d"
	MomentLikeCacheKey  = "moment:like:%d"
	MomentCommentCacheKey = "moment:comment:%d"
	GroupUserCacheKey = "group_user:%d_%d"
	GroupUserAllCacheKey = "group_user:all:%d"
	GroupCacheKey = "group:%d"
	GroupAllCacheKey = "group:all:%d"
	FriendCacheKey = "friend:%d_%d"
	FriendAllCacheKey = "friend:all:%d"
	CollectListCacheKey  = "collect:list:%d"
)

// newCache new一个cache
func newCache(obj func() interface{}) cache.Driver {
	encoding := cache.JSONEncoding{}
	return cache.NewRedisCache(redis.Client, prefixKey, encoding, obj)
}

// newCtxCache new一个上下文cache
func newCtxCache(ctx context.Context, obj func() interface{}) cache.Driver {
	encoding := cache.JSONEncoding{}
	return cache.NewRedisCache(redis.WrapRedisClient(ctx, redis.Client), prefixKey, encoding, obj)
}
