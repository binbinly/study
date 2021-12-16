package service

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"

	pb "common/proto/center"
)

// ICenter 中心服接口
type ICenter interface {
	UserRegister(ctx context.Context, username, password string, phone int64) (int64, error)
	UsernameLogin(ctx context.Context, username, password string) (*pb.UserToken, error)
	UserPhoneLogin(ctx context.Context, phone int64) (*pb.UserToken, error)
	UserEditPwd(ctx context.Context, id int64, oldPassword, password string) error
	UserEdit(ctx context.Context, id int64, userMap map[string]interface{}) error
	UserLogout(ctx context.Context, id int64) error
}

// UserRegister 注册用户
func (s *Service) UserRegister(ctx context.Context, username, password string, phone int64) (int64, error) {
	req := &pb.RegisterReq{
		Username: username,
		Password: password,
		Phone:    phone,
	}
	reply, err := s.centerService.Register(ctx, req)
	if err != nil {
		return 0, err
	}
	return reply.Id, nil
}

// UsernameLogin 用户名密码登录
func (s *Service) UsernameLogin(ctx context.Context, username, password string) (user *pb.UserToken, err error) {
	user, err = s.centerService.UsernameLogin(ctx, &pb.UsernameReq{
		Username: username,
		Password: password,
	})
	if err != nil {
		return nil, err
	}
	return
}

// UserPhoneLogin 邮箱登录
func (s *Service) UserPhoneLogin(ctx context.Context, phone int64) (user *pb.UserToken, err error) {
	user, err = s.centerService.PhoneLogin(ctx, &pb.PhoneReq{
		Phone: phone,
	})
	if err != nil {
		return nil, err
	}
	return
}

// UserEditPwd 修改用户密码
func (s *Service) UserEditPwd(ctx context.Context, id int64, oldPassword, password string) error {
	_, err := s.centerService.EditPwd(ctx, &pb.EditPwdReq{
		Id:     id,
		OldPwd: oldPassword,
		Pwd:    password,
	})
	if err != nil {
		return err
	}
	return nil
}

// UserEdit 修改用户信息
func (s *Service) UserEdit(ctx context.Context, id int64, userMap map[string]interface{}) error {
	bytes, err := json.Marshal(userMap)
	if err != nil {
		return errors.Wrapf(err, "[service.center] json marshal")
	}

	_, err = s.centerService.Edit(ctx, &pb.EditReq{
		Id:      id,
		Content: bytes,
	})
	if err != nil {
		return err
	}
	return nil
}

// UserLogout 用户登出
func (s *Service) UserLogout(ctx context.Context, id int64) error {
	_, err := s.centerService.Logout(ctx, &pb.UIDReq{Id: id})
	if err != nil {
		return err
	}
	return nil
}
