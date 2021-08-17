package chat

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"golang.org/x/sync/singleflight"

	"chat/app/center"
	"chat/app/center/server"
	"chat/app/chat/model"
	"chat/pkg/cache"
	"chat/pkg/log"
	"chat/pkg/redis"
	"chat/proto/base"
	pb "chat/proto/center"
)

// ICenter 中心服接口
type ICenter interface {
	UserRegister(ctx context.Context, username, password string, phone int64) error
	UsernameLogin(ctx context.Context, username, password string) (*pb.UserToken, error)
	UserPhoneLogin(ctx context.Context, phone int64) (*pb.UserToken, error)
	UserEditPwd(ctx context.Context, id uint32, password string) error
	UserEdit(ctx context.Context, id uint32, userMap map[string]interface{}) error
	GetUserByID(ctx context.Context, id uint32) (*base.UserInfo, error)
	UserLogout(ctx context.Context, id uint32) error
	SendSMS(ctx context.Context, phone int64) (string, error)
	CheckVCode(ctx context.Context, phone int64, code string) error
	ServerByUserID(ctx context.Context, id uint32) (string, error)
	ServersByUserIDs(ctx context.Context, ids []uint32) ([]string, error)
}

// UserRegister 注册用户
func (s *Service) UserRegister(ctx context.Context, username, password string, phone int64) error {
	req := &pb.RegisterReq{
		Username: username,
		Password: password,
		Phone:    phone,
	}
	reply, err := s.rpcClient.UserRegister(ctx, req)
	if err != nil {
		return server.HandleError(err)
	}
	//异步同步入es
	s.ec.PushUser(&model.UserEs{
		ID:       reply.Id,
		Username: username,
		Phone:    strconv.FormatInt(phone, 10),
	})
	return nil
}

// UsernameLogin 用户名密码登录
func (s *Service) UsernameLogin(ctx context.Context, username, password string) (user *pb.UserToken, err error) {
	user, err = s.rpcClient.UsernameLogin(ctx, &pb.UsernameReq{
		Username: username,
		Password: password,
	})
	if err != nil {
		return nil, server.HandleError(err)
	}

	return
}

// UserPhoneLogin 邮箱登录
func (s *Service) UserPhoneLogin(ctx context.Context, phone int64) (user *pb.UserToken, err error) {
	user, err = s.rpcClient.PhoneLogin(ctx, &pb.PhoneReq{
		Phone: phone,
	})
	if err != nil {
		return nil, server.HandleError(err)
	}

	return
}

// UserEditPwd 修改用户密码
func (s *Service) UserEditPwd(ctx context.Context, id uint32, password string) error {
	_, err := s.rpcClient.UserEditPwd(ctx, &pb.EditPwdReq{
		Id:  id,
		Pwd: password,
	})
	if err != nil {
		return server.HandleError(err)
	}

	return nil
}

// UserEdit 修改用户信息
func (s *Service) UserEdit(ctx context.Context, id uint32, userMap map[string]interface{}) error {
	bytes, err := json.Marshal(userMap)
	if err != nil {
		return errors.Wrapf(err, "[service.center] json marshal")
	}

	_, err = s.rpcClient.UserEdit(ctx, &pb.EditReq{
		Id:      id,
		Content: bytes,
	})
	if err != nil {
		return server.HandleError(err)
	}

	if u, ok := userMap["nickname"]; ok {
		//更新es数据
		err = s.ec.UserUpdate(ctx, id, map[string]interface{}{"nickname": u})
		if err != nil {
			return errors.Wrapf(err, "[service.center] save es by id:%d", id)
		}
	}
	if a, ok := userMap["avatar"]; ok {
		//更新es数据
		err = s.ec.UserUpdate(ctx, id, map[string]interface{}{"avatar": a})
		if err != nil {
			return errors.Wrapf(err, "[service.center] save es by id:%d", id)
		}
	}
	//清除缓存
	err = s.userCache.DelCache(ctx, id)
	if err != nil {
		log.Warnf("[service.center] del cache: %v", err)
	}

	return nil
}

