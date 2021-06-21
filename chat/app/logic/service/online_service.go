package service

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"

	"chat/app/message"
	"chat/pkg/app"
	"chat/pkg/errno"
	"chat/pkg/log"
	"chat/pkg/redis"
)

const (
	// onlinePrefix 在线key前缀
	onlinePrefix = "online:"
	// userPrefix 用户令牌标识 用于单点登录
	userPrefix = "user:"
)

//IOnline 用户在线服务接口
type IOnline interface {
	//踢出上次登录
	UserTickOut(ctx context.Context, userID uint32) error
	//检查用户是否在线
	CheckOnline(ctx context.Context, userID uint32) (bool, error)
	//用户上线
	UserOnline(ctx context.Context, token, serverID string) (uint32, error)
	//用户下线
	UserOffline(ctx context.Context, userID uint32) error
	//后去用户所在连接服务器ID
	ServerByUserID(ctx context.Context, userID uint32) string
	//批量获取用户所在的连接服务器id
	ServersByUserIds(ctx context.Context, userID []uint32) ([]string, error)
}

//UserTickOut 踢出上次登录
func (s *Service) UserTickOut(ctx context.Context, userID uint32) error {
	server := s.ServerByUserID(ctx, userID)
	if server == "" { // 不存在，不必踢出上一次用户
		return nil
	}
	err := s.UserOffline(ctx, userID)
	if err != nil {
		return err
	}
	//关闭客户端
	msg, err := app.NewMessagePack(message.EventClose, "账号已在其他地方登录了!")
	if err != nil {
		return err
	}
	err = s.CloseClient(ctx, userID, server, msg)
	if err != nil {
		return errors.Wrapf(err, "[service.online] push fail err")
	}
	return nil
}

// UserOnline 用户上线
func (s *Service) UserOnline(ctx context.Context, token, serverID string) (uint32, error) {
	p, err := app.Parse(token, s.c.App.JwtSecret)
	if err != nil {
		return 0, errors.Wrapf(err, "[service.online] token parse")
	}
	//获取当前合法用户token，检查此令牌是否已过期
	curToken := redis.Client.Get(ctx, s.getUserKey(uint32(p.UserID))).Val()
	if curToken != token {
		return 0, errno.ErrTokenTimeout
	}
	//设置用户在线状态数据
	err = redis.Client.Set(ctx, s.getOnlineKey(uint32(p.UserID)), serverID, time.Duration(s.c.App.JwtTimeout)*time.Second).Err()
	if err != nil {
		return 0, errors.Wrapf(err, "[service.online] user online hset err")
	}
	//发送历史信号
	err = s.PushHistory(ctx, uint32(p.UserID), serverID)
	if err != nil {
		log.Warnf("[service.online] push history err:%v", err)
	}
	return uint32(p.UserID), nil
}

// UserOffline 用户下线
func (s *Service) UserOffline(ctx context.Context, userID uint32) error {
	redis.Client.Del(ctx, s.getOnlineKey(userID))
	return nil
}

// CheckOnline 检查用户是否在线
func (s *Service) CheckOnline(ctx context.Context, userID uint32) (bool, error) {
	res, err := redis.Client.Exists(ctx, s.getOnlineKey(userID)).Result()
	if err != nil {
		return false, errors.Wrapf(err, "[service.online] check online id:%d", userID)
	}
	return res == redis.Success, nil
}

//ServerByUserID 获取用户所在服务器
func (s *Service) ServerByUserID(ctx context.Context, userID uint32) string {
	return redis.Client.Get(ctx, s.getOnlineKey(userID)).Val()
}

// ServersByUserIds 获取用户所在的连接服务器
func (s *Service) ServersByUserIds(ctx context.Context, userIds []uint32) ([]string, error) {
	if len(userIds) == 0 {
		return make([]string, 0), nil
	}
	var keys []string
	for _, userID := range userIds {
		keys = append(keys, s.getOnlineKey(userID))
	}
	list, err := redis.Client.MGet(ctx, keys...).Result()
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
func (s *Service) getOnlineKey(userID uint32) string {
	return fmt.Sprintf("%s%d", onlinePrefix, userID)
}

func (s *Service) getUserKey(userID uint32) string {
	return fmt.Sprintf("%s%d", userPrefix, userID)
}
