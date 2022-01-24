package service

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"

	"chat-micro/app/constvar"
	"chat-micro/app/message"
	"chat-micro/pkg/broker"
	"chat-micro/pkg/redis"
	pb "chat-micro/proto/task"
)

type pack struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}

// pushUserIds push a message by userIds.
func (s *Service) pushUserIds(ctx context.Context, userIds []uint32, event string, v interface{}) (err error) {
	data, _ := json.Marshal(&pack{
		Event: event,
		Data:  v,
	})
	servers, err := s.ServersByUserIds(ctx, userIds)
	if err != nil {
		return errors.Wrapf(err, "[service.event] get servers err")
	}
	pushKeys := make(map[string][]uint32)
	var off []uint32
	for i, userID := range userIds {
		serverID := servers[i]
		if serverID != "" && userID != 0 {
			pushKeys[serverID] = append(pushKeys[serverID], userID)
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
	for serverID := range pushKeys {
		if err = s.publish(ctx, &pb.Message{
			Type:    pb.Message_SEND,
			Server:  serverID,
			Event:   event,
			UserIds: userIds,
			Body:    data,
		}); err != nil {
			return err
		}
	}
	return nil
}

//pushHistory 发送离线消息，一般在用户刚登陆时
func (s *Service) pushHistory(ctx context.Context, userID uint32, serverID string) error {
	return s.publish(ctx, &pb.Message{
		Type:    pb.Message_History,
		Server:  serverID,
		UserIds: []uint32{userID},
	})
}

//pushClose 主动关闭客户端消息
func (s *Service) pushClose(ctx context.Context, userID uint32, serverID string) error {
	data, _ := json.Marshal(&pack{
		Event: message.EventClose,
		Data:  msgCloseClient,
	})
	return s.publish(ctx, &pb.Message{
		Type:    pb.Message_CLOSE,
		Server:  serverID,
		Event:   message.EventClose,
		UserIds: []uint32{userID},
		Body:    data,
	})
}

// pushAll 广播消息
func (s *Service) pushAll(ctx context.Context, event string, data []byte) (err error) {
	return s.publish(ctx, &pb.Message{
		Type:  pb.Message_BROADCAST,
		Event: event,
		Body:  data,
	})
}

// publish 入消息队列
func (s *Service) publish(ctx context.Context, msg *pb.Message) error {
	b, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return s.opts.broker.Publish(s.opts.topic, &broker.Message{
		Body: b,
	})
}
