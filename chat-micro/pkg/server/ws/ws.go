package ws

import (
	"context"
	"fmt"
	"math"
	"net"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"chat-micro/pkg/logger"
	"chat-micro/pkg/server"
)

type wsServer struct {
	*server.Server
}

// Start 开启网络服务
func (w *wsServer) Start(ctx context.Context) error {
	opts := w.Options()

	ln, err := net.Listen("tcp", opts.Address)
	if err != nil {
		return err
	}
	w.SetListener(ln)

	var cid uint32 = 1
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		logger.Info("[ws.accept] start: ", r.RemoteAddr)
		//设置服务器最大连接控制,如果超过最大连接，那么则关闭此新的连接
		if w.Bucket.ConnTotal() >= opts.MaxConn {
			logger.Warn("[ws.start] connection size limit")
			return
		}

		conn := NewConnect(w, rw, r, cid)
		if cid++; cid == math.MaxUint32 {
			logger.Infof("[ws.server] conn.acceptTcp num is:%d", cid)
			cid = 1
		}
		conn.Start()
	})
	// 启动worker工作池机制
	w.MsgHandler.StartWorkerPool(opts.MaxWorkerTaskLen)

	go http.Serve(ln, nil)
	fmt.Println("Websocket Server start success, now listening...")

	return w.RegisterKeep()
}

//String implementation
func (w *wsServer) String() string {
	return "websocket"
}

func newServer(opts ...server.Option) server.IServer {
	option := server.NewOptions(opts...)
	return &wsServer{
		&server.Server{
			Opts:       option,
			Bucket:     server.NewBucket(option.BucketSize),
			MsgHandler: server.NewMsgHandle(option.WorkerPoolSize, option.Router),
		},
	}
}

//NewServer 实例化tcp服务器
func NewServer(opts ...server.Option) server.IServer {
	return newServer(opts...)
}
