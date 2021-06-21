package registry

import (
	"context"
	"fmt"
	"sync"
)

var manager = &PluginsManager{
	plugins: make(map[string]Registry),
}

//PluginsManager 插件管理器
type PluginsManager struct {
	plugins map[string]Registry
	lock    sync.Mutex
}

//注册插件
func (p *PluginsManager) registerPlugin(plugin Registry) {
	p.lock.Lock()
	defer p.lock.Unlock()

	_, ok := p.plugins[plugin.Name()]
	if ok {
		return
	}

	p.plugins[plugin.Name()] = plugin
	return
}

func (p *PluginsManager) initRegistry(ctx context.Context, name string, opts ...Option) (registry Registry, err error) {
	//查找对应的插件是否存在
	p.lock.Lock()
	defer p.lock.Unlock()

	plugin, ok := p.plugins[name]
	if !ok {
		err = fmt.Errorf("plugin %s not exitsts", name)
		return
	}
	registry = plugin
	err = plugin.Init(ctx, opts...)
	return
}

//RegisterPlugin 注册插件
func RegisterPlugin(registry Registry) {
	manager.registerPlugin(registry)
}

//InitRegistry 初始化注册中心
func InitRegistry(ctx context.Context, name string, opts ...Option) (registry Registry, err error) {
	return manager.initRegistry(ctx, name, opts...)
}
