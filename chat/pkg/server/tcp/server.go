package tcp

import (
	"fmt"
	"net"
	"strconv"

	"chat/pkg/crypt"
	"chat/pkg/log"
	"chat/pkg/server"
)

// Server 接口实现，定义一个Server服务类
type Server struct {
	*server.Server
	//tcp4 or other
	IPVersion string
	//当前Server的消息管理模块，用来绑定MsgID和对应的处理方法
	msgHandler IMsgHandle
	//连接管理器
	ConnMgr []*ConnManager
	//配置
	Config *Config
}

//NewServer 创建一个服务器句柄
func NewServer(c *Config, r *Engine) *Server {
	s := &Server{
		Server: &server.Server{
			ID:         c.ServerId,
			Addr:       fmt.Sprintf(":%d", c.Port),
			BucketSize: c.BucketSize,
		},
		IPVersion:  "tcp4",
		msgHandler: NewMsgHandle(c.WorkerPoolSize, r),
		Config:     c,
	}
	s.ConnMgr = make([]*ConnManager, c.BucketSize)
	for i := uint32(0); i < c.BucketSize; i++ {
		s.ConnMgr[i] = NewConnManager()
	}
	return s
}

//Serve 运行服务
func (s *Server) Serve() {
	s.Start()
}

// Start 开启网络服务
func (s *Server) Start() {
	fmt.Printf("[tcp.server] Tcp Server listener at Addr:%v is starting\n", s.Addr)

	go func() {
		// 启动worker工作池机制
		s.msgHandler.StartWorkerPool(s.Config.MaxWorkerTaskLen)

		// 获取一个tcp的addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, s.Addr)
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
		fmt.Println("start tcp server  success, now listening...")

		s.accept(listener)
	}()
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
			log.Errorf("[tcp.server] Accept err:%v", err)
			continue
		}
		// 设置客户端服务器间的活动性
		if err = conn.SetKeepAlive(s.Config.Keepalive); err != nil {
			log.Errorf("[tcp.server] conn.SetKeepAlive() error:%v", err)
			continue
		}
		// 设置接收消息缓冲区
		if err = conn.SetReadBuffer(s.Config.ReceiveBuf); err != nil {
			log.Errorf("[tcp.server] conn.SetReadBuffer() error:%v", err)
			continue
		}
		// 设置发送消息缓冲区
		if err = conn.SetWriteBuffer(s.Config.SendBuf); err != nil {
			log.Errorf("[tcp.server] conn.SetWriteBuffer() error:%v", err)
			continue
		}
		log.Info("[tcp.server] Get conn remote addr = ", conn.RemoteAddr().String())

		//3.2 设置服务器最大连接控制,如果超过最大连接，那么则关闭此新的连接
		if s.ConnTotal() >= s.Config.MaxConn {
			log.Warn("[tcp.server] server connect limited")
			err = conn.Close()
			if err != nil {
				log.Errorf("[tcp.server] conn closed err:%v", err)
			}
			continue
		}

		//3.3 处理该新连接请求的 业务 方法， 此时应该有 handler 和 conn是绑定的
		dealConn := NewConnect(s, conn, cid, s.msgHandler)
		cid++
		if cid++; cid == maxInt {
			log.Infof("[tcp.server] conn.acceptTcp num is:%d", cid)
			cid = 1
		}

		//3.4 启动当前链接的处理业务
		go dealConn.Start()
	}
}

//Stop 停止服务
func (s *Server) Stop() {
	fmt.Println("[tcp.server] stop")

	//将其他需要清理的连接信息或者其他信息 也要一并停止或者清理
	for _, manager := range s.ConnMgr {
		manager.Clear()
	}
}

// ConnManager 当前连接的管理器
func (s *Server) GetConnMgr(userId uint32) *ConnManager {
	userIdStr := strconv.Itoa(int(userId))
	idx := crypt.CityHash32([]byte(userIdStr), uint32(len(userIdStr))) % s.BucketSize
	return s.ConnMgr[idx]
}

// ConnTotal 当前服务器的总连接数
func (s *Server) ConnTotal() int {
	var c int
	for _, manager := range s.ConnMgr {
		c += manager.Len()
	}
	return c
}

// Broadcast 广播消息
func (s *Server) Broadcast(msgID uint32, data []byte) {
	for _, manager := range s.ConnMgr {
		for _, conn := range manager.Connections {
			conn.SendBuffMsg(msgID, data)
		}
	}
}
