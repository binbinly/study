// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: member/member.proto

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

// Api Endpoints for Member service

func NewMemberEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Member service

type MemberService interface {
	/// 注册
	Register(ctx context.Context, in *RegisterReq, opts ...client.CallOption) (*emptypb.Empty, error)
	/// 用户名密码登录
	Login(ctx context.Context, in *LoginReq, opts ...client.CallOption) (*MemberTokenReply, error)
	/// 手机号登录
	PhoneLogin(ctx context.Context, in *PhoneLoginReq, opts ...client.CallOption) (*MemberTokenReply, error)
	/// 修改会员信息
	MemberEdit(ctx context.Context, in *MemberEditReq, opts ...client.CallOption) (*emptypb.Empty, error)
	/// 修改密码
	MemberPwdEdit(ctx context.Context, in *PwdEditReq, opts ...client.CallOption) (*emptypb.Empty, error)
	/// 获取会员信息
	MemberProfile(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*MemberInfoReply, error)
	/// 登出
	Logout(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*emptypb.Empty, error)
	/// 添加收货地址
	AddressAdd(ctx context.Context, in *AddressAddReq, opts ...client.CallOption) (*AddressIDReply, error)
	/// 修改收货地址
	AddressEdit(ctx context.Context, in *Address, opts ...client.CallOption) (*emptypb.Empty, error)
	/// 收货地址列表
	GetAddressList(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*AddressReply, error)
	/// 删除收货地址
	AddressDel(ctx context.Context, in *AddressIDReq, opts ...client.CallOption) (*emptypb.Empty, error)
	/// 发送短信验证码
	SendCode(ctx context.Context, in *PhoneReq, opts ...client.CallOption) (*CodeReply, error)
	/// ---- 以下内部调用 ----
	/// 获取收货地址信息
	GetAddressInfo(ctx context.Context, in *AddressInfoReq, opts ...client.CallOption) (*AddressInfoInternal, error)
}

type memberService struct {
	c    client.Client
	name string
}

func NewMemberService(name string, c client.Client) MemberService {
	return &memberService{
		c:    c,
		name: name,
	}
}

