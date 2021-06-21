package consul

import (
	"context"
	"time"

	"github.com/hashicorp/consul/api"
	"github.com/pkg/errors"

	"chat/pkg/registry"
)

const (
	//MaxSyncServiceInterval 健康检查间隔
	MaxSyncServiceInterval = time.Second * 15
	//Deregister 服务自动注销时间
	Deregister = 24 * time.Hour
)

var consul *Registry

//Registry 注册插件
type Registry struct {
	options *registry.Options
	client  *api.Client
	service *registry.Service
}

//NewConsul 实例化服务注册中心
func NewConsul() *Registry {
	consul = &Registry{}
	return consul
}

//Name 插件的名字
func (e *Registry) Name() string {
	return "consul"
}

//Init 初始化
func (e *Registry) Init(ctx context.Context, opts ...registry.Option) (err error) {

	e.options = &registry.Options{}
	for _, opt := range opts {
		opt(e.options)
	}

	config := api.DefaultConfig()
	config.Address = e.options.Addr[0]
	e.client, err = api.NewClient(config)

	return err
}

//Register 服务注册
func (e *Registry) Register(ctx context.Context, service *registry.Service) (err error) {
	reg := &api.AgentServiceRegistration{
		ID:      service.ID,
		Name:    service.Name,
		Tags:    []string{service.Name},
		Port:    service.Port,
		Address: service.IP,
		Check: &api.AgentServiceCheck{ // 健康检查
			Interval:                       MaxSyncServiceInterval.String(), // 健康检查间隔
			DeregisterCriticalServiceAfter: Deregister.String(),             // 注销时间，相当于过期时间
		},
	}
	if service.Check.GRPC != "" {
		reg.Check.GRPC = service.Check.GRPC
	} else if service.Check.HTTP != "" {
		reg.Check.HTTP = service.Check.HTTP
	} else if service.Check.TCP != "" {
		reg.Check.TCP = service.Check.TCP
	}
	if err = e.client.Agent().ServiceRegister(reg); err != nil {
		return errors.Wrapf(err, "consul service register fail: %s", err.Error())
	}
	return
}

//Unregister 服务反注册
func (e *Registry) Unregister(ctx context.Context, service *registry.Service) (err error) {
	return e.client.Agent().ServiceDeregister(service.ID)
}

//Find 服务发现
func (e *Registry) Find(ctx context.Context, name string) (service *registry.Service, err error) {
	return
}
