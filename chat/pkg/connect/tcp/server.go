package tcp

import (
	"fmt"
	"math"
	"net"
	"time"

	"chat/pkg/app"
	"chat/pkg/connect"
	"chat/pkg/log"
)

//MsgIDAuth 鉴权消息id
const MsgIDAuth = 1

const (
	statusInit   = iota + 1 //连接初始化完成
	statusFinish            //连接已完成鉴权
	statusClosed            //连接已关闭
)

// 连接状态
type status int8

//Config tcp服务基础配置
type Config struct {
	Port             int           //当前服务器绑定的端口号
	MaxIPLimit       int           //限制ip最大连接数
	Keepalive        bool          //是否保持连接
	HandshakeTimeout time.Duration //未鉴权连接超时
	SendBuf          int           //发送缓冲区
	ReceiveBuf       int           //接收缓冲区
	MaxPacketSize    int64         //数据包的最大值
	MaxConn          int           //当前服务器主机允许的最大连接数
	WorkerPoolSize   uint32        //业务工作Worker池的数量
	MaxWorkerTaskLen uint32        //业务工作Worker对应负责的任务队列最大任务存储数量
	MaxMsgChanLen    uint32        //SendBuffMsg发送消息的缓冲最大长度
	BucketSize       uint32        //连接管理器个数
}

// Server 接口实现，定义一个Server服务类
type Server struct {
	*connect.Server
	IPVersion string // tcp4 or other

	c *Config
}

//NewServer 创建一个服务器句柄
func NewServer(c *Config, r *connect.Engine, opts ...app.Option) *Server {
	s := &Server{
		Server: &connect.Server{
			Type:       connect.ServerTCP,
			Port:       c.Port,
			BucketSize: c.BucketSize,
			MsgHandler: connect.NewMsgHandle(c.WorkerPoolSize, r),
			Options:    &app.Options{},
			Limit:      connect.NewLimiter(c.MaxIPLimit),
		},
		IPVersion: "tcp4",
		c:         c,
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
	address := fmt.Sprintf(":%d", s.c.Port)
	fmt.Printf("Tcp Server listener at Addr:%v is starting\n", address)

	go func() {
		// 启动worker工作池机制
		s.MsgHandler.StartWorkerPool(s.c.MaxWorkerTaskLen)

		// 获取一个tcp的addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, address)
		if err != nil {
			log.Panicf("resolve tcp addr err: %v", err)
			return
		}
		// 监听服务器地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			log.Panicf("listen %v err:%v", s.IPVersion, err)
			return
		}
		//已经监听成功
		fmt.Println("tcp server success, now listening...")

		s.accept(listener)
	}()
	s.Server.Register()
}

//监听客户端连接
func (s *Server) accept(listener *net.TCPListener) {
	var (
		conn *net.TCPConn
		err  error
		cid  uint32
	)
	cid = 1
	// 启动server网络连接业务
	for {
		//3.1 阻塞等待客户端建立连接请求
		if conn, err = listener.AcceptTCP(); err != nil {
			log.Warnf("[tcp.server] Accept err:%v", err)
			continue
		}
		// 设置客户端服务器间的活动性
		if err = conn.SetKeepAlive(s.c.Keepalive); err != nil {
			log.Warnf("[tcp.server] conn.SetKeepAlive() error:%v", err)
			continue
		}
		// 设置接收消息缓冲区
		if err = conn.SetReadBuffer(s.c.ReceiveBuf); err != nil {
			log.Warnf("[tcp.server] conn.SetReadBuffer() error:%v", err)
			continue
		}
		// 设置发送消息缓冲区
		if err = conn.SetWriteBuffer(s.c.SendBuf); err != nil {
			log.Warnf("[tcp.server] conn.SetWriteBuffer() error:%v", err)
			continue
		}
		log.Debug("[tcp.server] Get conn remote addr = ", conn.RemoteAddr().String())

		//3.2 设置服务器最大连接控制,如果超过最大连接，那么则关闭此新的连接
		if s.ConnTotal() >= s.c.MaxConn {
			log.Warn("[tcp.server] server connect limited")
			err = conn.Close()
			if err != nil {
				log.Warnf("[tcp.server] conn closed err:%v", err)
			}
			continue
		}

		//3.3 处理该新连接请求的 业务 方法， 此时应该有 handler 和 conn是绑定的
		dealConn := NewConnect(s, conn, cid)
		if cid++; cid == math.MaxInt32 {
			log.Infof("[tcp.server] conn.acceptTcp num is:%d", cid)
			cid = 1
		}

		//3.4 启动当前链接的处理业务
		go dealConn.Start()
	}
}
