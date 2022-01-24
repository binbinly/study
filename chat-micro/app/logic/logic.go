package logic

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/google/uuid"

	"chat-micro/pkg/logger"
	"chat-micro/pkg/registry"
	"chat-micro/pkg/transport"
)

// App global app
type App struct {
	opts   options
	ctx    context.Context
	cancel func()

	svc    *registry.Service
}

// New create a app globally
func New(opts ...Option) *App {
	o := options{
		ctx:              context.Background(),
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
	return &App{
		opts:   o,
		ctx:    ctx,
		cancel: cancel,
	}
}

//Run 启动api网关服务
func (a *App) Run() error {
	// start server
	for _, srv := range a.opts.servers {
		s := srv
		go func() {
			<-a.ctx.Done()
			if err := s.Stop(a.ctx); err != nil {
				log.Printf("Server stop err: %v", err)
			}
		}()
		go func() {
			if err := s.Start(a.ctx); err != nil {
				log.Fatalf("Server Failed to listen and serve: %v", err)
			}
		}()
	}
	// register service
	if a.opts.registry != nil {
		var address string
		if r, ok := a.opts.servers[0].(transport.Endpoint); ok {
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
			Id:      a.opts.id,
			Address: address,
		}
		service := &registry.Service{
			Name:     a.opts.name,
			Version:  a.opts.version,
			Metadata: a.opts.metadata,
			Nodes:    []*registry.Node{node},
		}
		if err := a.opts.registry.Register(service); err != nil {
			return err
		}
		a.svc = service

		go func() {
			t := time.NewTicker(a.opts.registerInterval)

			for {
				select {
				case <-t.C:
					if err := a.opts.registry.Register(service); err != nil {
						logger.Warn("Server register error: ", err)
					}
				}
			}
		}()
	}
	return nil
}

// Stop stops the application gracefully.
func (a *App) Stop() error {
	if a.opts.registry != nil {
		if err := a.opts.registry.Deregister(a.svc); err != nil {
			return err
		}
	}
	// cancel app
	if a.cancel != nil {
		a.cancel()
	}
	return nil
}

//Ctx get ctx
func (a *App) Ctx() context.Context {
	return a.ctx
}

//parseAddress 解析 address
func parseAddress(endpoint string) (string, error) {
	raw, err := url.Parse(endpoint)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s:%s", raw.Hostname(), raw.Port()), nil
}
