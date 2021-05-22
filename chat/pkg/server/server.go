package server

// 用于触发编译期的接口的合理性检查机制
var _ IServer = (*Server)(nil)

//定义服务接口
type IServer interface {
	Start()                                //启动服务器方法
	Stop()                                 //停止服务器方法
	Serve()                                //开启业务服务方法
	SetOnConnStart(func(IConnection) bool) //设置该Server的连接创建时Hook函数
	SetOnConnFinish(func(IConnection))     //设置该Server的连接鉴权完成时Hook函数
	SetOnConnStop(func(IConnection))       //设置该Server的连接断开时的Hook函数
	CallOnConnStart(conn IConnection) bool //调用连接OnConnStart Hook函数
	CallOnConnFinish(conn IConnection)     //调用连接OnConnFinish Hook函数
	CallOnConnStop(conn IConnection)       //调用连接OnConnStop Hook函数
}

// Server 接口实现，定义一个Server服务类
type Server struct {
	//服务器ID
	ID string
	//服务绑定的地址
	Addr string
	//该Server的连接创建时Hook函数
	OnConnStart func(conn IConnection) bool
	//该Server的连接断开时的Hook函数
	OnConnStop func(conn IConnection)
	//该Server的连接鉴权完成的Hook函数
	OnConnFinish func(conn IConnection)
	//管理器数量
	BucketSize uint32
}

// Start 开启网络服务
func (s *Server) Start() {}

//Stop 停止服务
func (s *Server) Stop() {}

//Serve 运行服务
func (s *Server) Serve() {}

//SetOnConnStart 设置该Server的连接创建时Hook函数
func (s *Server) SetOnConnStart(hook func(IConnection) bool) {
	s.OnConnStart = hook
}

//SetOnConnFinish 设置该Server的连接鉴权完成时的Hook函数
func (s *Server) SetOnConnFinish(hook func(IConnection)) {
	s.OnConnFinish = hook
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

//CallOnConnStop 调用连接OnConnFinish Hook函数
func (s *Server) CallOnConnFinish(conn IConnection) {
	if s.OnConnFinish != nil {
		s.OnConnFinish(conn)
	}
}

//CallOnConnStop 调用连接OnConnStop Hook函数
func (s *Server) CallOnConnStop(conn IConnection) {
	if s.OnConnStop != nil {
		s.OnConnStop(conn)
	}
}
