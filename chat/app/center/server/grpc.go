package server

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"chat/app/center"
	"chat/app/center/conf"
	"chat/pkg/app"
	"chat/pkg/net/grpc"
	"chat/pkg/registry"
	"chat/proto/base"
	pb "chat/proto/center"
)

const (
	//grpc 返回错误code
	//UserExisted 用户已存在
	UserExisted codes.Code = 100
	//UserNotFound 用户不存在
	UserNotFound codes.Code = 101
	//UserNotMatch 用户账号不匹配
	UserNotMatch codes.Code = 102
	//UserFrozen 用户已冻结
	UserFrozen codes.Code = 103
	//UserTokenExpired 用户令牌已过期
	UserTokenExpired codes.Code = 104
	//UserTokenError 用户令牌不合法
	UserTokenError codes.Code = 105
	//VerifyCodeRuleMinute 验证码分钟限制
	VerifyCodeRuleMinute codes.Code = 106
	//VerifyCodeRuleHour 验证码小时限制
	VerifyCodeRuleHour codes.Code = 107
	//VerifyCodeRuleDay 验证码天限制
	VerifyCodeRuleDay codes.Code = 108
	//VerifyCodeNotMatch 验证码不匹配
	VerifyCodeNotMatch codes.Code = 109
)

// NewGRPCServer creates a gRPC server
func NewGRPCServer(c *conf.Config, rs registry.Registry, srv center.ICenter) *grpc.Server {
	// 启动grpc服务
	sv := grpc.NewServer(&c.GrpcServer, app.WithHost(c.App.Host), app.WithID(c.App.ServerID),
		app.WithName(c.App.Name), app.WithRegistry(rs))
	sv.Init()
	pb.RegisterCenterServer(sv.Server, &server{srv})
	sv.AddHook(grpc.NewTracingHook())
	//启动服务
	sv.Start()
	return sv
}

var _ pb.CenterServer = &server{}

type server struct {
	srv center.ICenter
}

// UserRegister 用户注册
func (s server) UserRegister(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterReply, error) {
	id, err := s.srv.UserRegister(ctx, req.Username, req.Password, req.Phone)
	if err != nil {
		return nil, reply(err)
	}
	return &pb.RegisterReply{Id: id}, nil
}

// UsernameLogic 用户名登录
func (s server) UsernameLogin(ctx context.Context, req *pb.UsernameReq) (*pb.UserToken, error) {
	user, err := s.srv.UsernameLogin(ctx, req.Username, req.Password)

	if err != nil {
		return nil, reply(err)
	}
	return user, nil
}

// PhoneLogin 手机号登录
func (s server) PhoneLogin(ctx context.Context, req *pb.PhoneReq) (*pb.UserToken, error) {
	user, err := s.srv.UserPhoneLogin(ctx, req.Phone)
	if err != nil {
		return nil, reply(err)
	}
	return user, nil
}

// UserEdit 修改用户信息
func (s server) UserEdit(ctx context.Context, req *pb.EditReq) (*emptypb.Empty, error) {
	data := make(map[string]interface{})
	err := json.Unmarshal(req.Content, &data)
	if err != nil {
		return nil, errors.Wrapf(err, "[grpc.center] json unmarshal with data:%s", req.Content)
	}
	err = s.srv.UserEdit(ctx, req.Id, data)
	if err != nil {
		return nil, reply(err)
	}
	return &emptypb.Empty{}, nil
}

// UserEditPwd 修改密码
func (s server) UserEditPwd(ctx context.Context, req *pb.EditPwdReq) (*emptypb.Empty, error) {
	err := s.srv.UserEditPwd(ctx, req.Id, req.Pwd)
	if err != nil {
		return nil, reply(err)
	}
	return &emptypb.Empty{}, nil
}

//UserInfo 用户详情
func (s server) UserInfo(ctx context.Context, req *pb.UIDReq) (*base.UserInfo, error) {
	user, err := s.srv.UserInfoByID(ctx, req.Id)
	if err != nil {
		return nil, reply(err)
	}
	return user, nil
}

//UserLogout 用户登出
func (s server) UserLogout(ctx context.Context, req *pb.UIDReq) (*emptypb.Empty, error) {
	err := s.srv.UserLogout(ctx, req.Id)
	if err != nil {
		return nil, reply(err)
	}
	return &emptypb.Empty{}, nil
}

