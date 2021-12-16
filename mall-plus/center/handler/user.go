package handler

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/emptypb"

	"center/service"
	"common/errno"
	pb "common/proto/center"
)

//User 用户中心处理器
type User struct {
	srv service.ICenter
}

//New 实例化用户中心
func New(srv service.ICenter) *User {
	return &User{srv: srv}
}

//Register 注册
func (u *User) Register(ctx context.Context, req *pb.RegisterReq, registerReply *pb.RegisterReply) error {
	id, err := u.srv.UserRegister(ctx, req.Username, req.Password, req.Phone)
	if err != nil {
		return errno.CenterReplyErr(err)
	}
	registerReply.Id = id
	return nil
}

// UsernameLogin 用户名登录
func (u *User) UsernameLogin(ctx context.Context, req *pb.UsernameReq, reply *pb.UserToken) error {
	user, token, err := u.srv.UsernameLogin(ctx, req.Username, req.Password)
	if err != nil {
		return errno.CenterReplyErr(err)
	}
	reply.User = user
	reply.Token = token
	return nil
}

// PhoneLogin 手机号登录
func (u *User) PhoneLogin(ctx context.Context, req *pb.PhoneReq, reply *pb.UserToken) error {
	user, token, err := u.srv.UserPhoneLogin(ctx, req.Phone)
	if err != nil {
		return errno.CenterReplyErr(err)
	}
	reply.User = user
	reply.Token = token
	return nil
}

//Edit 修改用户信息
func (u *User) Edit(ctx context.Context, req *pb.EditReq, empty *emptypb.Empty) error {
	data := make(map[string]interface{})
	err := json.Unmarshal(req.Content, &data)
	if err != nil {
		return errors.Wrapf(err, "[grpc.center] json unmarshal with data:%s", req.Content)
	}
	err = u.srv.UserEdit(ctx, req.Id, data)
	if err != nil {
		return errno.CenterReplyErr(err)
	}
	return nil
}

//EditPwd 修改用户密码
func (u *User) EditPwd(ctx context.Context, req *pb.EditPwdReq, empty *emptypb.Empty) error {
	err := u.srv.UserEditPwd(ctx, req.Id, req.OldPwd, req.Pwd)
	if err != nil {
		return errno.CenterReplyErr(err)
	}
	return nil
}

//Info 获取用户信息
func (u *User) Info(ctx context.Context, req *pb.UIDReq, userinfo *pb.Userinfo) (err error) {
	userinfo, err = u.srv.UserInfoByID(ctx, req.Id)
	if err != nil {
		return errno.CenterReplyErr(err)
	}
	return
}

//Logout 用户登出
func (u *User) Logout(ctx context.Context, req *pb.UIDReq, empty *emptypb.Empty) error {
	err := u.srv.UserLogout(ctx, req.Id)
	if err != nil {
		return errno.CenterReplyErr(err)
	}
	return nil
}

//Online 用户上线
func (u *User) Online(ctx context.Context, req *pb.OnlineReq, onlineReply *pb.OnlineReply) error {
	userID, err := u.srv.UserOnline(ctx, req.Token, req.Server)
	if err != nil {
		return errno.CenterReplyErr(err)
	}
	onlineReply.Uid = userID
	return nil
}

//Offline 用户下线
func (u *User) Offline(ctx context.Context, req *pb.OfflineReq, empty *emptypb.Empty) error {
	err := u.srv.UserOffline(ctx, req.Uid)
	if err != nil {
		return errno.CenterReplyErr(err)
	}
	return nil
}

//ServerID 获取用户长连接所在服务器ID
func (u *User) ServerID(ctx context.Context, req *pb.UIDReq, reply *pb.ServerIDReply) error {
	serverID := u.srv.GetServerID(ctx, req.Id)
	reply.ServerID = serverID
	return nil
}

//BatchServersIDs 批量获取用户长连接所在服务器ID
func (u *User) BatchServersIDs(ctx context.Context, req *pb.UIDsReq, reply *pb.ServerIDsReply) error {
	serverIDs, err := u.srv.BatchServerIds(ctx, req.Ids)
	if err != nil {
		return errno.CenterReplyErr(err)
	}
	reply.ServerIDs = serverIDs
	return nil
}


