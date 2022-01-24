package service

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"chat-micro/app/constvar"
	"chat-micro/pkg/app"
	"chat-micro/pkg/logger"
	"chat-micro/pkg/redis"
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
	serverID := s.ServerByUserID(ctx, userID)
	if serverID == "" { // 不存在，不必踢出上一次用户
		return nil
	}
	if err := s.UserOffline(ctx, userID); err != nil {
		return err
	}
	return s.pushClose(ctx, userID, serverID)
}

// UserOnline 用户上线
func (s *Service) UserOnline(ctx context.Context, token, serverID string) (uint32, error) {
	p, err := app.Parse(token, s.opts.jwtSecret)
	if err != nil {
		return 0, ErrUserTokenError
	}
	//获取当前合法用户token，检查此令牌是否已过期
	curToken := redis.Client.Get(ctx, constvar.BuildUserTokenKey(uint32(p.UserID))).Val()
	if curToken != token {
		return 0, ErrUserTokenExpired
	}
	//设置用户在线状态数据
	err = redis.Client.Set(ctx, constvar.BuildOnlineKey(uint32(p.UserID)), serverID, time.Duration(s.opts.jwtTimeout)*time.Second).Err()
	if err != nil {
		return 0, errors.Wrapf(err, "[service.online] user online hset err")
	}
	//发送当前用户的离线消息
	if err = s.pushHistory(ctx, uint32(p.UserID), serverID); err != nil {
		logger.Warnf("[service.online] push history err:%v", err)
	}
	return uint32(p.UserID), nil
}

// UserOffline 用户下线
func (s *Service) UserOffline(ctx context.Context, userID uint32) error {
	return redis.Client.Del(ctx, constvar.BuildOnlineKey(userID)).Err()
}

// CheckOnline 检查用户是否在线
func (s *Service) CheckOnline(ctx context.Context, userID uint32) (bool, error) {
	res, err := redis.Client.Exists(ctx, constvar.BuildOnlineKey(userID)).Result()
	if err != nil {
		return false, errors.Wrapf(err, "[service.online] check online id:%d", userID)
	}
	return res == redis.Success, nil
}

//ServerByUserID 获取用户所在服务器
func (s *Service) ServerByUserID(ctx context.Context, userID uint32) string {
	return redis.Client.Get(ctx, constvar.BuildOnlineKey(userID)).Val()
}

// ServersByUserIds 批量获取用户所在的连接服务器
func (s *Service) ServersByUserIds(ctx context.Context, userIds []uint32) ([]string, error) {
	if len(userIds) == 0 {
		return []string{}, nil
	}
	keys := make([]string, len(userIds))
	for i, userID := range userIds {
		keys[i] = constvar.BuildOnlineKey(userID)
	}
	list, err := redis.Client.MGet(ctx, keys...).Result()
	if err != nil {
		return nil, errors.Wrapf(err, "[service.online[ get keys:%v by redis", keys)
	}
	servers := make([]string, 0, len(list))
	for _, v := range list {
		if v == nil {
			servers = append(servers, "")
		} else {
			servers = append(servers, v.(string))
		}
	}
	return servers, nil
}
