package connect

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"chat/pkg/app"
	"chat/pkg/crypt"
	"chat/pkg/registry"
)

// 用于触发编译期的接口的合理性检查机制
var _ IServer = (*Server)(nil)

//IServer 定义服务接口
type IServer interface {
	Start()    //启动服务器方法
	Stop()     //停止服务器方法
	Serve()    //开启业务服务方法
	Register() //服务注册

	ConnTotal() int                        //已建立连接总数
	Broadcast(msgID uint32, msg []byte)    //广播
	GetConnMgr(userID uint32) IConnManager //连接管理器
	GetServerID() string                   //获取服务器ID
	GetServerType() int8                   //获取服务器类型

	SetOnConnStart(func(IConnection) bool)        //设置该Server的连接创建时Hook函数
	SetOnConnAuth(func(IConnection, []byte) bool) //设置该Server的连接鉴权Hook函数
	SetOnConnStop(func(IConnection))              //设置该Server的连接断开时的Hook函数
	CallOnConnStart(IConnection) bool             //调用连接OnConnStart Hook函数
	CallOnConnAuth(IConnection, []byte) bool      //调用连接OnConnAuth Hook函数
	CallOnConnStop(IConnection)                   //调用连接OnConnStop Hook函数
}

// Server 接口实现，定义一个Server服务类
type Server struct {
	//监听端口
	Port int
	//服务器类型
	Type int8
	//连接管理器
	ConnMgr []IConnManager
	//当前Server的消息管理模块，用来绑定MsgID和对应的处理方法
	MsgHandler IMsgHandle
	//该Server的连接创建时Hook函数
	OnConnStart func(conn IConnection) bool
	//该Server的连接断开时的Hook函数
	OnConnStop func(conn IConnection)
	//该Server的连接鉴权完成的Hook函数
	OnConnAuth func(IConnection, []byte) bool
	//管理器数量
	BucketSize uint32
	//ip限流器
	Limit   *Limiter
	Options *app.Options
}

// Start 开启网络服务
func (s *Server) Start() {}

//Stop 停止服务
func (s *Server) Stop() {
	fmt.Println("server stop")
	if s.Options.Register != nil { //注销服务
		s.Options.Register.Unregister(context.Background(), &registry.Service{ID: app.BuildServerID(s.Options.ID, s.Port)})
	}
	//将其他需要清理的连接信息或者其他信息 也要一并停止或者清理
	for _, manager := range s.ConnMgr {
		manager.Clear()
	}
}

//Serve 运行服务
func (s *Server) Serve() {}

//Register 服务注册
func (s *Server) Register() {
	if s.Options.Register != nil {
		err := s.Options.Register.Register(context.Background(), &registry.Service{
			ID:   app.BuildServerID(s.Options.ID, s.Port),
			Name: s.Options.Name,
			IP:   s.Options.Host,
			Port: s.Port,
			Check: registry.Check{
				TCP: fmt.Sprintf("%s:%d", s.Options.Host, s.Port),
			},
		})
		if err != nil {
			log.Fatalf("failed to serve register %s server: %v", s.Options.Name, err)
		}
		log.Printf("server register success id:%v, name:%v", app.BuildServerID(s.Options.ID, s.Port), s.Options.Name)
	}
}

// GetConnMgr 当前连接的管理器
func (s *Server) GetConnMgr(userID uint32) IConnManager {
	userIDStr := strconv.Itoa(int(userID))
	idx := crypt.CityHash32([]byte(userIDStr), uint32(len(userIDStr))) % s.BucketSize
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
func (s *Server) Broadcast(msgID uint32, msg []byte) {
	for _, manager := range s.ConnMgr {
		for _, conn := range manager.GetConnections() {
			conn.SendBuffMsg(msgID, msg)
		}
	}
}

//GetServerID 获取服务器ID
func (s *Server) GetServerID() string {
	return app.BuildServerID(s.Options.ID, s.Port)
}

//GetServerType 获取服务器类型
func (s *Server) GetServerType() int8 {
	return s.Type
}

//SetOnConnStart 设置该Server的连接创建时Hook函数
func (s *Server) SetOnConnStart(hook func(IConnection) bool) {
	s.OnConnStart = hook
}

//SetOnConnAuth 设置该Server的连接鉴权的Hook函数
func (s *Server) SetOnConnAuth(hook func(IConnection, []byte) bool) {
	s.OnConnAuth = hook
}

//SetOnConnStop 设置该Server的连接断开时的Hook函数
func (s *Server) SetOnConnStop(hook func(IConnection)) {
	s.OnConnStop = hook
}

//CallOnConnStart 调用连接OnConnStart Hook函数
func (s *Server) CallOnConnStart(conn IConnection) bool {
	if s.OnConnStart != nil {
		return s.OnConnStart(conn)
	}
	return true
}

//CallOnConnAuth 调用连接OnConnAuth Hook函数
func (s *Server) CallOnConnAuth(conn IConnection, data []byte) bool {
	if s.OnConnAuth != nil {
		return s.OnConnAuth(conn, data)
	}
	return true
}

//CallOnConnStop 调用连接OnConnStop Hook函数
func (s *Server) CallOnConnStop(conn IConnection) {
	if s.OnConnStop != nil {
		s.OnConnStop(conn)
	}
}
