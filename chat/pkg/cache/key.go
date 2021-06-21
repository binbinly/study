package cache

import (
	"strings"
)

// BuildCacheKey 构建一个带有前缀的缓存key
func BuildCacheKey(prefix string, key string) (cacheKey string) {
	cacheKey = key
	if prefix != "" {
		cacheKey = strings.Join([]string{prefix, key}, ":")
	}
	return
}