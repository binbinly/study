package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"chat/pkg/app"
	"chat/pkg/registry"
)

//ServerConfig 服务器配置
type ServerConfig struct {
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

//Server 服务器结构
type Server struct {
	*http.Server //http服务器

	c       *ServerConfig //配置
	options *app.Options
}

//NewServer 创建服务器
func NewServer(c *ServerConfig, opts ...app.Option) *Server {
	s := &Server{
		Server: &http.Server{
			Addr:         fmt.Sprintf(":%d", c.Port),
			ReadTimeout:  c.ReadTimeout,
			WriteTimeout: c.WriteTimeout,
		},
		c:       c,
		options: &app.Options{},
	}
	for _, o := range opts {
		o(s.options)
	}
	return s
}

// Start start http server
func (s *Server) Start() {
	log.Printf("Listening and serving HTTP on %d\n", s.c.Port)
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe, err: %s", err.Error())
		}
	}()
	s.register()
}

//Stop 停止服务
func (s *Server) Stop() error {
	if s.options.Register != nil {
		s.options.Register.Unregister(context.Background(), &registry.Service{ID: app.BuildServerID(s.options.ID, s.c.Port)})
	}
	return s.Shutdown(context.Background())
}

//Register 服务注册
func (s *Server) register() {
	if s.options.Register == nil {
		return
	}
	// 服务注册
	err := s.options.Register.Register(context.Background(), &registry.Service{
		ID:   app.BuildServerID(s.options.ID, s.c.Port),
		Name: s.options.Name,
		IP:   s.options.Host,
		Port: s.c.Port,
		Check: registry.Check{
			HTTP: fmt.Sprintf("http://%v:%v/health", s.options.Host, s.c.Port),
		},
	})
	if err != nil {
		log.Fatalf("failed to http register %s server: %v", s.options.Name, err)
	}
	log.Printf("server http register success id:%v, name:%v", app.BuildServerID(s.options.ID, s.c.Port), s.options.Name)
}
