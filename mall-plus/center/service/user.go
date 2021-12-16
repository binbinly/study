package service

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"center/idl"
	"center/model"
	"common/constvar"
	"common/errno"
	pb "common/proto/center"
	"pkg/app"
	"pkg/redis"
)

//IUser 用户服务接口
type IUser interface {
	// 用户注册
	UserRegister(ctx context.Context, username, password string, phone int64) (id int64, err error)
	// 用户名登录
	UsernameLogin(ctx context.Context, username, password string) (*pb.Userinfo, string, error)
	// 手机号登录
	UserPhoneLogin(ctx context.Context, phone int64) (*pb.Userinfo, string, error)
	// 修改密码
	UserEditPwd(ctx context.Context, id int64, oldPassword, password string) error
	// 修改用户信息
	UserEdit(ctx context.Context, id int64, userMap map[string]interface{}) error
	// 获取用户详情
	UserInfoByID(ctx context.Context, id int64) (*pb.Userinfo, error)
	// 用户登出
	UserLogout(ctx context.Context, id int64) error
}

// UserRegister 注册用户
func (c *Center) UserRegister(ctx context.Context, username, password string, phone int64) (id int64, err error) {
	u := &model.UserModel{
		Username: username,
		Password: password,
		Phone:    phone,
		Status:   model.StatusNormal,
	}
	exist, err := c.repo.UserExist(ctx, username, phone)
	if err != nil {
		return 0, errors.Wrapf(err, "[center.user] user exist")
	}
	if exist {
		return 0, errno.ErrUserExisted
	}
	id, err = c.repo.UserCreate(ctx, u)
	if err != nil {
		return 0, errors.Wrapf(err, "[center.user] create user")
	}
	return id, nil
}

// UsernameLogin 用户名密码登录
func (c *Center) UsernameLogin(ctx context.Context, username, password string) (*pb.Userinfo, string, error) {
	// 如果是已经注册用户，则通过用户名获取用户信息
	user, err := c.repo.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, "", errors.Wrapf(err, "[center.user] err from db by username: %s", username)
	}

	// 否则新建用户信息, 并取得用户信息
	if user.ID == 0 {
		return nil, "", errno.ErrUserNotFound
	}

	if user.Status != model.StatusNormal {
		return nil, "", errno.ErrUserFrozen
	}

	// Compare the login password with the user password.
	if err = user.Compare(password); err != nil {
		return nil, "", errno.ErrUserNotMatch
	}

	return c.transUserToken(ctx, user)
}

// UserPhoneLogin 邮箱登录
func (c *Center) UserPhoneLogin(ctx context.Context, phone int64) (*pb.Userinfo, string, error) {
	// 如果是已经注册用户，则通过手机号获取用户信息
	user, err := c.repo.GetUserByPhone(ctx, phone)
	if err != nil {
		return nil, "", errors.Wrapf(err, "[center.user] err from db by phone: %d", phone)
	}

	// 否则新建用户信息, 并取得用户信息
	if user.ID == 0 {
		return nil, "", errno.ErrUserNotFound
	}

	if user.Status != model.StatusNormal {
		return nil, "", errno.ErrUserFrozen
	}

	return c.transUserToken(ctx, user)
}

// UserEdit update user info
func (c *Center) UserEdit(ctx context.Context, id int64, userMap map[string]interface{}) error {
	if err := c.repo.UserUpdate(ctx, id, userMap); err != nil {
		return errors.Wrapf(err, "[center.user] update user by id: %d", id)
	}
	return nil
}

//UserEditPwd 修改用户密码
func (c *Center) UserEditPwd(ctx context.Context, id int64, oldPassword, password string) error {
	user, err := c.userinfo(ctx, id)
	if err != nil {
		return err
	}

	if err = user.Compare(oldPassword); err != nil {
		return errno.ErrUserPwd
	}

	user.Password = password
	if err = c.repo.UserUpdatePwd(ctx, user); err != nil {
		return errors.Wrapf(err, "[center.user] update user pwd by id:%v", id)
	}
	return nil
}

// UserInfoByID 获取用户信息
func (c *Center) UserInfoByID(ctx context.Context, id int64) (*pb.Userinfo, error) {
	user, err := c.userinfo(ctx, id)
	if err != nil {
		return nil, err
	}
	return idl.TransferUser(user), nil
}

// UserLogout 用户登出
func (c *Center) UserLogout(ctx context.Context, id int64) error {
	pipe := redis.Client.Pipeline()
	pipe.Del(ctx, constvar.BuildUserTokenKey(id))
	pipe.Del(ctx, constvar.BuildOnlineKey(id))
	_, err := pipe.Exec(ctx)
	return err
}

// transUserToken 转换输出用户登录后信息
func (c *Center) transUserToken(ctx context.Context, user *model.UserModel) (*pb.Userinfo, string, error) {
	// 签发签名 Sign the json web token.
	payload := map[string]interface{}{"user_id": user.ID}
	tokenStr, err := app.Sign(ctx, payload, c.c.App.JwtSecret, c.c.App.JwtTimeout)
	if err != nil {
		return nil, "", errors.Wrapf(err, "[center.user] gen token sign err")
	}

	//踢出上一次登录信息
	if err = c.UserTickOut(ctx, user.ID); err != nil {
		return nil, "", errors.Wrapf(err, "[center.user] tickout")
	}
	// 设置当前令牌，用户单点登录判断
	redis.Client.Set(ctx, constvar.BuildUserTokenKey(user.ID), tokenStr, time.Hour*24)
	return idl.TransferUser(user), tokenStr, nil
}

// userinfo 获取用户模型
func (c *Center) userinfo(ctx context.Context, id int64) (*model.UserModel, error) {
	user, err := c.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, errors.Wrapf(err, "[center.user] from repo by id: %d", id)
	}
	if user.ID == 0 {
		return nil, errno.ErrUserNotFound
	}
	if user.Status != model.StatusNormal {
		return nil, errno.ErrUserFrozen
	}
	return user, nil
}
