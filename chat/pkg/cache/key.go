package cache

import (
	"strings"
)

// BuildCacheKey 构建一个带有前缀的缓存key
func BuildCacheKey(prefix string, key string) string {
	var str strings.Builder

	str.WriteString(prefix)
	str.WriteString(key)
	return str.String()
}