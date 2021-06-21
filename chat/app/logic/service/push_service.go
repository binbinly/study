package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"

	"chat/app/constvar"
	"chat/app/message"
	"chat/pkg/log"
	"chat/pkg/redis"
	"chat/proto/logic"
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
func (s *Service) PushBatch(ctx context.Context, req []*PushReq) (err error) {
	userIds := make([]uint32, 0)
	for _, pushReq := range req {
		userIds = append(userIds, pushReq.UserID)
	}
	servers, err := s.ServersByUserIds(ctx, userIds)
	if err != nil {
		return errors.Wrapf(err, "[service.push] get servers err")
	}
	mServers := make(map[uint32]string, len(servers))
	for i, userID := range userIds {
		mServers[userID] = servers[i]
	}
	ms := make([][]byte, 0)
	for _, pushReq := range req {
		if serverID, ok := mServers[pushReq.UserID]; ok {
			if serverID == "" && pushReq.UserID > 0 { //用户不在线，保存历史消息
				redis.Client.RPush(ctx, s.getHistoryKey(pushReq.UserID), pushReq.Data)
				continue
			}
			m, err := json.Marshal(&logic.SendMsg{
				Type:    logic.SendMsg_SEND,
				Server:  serverID,
				Event:   pushReq.Event,
				UserIds: []uint32{pushReq.UserID},
				Msg:     pushReq.Data,
			})
			if err != nil {
				log.Warnf("[service.push] batch json marshal err:%v", err)
				continue
			}
			ms = append(ms, m)
		} else {
			log.Warnf("[service.push[ userID:%v not serverID", pushReq.UserID)
		}
	}
	return s.queue.MultiPublish(ctx, ms...)
}

// PushUserIds push a message by userIds.
func (s *Service) PushUserIds(ctx context.Context, userIds []uint32, event string, data []byte) (err error) {
	servers, err := s.ServersByUserIds(ctx, userIds)
	if err != nil {
		return errors.Wrapf(err, "[service.push] get servers err")
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
	//不在线用户保存历史消息
	for _, u := range off {
		if u == 0 {
			continue
		}
		redis.Client.RPush(ctx, s.getHistoryKey(u), data)
	}
	//在线用户发送消息
	for server := range pushKeys {
		msg := &logic.SendMsg{
			Type:    logic.SendMsg_SEND,
			Server:  server,
			Event:   event,
			UserIds: userIds,
			Msg:     data,
		}
		m, err := json.Marshal(msg)
		if err != nil {
			return errors.Wrapf(err, "[service.push] json marshal")
		}
		err = s.queue.Publish(ctx, m)
		if err != nil {
			return errors.Wrapf(err, "[service.push] publish queue err")
		}
	}
	return nil
}

//PushHistory push历史消息入队列
func (s *Service) PushHistory(ctx context.Context, userID uint32, server string) error {
	msg := &logic.SendMsg{
		Type:    logic.SendMsg_History,
		Server:  server,
		UserIds: []uint32{userID},
	}
	m, err := json.Marshal(msg)
	if err != nil {
		return errors.Wrapf(err, "[service.push] json marshal")
	}
	err = s.queue.Publish(ctx, m)
	if err != nil {
		return errors.Wrapf(err, "[service.push] publish queue err")
	}
	return nil
}

//CloseClient 主动关闭客户端
func (s *Service) CloseClient(ctx context.Context, userID uint32, server string, data []byte) (err error) {
	msg := &logic.SendMsg{
		Type:    logic.SendMsg_CLOSE,
		Server:  server,
		Event:   message.EventClose,
		UserIds: []uint32{userID},
		Msg:     data,
	}
	m, err := json.Marshal(msg)
	if err != nil {
		return errors.Wrapf(err, "[push.queue] json marshal")
	}
	return s.queue.Publish(ctx, m)
}

// PushAll push message all
func (s *Service) PushAll(ctx context.Context, event string, data []byte) (err error) {
	msg := &logic.SendMsg{
		Type:  logic.SendMsg_BROADCAST,
		Event: event,
		Msg:   data,
	}
	m, err := json.Marshal(msg)
	if err != nil {
		return errors.Wrapf(err, "[push.queue] json marshal")
	}
	return s.queue.Publish(ctx, m)
}

func (s *Service) getHistoryKey(userID uint32) string {
	return fmt.Sprintf(constvar.HistoryPrefix, userID)
}
