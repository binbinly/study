package app

import (
	"fmt"

	"chat/pkg/registry"
)

//Options 选项结构
type Options struct {
	ID       string
	Name     string
	Host     string
	Register registry.Registry //服务注册中心
}

//Option 选项回调
type Option func(s *Options)

//WithID 设置应用id
func WithID(id string) Option {
	return func(s *Options) {
		s.ID = id
	}
}

//WithName 设置应用名
func WithName(name string) Option {
	return func(s *Options) {
		s.Name = name
	}
}

//WithHost 设置绑定的ip
func WithHost(host string) Option {
	return func(s *Options) {
		s.Host = host
	}
}

//WithRegistry 设置注册中心
func WithRegistry(rs registry.Registry) Option {
	return func(s *Options) {
		s.Register = rs
	}
}

//BuildServerID 生成服务ID
func BuildServerID(ID string, port int) string {
	return fmt.Sprintf("%s%d", ID, port)
}
