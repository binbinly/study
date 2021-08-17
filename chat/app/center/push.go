package center

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"

	"chat/app/constvar"
	"chat/app/message"
	"chat/pkg/log"
	"chat/pkg/redis"
	"chat/proto/base"
)

//PushReq 消息结构
type PushReq struct {
	UserID uint32
	Event  string
	Data   []byte
}

//IPush 推送消息至消息队列接口
type IPush interface {
	//批量加入队列
	PushBatch(ctx context.Context, req []*PushReq) (err error)
	//多用户发送消息
	PushUserIds(ctx context.Context, userIds []uint32, event string, data []byte) (err error)
	//主动关闭客户端连接
	CloseClient(ctx context.Context, userID uint32, server string, data []byte) (err error)
	//发送消息给所有
	PushAll(ctx context.Context, event string, data []byte) (err error)
	//发送用户历史消息
	PushHistory(ctx context.Context, userID uint32, server string) error
}

//PushBatch 批量加入队列
func (c *Center) PushBatch(ctx context.Context, reqs []*PushReq) (err error) {
	userIds := make([]uint32, 0, len(reqs))
	for _, pushReq := range reqs {
		userIds = append(userIds, pushReq.UserID)
	}
	servers, err := c.ServersByUserIds(ctx, userIds)
	if err != nil {
		return errors.Wrapf(err, "[center.push] get servers")
	}
	mServers := make(map[uint32]string, len(servers))
	for i, userID := range userIds {
		mServers[userID] = servers[i]
	}
	ms := make([][]byte, 0, len(reqs))
	for _, pushReq := range reqs {
		if serverID, ok := mServers[pushReq.UserID]; ok {
			if serverID == "" && pushReq.UserID > 0 { //用户不在线，保存离线消息
				redis.Client.RPush(ctx, constvar.BuildHistoryKey(pushReq.UserID), pushReq.Data)
				continue
			}
			m, err := json.Marshal(&base.SendMsg{
				Type:    base.SendMsg_SEND,
				Server:  serverID,
				Event:   pushReq.Event,
				UserIds: []uint32{pushReq.UserID},
				Msg:     pushReq.Data,
			})
			if err != nil {
				log.Warnf("[center.push] batch json marshal err:%v", err)
				continue
			}
			ms = append(ms, m)
		} else {
			log.Warnf("[center.push[ userID:%v not serverID", pushReq.UserID)
		}
	}
	return c.queue.MultiPublish(ctx, ms...)
}

// PushUserIds push a message by userIds.
func (c *Center) PushUserIds(ctx context.Context, userIds []uint32, event string, data []byte) (err error) {
	servers, err := c.ServersByUserIds(ctx, userIds)
	if err != nil {
		return errors.Wrapf(err, "[center.push] get servers err")
	}
	pushKeys := make(map[string][]uint32)
	var off []uint32
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
		msg := &base.SendMsg{
			Type:    base.SendMsg_SEND,
			Server:  server,
			Event:   event,
			UserIds: userIds,
			Msg:     data,
		}
		err = c.publish(ctx, msg)
		if err != nil {
			return err
		}
	}
	return nil
}

//PushHistory 发送离线消息，一般在用户刚登陆时
func (c *Center) PushHistory(ctx context.Context, userID uint32, server string) error {
	msg := &base.SendMsg{
		Type:    base.SendMsg_History,
		Server:  server,
		UserIds: []uint32{userID},
	}
	return c.publish(ctx, msg)
}

//CloseClient 主动关闭客户端消息
func (c *Center) CloseClient(ctx context.Context, userID uint32, server string, data []byte) (err error) {
	msg := &base.SendMsg{
		Type:    base.SendMsg_CLOSE,
		Server:  server,
		Event:   message.EventClose,
		UserIds: []uint32{userID},
		Msg:     data,
	}
	return c.publish(ctx, msg)
}

// PushAll 广播消息
func (c *Center) PushAll(ctx context.Context, event string, data []byte) (err error) {
	msg := &base.SendMsg{
		Type:  base.SendMsg_BROADCAST,
		Event: event,
		Msg:   data,
	}
	return c.publish(ctx, msg)
}

// publish 入消息队列
func (c *Center) publish(ctx context.Context, msg *base.SendMsg) error {
	m, err := json.Marshal(msg)
	if err != nil {
		return errors.Wrapf(err, "[center.push] json marshal")
	}
	return c.queue.Publish(ctx, m)
}
