package service

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"go-micro.dev/v4/logger"

	"common/constvar"
	"common/errno"
	"common/message"
	"pkg/app"
	"pkg/redis"
)

//IOnline 用户在线服务接口
type IOnline interface {
	//踢出上次登录
	UserTickOut(ctx context.Context, userID int64) error
	//检查用户是否在线
	CheckOnline(ctx context.Context, userID int64) (bool, error)
	//用户上线
	UserOnline(ctx context.Context, token, serverID string) (int64, error)
	//用户下线
	UserOffline(ctx context.Context, userID int64) error
	//获取用户所在连接服务器ID
	GetServerID(ctx context.Context, userID int64) string
	//批量获取用户所在的连接服务器id
	BatchServerIds(ctx context.Context, ids []int64) ([]string, error)
}

//UserTickOut 踢出上次登录
func (c *Center) UserTickOut(ctx context.Context, userID int64) error {
	server := c.GetServerID(ctx, userID)
	if server == "" { // 不存在，不必踢出上一次用户
		return nil
	}
	if err := c.UserOffline(ctx, userID); err != nil {
		return err
	}
	//包装关闭客户端连接消息
	msg, err := message.NewMessagePack(message.EventClose, "账号已在其他地方登录了!")
	if err != nil {
		return err
	}
	return c.CloseClient(ctx, userID, server, msg)
}

// UserOnline 用户上线
func (c *Center) UserOnline(ctx context.Context, token, serverID string) (int64, error) {
	p, err := app.Parse(token, c.c.App.JwtSecret)
	if err != nil {
		return 0, errno.ErrUserTokenEmpty
	}
	//获取当前合法用户token，检查此令牌是否已过期
	curToken := redis.Client.Get(ctx, constvar.BuildUserTokenKey(int64(p.UserID))).Val()
	if curToken != token {
		return 0, errno.ErrUserTokenExpired
	}
	//设置用户在线状态数据
	err = redis.Client.Set(ctx, constvar.BuildOnlineKey(int64(p.UserID)), serverID, time.Duration(c.c.App.JwtTimeout)*time.Second).Err()
	if err != nil {
		return 0, errors.Wrapf(err, "[center.online] user online hset err")
	}
	//发送当前用户的离线消息
	if err = c.PushHistory(ctx, int64(p.UserID), serverID); err != nil {
		logger.Warnf("[center.online] push history err:%v", err)
	}
	return int64(p.UserID), nil
}

// UserOffline 用户下线
func (c *Center) UserOffline(ctx context.Context, userID int64) error {
	return redis.Client.Del(ctx, constvar.BuildOnlineKey(userID)).Err()
}

// CheckOnline 检查用户是否在线
func (c *Center) CheckOnline(ctx context.Context, userID int64) (bool, error) {
	res, err := redis.Client.Exists(ctx, constvar.BuildOnlineKey(userID)).Result()
	if err != nil {
		return false, errors.Wrapf(err, "[center.online] check online id:%d", userID)
	}
	return res == redis.Success, nil
}

//GetServerID 获取用户所在服务器
func (c *Center) GetServerID(ctx context.Context, userID int64) string {
	id, err := redis.Client.Get(ctx, constvar.BuildOnlineKey(userID)).Result()
	if err != nil && err != redis.Nil {
		logger.Warnf("[center.online] get serverId by uid: %v err: %v", userID, err)
	}
	return id
}

// BatchServerIds 批量获取用户所在的连接服务器
func (c *Center) BatchServerIds(ctx context.Context, userIds []int64) ([]string, error) {
	if len(userIds) == 0 {
		return []string{}, nil
	}
	keys := make([]string, len(userIds))
	for i, userID := range userIds {
		keys[i] = constvar.BuildOnlineKey(userID)
	}
	list, err := redis.Client.MGet(ctx, keys...).Result()
	if err != nil {
		return nil, errors.Wrapf(err, "[center.online[ get keys:%v by redis", keys)
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
