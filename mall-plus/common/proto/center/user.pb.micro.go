// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: center/user.proto

package common

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "google.golang.org/protobuf/proto"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for User service

func NewUserEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for User service

type UserService interface {
	/// 用户注册
	Register(ctx context.Context, in *RegisterReq, opts ...client.CallOption) (*RegisterReply, error)
	/// 用户名密码登录
	UsernameLogin(ctx context.Context, in *UsernameReq, opts ...client.CallOption) (*UserToken, error)
	/// 手机号登录
	PhoneLogin(ctx context.Context, in *PhoneReq, opts ...client.CallOption) (*UserToken, error)
	/// 修改用户信息
	Edit(ctx context.Context, in *EditReq, opts ...client.CallOption) (*emptypb.Empty, error)
	/// 修改密码
	EditPwd(ctx context.Context, in *EditPwdReq, opts ...client.CallOption) (*emptypb.Empty, error)
	/// 获取用户信息
	Info(ctx context.Context, in *UIDReq, opts ...client.CallOption) (*Userinfo, error)
	/// 用户登出
	Logout(ctx context.Context, in *UIDReq, opts ...client.CallOption) (*emptypb.Empty, error)
	/// 用户上线，建立长连接
	Online(ctx context.Context, in *OnlineReq, opts ...client.CallOption) (*OnlineReply, error)
	/// 用户下线，断开长连接
	Offline(ctx context.Context, in *OfflineReq, opts ...client.CallOption) (*emptypb.Empty, error)
	/// 获取用户长连接所在的服务器ID
	ServerID(ctx context.Context, in *UIDReq, opts ...client.CallOption) (*ServerIDReply, error)
	/// 批量获取长连接所在的服务器ID
	BatchServersIDs(ctx context.Context, in *UIDsReq, opts ...client.CallOption) (*ServerIDsReply, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) Register(ctx context.Context, in *RegisterReq, opts ...client.CallOption) (*RegisterReply, error) {
	req := c.c.NewRequest(c.name, "User.Register", in)
	out := new(RegisterReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UsernameLogin(ctx context.Context, in *UsernameReq, opts ...client.CallOption) (*UserToken, error) {
	req := c.c.NewRequest(c.name, "User.UsernameLogin", in)
	out := new(UserToken)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) PhoneLogin(ctx context.Context, in *PhoneReq, opts ...client.CallOption) (*UserToken, error) {
	req := c.c.NewRequest(c.name, "User.PhoneLogin", in)
	out := new(UserToken)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Edit(ctx context.Context, in *EditReq, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "User.Edit", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) EditPwd(ctx context.Context, in *EditPwdReq, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "User.EditPwd", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Info(ctx context.Context, in *UIDReq, opts ...client.CallOption) (*Userinfo, error) {
	req := c.c.NewRequest(c.name, "User.Info", in)
	out := new(Userinfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Logout(ctx context.Context, in *UIDReq, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "User.Logout", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Online(ctx context.Context, in *OnlineReq, opts ...client.CallOption) (*OnlineReply, error) {
	req := c.c.NewRequest(c.name, "User.Online", in)
	out := new(OnlineReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Offline(ctx context.Context, in *OfflineReq, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "User.Offline", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) ServerID(ctx context.Context, in *UIDReq, opts ...client.CallOption) (*ServerIDReply, error) {
	req := c.c.NewRequest(c.name, "User.ServerID", in)
	out := new(ServerIDReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) BatchServersIDs(ctx context.Context, in *UIDsReq, opts ...client.CallOption) (*ServerIDsReply, error) {
	req := c.c.NewRequest(c.name, "User.BatchServersIDs", in)
	out := new(ServerIDsReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserHandler interface {
	/// 用户注册
	Register(context.Context, *RegisterReq, *RegisterReply) error
	/// 用户名密码登录
	UsernameLogin(context.Context, *UsernameReq, *UserToken) error
	/// 手机号登录
	PhoneLogin(context.Context, *PhoneReq, *UserToken) error
	/// 修改用户信息
	Edit(context.Context, *EditReq, *emptypb.Empty) error
	/// 修改密码
	EditPwd(context.Context, *EditPwdReq, *emptypb.Empty) error
	/// 获取用户信息
	Info(context.Context, *UIDReq, *Userinfo) error
	/// 用户登出
	Logout(context.Context, *UIDReq, *emptypb.Empty) error
	/// 用户上线，建立长连接
	Online(context.Context, *OnlineReq, *OnlineReply) error
	/// 用户下线，断开长连接
	Offline(context.Context, *OfflineReq, *emptypb.Empty) error
	/// 获取用户长连接所在的服务器ID
	ServerID(context.Context, *UIDReq, *ServerIDReply) error
	/// 批量获取长连接所在的服务器ID
	BatchServersIDs(context.Context, *UIDsReq, *ServerIDsReply) error
}

func RegisterUserHandler(s server.Server, hdlr UserHandler, opts ...server.HandlerOption) error {
	type user interface {
		Register(ctx context.Context, in *RegisterReq, out *RegisterReply) error
		UsernameLogin(ctx context.Context, in *UsernameReq, out *UserToken) error
		PhoneLogin(ctx context.Context, in *PhoneReq, out *UserToken) error
		Edit(ctx context.Context, in *EditReq, out *emptypb.Empty) error
		EditPwd(ctx context.Context, in *EditPwdReq, out *emptypb.Empty) error
		Info(ctx context.Context, in *UIDReq, out *Userinfo) error
		Logout(ctx context.Context, in *UIDReq, out *emptypb.Empty) error
		Online(ctx context.Context, in *OnlineReq, out *OnlineReply) error
		Offline(ctx context.Context, in *OfflineReq, out *emptypb.Empty) error
		ServerID(ctx context.Context, in *UIDReq, out *ServerIDReply) error
		BatchServersIDs(ctx context.Context, in *UIDsReq, out *ServerIDsReply) error
	}
	type User struct {
		user
	}
	h := &userHandler{hdlr}
	return s.Handle(s.NewHandler(&User{h}, opts...))
}

type userHandler struct {
	UserHandler
}

func (h *userHandler) Register(ctx context.Context, in *RegisterReq, out *RegisterReply) error {
	return h.UserHandler.Register(ctx, in, out)
}

func (h *userHandler) UsernameLogin(ctx context.Context, in *UsernameReq, out *UserToken) error {
	return h.UserHandler.UsernameLogin(ctx, in, out)
}

func (h *userHandler) PhoneLogin(ctx context.Context, in *PhoneReq, out *UserToken) error {
	return h.UserHandler.PhoneLogin(ctx, in, out)
}

func (h *userHandler) Edit(ctx context.Context, in *EditReq, out *emptypb.Empty) error {
	return h.UserHandler.Edit(ctx, in, out)
}

func (h *userHandler) EditPwd(ctx context.Context, in *EditPwdReq, out *emptypb.Empty) error {
	return h.UserHandler.EditPwd(ctx, in, out)
}

func (h *userHandler) Info(ctx context.Context, in *UIDReq, out *Userinfo) error {
	return h.UserHandler.Info(ctx, in, out)
}

func (h *userHandler) Logout(ctx context.Context, in *UIDReq, out *emptypb.Empty) error {
	return h.UserHandler.Logout(ctx, in, out)
}

func (h *userHandler) Online(ctx context.Context, in *OnlineReq, out *OnlineReply) error {
	return h.UserHandler.Online(ctx, in, out)
}

func (h *userHandler) Offline(ctx context.Context, in *OfflineReq, out *emptypb.Empty) error {
	return h.UserHandler.Offline(ctx, in, out)
}

func (h *userHandler) ServerID(ctx context.Context, in *UIDReq, out *ServerIDReply) error {
	return h.UserHandler.ServerID(ctx, in, out)
}

func (h *userHandler) BatchServersIDs(ctx context.Context, in *UIDsReq, out *ServerIDsReply) error {
	return h.UserHandler.BatchServersIDs(ctx, in, out)
}
