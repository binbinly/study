package constvar

// 用来定义一些公共的常量
const (
	// DefaultLimit 默认分页数
	DefaultLimit = 20

	// MaxID 最大id
	MaxID = 0xffffffffffff

	//HistoryPrefix 历史消息键
	HistoryPrefix = "history:message:%d"

	// ModeDebug debug mode
	ModeDebug string = "debug"
	// ModeRelease release mode
	ModeRelease string = "release"
	// ModeTest test mode
	ModeTest string = "test"
)
