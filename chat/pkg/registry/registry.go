package registry

import (
	"context"
)

//Registry 注册器
type Registry interface {
	//插件名字
	Name() string
	//初始化
	Init(ctx context.Context, opts ...Option) (err error)
	//服务注册
	Register(ctx context.Context, service *Service) (err error)
	//服务反注册
	Unregister(ctx context.Context, service *Service) (err error)
	//服务发现：通过服务的名字获取服务的位置信息（ip和port列表）
	Find(ctx context.Context, name string) (service *Service, err error)
}

//Config 服务注册中心
type Config struct {
	Name string
	Host string
}
