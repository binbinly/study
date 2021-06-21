package grpc

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/emptypb"

	"chat/app/connect"
	"chat/app/connect/conf"
	"chat/pkg/app"
	"chat/pkg/log"
	"chat/pkg/net/grpc"
	"chat/pkg/registry"
	pb "chat/proto/connect"
)

// New logic grpc server
func New(c *conf.Config, srv *connect.Server, rs registry.Registry) *grpc.Server {
	// 启动grpc服务
	sv := grpc.NewServer(&c.GrpcServer, app.WithHost(c.App.Host), app.WithID(c.App.ServerID),
		app.WithName(c.App.Name), app.WithRegistry(rs))
	sv.Init()
	pb.RegisterConnectServer(sv.Server, &server{srv})
	sv.AddHook(grpc.NewTracingHook())
	//启动服务
	sv.Start()
	return sv
}

var _ pb.ConnectServer = &server{}

type server struct {
	srv *connect.Server
}

//Close 主动关闭用户连接
func (s *server) Close(ctx context.Context, req *pb.CloseReq) (*emptypb.Empty, error) {
	log.Debugf("[grpc.close] begin req:%v", req)
	conn, err := s.srv.Server.GetConnMgr(req.UserId).Get(req.UserId)
	if err != nil {
		return nil, errors.Wrapf(err, "[connect.close] get conn err uid:%v", req.UserId)
	}
	// 发送提示消息后关闭客户端
	conn.SendMsg(uint32(req.Proto.GetMsgId()), req.Proto.GetData())
	conn.Stop()
	return &emptypb.Empty{}, nil
}

//Send 发送消息给指定用户
func (s *server) Send(ctx context.Context, req *pb.SendReq) (*emptypb.Empty, error) {
	log.Debugf("[grpc.send] begin req:%v", req)
	if len(req.UserIds) == 0 || req.Proto == nil {
		return nil, errors.New("rpc send msg arg error")
	}
	for _, uid := range req.UserIds {
		conn, err := s.srv.Server.GetConnMgr(uid).Get(uid)
		if err != nil {
			log.Warnf("[connect.send] get conn err uid:%v", uid)
			continue
		}
		if err = conn.SendBuffMsg(uint32(req.Proto.MsgId), req.Proto.GetData()); err != nil {
			log.Warnf("[connect.send] send msg err uid:%v", uid)
			return nil, err
		}
	}
	return &emptypb.Empty{}, nil
}

//Broadcast 广播消息至所有用户
func (s *server) Broadcast(ctx context.Context, req *pb.BroadcastReq) (reply *emptypb.Empty, err error) {
	log.Debugf("[grpc.broadcast] begin req:%v", req)
	go s.srv.Server.Broadcast(uint32(req.Proto.MsgId), req.Proto.GetData())
	return
}