func (c *memberService) Register(ctx context.Context, in *RegisterReq, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "Member.Register", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberService) Login(ctx context.Context, in *LoginReq, opts ...client.CallOption) (*MemberTokenReply, error) {
	req := c.c.NewRequest(c.name, "Member.Login", in)
	out := new(MemberTokenReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberService) PhoneLogin(ctx context.Context, in *PhoneLoginReq, opts ...client.CallOption) (*MemberTokenReply, error) {
	req := c.c.NewRequest(c.name, "Member.PhoneLogin", in)
	out := new(MemberTokenReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberService) MemberEdit(ctx context.Context, in *MemberEditReq, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "Member.MemberEdit", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberService) MemberPwdEdit(ctx context.Context, in *PwdEditReq, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "Member.MemberPwdEdit", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberService) MemberProfile(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*MemberInfoReply, error) {
	req := c.c.NewRequest(c.name, "Member.MemberProfile", in)
	out := new(MemberInfoReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberService) Logout(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "Member.Logout", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberService) AddressAdd(ctx context.Context, in *AddressAddReq, opts ...client.CallOption) (*AddressIDReply, error) {
	req := c.c.NewRequest(c.name, "Member.AddressAdd", in)
	out := new(AddressIDReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberService) AddressEdit(ctx context.Context, in *Address, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "Member.AddressEdit", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberService) GetAddressList(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*AddressReply, error) {
	req := c.c.NewRequest(c.name, "Member.GetAddressList", in)
	out := new(AddressReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberService) AddressDel(ctx context.Context, in *AddressIDReq, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "Member.AddressDel", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberService) SendCode(ctx context.Context, in *PhoneReq, opts ...client.CallOption) (*CodeReply, error) {
	req := c.c.NewRequest(c.name, "Member.SendCode", in)
	out := new(CodeReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberService) GetAddressInfo(ctx context.Context, in *AddressInfoReq, opts ...client.CallOption) (*AddressInfoInternal, error) {
	req := c.c.NewRequest(c.name, "Member.GetAddressInfo", in)
	out := new(AddressInfoInternal)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Member service

type MemberHandler interface {
	/// 注册
	Register(context.Context, *RegisterReq, *emptypb.Empty) error
	/// 用户名密码登录
	Login(context.Context, *LoginReq, *MemberTokenReply) error
	/// 手机号登录
	PhoneLogin(context.Context, *PhoneLoginReq, *MemberTokenReply) error
	/// 修改会员信息
	MemberEdit(context.Context, *MemberEditReq, *emptypb.Empty) error
	/// 修改密码
	MemberPwdEdit(context.Context, *PwdEditReq, *emptypb.Empty) error
	/// 获取会员信息
	MemberProfile(context.Context, *emptypb.Empty, *MemberInfoReply) error
	/// 登出
	Logout(context.Context, *emptypb.Empty, *emptypb.Empty) error
	/// 添加收货地址
	AddressAdd(context.Context, *AddressAddReq, *AddressIDReply) error
	/// 修改收货地址
	AddressEdit(context.Context, *Address, *emptypb.Empty) error
	/// 收货地址列表
	GetAddressList(context.Context, *emptypb.Empty, *AddressReply) error
	/// 删除收货地址
	AddressDel(context.Context, *AddressIDReq, *emptypb.Empty) error
	/// 发送短信验证码
	SendCode(context.Context, *PhoneReq, *CodeReply) error
	/// ---- 以下内部调用 ----
	/// 获取收货地址信息
	GetAddressInfo(context.Context, *AddressInfoReq, *AddressInfoInternal) error
}

func RegisterMemberHandler(s server.Server, hdlr MemberHandler, opts ...server.HandlerOption) error {
	type member interface {
		Register(ctx context.Context, in *RegisterReq, out *emptypb.Empty) error
		Login(ctx context.Context, in *LoginReq, out *MemberTokenReply) error
		PhoneLogin(ctx context.Context, in *PhoneLoginReq, out *MemberTokenReply) error
		MemberEdit(ctx context.Context, in *MemberEditReq, out *emptypb.Empty) error
		MemberPwdEdit(ctx context.Context, in *PwdEditReq, out *emptypb.Empty) error
		MemberProfile(ctx context.Context, in *emptypb.Empty, out *MemberInfoReply) error
		Logout(ctx context.Context, in *emptypb.Empty, out *emptypb.Empty) error
		AddressAdd(ctx context.Context, in *AddressAddReq, out *AddressIDReply) error
		AddressEdit(ctx context.Context, in *Address, out *emptypb.Empty) error
		GetAddressList(ctx context.Context, in *emptypb.Empty, out *AddressReply) error
		AddressDel(ctx context.Context, in *AddressIDReq, out *emptypb.Empty) error
		SendCode(ctx context.Context, in *PhoneReq, out *CodeReply) error
		GetAddressInfo(ctx context.Context, in *AddressInfoReq, out *AddressInfoInternal) error
	}
	type Member struct {
		member
	}
	h := &memberHandler{hdlr}
	return s.Handle(s.NewHandler(&Member{h}, opts...))
}

type memberHandler struct {
	MemberHandler
}

func (h *memberHandler) Register(ctx context.Context, in *RegisterReq, out *emptypb.Empty) error {
	return h.MemberHandler.Register(ctx, in, out)
}

func (h *memberHandler) Login(ctx context.Context, in *LoginReq, out *MemberTokenReply) error {
	return h.MemberHandler.Login(ctx, in, out)
}

func (h *memberHandler) PhoneLogin(ctx context.Context, in *PhoneLoginReq, out *MemberTokenReply) error {
	return h.MemberHandler.PhoneLogin(ctx, in, out)
}

func (h *memberHandler) MemberEdit(ctx context.Context, in *MemberEditReq, out *emptypb.Empty) error {
	return h.MemberHandler.MemberEdit(ctx, in, out)
}

func (h *memberHandler) MemberPwdEdit(ctx context.Context, in *PwdEditReq, out *emptypb.Empty) error {
	return h.MemberHandler.MemberPwdEdit(ctx, in, out)
}

func (h *memberHandler) MemberProfile(ctx context.Context, in *emptypb.Empty, out *MemberInfoReply) error {
	return h.MemberHandler.MemberProfile(ctx, in, out)
}

func (h *memberHandler) Logout(ctx context.Context, in *emptypb.Empty, out *emptypb.Empty) error {
	return h.MemberHandler.Logout(ctx, in, out)
}

func (h *memberHandler) AddressAdd(ctx context.Context, in *AddressAddReq, out *AddressIDReply) error {
	return h.MemberHandler.AddressAdd(ctx, in, out)
}

func (h *memberHandler) AddressEdit(ctx context.Context, in *Address, out *emptypb.Empty) error {
	return h.MemberHandler.AddressEdit(ctx, in, out)
}

func (h *memberHandler) GetAddressList(ctx context.Context, in *emptypb.Empty, out *AddressReply) error {
	return h.MemberHandler.GetAddressList(ctx, in, out)
}

func (h *memberHandler) AddressDel(ctx context.Context, in *AddressIDReq, out *emptypb.Empty) error {
	return h.MemberHandler.AddressDel(ctx, in, out)
}

func (h *memberHandler) SendCode(ctx context.Context, in *PhoneReq, out *CodeReply) error {
	return h.MemberHandler.SendCode(ctx, in, out)
}

func (h *memberHandler) GetAddressInfo(ctx context.Context, in *AddressInfoReq, out *AddressInfoInternal) error {
	return h.MemberHandler.GetAddressInfo(ctx, in, out)
}