package grpc

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	"chat/app/connect"
	"chat/app/connect/conf"
	"chat/pkg/log"
	"chat/pkg/registry"
	grpc2 "chat/pkg/server/grpc"
	pb "chat/proto/connect"
)

// New logic grpc server
func New(c *conf.Config, srv *connect.Server, rs registry.Registry) *grpc2.Server {
	// 启动grpc服务
	sv := grpc2.NewServer(&c.GrpcServer, c.Host, func(s *grpc.Server) {
		pb.RegisterConnectServer(s, &server{srv})
	})
	sv.SetOnRequestStart(DoRequestStart)
	sv.SetOnRequestEnd(DoRequestEnd)
	// 服务注册
	err := rs.Register(context.Background(), &registry.Service{
		Id:   conf.Conf.ServerId,
		Name: conf.Conf.Name,
		IP:   conf.Conf.Host,
		Port: conf.Conf.GrpcServer.Port,
		Check: registry.Check{
			GRPC: fmt.Sprintf("%v:%v/%v", conf.Conf.Host, conf.Conf.GrpcServer.Port, conf.Conf.Name),
		},
	})
	if err != nil {
		log.Fatalf("[RegisterServer] failed to register %s server: %v", conf.Conf.Name, err)
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

var _ pb.ConnectServer = &server{}

type server struct {
	srv *connect.Server
}

//Close 主动关闭用户连接
func (s *server) Close(ctx context.Context, req *pb.CloseReq) (*emptypb.Empty, error) {
	log.Infof("[grpc.close] begin req:%v", req)
	conn, err := s.srv.WsServer.GetConnMgr(req.UserId).Get(req.UserId)
	if err != nil {
		return nil, errors.Wrapf(err, "[connect.close] get conn err uid:%v", req.UserId)
	}
	// 发送提示消息后关闭客户端
	conn.SendMsg(req.Proto.GetData())
	conn.Stop()
	return &emptypb.Empty{}, nil
}

//Send 发送消息给指定用户
func (s *server) Send(ctx context.Context, req *pb.SendReq) (*emptypb.Empty, error) {
	log.Infof("[grpc.send] begin req:%v", req)
	if len(req.UserIds) == 0 || req.Proto == nil {
		return nil, errors.New("rpc send msg arg error")
	}
	for _, uid := range req.UserIds {
		conn, err := s.srv.WsServer.GetConnMgr(uid).Get(uid)
		if err != nil {
			log.Errorf("[connect.send] get conn err uid:%v", uid)
			continue
		}
		if err = conn.SendBuffMsg(req.Proto.GetData()); err != nil {
			log.Errorf("[connect.send] send msg err uid:%v", uid)
			return nil, err
		}
	}
	return &emptypb.Empty{}, nil
}

//Broadcast 广播消息至所有用户
func (s *server) Broadcast(ctx context.Context, req *pb.BroadcastReq) (reply *emptypb.Empty, err error) {
	log.Infof("[grpc.broadcast] begin req:%v", req)
	go s.srv.WsServer.Broadcast(req.Proto.GetData())
	return
}
