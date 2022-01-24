package tcp

import (
	"context"
	"fmt"
	"math"
	"net"

	"chat-micro/pkg/logger"
	"chat-micro/pkg/server"
)

type tcpServer struct {
	*server.Server

	network string
}

// Start 开启网络服务
func (t *tcpServer) Start(ctx context.Context) error {
	opts := t.Options()
	fmt.Printf("[TCP] Server listener at Addr:%v is starting\n", opts.Address)

	// 启动worker工作池机制
	t.MsgHandler.StartWorkerPool(opts.MaxWorkerTaskLen)

	//1 获取一个TCP的Addr
	addr, err := net.ResolveTCPAddr(t.network, opts.Address)
	if err != nil {
		return err
	}

	// 监听服务器地址
	lis, err := net.ListenTCP(t.network, addr)
	if err != nil {
		return err
	}
	t.SetListener(lis)

	//已经监听成功
	fmt.Println("TCP Server start success, now listening...")

	go t.accept(lis)

	return t.RegisterKeep()
}

//String implementation
func (t *tcpServer) String() string {
	return "tcp"
}

//accept 监听客户端连接
func (t *tcpServer) accept(lis *net.TCPListener) {
	var (
		conn *net.TCPConn
		err  error
		cid  uint32
	)
	cid = 1
	// 启动server网络连接业务
	for {
		//3.1 阻塞等待客户端建立连接请求
		if conn, err = lis.AcceptTCP(); err != nil {
			logger.Warnf("[tcp.server] Accept err:%v", err)
			continue
		}
		opts := t.Options()
		// 设置接收消息缓冲区大小
		if err = conn.SetReadBuffer(opts.ReadBufferSize); err != nil {
			logger.Warnf("[tcp.server] conn.SetReadBuffer() err:%v", err)
			continue
		}
		// 设置发送消息缓冲区大小
		if err = conn.SetWriteBuffer(opts.WriteBufferSize); err != nil {
			logger.Warnf("[tcp.server] conn.SetWriteBuffer() err:%v", err)
			continue
		}
		if opts.Context != nil {
			if v, ok := opts.Context.Value(keepaliveKey{}).(bool); ok && v != true {
				// 设置客户端服务器间的活动性
				if err = conn.SetKeepAlive(true); err != nil {
					logger.Warnf("[tcp.server] conn.SetKeepAlive() err:%v", err)
					continue
				}
			}
		}
		logger.Debug("[tcp.server] Get conn remote addr = ", conn.RemoteAddr().String())

		//3.2 设置服务器最大连接控制,如果超过最大连接，那么则关闭此新的连接
		if t.Bucket.ConnTotal() >= opts.MaxConn {
			logger.Warnf("[tcp.server] server connect limited addr: %v", conn.RemoteAddr().String())
			if err = conn.Close(); err != nil {
				logger.Warnf("[tcp.server] conn closed addr: %v err: %v", conn.RemoteAddr().String(), err)
			}
			continue
		}

		//3.3 处理该新连接请求的 业务 方法， 此时应该有 handler 和 conn是绑定的
		dealConn := NewConnect(t, conn, cid)
		if cid++; cid == math.MaxUint32 {
			logger.Infof("[tcp.server] conn.acceptTcp num is:%d", cid)
			cid = 1
		}

		//3.4 启动当前链接的处理业务
		go dealConn.Start()
	}
}

func newServer(opts ...server.Option) server.IServer {
	option := server.NewOptions(opts...)
	return &tcpServer{
		Server: &server.Server{
			Opts:       option,
			Bucket:     server.NewBucket(option.BucketSize),
			MsgHandler: server.NewMsgHandle(option.WorkerPoolSize, option.Router),
		},
		network: "tcp4",
	}
}

//NewServer 实例化tcp服务器
func NewServer(opts ...server.Option) server.IServer {
	return newServer(opts...)
}
