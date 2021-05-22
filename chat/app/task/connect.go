package task

import (
	"chat/app/task/conf"
	"chat/pkg/log"
	"chat/pkg/server/grpc"
	"chat/proto/connect"
	"context"
	"sync/atomic"
)

// Connect
type Connect struct {
	serverId      string                     //连接层服务器id
	client        connect.ConnectClient      //调用连接层client
	sendChan      []chan *connect.SendReq    //推送至连接层队列
	broadcastChan chan *connect.BroadcastReq //广播消息队列
	pushChanNum   uint64
	routineSize   uint64

	ctx    context.Context
	cancel context.CancelFunc
}

// NewConnect new a connect.
func NewConnect(c *conf.Config, id, addr string) (*Connect, error) {
	ct := &Connect{
		serverId:      id,
		sendChan:      make([]chan *connect.SendReq, c.Connect.RoutineSize),
		broadcastChan: make(chan *connect.BroadcastReq, c.Connect.RoutineSize),
		routineSize:   uint64(c.Connect.RoutineSize),
	}
	var err error
	if ct.client, err = newConnectClient(addr, &c.GrpcClient); err != nil {
		return nil, err
	}
	ct.ctx, ct.cancel = context.WithCancel(context.Background())

	for i := 0; i < c.Connect.RoutineSize; i++ {
		ct.sendChan[i] = make(chan *connect.SendReq, c.Connect.RoutineChan)
		go ct.run(ct.sendChan[i], ct.broadcastChan)
	}
	return ct, nil
}

//Push 推送一条消息
func (c *Connect) Push(arg *connect.SendReq) (err error) {
	idx := atomic.AddUint64(&c.pushChanNum, 1) % c.routineSize
	c.sendChan[idx] <- arg
	return
}

//Broadcast 广播消息
func (c *Connect) Broadcast(arg *connect.BroadcastReq) (err error) {
	c.broadcastChan <- arg
	return
}

//Close 关闭连接
func (c *Connect) Close(arg *connect.CloseReq) (err error) {
	_, err = c.client.Close(context.Background(), arg)
	if err != nil {
		log.Warnf("[connect.close] arg:%v, id:%v, err:%v", arg, c.serverId, err)
	}
	return nil
}

func (c *Connect) run(sendChan chan *connect.SendReq, broadcastChan chan *connect.BroadcastReq) {
	for {
		select {
		case arg := <-broadcastChan:
			_, err := c.client.Broadcast(context.Background(), arg)
			if err != nil {
				log.Warnf("[connect.broadcast] arg:%v, id:%v, err:%v", arg, c.serverId, err)
			}
		case arg := <-sendChan:
			_, err := c.client.Send(context.Background(), arg)
			if err != nil {
				log.Warnf("[connect.send] arg:%v, id:%v, err:%v", arg, c.serverId, err)
			}
		case <-c.ctx.Done():
			return
		}
	}
}

//newConnectClient 创建连接层客户端
func newConnectClient(addr string, c *grpc.ClientConfig) (connect.ConnectClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()
	conn, err := grpc.NewRpcClientConn(c, ctx, addr)
	if err != nil {
		return nil, err
	}
	return connect.NewConnectClient(conn), err
}
