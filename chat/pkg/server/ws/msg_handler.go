package ws

import (
	"fmt"
)

var _ IMsgHandle = (*MsgHandle)(nil)

// 消息管理器
type IMsgHandle interface {
	DoMsgHandler(request *Request)       //马上以非阻塞方式处理消息
	StartWorkerPool(taskLen uint32)      //启动worker工作池
	SendMsgToTaskQueue(request *Request) //将消息交给TaskQueue,由worker进行处理
}

type MsgHandle struct {
	Handlers       *Engine                 //路由处理器
	WorkerPoolSize uint32                  //业务工作Worker池的数量
	TaskQueue      []chan *Request //Worker负责取任务的消息队列
}

func NewMsgHandle(poolSize uint32, r *Engine) *MsgHandle {
	return &MsgHandle{
		Handlers:       r,
		WorkerPoolSize: poolSize,
		TaskQueue:      make([]chan *Request, poolSize),
	}
}

//DoMsgHandler 马上以非阻塞方式处理消息
func (m *MsgHandle) DoMsgHandler(req *Request) {
	m.Handlers.Start(req)
}

//StartWorkerPool 启动worker工作池
func (m *MsgHandle) StartWorkerPool(taskLen uint32) {
	//遍历需要启动worker的数量，依此启动
	for i := 0; i < int(m.WorkerPoolSize); i++ {
		//一个worker被启动
		//给当前worker对应的任务队列开辟空间
		m.TaskQueue[i] = make(chan *Request, taskLen)
		//启动当前Worker，阻塞的等待对应的任务队列是否有消息传递进来
		go m.startWorker(i, m.TaskQueue[i])
	}
}

//SendMsgToTaskQueue 将消息交给TaskQueue,由worker进行处理
func (m *MsgHandle) SendMsgToTaskQueue(request *Request) {
	//根据ConnID来分配当前的连接应该由哪个worker负责处理
	//轮询的平均分配法则

	//得到需要处理此条连接的workerID
	workerID := request.GetConnection().GetConnID() % m.WorkerPoolSize
	//将请求消息发送给任务队列
	m.TaskQueue[workerID] <- request
}

//startWorker 启动一个Worker工作流程
func (m *MsgHandle) startWorker(workerID int, taskQueue chan *Request) {
	fmt.Printf("[tcp.msgHandle] worker id %v started\n", workerID)
	//不断的等待队列中的消息
	for {
		select {
		//有消息则取出队列的Request，并执行绑定的业务方法
		case request := <-taskQueue:
			m.DoMsgHandler(request)
		}
	}
}
