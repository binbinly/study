package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
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

	c *ServerConfig //配置
}

//NewServer 创建服务器
func NewServer(c *ServerConfig) *Server {
	s := &Server{
		Server: &http.Server{
			Addr:         fmt.Sprintf(":%d", c.Port),
			ReadTimeout:  c.ReadTimeout,
			WriteTimeout: c.WriteTimeout,
		},
		c: c,
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
}

//Stop 停止服务
func (s *Server) Stop() error {
	return s.Shutdown(context.Background())
}