// SendSMS 发送短信验证码
func (s server) SendSMS(ctx context.Context, req *pb.PhoneReq) (*pb.CodeReply, error) {
	code, err := s.srv.SendSMS(ctx, strconv.FormatInt(req.Phone, 10))
	if err != nil {
		return nil, reply(err)
	}
	return &pb.CodeReply{Code: code}, nil
}

// CheckVCode 短信验证
func (s server) CheckVCode(ctx context.Context, req *pb.CheckCodeReq) (*emptypb.Empty, error) {
	err := s.srv.CheckVCode(ctx, req.Phone, req.Code)
	if err != nil {
		return nil, reply(err)
	}
	return &emptypb.Empty{}, nil
}

//Online 用户上线
func (s server) Online(ctx context.Context, req *pb.OnlineReq) (*pb.OnlineReply, error) {
	userID, err := s.srv.UserOnline(ctx, req.Token, req.Server)
	if err != nil {
		return nil, reply(err)
	}
	return &pb.OnlineReply{Uid: userID}, nil
}

//Offline 用户下线
func (s server) Offline(ctx context.Context, req *pb.OfflineReq) (*emptypb.Empty, error) {
	err := s.srv.UserOffline(ctx, req.Uid)
	if err != nil {
		return nil, reply(err)
	}
	return &emptypb.Empty{}, nil
}

//ServerByUserID 获取用户长连接所在服务器ID
func (s server) ServerByUserID(ctx context.Context, req *pb.UIDReq) (*pb.ServerIDReply, error) {
	serverID := s.srv.ServerByUserID(ctx, req.Id)
	return &pb.ServerIDReply{ServerID: serverID}, nil
}

//BatchServersByUserIDs 批量获取用户长连接所在服务器ID
func (s server) BatchServersByUserIDs(ctx context.Context, req *pb.UIDsReq) (*pb.ServerIDsReply, error) {
	serverIDs, err := s.srv.ServersByUserIds(ctx, req.Ids)
	if err != nil {
		return nil, reply(err)
	}
	return &pb.ServerIDsReply{ServerIDs: serverIDs}, nil
}

// CheckOnline 用户是否在线
func (s server) CheckOnline(ctx context.Context, req *pb.UIDReq) (*pb.BoolReply, error) {
	is, err := s.srv.CheckOnline(ctx, req.Id)
	if err != nil {
		return nil, reply(err)
	}
	return &pb.BoolReply{Is: is}, nil
}

// reply 错误处理
func reply(err error) error {
	switch err {
	case center.ErrUserExisted:
		return status.Errorf(UserExisted, "User already exists")
	case center.ErrUserNotFound:
		return status.Errorf(UserNotFound, "Account not found")
	case center.ErrUserNotMatch:
		return status.Errorf(UserNotMatch, "Account password does not match")
	case center.ErrUserFrozen:
		return status.Errorf(UserFrozen, "Account has been frozen")
	case center.ErrUserTokenExpired:
		return status.Errorf(UserTokenExpired, "Token has expired")
	case center.ErrUserTokenError:
		return status.Errorf(UserTokenError, "Token error")
	case center.ErrVerifyCodeNotMatch:
		return status.Errorf(VerifyCodeNotMatch, "One minute limit")
	case center.ErrVerifyCodeRuleDay:
		return status.Errorf(VerifyCodeRuleDay, "One day limit")
	case center.ErrVerifyCodeRuleHour:
		return status.Errorf(VerifyCodeRuleHour, "One hour limit")
	case center.ErrVerifyCodeRuleMinute:
		return status.Errorf(VerifyCodeRuleMinute, "One minute limit")
	}
	return err
}

//HandleError 处理中心服错误
func HandleError(err error) error {
	if st, ok := status.FromError(err); ok {
		switch st.Code() {
		case UserExisted:
			return center.ErrUserExisted
		case UserNotFound:
			return center.ErrUserNotFound
		case UserNotMatch:
			return center.ErrUserNotMatch
		case UserFrozen:
			return center.ErrUserFrozen
		case UserTokenExpired:
			return center.ErrUserTokenExpired
		case UserTokenError:
			return center.ErrUserTokenError
		case VerifyCodeRuleMinute:
			return center.ErrVerifyCodeRuleMinute
		case VerifyCodeRuleHour:
			return center.ErrVerifyCodeRuleHour
		case VerifyCodeRuleDay:
			return center.ErrVerifyCodeRuleDay
		case VerifyCodeNotMatch:
			return center.ErrVerifyCodeNotMatch
		}
	}
	return errors.Wrapf(err, "[grpc.center] call")
}
