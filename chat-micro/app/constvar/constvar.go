package constvar

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
)

const (
	//ServiceConnect 连接服务
	ServiceConnect = "chat.connect"
	//ServiceLogic 逻辑业务服务
	ServiceLogic = "chat.logic"
	//ServiceTask 工作服务
	ServiceTask = "chat.task"
)

const (
	//TaskTopic 工作任务服务topic
	TaskTopic = "chat.task.message"
)
