package service

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"chat-micro/app/constvar"
	"chat-micro/app/logic/idl"
	"chat-micro/app/logic/model"
	"chat-micro/pkg/app"
	"chat-micro/pkg/redis"
)

//IUser 用户服务接口
type IUser interface {
	// 用户注册
	UserRegister(ctx context.Context, username, password string, phone int64) (uint32, error)
	// 用户名登录
	UsernameLogin(ctx context.Context, username, password string) (*model.User, string, error)
	// 手机号登录
	UserPhoneLogin(ctx context.Context, phone int64) (*model.User, string, error)
	// 修改密码
	UserEditPwd(ctx context.Context, id uint32, password string) error
	// 修改用户信息
	UserEdit(ctx context.Context, id uint32, userMap map[string]interface{}) error
	// 获取用户详情
	UserInfoByID(ctx context.Context, id uint32) (*model.User, error)
	// 用户登出
	UserLogout(ctx context.Context, id uint32) error
	// 搜索用户
	UserSearch(ctx context.Context, keyword string) ([]*model.User, error)
	// 标签用户列表
	UserTagList(ctx context.Context, uid uint32) ([]*model.UserTag, error)
}

// UserRegister 注册用户
func (s *Service) UserRegister(ctx context.Context, username, password string, phone int64) (id uint32, err error) {
	exist, err := s.repo.UserExist(ctx, username, phone)
	if err != nil {
		return 0, errors.Wrapf(err, "[service.user] user exist")
	}
	if exist {
		return 0, ErrUserExisted
	}
	u := &model.UserModel{
		Username: username,
		Password: password,
		Phone:    phone,
		Status:   model.UserStatusNormal,
	}
	id, err = s.repo.UserCreate(ctx, u)
	if err != nil {
		return 0, errors.Wrapf(err, "[service.user] create user")
	}
	return id, nil
}

// UsernameLogin 用户名密码登录
func (s *Service) UsernameLogin(ctx context.Context, username, password string) (user *model.User, token string, err error) {
	// 如果是已经注册用户，则通过用户名获取用户信息
	userModel, err := s.repo.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, "", errors.Wrapf(err, "[service.user] err from db by username: %s", username)
	}

	// 否则新建用户信息, 并取得用户信息
	if userModel.ID == 0 {
		return nil, "", ErrUserNotFound
	}

	if userModel.Status != model.UserStatusNormal {
		return nil, "", ErrUserFrozen
	}

	// Compare the login password with the user password.
	err = userModel.Compare(password)
	if err != nil {
		return nil, "", ErrUserNotMatch
	}

	return s.transUserToken(ctx, userModel)
}

// UserPhoneLogin 手机登录
func (s *Service) UserPhoneLogin(ctx context.Context, phone int64) (*model.User, string, error) {
	// 如果是已经注册用户，则通过手机号获取用户信息
	userModel, err := s.repo.GetUserByPhone(ctx, phone)
	if err != nil {
		return nil, "", errors.Wrapf(err, "[service.user] err from db by phone: %d", phone)
	}

	// 否则新建用户信息, 并取得用户信息
	if userModel.ID == 0 {
		return nil, "", ErrUserNotFound
	}

	if userModel.Status != model.UserStatusNormal {
		return nil, "", ErrUserFrozen
	}

	return s.transUserToken(ctx, userModel)
}

// UserEdit update user info
func (s *Service) UserEdit(ctx context.Context, id uint32, userMap map[string]interface{}) error {
	err := s.repo.UserUpdate(ctx, id, userMap)

	if err != nil {
		return errors.Wrapf(err, "[service.user] update user by id: %d", id)
	}
	return nil
}

//UserEditPwd 修改用户密码
func (s *Service) UserEditPwd(ctx context.Context, id uint32, password string) error {
	user, err := s.userinfo(ctx, id)
	if err != nil {
		return err
	}
	user.Password = password
	if err = s.repo.UserUpdatePwd(ctx, user); err != nil {
		return errors.Wrapf(err, "[service.user] update user pwd by id:%v", id)
	}
	return nil
}

// UserInfoByID 获取用户信息
func (s *Service) UserInfoByID(ctx context.Context, id uint32) (*model.User, error) {
	user, err := s.userinfo(ctx, id)
	if err != nil {
		return nil, err
	}
	return idl.TransferUser(user), nil
}

// UserLogout 用户登出
func (s *Service) UserLogout(ctx context.Context, id uint32) error {
	return redis.Client.Del(ctx, constvar.BuildUserTokenKey(id)).Err()
}

// UserSearch 搜索用户
func (s *Service) UserSearch(ctx context.Context, keyword string) ([]*model.User, error) {
	list, err := s.repo.GetUsersByKeyword(ctx, keyword)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.user] search user keyword: %s", keyword)
	}
	return idl.TransferUsers(list), nil
}

// UserTagList 用户标签列表
func (s *Service) UserTagList(ctx context.Context, uid uint32) (list []*model.UserTag, err error) {
	list, err = s.repo.GetTagsByUserID(ctx, uid)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.user] UserTagList id:%d", uid)
	}
	return
}

// transUserToken 转换输出用户登录后信息
func (s *Service) transUserToken(ctx context.Context, user *model.UserModel) (*model.User, string, error) {
	// 签发签名 Sign the json web token.
	payload := map[string]interface{}{"user_id": user.ID}
	token, err := app.Sign(ctx, payload, s.opts.jwtSecret, s.opts.jwtTimeout)
	if err != nil {
		return nil, "", errors.Wrapf(err, "[service.user] gen token sign err")
	}
	// 设置当前令牌，用户单点登录判断
	redis.Client.Set(ctx, constvar.BuildUserTokenKey(user.ID), token, time.Hour*24)
	return idl.TransferUser(user), token, nil
}

// userinfo 获取用户模型
func (s *Service) userinfo(ctx context.Context, id uint32) (*model.UserModel, error) {
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.user] from repo by id: %d", id)
	}
	if user.ID == 0 {
		return nil, ErrUserNotFound
	}
	if user.Status != model.UserStatusNormal {
		return nil, ErrUserFrozen
	}
	return user, nil
}
