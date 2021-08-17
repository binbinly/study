package v1

import (
	service "chat/app/chat"
	"chat/app/chat/server"
	"chat/proto"
	"chat/proto/chat"
)

//Register 用户注册
func Register(c *server.Context) (*chat.ReceiveReply, error) {
	data := &proto.ReqRegister{}
	err := unmarshal(c, data)
	if err != nil {
		return &chat.ReceiveReply{Code: chat.ReceiveReply_ErrJsonUnmarshal}, nil
	}
	return response(nil, service.Svc.UserRegister(c, data.Username, data.Password, data.Phone))
}

//Login 用户登录
func Login(c *server.Context) (*chat.ReceiveReply, error) {
	data := &proto.ReqLogin{}
	err := unmarshal(c, data)
	if err != nil {
		return &chat.ReceiveReply{Code: chat.ReceiveReply_ErrJsonUnmarshal}, nil
	}
	return response(service.Svc.UsernameLogin(c, data.Username, data.Password))
}

//LoginPhone 手机号验证码登录
func LoginPhone(c *server.Context) (*chat.ReceiveReply, error) {
	data := &proto.ReqPhoneLogin{}
	err := unmarshal(c, data)
	if err != nil {
		return &chat.ReceiveReply{Code: chat.ReceiveReply_ErrJsonUnmarshal}, nil
	}
	// 验证校验码
	if err = service.Svc.CheckVCode(c, data.Phone, data.VerifyCode); err != nil {
		return &chat.ReceiveReply{Code: chat.ReceiveReply_ErrVerifyCode}, nil
	}

	// 登录
	return response(service.Svc.UserPhoneLogin(c, data.Phone))
}

//Logout 登出
func Logout(c *server.Context) (*chat.ReceiveReply, error) {
	return response(nil, service.Svc.UserLogout(c, c.Req.GetUid()))
}

//UserProfile 用户信息
func UserProfile(c *server.Context) (*chat.ReceiveReply, error) {
	return response(service.Svc.GetUserByID(c, c.Req.GetUid()))
}

//Search 搜索
func Search(c *server.Context) (*chat.ReceiveReply, error) {
	return response(service.Svc.UserSearch(c, string(c.Req.Body)))
}

//UserTags 用户标签列表
func UserTags(c *server.Context) (*chat.ReceiveReply, error) {
	return response(service.Svc.UserTagList(c, c.Req.GetUid()))
}

//UserEdit 用户信息修改
func UserEdit(c *server.Context) (*chat.ReceiveReply, error) {
	data := make(map[string]interface{})
	err := unmarshal(c, &data)
	if err != nil {
		return &chat.ReceiveReply{Code: chat.ReceiveReply_ErrJsonUnmarshal}, nil
	}
	return response(nil, service.Svc.UserEdit(c, c.Req.GetUid(), data))
}

//UserReport 用户举报
func UserReport(c *server.Context) (*chat.ReceiveReply, error) {
	data := &proto.ReqReport{}
	err := unmarshal(c, data)
	if err != nil {
		return &chat.ReceiveReply{Code: chat.ReceiveReply_ErrJsonUnmarshal}, nil
	}
	return response(nil, service.Svc.ReportCreate(c, c.Req.GetUid(), data.UserID, data.Type, data.Category, data.Content))
}
