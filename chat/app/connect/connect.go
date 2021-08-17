package connect

import (
	grpc2 "google.golang.org/grpc"

	"chat/app/connect/conf"
	"chat/pkg/connect"
	"chat/pkg/net/grpc"
	"chat/pkg/registry"
	"chat/pkg/registry/consul"
	"chat/proto/center"
)

//Svc 全局服务
var Svc *Server

// Server is connect server.
type Server struct {
	Server connect.IServer

	c         *conf.Config
	rpcConn   *grpc2.ClientConn
	rpcClient center.CenterClient //grpc客户端
}

// NewServer returns a new Server.
func NewServer(c *conf.Config) *Server {
	// 初始化consul.resolver
	consul.Init()
	target := "consul://" + c.Registry.Host + "/" + c.GrpcClient.ServiceName
	conn := grpc.NewRPCClientConn(&c.GrpcClient, target)
	s := &Server{
		c:         c,
		rpcConn:   conn,
		rpcClient: center.NewCenterClient(conn),
	}
	Svc = s
	return s
}

// StartWs 开启websocket服务器 监听ws连接
func (s *Server) StartWs(rs registry.Registry, serverID string) {
	s.Server = StartWsServer(s.c, rs, serverID)
}

// StartTCP 开启tcp服务器 监听tcp连接
func (s *Server) StartTCP(rs registry.Registry, serverID string) {
	s.Server = StartTCPServer(s.c, rs, serverID)
}

// Close close the server.
func (s *Server) Close() (err error) {
	if s.Server != nil {
		s.Server.Stop()
	}
	return s.rpcConn.Close()
}
