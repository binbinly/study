package service

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"go-micro.dev/v4/logger"

	"common/constvar"
	"common/message"
	pb "common/proto/task"
	"pkg/redis"
)

//PushReq 消息结构
type PushReq struct {
	UserID int64
	Event  string
	Data   []byte
}

//IPush 推送消息至消息队列接口
type IPush interface {
	//批量加入队列
	PushBatch(ctx context.Context, req []*PushReq) (err error)
	//多用户发送消息
	PushUserIds(ctx context.Context, userIds []int64, event string, data []byte) (err error)
	//主动关闭客户端连接
	CloseClient(ctx context.Context, userID int64, server string, data []byte) (err error)
	//发送消息给所有
	PushAll(ctx context.Context, event string, data []byte) (err error)
	//发送用户历史消息
	PushHistory(ctx context.Context, userID int64, server string) error
}

//PushBatch 批量加入队列
func (c *Center) PushBatch(ctx context.Context, reqs []*PushReq) (err error) {
	userIds := make([]int64, 0, len(reqs))
	for _, pushReq := range reqs {
		userIds = append(userIds, pushReq.UserID)
	}
	servers, err := c.BatchServerIds(ctx, userIds)
	if err != nil {
		return errors.Wrapf(err, "[center.push] get servers")
	}
	mServers := make(map[int64]string, len(servers))
	for i, userID := range userIds {
		mServers[userID] = servers[i]
	}
	for _, pushReq := range reqs {
		if serverID, ok := mServers[pushReq.UserID]; ok {
			if serverID == "" && pushReq.UserID > 0 { //用户不在线，保存离线消息
				redis.Client.RPush(ctx, constvar.BuildHistoryKey(pushReq.UserID), pushReq.Data)
				continue
			}
			if err = c.event.Publish(ctx, &pb.Msg{
				Type:      pb.Msg_SEND,
				Server:    serverID,
				Event:     pushReq.Event,
				UserIds:   []int64{pushReq.UserID},
				Timestamp: time.Now().Unix(),
				Body:      pushReq.Data,
			}); err != nil {
				logger.Warnf("[center.push] publish err:%v", err)
				continue
			}
		} else {
			logger.Warnf("[center.push[ userID:%v not serverID", pushReq.UserID)
		}
	}
	return nil
}

// PushUserIds push a message by userIds.
func (c *Center) PushUserIds(ctx context.Context, userIds []int64, event string, data []byte) (err error) {
	servers, err := c.BatchServerIds(ctx, userIds)
	if err != nil {
		return errors.Wrapf(err, "[center.push] get servers err")
	}
	pushKeys := make(map[string][]int64)
	var off []int64
	for i, userID := range userIds {
		server := servers[i]
		if server != "" && userID != 0 {
			pushKeys[server] = append(pushKeys[server], userID)
		} else { //不在线
			off = append(off, userID)
		}
	}
	//不在线用户保存离线消息
	for _, u := range off {
		if u == 0 {
			continue
		}
		redis.Client.RPush(ctx, constvar.BuildHistoryKey(u), data)
	}
	//在线用户发送消息
	for server := range pushKeys {
		msg := &pb.Msg{
			Type:      pb.Msg_SEND,
			Server:    server,
			Event:     event,
			UserIds:   userIds,
			Timestamp: time.Now().Unix(),
			Body:      data,
		}
		err = c.event.Publish(ctx, msg)
		if err != nil {
			return err
		}
	}
	return nil
}

//PushHistory 发送离线消息，一般在用户刚登陆时
func (c *Center) PushHistory(ctx context.Context, userID int64, server string) error {
	msg := &pb.Msg{
		Type:      pb.Msg_History,
		Server:    server,
		Timestamp: time.Now().Unix(),
		UserIds:   []int64{userID},
	}
	return c.event.Publish(ctx, msg)
}

//CloseClient 主动关闭客户端消息
func (c *Center) CloseClient(ctx context.Context, userID int64, server string, data []byte) (err error) {
	msg := &pb.Msg{
		Type:      pb.Msg_CLOSE,
		Server:    server,
		Event:     message.EventClose,
		UserIds:   []int64{userID},
		Timestamp: time.Now().Unix(),
		Body:      data,
	}
	return c.event.Publish(ctx, msg)
}

// PushAll 广播消息
func (c *Center) PushAll(ctx context.Context, event string, data []byte) (err error) {
	msg := &pb.Msg{
		Type:      pb.Msg_BROADCAST,
		Event:     event,
		Timestamp: time.Now().Unix(),
		Body:      data,
	}
	return c.event.Publish(ctx, msg)
}
