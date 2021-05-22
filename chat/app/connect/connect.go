package connect

import (
	"context"
	"fmt"
	"log"

	grpc2 "google.golang.org/grpc"

	"chat/app/connect/conf"
	"chat/pkg/registry"
	"chat/pkg/registry/consul"
	"chat/pkg/server/grpc"
	"chat/pkg/server/tcp"
	"chat/pkg/server/ws"
	"chat/proto/logic"
)

var Svc *Server

// Server is comet server.
type Server struct {
	c         *conf.Config
	rs        registry.Registry
	rpcConn   *grpc2.ClientConn
	rpcClient logic.LogicClient
	WsServer  *ws.Server
	TcpServer *tcp.Server
}

// NewServer returns a new Server.
func NewServer(c *conf.Config, rs registry.Registry) *Server {
	s := &Server{
		c:       c,
		rs:      rs,
		rpcConn: newLogicClient(c),
	}
	s.rpcClient = logic.NewLogicClient(s.rpcConn)
	Svc = s
	return s
}

// StartWsServer 开启websocket服务器 监听ws连接
func (s *Server) StartWsServer() {
	s.WsServer = NewWsServer(&s.c.Ws)
	// 服务注册
	err := s.rs.Register(context.Background(), &registry.Service{
		Id:   "w-" + conf.Conf.ServerId,
		Name: conf.Conf.Name + "-ws",
		IP:   conf.Conf.Host,
		Port: conf.Conf.Ws.Port,
		Check: registry.Check{
			TCP: fmt.Sprintf("%s:%d", conf.Conf.Host, conf.Conf.Ws.Port),
		},
	})
	if err != nil {
		log.Fatalf("registry failed to websocket register %s server: %v", conf.Conf.Name, err)
	}
}

// StartTcpServer 开启tcp服务器 监听tcp连接
func (s *Server) StartTcpServer() {
	s.TcpServer = NewTcpServer(&s.c.Tcp)
	// 服务注册
	err := s.rs.Register(context.Background(), &registry.Service{
		Id:   "t-" + conf.Conf.ServerId,
		Name: conf.Conf.Name + "-tcp",
		IP:   conf.Conf.Host,
		Port: conf.Conf.Tcp.Port,
		Check: registry.Check{
			TCP: fmt.Sprintf("%s:%d", conf.Conf.Host, conf.Conf.Tcp.Port),
		},
	})
	if err != nil {
		log.Fatalf("registry failed to tcp register %s server: %v", conf.Conf.Name, err)
	}
}

// Close close the server.
func (s *Server) Close() (err error) {
	if s.WsServer != nil {
		s.WsServer.Stop()
		s.rs.Unregister(context.Background(), &registry.Service{Id: "w-" + conf.Conf.ServerId})
	}
	if s.TcpServer != nil {
		s.TcpServer.Stop()
		s.rs.Unregister(context.Background(), &registry.Service{Id: "t-" + conf.Conf.ServerId})
	}
	return s.rpcConn.Close()
}

//newLogicClient 创建逻辑层客户端
func newLogicClient(c *conf.Config) *grpc2.ClientConn {
	// 初始化consul.resolver
	consul.Init()
	ctx, cancel := context.WithTimeout(context.Background(), c.GrpcClient.Timeout)
	defer cancel()
	target := "consul://" + c.Registry.Host + "/" + c.LogicName
	conn, err := grpc.NewRpcClientConn(&c.GrpcClient, ctx, target)
	if err != nil {
		log.Fatalf("new grpc client conn failed err:%v", err)
	}
	return conn
}
