package consul

import (
	"fmt"
	"sync"
	"time"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc/resolver"

	"chat/pkg/log"
)

//Init 返回一个resolver.Builder的实例
func Init() {
	resolver.Register(NewBuilder())
}

type consulResolver struct {
	wg                   sync.WaitGroup
	cc                   resolver.ClientConn
	name                 string
	disableServiceConfig bool
	lastIndex            uint64
}

// watcher 更新服务变化
func (cr *consulResolver) watcher() {
	defer cr.wg.Done()

	for {
		services, metaInfo, err := consul.client.Health().Service(cr.name, "", true, &api.QueryOptions{WaitIndex: cr.lastIndex})
		if err != nil {
			log.Warnf("[registry.consul] watcher services name:%v, err:%v", cr.name, err)
			time.Sleep(time.Second)
			continue
		}
		cr.lastIndex = metaInfo.LastIndex
		if len(services) == 0 {
			time.Sleep(time.Second)
			continue
		}
		var adds []resolver.Address

		for _, v := range services {
			adds = append(adds, resolver.Address{Addr: fmt.Sprintf("%v:%v", v.Service.Address, v.Service.Port)})
		}

		state := &resolver.State{
			Addresses: adds,
		}
		cr.cc.UpdateState(*state)
	}
}

//ResolverNow方法什么也不做，因为和consul保持了发布订阅的关系
//不需要像dns_resolver那个定时的去刷新
func (cr *consulResolver) ResolveNow(opt resolver.ResolveNowOptions) {}

func (cr *consulResolver) Close() {}
