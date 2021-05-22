package service

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"

	"chat/app/logic/message"
	"chat/pkg/app"
	"chat/pkg/errno"
	"chat/pkg/redis"
)

const (
	// onlinePrefix 在线key前缀
	onlinePrefix = "online:"
	// userPrefix 用户令牌标识 用于单点登录
	userPrefix = "user:"
)

type IOnline interface {
	//踢出上次登录
	UserTickOut(ctx context.Context, userId uint32) error
	//检查用户是否在线
	CheckOnline(ctx context.Context, userId uint32) (bool, error)
	//用户上线
	UserOnline(ctx context.Context, token, serverId string) (uint32, error)
	//用户下线
	UserOffline(ctx context.Context, userId uint32) error
	//后去用户所在连接服务器ID
	ServerByUserId(ctx context.Context, userId uint32) string
	//批量获取用户所在的连接服务器id
	ServersByUserIds(ctx context.Context, userId []uint32) ([]string, error)
}

func (s *Service) UserTickOut(ctx context.Context, userId uint32) error {
	server := s.ServerByUserId(ctx, userId)
	if server == "" { // 不存在，不必踢出上一次用户
		return nil
	}
	err := s.UserOffline(ctx, userId)
	if err != nil {
		return err
	}
	//关闭客户端
	msg, err := app.NewMessagePack(message.EventClose, "账号已在其他地方登录了!")
	if err != nil {
		return err
	}
	err = s.CloseClient(ctx, userId, server, msg)
	if err != nil {
		return errors.Wrapf(err, "[service.online] push fail err")
	}
	return nil
}

// UserOnline 用户上线
func (s *Service) UserOnline(ctx context.Context, token, serverId string) (uint32, error) {
	p, err := app.Parse(token, s.c.App.JwtSecret)
	if err != nil {
		return 0, errors.Wrapf(err, "[service.online] token parse")
	}
	//获取当前合法用户token，检查此令牌是否已过期
	curToken := redis.Client.Get(s.getUserKey(uint32(p.UserId))).Val()
	if curToken != token {
		return 0, errno.ErrTokenTimeout
	}
	//设置用户在线状态数据
	err = redis.Client.Set(s.getOnlineKey(uint32(p.UserId)), serverId, time.Duration(s.c.App.JwtTimeout)*time.Second).Err()
	if err != nil {
		return 0, errors.Wrapf(err, "[service.online] user online hset err")
	}
	//发送历史信号
	s.PushHistory(ctx, uint32(p.UserId), serverId)

	return uint32(p.UserId), nil
}

// UserOffline 用户下线
func (s *Service) UserOffline(ctx context.Context, userId uint32) error {
	redis.Client.Del(s.getOnlineKey(userId))
	return nil
}

// CheckOnline 检查用户是否在线
func (s *Service) CheckOnline(ctx context.Context, userId uint32) (bool, error) {
	res, err := redis.Client.Exists(s.getOnlineKey(userId)).Result()
	if err != nil {
		return false, errors.Wrapf(err, "[service.online] check online id:%d", userId)
	}
	return res == redis.Success, nil
}

//ServerByUserId 获取用户所在服务器
func (s *Service) ServerByUserId(ctx context.Context, userId uint32) string {
	return redis.Client.Get(s.getOnlineKey(userId)).Val()
}

// ServersByUserIds 获取用户所在的连接服务器
func (s *Service) ServersByUserIds(ctx context.Context, userIds []uint32) ([]string, error) {
	if len(userIds) == 0 {
		return make([]string, 0), nil
	}
	var keys []string
	for _, userId := range userIds {
		keys = append(keys, s.getOnlineKey(userId))
	}
	list, err := redis.Client.MGet(keys...).Result()
	if err != nil {
		return nil, errors.Wrapf(err, "[service.online[ get keys:%v by redis", keys)
	}
	var servers []string
	for _, v := range list {
		if v == nil {
			servers = append(servers, "")
		} else {
			servers = append(servers, v.(string))
		}
	}
	return servers, nil
}

// getKey 获取用户键值
func (s *Service) getOnlineKey(userId uint32) string {
	return fmt.Sprintf("%s%d", onlinePrefix, userId)
}

func (s *Service) getUserKey(userId uint32) string {
	return fmt.Sprintf("%s%d", userPrefix, userId)
}
