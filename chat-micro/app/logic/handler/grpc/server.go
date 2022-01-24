package grpc

import (
	"chat-micro/app/logic/service"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "chat-micro/proto/logic"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	//grpc 返回错误code
	//UserFrozen 用户已冻结
	UserFrozen codes.Code = 103
	//UserTokenExpired 用户令牌已过期
	UserTokenExpired codes.Code = 104
	//UserTokenError 用户令牌不合法
	UserTokenError codes.Code = 105
)

var _ pb.LogicServer = &server{}

func New(svc *service.Service) *server {
	return &server{
		svc: svc,
	}
}

type server struct {
	svc *service.Service
	pb.UnimplementedLogicServer
}

//Online 用户上线
func (s server) Online(ctx context.Context, req *pb.OnlineReq) (*pb.OnlineReply, error) {
	userID, err := s.svc.UserOnline(ctx, req.Token, req.ServerId)
	if err != nil {
		return nil, reply(err)
	}
	return &pb.OnlineReply{Uid: userID}, nil
}

//Offline 用户下线
func (s server) Offline(ctx context.Context, req *pb.OfflineReq) (*emptypb.Empty, error) {
	err := s.svc.UserOffline(ctx, req.Uid)
	if err != nil {
		return nil, reply(err)
	}
	return &emptypb.Empty{}, nil
}

//ServerByUserID 获取用户所有的服务器id
func (s server) ServerByUserID(ctx context.Context, req *pb.UIDReq) (*pb.ServerIDReply, error) {
	serverID := s.svc.ServerByUserID(ctx, req.Id)
	return &pb.ServerIDReply{ServerId: serverID}, nil
}

//BatchServersByUserIDs 批量获取用户的服务器id
func (s server) BatchServersByUserIDs(ctx context.Context, req *pb.UIDsReq) (*pb.ServerIDsReply, error) {
	serverIDs, err := s.svc.ServersByUserIds(ctx, req.Ids)
	if err != nil {
		return nil, reply(err)
	}
	return &pb.ServerIDsReply{ServerIds: serverIDs}, nil
}

// reply 错误处理
func reply(err error) error {
	switch err {
	case service.ErrUserFrozen:
		return status.Errorf(UserFrozen, "Account has been frozen")
	case service.ErrUserTokenExpired:
		return status.Errorf(UserTokenExpired, "Token has expired")
	case service.ErrUserTokenError:
		return status.Errorf(UserTokenError, "Token error")
	}
	return err
}