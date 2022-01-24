package task

import (
	"context"
	"sync/atomic"

	"github.com/pkg/errors"

	"chat-micro/internal/grpc"
	"chat-micro/pkg/logger"
	pb "chat-micro/proto/connect"
)

// Connect 连接层结构
type Connect struct {
	serverID      string                //连接层服务器id
	client        pb.ConnectClient      //调用连接层client
	sendChan      []chan *pb.SendReq    //推送至连接层队列
	broadcastChan chan *pb.BroadcastReq //广播消息队列
	sendIndex     uint64
	routineNum    int

	ctx    context.Context
	cancel context.CancelFunc
}

// NewConnect new a connect
func NewConnect(opts options, id, addr string) (*Connect, error) {
	ctx, cancel := context.WithCancel(context.Background())

	conn := grpc.NewClientConn(ctx, addr)
	c := &Connect{
		ctx:           ctx,
		cancel:        cancel,
		serverID:      id,
		routineNum:    opts.routineNum,
		sendChan:      make([]chan *pb.SendReq, opts.routineNum),
		broadcastChan: make(chan *pb.BroadcastReq, opts.routineSize),
		client:        pb.NewConnectClient(conn),
	}

	for i := 0; i < opts.routineNum; i++ {
		c.sendChan[i] = make(chan *pb.SendReq, opts.routineSize)
		go c.run(c.sendChan[i], c.broadcastChan)
	}
	return c, nil
}

//Push 推送一条消息
func (c *Connect) Push(arg *pb.SendReq) (err error) {
	idx := atomic.AddUint64(&c.sendIndex, 1) % uint64(c.routineNum)
	c.sendChan[idx] <- arg
	return
}

//Broadcast 广播消息
func (c *Connect) Broadcast(arg *pb.BroadcastReq) (err error) {
	c.broadcastChan <- arg
	return
}

//Close 关闭连接
func (c *Connect) Close(ctx context.Context, arg *pb.CloseReq) (err error) {
	_, err = c.client.Close(ctx, arg)
	if err != nil {
		return errors.Wrapf(err, "[task.connect] close serverId: %v, arg: %v", c.serverID, arg)
	}
	return nil
}

func (c *Connect) run(sendChan chan *pb.SendReq, broadcastChan chan *pb.BroadcastReq) {
	for {
		select {
		case arg := <-broadcastChan:
			_, err := c.client.Broadcast(context.Background(), arg)
			if err != nil {
				logger.Warnf("[task.connect] broadcast arg:%v, serverId:%v, err:%v", arg, c.serverID, err)
			}
		case arg := <-sendChan:
			_, err := c.client.Send(context.Background(), arg)
			if err != nil {
				logger.Warnf("[task.connect] send arg:%v, serverId:%v, err:%v", arg, c.serverID, err)
			}
		case <-c.ctx.Done():
			logger.Infof("[task.connect] serverId: %v close", c.serverID)
			return
		}
	}
}
