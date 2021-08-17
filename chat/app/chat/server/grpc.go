package server

import (
	"chat/app/chat"
	"context"

	"chat/app/chat/conf"
	"chat/pkg/app"
	"chat/pkg/net/grpc"
	"chat/pkg/registry"
	pb "chat/proto/chat"
)

// NewGRPCServer creates a gRPC server
func NewGRPCServer(c *conf.Config, rs registry.Registry, srv chat.IService, engine *Engine) *grpc.Server {
	// 启动grpc服务
	sv := grpc.NewServer(&c.GrpcServer, app.WithHost(c.App.Host), app.WithID(c.App.ServerID),
		app.WithName(c.App.Name), app.WithRegistry(rs))
	sv.Init()
	pb.RegisterChatServer(sv.Server, &server{srv, engine})
	sv.AddHook(grpc.NewTracingHook())
	//启动服务
	sv.Start()
	return sv
}

var _ pb.ChatServer = &server{}

type server struct {
	srv    chat.IService
	engine *Engine
}

//Receive 其他消息路由
func (s server) Receive(ctx context.Context, req *pb.ReceiveReq) (*pb.ReceiveReply, error) {
	return s.engine.Start(ctx, req)
}
