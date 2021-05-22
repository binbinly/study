package grpc

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	"chat/app/logic/conf"
	"chat/app/logic/service"
	"chat/pkg/errno"
	"chat/pkg/log"
	"chat/pkg/registry"
	grpc2 "chat/pkg/server/grpc"
	pb "chat/proto/logic"
)

// New logic grpc server
func New(c *conf.Config, rs registry.Registry, srv service.IService, engine *Engine) *grpc2.Server {
	// 启动grpc服务
	sv := grpc2.NewServer(&c.GrpcServer, c.App.Host, func(s *grpc.Server) {
		pb.RegisterLogicServer(s, &server{srv, engine})
	})
	sv.SetOnRequestStart(DoRequestStart)
	sv.SetOnRequestEnd(DoRequestEnd)
	// 服务注册
	err := rs.Register(context.Background(), &registry.Service{
		Id:   c.App.ServerId,
		Name: c.App.Name,
		IP:   c.App.Host,
		Port: c.GrpcServer.Port,
		Check: registry.Check{
			GRPC: fmt.Sprintf("%v:%v/%v", c.App.Host, c.GrpcServer.Port, c.App.Name),
		},
	})
	if err != nil {
		log.Fatalf("[RegisterServer] failed to register %s server: %v", c.App.Name, err)
	}
	return sv
}

//DoRequestStart 请求开始执行
func DoRequestStart(ctx context.Context, req interface{}) {
	log.Infof("do request start req:%v", req)
}

//DoRequestEnd 请求结束执行
func DoRequestEnd(ctx context.Context, req interface{}) {
	log.Infof("do request end req:%v", req)
}

var _ pb.LogicServer = &server{}

type server struct {
	srv    service.IService
	engine *Engine
}

//Online 用户上线
func (s server) Online(ctx context.Context, req *pb.OnlineReq) (*pb.OnlineReply, error) {

	userId, err := s.srv.UserOnline(ctx, req.Token, req.Server)
	if errors.Is(err, errno.ErrTokenTimeout) {
		return nil, errno.ErrTokenTimeout
	} else if err != nil {
		return nil, errors.Wrapf(err, "[grpc.online] set redis err")
	}
	return &pb.OnlineReply{
		Uid: userId,
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
	return s.engine.Start(req)
}
