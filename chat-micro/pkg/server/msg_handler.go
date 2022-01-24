package server

import (
	"fmt"
)

var _ IMsgHandle = (*MsgHandle)(nil)

//IMsgHandle 消息管理器
type IMsgHandle interface {
	DoMsgHandler(IRequest IRequest)       //马上以非阻塞方式处理消息
	StartWorkerPool(taskLen int)          //启动worker工作池
	SendMsgToTaskQueue(IRequest IRequest) //将消息交给TaskQueue,由worker进行处理
}

//MsgHandle 消息处理器结构
type MsgHandle struct {
	Handlers       *Engine         //路由处理器
	WorkerPoolSize int             //业务工作Worker池的数量
	TaskQueue      []chan IRequest //Worker负责取任务的消息队列
}

//NewMsgHandle 实例化消息处理器
func NewMsgHandle(poolSize int, r *Engine) IMsgHandle {
	return &MsgHandle{
		Handlers:       r,
		WorkerPoolSize: poolSize,
		TaskQueue:      make([]chan IRequest, poolSize),
	}
}

//DoMsgHandler 马上以非阻塞方式处理消息
func (m *MsgHandle) DoMsgHandler(req IRequest) {
	m.Handlers.Start(req)
}

//StartWorkerPool 启动worker工作池
func (m *MsgHandle) StartWorkerPool(taskLen int) {
	//遍历需要启动worker的数量，依此启动
	for i := 0; i < m.WorkerPoolSize; i++ {
		//一个worker被启动
		//给当前worker对应的任务队列开辟空间
		m.TaskQueue[i] = make(chan IRequest, taskLen)
		//启动当前Worker，阻塞的等待对应的任务队列是否有消息传递进来
		go m.startWorker(i, m.TaskQueue[i])
	}
}

//SendMsgToTaskQueue 将消息交给TaskQueue,由worker进行处理
func (m *MsgHandle) SendMsgToTaskQueue(IRequest IRequest) {
	//根据ConnID来分配当前的连接应该由哪个worker负责处理
	//轮询的平均分配法则

	//得到需要处理此条连接的workerID
	workerID := IRequest.Conn().GetID() % uint32(m.WorkerPoolSize)
	//将请求消息发送给任务队列
	m.TaskQueue[workerID] <- IRequest
}

//startWorker 启动一个Worker工作流程
func (m *MsgHandle) startWorker(workerID int, taskQueue chan IRequest) {
	fmt.Printf("[tcp.msgHandle] worker id %v started\n", workerID)
	//不断的等待队列中的消息
	for {
		select {
		//有消息则取出队列的IRequest，并执行绑定的业务方法
		case IRequest := <-taskQueue:
			m.DoMsgHandler(IRequest)
		}
	}
}
