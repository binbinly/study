package app

import (
	"github.com/asim/go-micro/plugins/client/grpc/v4"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/selector"
)

//NewTestClient 实例化测试客户端
func NewTestClient(registryAddress string) client.Client {
	// create a new service
	service := micro.NewService(
		micro.Client(grpc.NewClient()),
		micro.Selector(selector.NewSelector(selector.Registry(
			consul.NewRegistry(registry.Addrs(registryAddress)),
		))))
	return service.Client()
}
