package server

import (
	"context"
	"errors"
	"log"
	"net"
	"net/url"
	"sync"
	"time"

	"chat-micro/pkg/logger"
	"chat-micro/pkg/registry"
	"chat-micro/pkg/util"
)

var (
	//ErrConnNotFound 连接未找到
	ErrConnNotFound = errors.New("connection not found")
	//ErrConnectNotFinish 连接未完成，不可以发送消息
	ErrConnectNotFinish = errors.New("connection not finish when send msg")
)

// IServer is a simple micro server abstraction
type IServer interface {
	// Initialise options
	Init(...Option) error
	// Retrieve the options
	Options() Options
	// Start the server
	Start(ctx context.Context) error
	// Stop the server
	Stop(ctx context.Context) error
	// Server implementation
	String() string
	// Endpoint return a real address to registry endpoint.
	Endpoint() (*url.URL, error)
	// Bucket 所有连接管理
	GetBucket() IBucket
}

//Server 基础服务
type Server struct {
	//连接管理器
	Bucket IBucket
	//当前Server的消息管理模块，用来绑定MsgID和对应的处理方法
	MsgHandler IMsgHandle
	Opts       Options

	//监听器
	listener net.Listener
	sync.Mutex
}

// Options 服务选项
func (s *Server) Options() Options {
	s.Lock()
	opts := s.Opts
	s.Unlock()
	return opts
}

//Init 初始化
func (s *Server) Init(opts ...Option) error {
	s.Lock()
	for _, o := range opts {
		o(&s.Opts)
	}
	s.Unlock()
	return nil
}

// Endpoint return a real address to registry endpoint.
func (s *Server) Endpoint() (*url.URL, error) {
	addr, err := util.Extract(s.Opts.Address, s.listener)
	if err != nil {
		return nil, err
	}
	return &url.URL{Scheme: "http", Host: addr}, nil
}

//RegisterKeep 保持服务注册活性
func (s *Server) RegisterKeep() error {
	// register
	if err := s.Register(); err != nil {
		return err
	}

	go func() {
		tk := time.NewTicker(s.Opts.RegisterInterval)

		for {
			select {
			// register self on interval
			case <-tk.C:
				if err := s.Register(); err != nil {
					logger.Error("Server register error: ", err)
				}
			}
		}
	}()

	return nil
}

//SetListener 设置监听器
func (s *Server) SetListener(l net.Listener) {
	log.Printf("Listening on %s", l.Addr().String())
	s.Lock()
	s.Opts.Address = l.Addr().String()
	s.Unlock()
	s.listener = l
}

//Stop 关闭服务器
func (s *Server) Stop(ctx context.Context) error {
	s.Bucket.Clear()

	s.Deregister()
	return s.listener.Close()
}

//Register 服务注册
func (s *Server) Register() error {
	if s.Opts.Registry == nil {
		return nil
	}

	s.Lock()
	opts := s.Opts
	s.Unlock()

	service := serviceDef(opts)
	if err := s.Opts.Registry.Register(service, registry.RegisterTTL(opts.RegisterTTL)); err != nil {
		return err
	}
	logger.Infof("[TCP] Registering node: %s", opts.Name+"-"+opts.Id)

	return nil
}

//Deregister 服务注销
func (s *Server) Deregister() error {
	if s.Opts.Registry == nil {
		return nil
	}

	s.Lock()
	opts := s.Opts
	s.Unlock()

	service := serviceDef(opts)
	if err := opts.Registry.Deregister(service); err != nil {
		return err
	}

	logger.Infof("[TCP] Deregistering node: %s", opts.Name+"-"+opts.Id)
	return nil
}

// GetBucket 获取所有连接桶
func (s *Server) GetBucket() IBucket {
	return s.Bucket
}
