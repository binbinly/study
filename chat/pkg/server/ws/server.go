package ws

import (
	"fmt"
	"net/http"
	"strconv"

	"chat/pkg/crypt"
	"chat/pkg/log"
	"chat/pkg/server"
)

type Server struct {
	*server.Server
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
	fmt.Println("[server.ws] start!!!")
	s.Start()
}

// Start 开启网络服务
func (s *Server) Start() {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		log.Info("[ws.accept] start", r.RemoteAddr)
		//设置服务器最大连接控制,如果超过最大连接，那么则关闭此新的连接
		if s.ConnTotal() >= s.Config.MaxConn {
			log.Warn("[ws.server] connection size limit")
			return
		}
		conn := NewConnect(s, w, r)
		conn.Start()
	})
	// 启动worker工作池机制
	s.msgHandler.StartWorkerPool(s.Config.MaxWorkerTaskLen)
	go func() {
		if err := http.ListenAndServe(s.Addr, nil); err != nil {
			log.Fatalf("ListenAndServe, err: %s", err.Error())
		}
	}()
	fmt.Printf("[ws.server] Websocket Server listener at Addr:%v is starting\n", s.Addr)
}

//Stop 停止服务
func (s *Server) Stop() {
	fmt.Println("[connect.server] stop")

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
func (s *Server) Broadcast(msg []byte) {
	for _, manager := range s.ConnMgr {
		for _, conn := range manager.Connections {
			conn.SendBuffMsg(msg)
		}
	}
}
