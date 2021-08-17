package task

import (
	"context"

	"github.com/pkg/errors"

	"chat/app/constvar"
	"chat/pkg/log"
	"chat/pkg/redis"
	"chat/proto/base"
	"chat/proto/connect"
)

//send 发送指定消息
func (t *Task) send(ctx context.Context, msg *base.SendMsg) (err error) {
	req := &connect.SendReq{
		UserIds: msg.UserIds,
		Proto: &base.Proto{
			Data: msg.Msg,
		},
	}
	if c, ok := t.connects[msg.Server]; ok {
		if err = c.Push(req); err != nil {
			return errors.Wrapf(err, "[task.send] serverId:%v userIds:%v event:%s", msg.Server, msg.UserIds, msg.Event)
		}
	}
	return
}

//history 发送历史消息
func (t *Task) history(ctx context.Context, msg *base.SendMsg) (err error) {
	go func() {
		for {
			val := redis.Client.LPop(ctx, constvar.BuildHistoryKey(msg.UserIds[0])).Val()
			if val == "" {
				break
			}
			req := &connect.SendReq{
				UserIds: msg.UserIds,
				Proto: &base.Proto{
					Data: []byte(val),
				},
			}
			if c, ok := t.connects[msg.Server]; ok {
				if err = c.Push(req); err != nil {
					log.Warnf("[task.history] serverId:%v userIds:%v event:%s", msg.Server, msg.UserIds, msg.Event)
				}
				log.Infof("[task.history] serverId:%v userIds:%v event:%s", msg.Server, msg.UserIds, msg.Event)
			}
		}
	}()

	return nil
}

//close 发送关闭连接消息
func (t *Task) close(ctx context.Context, msg *base.SendMsg) (err error) {
	req := &connect.CloseReq{
		UserId: msg.UserIds[0],
		Proto: &base.Proto{
			Data: msg.Msg,
		},
	}
	if c, ok := t.connects[msg.Server]; ok {
		if err = c.Close(ctx, req); err != nil {
			return errors.Wrapf(err, "[task.send] serverId:%v userIds:%v event:%s", msg.Server, msg.UserIds, msg.Event)
		}
		log.Infof("[task.send] serverId:%s userId:%v event:%d", msg.Server, msg.UserIds, msg.Event)
	}
	return nil
}

//broadcast 广播消息
func (t *Task) broadcast(ctx context.Context, msg *base.SendMsg) (err error) {
	req := &connect.BroadcastReq{
		Proto: &base.Proto{
			Data: msg.Msg,
		},
	}
	for id, c := range t.connects {
		if err = c.Broadcast(req); err != nil {
			return errors.Wrapf(err, "[task.broadcast] serverId:%v userIds:%v event:%s", id, msg.UserIds, msg.Event)
		}
	}
	log.Infof("[task.broadcast] connect len:%v", len(t.connects))
	return
}
