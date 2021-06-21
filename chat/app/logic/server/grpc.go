package server

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/emptypb"

	"chat/app/logic/conf"
	"chat/app/logic/service"
	"chat/pkg/app"
	"chat/pkg/errno"
	"chat/pkg/net/grpc"
	"chat/pkg/registry"
	pb "chat/proto/logic"
)

// NewGRPCServer creates a gRPC server
func NewGRPCServer(c *conf.Config, rs registry.Registry, srv service.IService, engine *Engine) *grpc.Server {
	// 启动grpc服务
	sv := grpc.NewServer(&c.GrpcServer, app.WithHost(c.App.Host), app.WithID(c.App.ServerID),
		app.WithName(c.App.Name), app.WithRegistry(rs))
	sv.Init()
	pb.RegisterLogicServer(sv.Server, &server{srv, engine})
	sv.AddHook(grpc.NewTracingHook())
	//启动服务
	sv.Start()
	return sv
}

var _ pb.LogicServer = &server{}

type server struct {
	srv    service.IService
	engine *Engine
}

//Online 用户上线
func (s server) Online(ctx context.Context, req *pb.OnlineReq) (*pb.OnlineReply, error) {

	userID, err := s.srv.UserOnline(ctx, req.Token, req.Server)
	if errors.Is(err, errno.ErrTokenTimeout) {
		return nil, errno.ErrTokenTimeout
	} else if err != nil {
		return nil, errors.Wrapf(err, "[grpc.online] set redis err")
	}
	return &pb.OnlineReply{
		Uid: userID,
	}, nil
}

//Offline 用户下线
func (s server) Offline(ctx context.Context, req *pb.OfflineReq) (reply *emptypb.Empty, err error) {
	err = s.srv.UserOffline(ctx, req.Uid)
	if err != nil {
		return nil, errors.Wrapf(err, "[grpc.offline] del redis err")
	}
	return &emptypb.Empty{}, nil
}

//Receive 其他消息路由
func (s server) Receive(ctx context.Context, req *pb.ReceiveReq) (*pb.ReceiveReply, error) {
	return s.engine.Start(ctx, req)
}
