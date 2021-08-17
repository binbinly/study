package constvar

import "time"

// 用来定义一些公共的常量
const (
	// DefaultLimit 默认分页数
	DefaultLimit = 20

	// MaxID 最大id
	MaxID = 0xffffffffffff

	// ModeDebug debug mode
	ModeDebug string = "debug"
	// ModeRelease release mode
	ModeRelease string = "release"
	// ModeTest test mode
	ModeTest string = "test"

	// CacheExpireTime 缓存过期时间
	CacheExpireTime = time.Hour * 24
)
