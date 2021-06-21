package connect

import (
	"context"
	"log"

	grpc2 "google.golang.org/grpc"

	"chat/app/connect/conf"
	"chat/pkg/connect"
	"chat/pkg/net/grpc"
	"chat/pkg/registry"
	"chat/pkg/registry/consul"
	"chat/proto/logic"
)

//Svc 全局服务
var Svc *Server

// Server is connect server.
type Server struct {
	Server   connect.IServer

	c         *conf.Config
	rpcConn   *grpc2.ClientConn
	rpcClient logic.LogicClient //grpc客户端
}

// NewServer returns a new Server.
func NewServer(c *conf.Config) *Server {
	s := &Server{
		c:       c,
		rpcConn: newLogicClient(c),
	}
	s.rpcClient = logic.NewLogicClient(s.rpcConn)
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

//newLogicClient 创建逻辑层客户端
func newLogicClient(c *conf.Config) *grpc2.ClientConn {
	// 初始化consul.resolver
	consul.Init()
	ctx, cancel := context.WithTimeout(context.Background(), c.GrpcClient.Timeout)
	defer cancel()
	target := "consul://" + c.Registry.Host + "/" + c.GrpcClient.ServiceName
	conn, err := grpc.NewRPCClientConn(ctx, &c.GrpcClient, target)
	if err != nil {
		log.Fatalf("new grpc client conn failed err:%v", err)
	}
	return conn
}
