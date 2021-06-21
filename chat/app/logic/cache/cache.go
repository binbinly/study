package cache

import (
	"time"

	"chat/pkg/cache"
	"chat/pkg/redis"
)

const (
	// defaultExpireTime 默认过期时间
	defaultExpireTime = time.Hour * 24
	// prefixKey 缓存前缀
	prefixKey = "cache"

	userCacheKey           = "user:%d"
	applyListCacheKey      = "apply:list:%d"
	tagAllCacheKey         = "tag:all:%d"
	momentCacheKey         = "moment:%d"
	momentTimelineCacheKey = "moment:timeline:count:%d_%d"
	momentLikeCacheKey     = "moment:like:%d"
	momentCommentCacheKey  = "moment:comment:%d"
	groupUserCacheKey      = "group_user:%d_%d"
	groupUserAllCacheKey   = "group_user:all:%d"
	groupCacheKey          = "group:%d"
	groupAllCacheKey       = "group:all:%d"
	friendCacheKey         = "friend:%d_%d"
	friendAllCacheKey      = "friend:all:%d"
	collectListCacheKey    = "collect:list:%d"
)

// newCache new一个cache
func newCache(obj func() interface{}) cache.Driver {
	encoding := cache.JSONEncoding{}
	return cache.NewRedisCache(redis.Client, prefixKey, encoding, obj)
}