// GetUserByID 获取用户信息
// 缓存的更新策略使用 Cache Aside Pattern
// see: https://coolshell.cn/articles/17416.html
func (s *Service) GetUserByID(ctx context.Context, id uint32) (user *base.UserInfo, err error) {
	// 从cache获取
	user, err = s.userCache.GetCache(ctx, id)
	if err != nil {
		if err == cache.ErrPlaceholder {
			return new(base.UserInfo), nil
		} else if err != redis.Nil {
			// fail fast, if cache error return, don't request to db
			return nil, errors.Wrapf(err, "[service.center] get user by uid: %d", id)
		}
	}
	// hit cache
	if user != nil {
		log.Debugf("[service.center] get user data from cache, uid: %d", id)
		return
	}

	// use sync/singleflight mode to get data
	// why not use redis lock? see this topic: https://redis.io/topics/distlock
	// demo see: https://github.com/go-demo/singleflight-demo/blob/master/main.go
	// https://juejin.cn/post/6844904084445593613
	getDataFn := func() (interface{}, error) {
		// 从中心服中获取
		user, err = s.rpcClient.UserInfo(ctx, &pb.UIDReq{Id: id})
		// if data is empty, set not found cache to prevent cache penetration(缓存穿透)
		if err != nil {
			err = server.HandleError(err)
			if errors.Is(err, center.ErrUserNotFound) { //不存在，缓存空
				err = s.userCache.SetCacheWithNotFound(ctx, id)
				if err != nil {
					log.Warnf("[service.center] SetCacheWithNotFound err, uid: %d", id)
				}
				return user, nil
			}
			return nil, errors.Wrapf(err, "[service.center] query db err")
		}

		// set cache
		err = s.userCache.SetCache(ctx, id, user)
		if err != nil {
			return nil, errors.Wrap(err, "[service.center] set cache data err")
		}
		return user, nil
	}

	g := singleflight.Group{}
	doKey := fmt.Sprintf("get_user_%d", id)
	val, err, _ := g.Do(doKey, getDataFn)
	if err != nil {
		return nil, errors.Wrap(err, "[service.center] get user err via single flight do")
	}
	data := val.(*base.UserInfo)

	return data, nil
}

// GetUsersByIds 批量获取用户
func (s *Service) GetUsersByIds(ctx context.Context, ids []uint32) ([]*base.UserInfo, error) {
	users := make([]*base.UserInfo, 0)

	// 从cache批量获取
	userCacheMap, err := s.userCache.MultiGetCache(ctx, ids)
	if err != nil {
		return users, errors.Wrap(err, "[service.center] multi get user cache data err")
	}

	// 查询未命中
	for _, uid := range ids {
		idx := s.userCache.BuildCacheKey([]interface{}{uid})
		userModel, ok := userCacheMap[idx]
		if !ok {
			userModel, err = s.GetUserByID(ctx, uid)
			if err != nil {
				log.Warnf("[service.center] get user err: %v", err)
				continue
			}
		}
		users = append(users, userModel)
	}
	return users, nil
}

// UserLogout 用户登出
func (s *Service) UserLogout(ctx context.Context, id uint32) error {
	_, err := s.rpcClient.UserLogout(ctx, &pb.UIDReq{Id: id})
	if err != nil {
		return server.HandleError(err)
	}

	return nil
}

// SendSMS 发送短信
func (s *Service) SendSMS(ctx context.Context, phone int64) (string, error) {
	reply, err := s.rpcClient.SendSMS(ctx, &pb.PhoneReq{Phone: phone})
	if err != nil {
		return "", server.HandleError(err)
	}

	return reply.Code, nil
}

// CheckVCode 验证校验码是否正确
func (s *Service) CheckVCode(ctx context.Context, phone int64, code string) error {
	_, err := s.rpcClient.CheckVCode(ctx, &pb.CheckCodeReq{
		Phone: phone,
		Code:  code,
	})
	if err != nil {
		return server.HandleError(err)
	}

	return nil
}

// ServerByUserID 用户长链接所在服务器ID
func (s *Service) ServerByUserID(ctx context.Context, id uint32) (string, error) {
	res, err := s.rpcClient.ServerByUserID(ctx, &pb.UIDReq{Id: id})
	if err != nil {
		return "", server.HandleError(err)
	}
	return res.ServerID, nil
}

// ServerByUserID 批量获取用户长链接所在服务器ID
func (s *Service) ServersByUserIDs(ctx context.Context, ids []uint32) ([]string, error) {
	res, err := s.rpcClient.BatchServersByUserIDs(ctx, &pb.UIDsReq{Ids: ids})
	if err != nil {
		return nil, server.HandleError(err)
	}

	return res.ServerIDs, nil
}
