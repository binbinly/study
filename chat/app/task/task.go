package task

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/consul/api"
	"github.com/pkg/errors"

	"chat/app/task/conf"
	"chat/pkg/log"
	"chat/pkg/queue"
	"chat/pkg/queue/iqueue"
	"chat/proto/logic"
)

//Task 任务层结构
type Task struct {
	c         *conf.Config
	consumer  iqueue.Consumer
	lastIndex uint64
	connects  map[string]*Connect
}

//New 实例化任务层
func New(c *conf.Config) *Task {
	t := &Task{
		c:        c,
		connects: make(map[string]*Connect),
	}
	return t
}

//Start 开启task服务
func (t *Task) Start() {
	fmt.Println("task server start !!!")
	//开启队列消费者
	t.consumer = queue.NewConsumer(&t.c.Queue, t.handlerMessage)
	t.watchConnect()
}

// handlerMessage 处理消息
func (t *Task) handlerMessage(body []byte) (err error) {
	log.Debugf("[task.message] begin body:%v", string(body))
	msg := &logic.SendMsg{}
	err = json.Unmarshal(body, msg)
	if err != nil {
		return errors.Wrapf(err, "[task.message] json unmarshal")
	}
	ctx := context.Background()
	switch msg.Type {
	case logic.SendMsg_CLOSE:
		err = t.close(ctx, msg)
	case logic.SendMsg_SEND:
		err = t.send(ctx, msg)
	case logic.SendMsg_BROADCAST:
		err = t.broadcast(ctx, msg)
	case logic.SendMsg_History:
		err = t.history(ctx, msg)
	default:
		log.Warnf("nsq msg no match push type: %s", msg.Type)
	}
	return
}

func (t *Task) watchConnect() {
	config := api.DefaultConfig()
	config.Address = t.c.Consul
	consul, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			services, metaInfo, err := consul.Health().Service(t.c.GrpcClient.ServiceName, "", true, &api.QueryOptions{WaitIndex: t.lastIndex})
			if err != nil {
				log.Warnf("[task.watcher] get services err:%v", err)
				time.Sleep(time.Second)
				continue
			}
			t.lastIndex = metaInfo.LastIndex
			if len(services) == 0 {
				time.Sleep(time.Second)
				continue
			}
			log.Debugf("[task.watch] services size:%v", len(services))

			adds := map[string]string{}

			for _, v := range services {
				adds[v.Service.ID] = fmt.Sprintf("%v:%v", v.Service.Address, v.Service.Port)
			}
			if err = t.updateConnects(adds); err != nil {
				log.Warnf("[task.watcher] update connects err:%v", err)
				time.Sleep(time.Second)
				continue
			}
		}
	}()
}

//updateConnects 更新连接层对象
func (t *Task) updateConnects(adds map[string]string) error {
	connects := map[string]*Connect{}
	for id, addr := range adds {
		if old, ok := t.connects[id]; ok {
			connects[id] = old
			continue
		}
		c, err := NewConnect(t.c, id, addr)
		if err != nil {
			return errors.Wrapf(err, "[task.newAddress] NewConnect id:%v, addr:%v", id, addr)
		}
		connects[id] = c
		log.Debugf("watchConnect add connect id:%v, addr:%v", id, addr)
	}
	for id, old := range t.connects {
		if _, ok := connects[id]; !ok {
			old.cancel()
			log.Debugf("watchConnect del connect id:%v", id)
		}
	}
	t.connects = connects
	return nil
}

//Close 关闭服务
func (t *Task) Close() {
	for _, connect := range t.connects {
		connect.cancel()
	}
	t.consumer.Stop()
}
