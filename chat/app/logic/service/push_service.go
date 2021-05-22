package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"

	"chat/app/constvar"
	"chat/app/logic/message"
	"chat/pkg/log"
	"chat/pkg/redis"
	"chat/proto/logic"
)

type PushReq struct {
	UserId uint32
	Event  string
	Data   []byte
}

type IPush interface {
	//批量加入队列
	PushBatch(ctx context.Context, req []*PushReq) (err error)
	//多用户发送消息
	PushUserIds(ctx context.Context, userIds []uint32, event string, data []byte) (err error)
	//主动关闭客户端连接
	CloseClient(ctx context.Context, userId uint32, server string, data []byte) (err error)
	//发送消息给所有
	PushAll(ctx context.Context, event string, data []byte) (err error)
	//发送用户历史消息
	PushHistory(ctx context.Context, userId uint32, server string) error
}

//PushBatch 批量加入队列
func (s *Service) PushBatch(ctx context.Context, req []*PushReq) (err error) {
	userIds := make([]uint32, 0)
	for _, pushReq := range req {
		userIds = append(userIds, pushReq.UserId)
	}
	servers, err := s.ServersByUserIds(ctx, userIds)
	if err != nil {
		return errors.Wrapf(err, "[service.push] get servers err")
	}
	mServers := make(map[uint32]string, len(servers))
	for i, userId := range userIds {
		mServers[userId] = servers[i]
	}
	ms := make([][]byte, 0)
	for _, pushReq := range req {
		if serverId, ok := mServers[pushReq.UserId]; ok {
			if serverId == "" && pushReq.UserId > 0 { //用户不在线，保存历史消息
				redis.Client.RPush(s.getHistoryKey(pushReq.UserId), pushReq.Data)
				continue
			}
			m, err := json.Marshal(&logic.SendMsg{
				Type:    logic.SendMsg_SEND,
				Server:  serverId,
				Event:   pushReq.Event,
				UserIds: []uint32{pushReq.UserId},
				Msg:     pushReq.Data,
			})
			if err != nil {
				log.Warnf("[service.push] batch json marshal err:%v", err)
				continue
			}
			ms = append(ms, m)
		} else {
			log.Warnf("[service.push[ userId:%v not serverId", pushReq.UserId)
		}
	}
	return s.queue.MultiPublish(ms...)
}

// PushUserIds push a message by userIds.
func (s *Service) PushUserIds(ctx context.Context, userIds []uint32, event string, data []byte) (err error) {
	servers, err := s.ServersByUserIds(ctx, userIds)
	if err != nil {
		return errors.Wrapf(err, "[service.push] get servers err")
	}
	pushKeys := make(map[string][]uint32)
	var off []uint32
	for i, userId := range userIds {
		server := servers[i]
		if server != "" && userId != 0 {
			pushKeys[server] = append(pushKeys[server], userId)
		} else { //不在线
			off = append(off, userId)
		}
	}
	//不在线用户保存历史消息
	for _, u := range off {
		if u == 0 {
			continue
		}
		redis.Client.RPush(s.getHistoryKey(u), data)
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
		err = s.queue.Publish(m)
		if err != nil {
			return errors.Wrapf(err, "[service.push] publish queue err")
		}
	}
	return nil
}

//PushHistory push历史消息入队列
func (s *Service) PushHistory(ctx context.Context, userId uint32, server string) error {
	msg := &logic.SendMsg{
		Type:    logic.SendMsg_History,
		Server:  server,
		UserIds: []uint32{userId},
	}
	m, err := json.Marshal(msg)
	if err != nil {
		return errors.Wrapf(err, "[service.push] json marshal")
	}
	err = s.queue.Publish(m)
	if err != nil {
		return errors.Wrapf(err, "[service.push] publish queue err")
	}
	return nil
}

//CloseClient 主动关闭客户端
func (s *Service) CloseClient(ctx context.Context, userId uint32, server string, data []byte) (err error) {
	msg := &logic.SendMsg{
		Type:    logic.SendMsg_CLOSE,
		Server:  server,
		Event:   message.EventClose,
		UserIds: []uint32{userId},
		Msg:     data,
	}
	m, err := json.Marshal(msg)
	if err != nil {
		return errors.Wrapf(err, "[push.queue] json marshal")
	}
	return s.queue.Publish(m)
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
	return s.queue.Publish(m)
}

func (s *Service) getHistoryKey(userId uint32) string {
	return fmt.Sprintf(constvar.HistoryPrefix, userId)
}
