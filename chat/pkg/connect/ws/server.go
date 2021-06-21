package ws

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"math"
	"net/http"
	"time"

	"chat/pkg/app"
	"chat/pkg/connect"
	"chat/pkg/log"
)

//Config websocket配置
type Config struct {
	Port             int           //当前服务器主机监听端口号
	MaxIPLimit       int           //限制ip最大连接数
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

//Server websocket服务器
type Server struct {
	*connect.Server

	c *Config
}

//NewServer 创建一个服务器句柄
func NewServer(c *Config, r *connect.Engine, opts ...app.Option) *Server {
	s := &Server{
		Server: &connect.Server{
			Type:       connect.ServerWs,
			Port:       c.Port,
			BucketSize: c.BucketSize,
			MsgHandler: connect.NewMsgHandle(c.WorkerPoolSize, r),
			Options:    &app.Options{},
			Limit:      connect.NewLimiter(c.MaxIPLimit),
		},
		c: c,
	}
	for _, o := range opts {
		o(s.Options)
	}
	s.ConnMgr = make([]connect.IConnManager, c.BucketSize)
	for i := uint32(0); i < c.BucketSize; i++ {
		s.ConnMgr[i] = connect.NewConnManager()
	}
	if c.MaxPacketSize > 0 {
		connect.MaxPacketSize = c.MaxPacketSize
	}
	return s
}

//Serve 运行服务
func (s *Server) Serve() {
	s.Start()
}

// Start 开启网络服务
func (s *Server) Start() {
	var cid uint32 = 1
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		log.Debug("[ws.accept] start", r.RemoteAddr)
		//设置服务器最大连接控制,如果超过最大连接，那么则关闭此新的连接
		if s.ConnTotal() >= s.c.MaxConn {
			log.Warn("[ws.server] connection size limit")
			return
		}

		conn := NewConnect(s, w, r, cid)
		if cid++; cid == math.MaxInt32 {
			log.Infof("[tcp.server] conn.acceptTcp num is:%d", cid)
			cid = 1
		}
		conn.Start()
	})
	// 启动worker工作池机制
	addr := fmt.Sprintf(":%d", s.Port)
	s.MsgHandler.StartWorkerPool(s.c.MaxWorkerTaskLen)
	go func() {
		if err := http.ListenAndServe(addr, nil); err != nil {
			log.Fatalf("ListenAndServe, err: %s", err.Error())
		}
	}()
	fmt.Printf("Websocket Server listener at Addr:%v is starting\n", addr)
	s.Server.Register()
}
