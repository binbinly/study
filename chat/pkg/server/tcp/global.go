package tcp

import (
	"time"
)

const maxInt = 1<<31 - 1 // 连接id最大值

const maxPacketSize = 4096 //数据包的最大值

const MsgIdAuth = 0 //鉴权消息id

const (
	statusInit   = 1 //连接初始化完成
	statusFinish = 2 //连接已完成鉴权
	statusClosed = 3 //连接已关闭
)

// 连接状态
type status int8

// tcp服务基础配置
type Config struct {
	ServerId         string        //服务器id
	Port             int           //当前服务器绑定的端口号
	Keepalive        bool          //是否保持连接
	HandshakeTimeout time.Duration //未鉴权连接超时
	SendBuf          int           //发送缓冲区
	ReceiveBuf       int           //接收缓冲区
	MaxPacketSize    uint32        //数据包的最大值
	MaxConn          int           //当前服务器主机允许的最大连接数
	WorkerPoolSize   uint32        //业务工作Worker池的数量
	MaxWorkerTaskLen uint32        //业务工作Worker对应负责的任务队列最大任务存储数量
	MaxMsgChanLen    uint32        //SendBuffMsg发送消息的缓冲最大长度
	BucketSize       uint32        //连接管理器个数
}
