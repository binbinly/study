package task

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/consul/api"
	"github.com/pkg/errors"

	"chat-micro/pkg/broker"
	"chat-micro/pkg/logger"
	pb "chat-micro/proto/task"
)

//Task 任务层结构
type Task struct {
	opts      options
	connects  map[string]*Connect
	lastIndex uint64
}

//New 实例化任务层
func New(opts ...Option) *Task {
	return &Task{
		opts:     newOptions(opts...),
		connects: make(map[string]*Connect),
	}
}

//Run 启动task服务
func (t *Task) Run() {
	fmt.Println("Task Server Run !!!")
	//开启队列消费者
	_, err := t.opts.broker.Subscribe(t.opts.topic, t.handlerMessage)
	if err != nil {
		log.Fatalf("Broker Subscribe err: %v", err)
	}
	t.watchConnect()
}

// handlerMessage 处理消息
func (t *Task) handlerMessage(event broker.Event) (err error) {
	logger.Infof("[task.message] begin body:%s", event.Message().Body)

	msg := &pb.Message{}
	if err = json.Unmarshal(event.Message().Body, msg); err != nil {
		logger.Warnf("[task.message] json unmarshal err: %v", err)
		return nil
	}
	ctx := context.Background()
	switch msg.Type {
	case pb.Message_CLOSE:
		err = t.close(ctx, msg)
	case pb.Message_SEND:
		err = t.send(msg)
	case pb.Message_BROADCAST:
		err = t.broadcast(msg)
	case pb.Message_History:
		err = t.history(ctx, msg)
	default:
		logger.Warnf("[task.message] not found message type: %s", msg.Type)
	}
	return
}

func (t *Task) watchConnect() {
	config := api.DefaultConfig()
	config.Address = t.opts.registryAddress
	consul, err := api.NewClient(config)
	if err != nil {
		log.Fatalf("consul client new err: %v", err)
	}
	go func() {
		for {
			services, metaInfo, err := consul.Health().Service(t.opts.connectName, "", true, &api.QueryOptions{WaitIndex: t.lastIndex})
			if err != nil {
				logger.Warnf("[task.watcher] get services err:%v", err)
				time.Sleep(time.Second)
				continue
			}
			t.lastIndex = metaInfo.LastIndex
			if len(services) == 0 {
				time.Sleep(time.Second)
				continue
			}
			logger.Infof("[task.watch] services size:%v", len(services))

			adds := map[string]string{}

			for _, v := range services {
				adds[v.Service.ID] = fmt.Sprintf("%v:%v", v.Service.Address, v.Service.Port)
			}
			if err = t.updateConnects(adds); err != nil {
				logger.Warnf("[task.watcher] update connects err:%v", err)
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
		c, err := NewConnect(t.opts, id, addr)
		if err != nil {
			return errors.Wrapf(err, "[task.newAddress] NewConnect id:%v, addr:%v", id, addr)
		}
		connects[id] = c
		logger.Debugf("watchConnect add connect id:%v, addr:%v", id, addr)
	}
	for id, old := range t.connects {
		if _, ok := connects[id]; !ok {
			old.cancel()
			logger.Debugf("watchConnect del connect id:%v", id)
		}
	}
	t.connects = connects
	return nil
}

//Stop 停止服务
func (t *Task) Stop() error {
	for _, connect := range t.connects {
		connect.cancel()
	}
	return t.opts.broker.Disconnect()
}
