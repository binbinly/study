package ws

import (
	"time"
)

const (
	maxMsgChanLen = 1024 //SendBuffMsg发送消息的缓冲最大长度
)

type Config struct {
	ServerId         string        //服务器ID
	Port             int           //当前服务器主机监听端口号
	WriteWait        time.Duration //写入客户端超时
	PongWait         time.Duration //读取下一个Pong消息超时
	PingPeriod       time.Duration //ping时间间隔
	MaxPacketSize    int64         //都需数据包的最大值
	ReadBufferSize   int           //接受缓冲区
	WriteBufferSize  int           //发送缓冲区
	MaxConn          int           //当前服务器主机允许的最大连接数
	WorkerPoolSize   uint32        //业务工作Worker池的数量
	MaxWorkerTaskLen uint32        //业务工作Worker对应负责的任务队列最大任务存储数量
	MaxMsgChanLen    uint32        //SendBuffMsg发送消息的缓冲最大长度
	BucketSize       uint32        //连接管理器个数
}
