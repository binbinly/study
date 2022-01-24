package grpc

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/emptypb"

	"chat-micro/app/connect"
	"chat-micro/pkg/logger"
	pb "chat-micro/proto/connect"
)

var _ pb.ConnectServer = &server{}

func New(svc *connect.Connect) *server {
	return &server{
		svc: svc,
	}
}

type server struct {
	svc *connect.Connect
	pb.UnimplementedConnectServer
}

//Close 主动关闭用户连接
func (s *server) Close(ctx context.Context, req *pb.CloseReq) (*emptypb.Empty, error) {
	conn, err := s.svc.Server().GetBucket().GetConnMgr(req.UserId).Get(req.UserId)
	if err != nil {
		return nil, errors.Wrapf(err, "[connect.close] get conn err uid:%v", req.UserId)
	}
	// 发送提示消息后关闭客户端
	conn.SendMsg(req.MsgId, req.Data)
	conn.Stop()
	return &emptypb.Empty{}, nil
}

//Send 发送消息给指定用户
func (s *server) Send(ctx context.Context, req *pb.SendReq) (*emptypb.Empty, error) {
	for _, id := range req.UserIds {
		conn, err := s.svc.Server().GetBucket().GetConnMgr(id).Get(id)
		if err != nil {
			logger.Warnf("[connect.send] get conn uid:%v err: %v", id, err)
			continue
		}
		if err = conn.SendBuffMsg(req.MsgId, req.Data); err != nil {
			logger.Warnf("[connect.send] send msg uid:%v err: %v", id, err)
			return nil, err
		}
	}
	return &emptypb.Empty{}, nil
}

//Broadcast 广播消息至所有用户
func (s *server) Broadcast(ctx context.Context, req *pb.BroadcastReq) (*emptypb.Empty, error) {
	s.svc.Server().GetBucket().Broadcast(req.MsgId, req.Data)
	return &emptypb.Empty{}, nil
}
