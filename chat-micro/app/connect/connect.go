package connect

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/google/uuid"

	"chat-micro/app/constvar"
	"chat-micro/internal/grpc"
	"chat-micro/pkg/logger"
	"chat-micro/pkg/registry"
	"chat-micro/pkg/registry/consul"
	"chat-micro/pkg/server"
	"chat-micro/pkg/transport"
	pb "chat-micro/proto/logic"
)

// Connect global app
type Connect struct {
	opts   options
	ctx    context.Context
	cancel func()

	logic pb.LogicClient
	svc   *registry.Service
}

// New create a app globally
func New(opts ...Option) *Connect {
	o := options{
		ctx:              context.Background(),
		name:             constvar.ServiceConnect,
		version:          "latest",
		registerTTL:      registerTTL,
		registerInterval: registerInterval,
	}
	if id, err := uuid.NewUUID(); err == nil {
		o.id = id.String()
	}
	for _, opt := range opts {
		opt(&o)
	}
	ctx, cancel := context.WithCancel(o.ctx)

	// 初始化consul.resolver
	consul.Init(o.registry)
	target := "consul:///" + constvar.ServiceLogic
	return &Connect{
		opts:   o,
		ctx:    ctx,
		cancel: cancel,
		logic:  pb.NewLogicClient(grpc.NewClientConn(ctx, target)),
	}
}

//Init 初始化
func (c *Connect) Init(opts ...Option) {
	for _, o := range opts {
		o(&c.opts)
	}
}

//ServerID 服务器id
func (c *Connect) ServerID() string {
	return c.opts.id
}

//Run 启动服务
func (c *Connect) Run() error {
	// start server
	servers := []transport.Server{c.opts.transport, c.opts.server}
	for _, srv := range servers {
		s := srv
		go func() {
			<-c.ctx.Done()
			if err := s.Stop(c.ctx); err != nil {
				log.Printf("Server stop err: %v", err)
			}
		}()
		go func() {
			if err := s.Start(c.ctx); err != nil {
				log.Fatalf("Server Failed to listen and serve: %v", err)
			}
		}()
	}

	//register service
	if c.opts.registry != nil {
		var address string
		if r, ok := c.opts.transport.(transport.Endpoint); ok {
			e, err := r.Endpoint()
			if err != nil {
				return err
			}
			address, err = parseAddress(e.String())
			if err != nil {
				return err
			}
		}

		node := &registry.Node{
			Id:      c.opts.id,
			Address: address,
		}
		service := &registry.Service{
			Name:     c.opts.name,
			Version:  c.opts.version,
			Metadata: c.opts.metadata,
			Nodes:    []*registry.Node{node},
		}
		if err := c.opts.registry.Register(service); err != nil {
			return err
		}
		c.svc = service

		go func() {
			t := time.NewTicker(c.opts.registerInterval)

			for {
				select {
				case <-t.C:
					if err := c.opts.registry.Register(service); err != nil {
						logger.Warn("Server register error: ", err)
					}
				}
			}
		}()
	}
	return nil
}

//Server get Server
func (c *Connect) Server() server.IServer {
	return c.opts.server
}

// Stop stops the application gracefully.
func (c *Connect) Stop() error {
	if c.opts.registry != nil {
		if err := c.opts.registry.Deregister(c.svc); err != nil {
			return err
		}
	}
	// cancel app
	if c.cancel != nil {
		c.cancel()
	}
	return nil
}

//parseAddress 解析 address
func parseAddress(endpoint string) (string, error) {
	raw, err := url.Parse(endpoint)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s:%s", raw.Hostname(), raw.Port()), nil
}
