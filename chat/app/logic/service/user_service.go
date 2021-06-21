package service

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"chat/app/logic/idl"
	"chat/app/logic/model"
	"chat/pkg/app"
	"chat/pkg/redis"
)

//IUser 用户服务接口
type IUser interface {
	// 用户注册
	UserRegister(ctx context.Context, username, password string, phone int64) error
	// 用户名登录
	UsernameLogin(ctx context.Context, username, password string) (userToken *model.UserToken, err error)
	// 手机号登录
	UserPhoneLogin(ctx context.Context, phone int64) (userToken *model.UserToken, err error)
	// 修改用户信息
	UserEdit(ctx context.Context, id uint32, userMap map[string]interface{}) error
	// 搜索用户
	UserSearch(ctx context.Context, keyword string) (users []*model.UserEs, err error)
	// 获取用户详情
	UserInfoByID(ctx context.Context, id uint32) (*model.UserInfo, error)
	// 标签用户列表
	UserTagList(ctx context.Context, UserID uint32) (list []*model.UserTag, err error)
	// 检查用户是否正常
	UserCheck(ctx context.Context, id uint32) (bool, error)
	// 用户登出
	UserLogout(ctx context.Context, UserID uint32) error
}

// UserRegister 注册用户
func (s *Service) UserRegister(ctx context.Context, username, password string, phone int64) error {
	u := &model.UserModel{
		Username: username,
		Password: password,
		Phone:    phone,
	}
	is := s.repo.UserCheckExist(ctx, username, phone)
	if is {
		return ErrUserKeyExisted
	}
	_, err := s.repo.UserCreate(ctx, u)
	if err != nil {
		return errors.Wrapf(err, "[service.user] create user")
	}
	//异步同步入es
	s.ec.PushUser(u)
	return nil
}

// UsernameLogin 用户名密码登录
func (s *Service) UsernameLogin(ctx context.Context, username, password string) (userToken *model.UserToken, err error) {
	// 如果是已经注册用户，则通过用户名获取用户信息
	userModel, err := s.repo.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.user] err from db by username: %s", username)
	}

	// Compare the login password with the user password.
	err = userModel.Compare(password)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.user] password compare :%v", password)
	}

	return s.transUserToken(ctx, userModel)
}

// UserPhoneLogin 邮箱登录
func (s *Service) UserPhoneLogin(ctx context.Context, phone int64) (userToken *model.UserToken, err error) {
	// 如果是已经注册用户，则通过手机号获取用户信息
	userModel, err := s.repo.GetUserByPhone(ctx, phone)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.user] err from db by phone: %d", phone)
	}

	// 否则新建用户信息, 并取得用户信息
	if userModel.ID == 0 {
		return nil, errors.Wrapf(err, "[service.user] not found phone:%v", phone)
	}

	return s.transUserToken(ctx, userModel)
}

// UserEdit update user info
func (s *Service) UserEdit(ctx context.Context, id uint32, userMap map[string]interface{}) error {
	err := s.repo.UserUpdate(ctx, id, userMap)

	if err != nil {
		return errors.Wrapf(err, "[service.user] update user by id: %d", id)
	}
	if u, ok := userMap["nickname"]; ok {
		err = s.ec.UserUpdate(ctx, id, map[string]interface{}{"nickname": u})
		if err != nil {
			return errors.Wrapf(err, "[service.user] save es by id:%d", id)
		}
	}

	return nil
}

// UserSearch 搜索用户
func (s *Service) UserSearch(ctx context.Context, keyword string) (users []*model.UserEs, err error) {
	users, err = s.ec.UserSearch(ctx, keyword)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.user] search user keyword: %s", keyword)
	}
	if len(users) == 0 {
		return make([]*model.UserEs, 0), nil
	}
	return users, nil
}

// UserInfoByID 获取用户信息
func (s *Service) UserInfoByID(ctx context.Context, id uint32) (*model.UserInfo, error) {
	userModel, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.user] err from db by id: %d", id)
	}
	return idl.TransferUser(userModel), nil
}

// UserTagList 用户标签列表
func (s *Service) UserTagList(ctx context.Context, UserID uint32) (list []*model.UserTag, err error) {
	list, err = s.repo.GetTagsByUserID(ctx, UserID)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.user] UserTagList id:%d", UserID)
	}
	return
}

// UserCheck 检查用户是否正常
func (s *Service) UserCheck(ctx context.Context, id uint32) (bool, error) {
	u, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return false, errors.Wrapf(err, "[service.user] check id: %d", id)
	}
	if u.Status == model.StatusNormal {
		return true, nil
	}
	return false, nil
}

// UserLogout 用户登出
func (s *Service) UserLogout(ctx context.Context, UserID uint32) error {
	pipe := redis.Client.Pipeline()
	pipe.Del(ctx, s.getUserKey(UserID))
	pipe.Del(ctx, s.getOnlineKey(UserID))
	_, err := pipe.Exec(ctx)
	return err
}

// transUserToken 转换输出用户登录后信息
func (s *Service) transUserToken(ctx context.Context, user *model.UserModel) (*model.UserToken, error) {
	// 签发签名 Sign the json web token.
	payload := map[string]interface{}{"user_id": user.ID}
	tokenStr, err := app.Sign(ctx, payload, s.c.App.JwtSecret, s.c.App.JwtTimeout)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.user] gen token sign err")
	}
	//踢出上一次登录信息
	err = s.UserTickOut(ctx, user.ID)
	if err != nil {
		return nil, ErrUserLogin
	}
	transInput := &idl.TransferUserInput{
		User:  user,
		Token: tokenStr,
	}
	// 设置当前令牌，用户单点登录判断
	redis.Client.Set(ctx, s.getUserKey(user.ID), tokenStr, time.Hour*24)
	return idl.TransferAuth(transInput), nil
}
