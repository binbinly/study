package task

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"chat-micro/app/constvar"
	"chat-micro/pkg/logger"
	"chat-micro/pkg/redis"
	pb "chat-micro/proto/connect"
	task "chat-micro/proto/task"
)

//send 发送指定消息
func (t *Task) send(msg *task.Message) (err error) {
	req := &pb.SendReq{
		UserIds: msg.UserIds,
		Data:    msg.Body,
	}
	conn, ok := t.connects[msg.Server]
	if !ok {
		return fmt.Errorf("[task.send] not found serverId: %v", msg.Server)
	}
	if err = conn.Push(req); err != nil {
		return errors.Wrapf(err, "[task.send] serverId:%v userIds:%v event:%s", msg.Server, msg.UserIds, msg.Event)
	}
	return
}

//history 发送历史消息
func (t *Task) history(ctx context.Context, msg *task.Message) error {
	conn, ok := t.connects[msg.Server]
	if !ok {
		return fmt.Errorf("[task.history] not found serverId: %v", msg.Server)
	}
	list, err := redis.Client.LRange(ctx, constvar.BuildHistoryKey(msg.UserIds[0]), 0, -1).Result()
	if err != nil {
		return errors.Wrapf(err, "[task.history] lrange err: %v", err)
	}
	go func() {
		for _, m := range list {
			req := &pb.SendReq{
				UserIds: msg.UserIds,
				Data:    []byte(m),
			}
			if err = conn.Push(req); err != nil {
				logger.Warnf("[task.history] serverId:%v userIds:%v event:%s", msg.Server, msg.UserIds, msg.Event)
			}
		}
	}()

	return nil
}

//close 发送关闭连接消息
func (t *Task) close(ctx context.Context, msg *task.Message) (err error) {
	conn, ok := t.connects[msg.Server]
	if !ok {
		return fmt.Errorf("[task.history] not found serverId: %v", msg.Server)
	}
	req := &pb.CloseReq{
		UserId: msg.UserIds[0],
		Data:   msg.Body,
	}
	if err = conn.Close(ctx, req); err != nil {
		return errors.Wrapf(err, "[task.send] serverId:%v userIds:%v event:%s", msg.Server, msg.UserIds, msg.Event)
	}
	return nil
}

//broadcast 广播消息
func (t *Task) broadcast(msg *task.Message) (err error) {
	req := &pb.BroadcastReq{Data: msg.Body}
	for id, c := range t.connects {
		if err = c.Broadcast(req); err != nil {
			return errors.Wrapf(err, "[task.broadcast] serverId:%v userIds:%v event:%s", id, msg.UserIds, msg.Event)
		}
	}
	return
}
