package handler

import (
	"context"

	"github.com/spf13/cast"
	"google.golang.org/protobuf/types/known/emptypb"

	"common/errno"
	pb "common/proto/member"
	"common/util"
	"member/service"
	"pkg/utils"
)

//Auth 会员服身份验证
func Auth(method string) bool {
	switch method {
	case "Member.Register", "Member.Login", "Member.PhoneLogin", "Member.SendCode":
		//这些路由不需要身份验证
		return false
	}
	return true
}

//Member 会员服务处理器
type Member struct {
	srv service.IService
}

//New 实例化会员服务处理器
func New(srv service.IService) *Member {
	return &Member{
		srv: srv,
	}
}

//Register 注册
func (m *Member) Register(ctx context.Context, req *pb.RegisterReq, empty *emptypb.Empty) error {
	is := utils.ValidateMobile(req.Phone)
	phone := cast.ToInt64(req.Phone)
	if !is || phone == 0 {
		return errno.ErrMemberPhoneValid
	}
	// 验证验证码是否正确
	err := m.srv.CheckVCode(ctx, phone, req.Code)
	if err != nil {
		return err
	}
	_, err = m.srv.MemberRegister(ctx, req.Username, req.Password, phone)
	if err != nil {
		return errno.MemberReplyErr(err)
	}
	return nil
}

//Login 登录
func (m *Member) Login(ctx context.Context, req *pb.LoginReq, reply *pb.MemberTokenReply) error {
	data, err := m.srv.MemberUsernameLogin(ctx, req.Username, req.Password)
	if err != nil {
		return errno.MemberReplyErr(err)
	}
	reply.Data = data
	return nil
}

//PhoneLogin 手机号登录
func (m *Member) PhoneLogin(ctx context.Context, req *pb.PhoneLoginReq, reply *pb.MemberTokenReply) error {
	is := utils.ValidateMobile(req.Phone)
	phone := cast.ToInt64(req.Phone)
	if !is || phone == 0 {
		return errno.ErrMemberPhoneValid
	}
	// 验证验证码是否正确
	err := m.srv.CheckVCode(ctx, phone, req.Code)
	if err != nil {
		return err
	}
	data, err := m.srv.MemberPhoneLogin(ctx, phone)
	if err != nil {
		return errno.MemberReplyErr(err)
	}
	reply.Data = data
	return nil
}

//MemberEdit 修改会员信息
func (m *Member) MemberEdit(ctx context.Context, req *pb.MemberEditReq, empty *emptypb.Empty) error {
	userMap := make(map[string]interface{})
	if req.Avatar != "" {
		userMap["avatar"] = req.Avatar
	}
	if req.Nickname != "" {
		userMap["nickname"] = req.Nickname
	}
	if req.Sign != "" {
		userMap["sign"] = req.Sign
	}
	if len(userMap) == 0 { //数据为空
		return errno.ErrParamsCheckInvalid
	}

	err := m.srv.MemberEdit(ctx, util.GetUserID(ctx), userMap)
	if err != nil {
		return errno.MemberReplyErr(err)
	}
	return nil
}

//MemberPwdEdit 修改密码
func (m *Member) MemberPwdEdit(ctx context.Context, req *pb.PwdEditReq, empty *emptypb.Empty) error {
	err := m.srv.UserEditPwd(ctx, util.GetUserID(ctx), req.OldPassword, req.Password)
	if err != nil {
		return err
	}
	return nil
}

//MemberProfile 获取会员信息
func (m *Member) MemberProfile(ctx context.Context, empty *emptypb.Empty, reply *pb.MemberInfoReply) error {
	member, err := m.srv.MemberInfo(ctx, util.GetUserID(ctx))
	if err != nil {
		return errno.MemberReplyErr(err)
	}
	reply.Data = member
	return nil
}

//Logout 登出
func (m *Member) Logout(ctx context.Context, empty *emptypb.Empty, empty2 *emptypb.Empty) error {
	err := m.srv.UserLogout(ctx, util.GetUserID(ctx))
	if err != nil {
		return err
	}
	return nil
}

//AddressAdd 添加收货地址
func (m *Member) AddressAdd(ctx context.Context, req *pb.AddressAddReq, reply *pb.AddressIDReply) error {
	id, err := m.srv.MemberAddressAdd(ctx, util.GetUserID(ctx), req.Name, req.Phone, req.Province,
		req.City, req.County, req.Detail, int(req.AreaCode), int8(req.IsDefault))
	if err != nil {
		return errno.MemberReplyErr(err)
	}
	reply.Data = id
	return nil
}

//AddressEdit 修改收货地址
func (m *Member) AddressEdit(ctx context.Context, req *pb.Address, empty *emptypb.Empty) error {
	err := m.srv.MemberAddressEdit(ctx, req.Id, util.GetUserID(ctx), req.Name, req.Phone, req.Province,
		req.City, req.County, req.Detail, int(req.AreaCode), int8(req.IsDefault))
	if err != nil {
		return errno.MemberReplyErr(err)
	}
	return nil
}

//GetAddressList 收货地址列表
func (m *Member) GetAddressList(ctx context.Context, empty *emptypb.Empty, reply *pb.AddressReply) error {
	list, err := m.srv.MemberAddressList(ctx, util.GetUserID(ctx))
	if err != nil {
		return errno.MemberReplyErr(err)
	}
	reply.Data = list
	return nil
}

//AddressDel 删除收货地址
func (m *Member) AddressDel(ctx context.Context, req *pb.AddressIDReq, empty *emptypb.Empty) error {
	err := m.srv.MemberAddressDel(ctx, req.Id, util.GetUserID(ctx))
	if err != nil {
		return errno.MemberReplyErr(err)
	}
	return nil
}

//GetAddressInfo 获取收货地址信息
func (m *Member) GetAddressInfo(ctx context.Context, req *pb.AddressInfoReq, reply *pb.AddressInfoInternal) error {
	info, err := m.srv.GetMemberAddressInfo(ctx, req.AddressId, req.UserId)
	if err != nil {
		return errno.MemberReplyErr(err)
	}
	reply.Info = info
	return nil
}

//SendCode 发送短信验证码
func (m *Member) SendCode(ctx context.Context, req *pb.PhoneReq, reply *pb.CodeReply) error {
	code, err := m.srv.SendCode(ctx, req.Phone)
	if err != nil {
		return err
	}
	reply.Data = code
	return nil
}
