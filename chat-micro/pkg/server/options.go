package server

import (
	"context"
	"crypto/tls"
	"sync"
	"time"

	"github.com/google/uuid"

	"chat-micro/pkg/registry"
)

//AuthHandler 鉴权处理
type AuthHandler = func(Connection, IRequest) (bool, uint32)

//Option 服务器选项
type Option func(*Options)

//Options 服务器选项结构
type Options struct {
	Registry registry.Registry
	Metadata map[string]string
	Name     string
	Address  string
	Id       string
	Version  string

	MaxPacketSize    int           //都需数据包的最大值
	MaxConn          int           //当前服务器主机允许的最大链接个数
	WorkerPoolSize   int           //业务工作Worker池的数量
	MaxWorkerTaskLen int           //业务工作Worker对应负责的任务队列最大任务存储数量
	MaxMsgChanLen    int           //SendBuffMsg发送消息的缓冲最大长度
	BucketSize       int           //连接管理器个数
	ReadBufferSize   int           //接收缓冲区
	WriteBufferSize  int           //发送缓冲区
	WriteWait        time.Duration //写入客户端超时
	Timeout          time.Duration //连接超时

	// RegisterCheck runs a check function before registering the service
	RegisterCheck func(context.Context) error
	// The register expiry time
	RegisterTTL time.Duration
	// The interval on which to register
	RegisterInterval time.Duration

	// The router for requests
	Router *Engine

	// TLSConfig specifies tls.Config for secure serving
	TLSConfig *tls.Config

	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context

	//该Server的连接创建开始时Hook函数
	OnConnStarting func(conn Connection)
	//该Server的连接创建完成时Hook函数
	OnConnStarted func(conn Connection)
	//该Server的连接断开时的Hook函数
	OnConnStop func(conn Connection)
	//该Server的连接鉴权完成的Hook函数
	OnConnAuth AuthHandler
}

func NewOptions(opt ...Option) Options {
	opts := Options{
		Id:               uuid.New().String(),
		Name:             "server",
		Address:          ":0",
		Version:          "latest",
		Metadata:         map[string]string{},
		RegisterInterval: time.Second * 30,
		RegisterTTL:      time.Second * 90,
		MaxPacketSize:    4096,
		MaxConn:          36000,
		WorkerPoolSize:   4,
		MaxWorkerTaskLen: 128,
		MaxMsgChanLen:    128,
		ReadBufferSize:   4096,
		WriteBufferSize:  4096,
		BucketSize:       1,
		Timeout:          time.Second * 5,
		WriteWait:        time.Second * 10,
		OnConnStarting:   func(conn Connection) {},
		OnConnStarted:    func(conn Connection) {},
		OnConnStop:       func(conn Connection) {},
		OnConnAuth:       func(conn Connection, r IRequest) (bool, uint32) { return true, 0 },
	}

	for _, o := range opt {
		o(&opts)
	}

	return opts
}

// Server name
func Name(n string) Option {
	return func(o *Options) {
		o.Name = n
	}
}

// Unique server id
func Id(id string) Option {
	return func(o *Options) {
		o.Id = id
	}
}

// Version of the service
func Version(v string) Option {
	return func(o *Options) {
		o.Version = v
	}
}

// Address to bind to - host:port
func Address(a string) Option {
	return func(o *Options) {
		o.Address = a
	}
}

// Registry used for discovery
func Registry(r registry.Registry) Option {
	return func(o *Options) {
		o.Registry = r
	}
}

// Metadata associated with the server
func Metadata(md map[string]string) Option {
	return func(o *Options) {
		o.Metadata = md
	}
}

// RegisterCheck run func before registry service
func RegisterCheck(fn func(context.Context) error) Option {
	return func(o *Options) {
		o.RegisterCheck = fn
	}
}

// Register the service with a TTL
func RegisterTTL(t time.Duration) Option {
	return func(o *Options) {
		o.RegisterTTL = t
	}
}

// Register the service with at interval
func RegisterInterval(t time.Duration) Option {
	return func(o *Options) {
		o.RegisterInterval = t
	}
}

// WithRouter sets the request router
func WithRouter(r *Engine) Option {
	return func(o *Options) {
		o.Router = r
	}
}

//WithMaxPacketSize set
func WithMaxPacketSize(size int) Option {
	return func(o *Options) {
		o.MaxPacketSize = size
	}
}

//WithMaxConn set
func WithMaxConn(size int) Option {
	return func(o *Options) {
		o.MaxConn = size
	}
}

//WithWorkerPoolSize set
func WithWorkerPoolSize(size int) Option {
	return func(o *Options) {
		o.WorkerPoolSize = size
	}
}

//WithMaxWorkerTaskLen set
func WithMaxWorkerTaskLen(size int) Option {
	return func(o *Options) {
		o.MaxWorkerTaskLen = size
	}
}

//WithMaxMsgChanLen set
func WithMaxMsgChanLen(size int) Option {
	return func(o *Options) {
		o.MaxMsgChanLen = size
	}
}

//WithBucketSize set
func WithBucketSize(size int) Option {
	return func(o *Options) {
		o.BucketSize = size
	}
}

//WithReadBufferSize set
func WithReadBufferSize(size int) Option {
	return func(o *Options) {
		o.ReadBufferSize = size
	}
}

//WithWriteBufferSize set
func WithWriteBufferSize(size int) Option {
	return func(o *Options) {
		o.WriteBufferSize = size
	}
}

//WithWriteWait set
func WithWriteWait(d time.Duration) Option {
	return func(o *Options) {
		o.WriteWait = d
	}
}

//WithTimeout set
func WithTimeout(d time.Duration) Option {
	return func(o *Options) {
		o.Timeout = d
	}
}

//WithOnConnStarting set
func WithOnConnStarting(f func(conn Connection)) Option {
	return func(o *Options) {
		o.OnConnStarting = f
	}
}

//WithOnConnStarted set
func WithOnConnStarted(f func(conn Connection)) Option {
	return func(o *Options) {
		o.OnConnStarted = f
	}
}

//WithOnConnStop set
func WithOnConnStop(f func(conn Connection)) Option {
	return func(o *Options) {
		o.OnConnStop = f
	}
}

//WithOnConnAuth set
func WithOnConnAuth(f AuthHandler) Option {
	return func(o *Options) {
		o.OnConnAuth = f
	}
}

// Wait tells the server to wait for requests to finish before exiting
// If `wg` is nil, server only wait for completion of rpc handler.
// For user need finer grained control, pass a concrete `wg` here, server will
// wait against it on stop.
func Wait(wg *sync.WaitGroup) Option {
	return func(o *Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		if wg == nil {
			wg = new(sync.WaitGroup)
		}
		o.Context = context.WithValue(o.Context, "wait", wg)
	}
}
