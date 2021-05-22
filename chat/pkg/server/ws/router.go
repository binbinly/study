package ws

import (
	"math"
	"sync"
)

const abortIndex int8 = math.MaxInt8 / 2 // 路由中间件嵌套最大值

type IEngine interface {
	// 获取当前上下文
	allocateContext() *Context
	// 执行入口
	Start(req *Request)
	// 添加路由
	addRoute(event string, handlers HandlerChain)
	// 添加中间件
	Use(middleware ...HandlerFunc)
}

//路由处理方法
type HandlerFunc func(c *Context)

type HandlerChain []HandlerFunc

//路由组
type RouterGroup struct {
	Handlers HandlerChain
	engine   IEngine
}

// implement http.Handler
type Engine struct {
	tree map[string]HandlerChain
	RouterGroup
	pool sync.Pool
}

type Context struct {
	Req      *Request
	handlers HandlerChain
	index    int8
}

// get a new engine(tcp.Handler)
func NewEngine() *Engine {
	engine := &Engine{
		RouterGroup: RouterGroup{
			Handlers: nil,
		},
		tree: make(map[string]HandlerChain),
	}
	engine.RouterGroup.engine = engine
	engine.pool.New = func() interface{} {
		return engine.allocateContext()
	}
	return engine
}

func (c *Context) Next() {
	c.index++
	// 调用 c.Abort() 的时候不会往后执行
	for c.index < int8(len(c.handlers)) {
		c.handlers[c.index](c)
		c.index++
	}
}

//退出路由
func (c *Context) Abort() {
	c.index = abortIndex
}

func (c *Context) Reset() {
	c.handlers = nil
	c.index = -1
}

// Done always returns nil (chan which will wait forever),
// if you want to abort your work when the connection was closed
// you should use Request.Context().Done() instead.
func (c *Context) Done() <-chan struct{} {
	return nil
}

// Err always returns nil, maybe you want to use Request.Context().Err() instead.
func (c *Context) Err() error {
	return nil
}

// Value returns the value associated with this context for key, or nil
// if no value is associated with key. Successive calls to Value with
// the same key returns the same result.
func (c *Context) Value(key interface{}) interface{} {
	if key == 0 {
		return c.Req
	}
	return c.Req.GetData()
}

func (engine *Engine) allocateContext() *Context {
	return &Context{}
}

//执行入口
func (engine *Engine) Start(req *Request) {
	c := engine.pool.Get().(*Context)
	c.Req = req
	c.Reset()
	handlers := engine.getHandlers(c.Req.GetEvent())
	if handlers != nil {
		c.handlers = handlers
		c.Next()
	}
}

func (engine *Engine) addRoute(event string, handlers HandlerChain) {
	engine.tree[event] = handlers
}

func (engine *Engine) getHandlers(event string) (handlers HandlerChain) {
	handlers, ok := engine.tree[event]
	if !ok {
		return nil
	}
	return
}

// set common middleware
func (engine *Engine) Use(middleware ...HandlerFunc) {
	engine.RouterGroup.Use(middleware...)
}

func (group *RouterGroup) Use(middleware ...HandlerFunc) {
	group.Handlers = append(group.Handlers, middleware...)
}

// specific middleware
func (group *RouterGroup) AddRoute(event string, handlers ...HandlerFunc) {
	handlers = group.mergeHandlers(handlers)
	group.engine.addRoute(event, handlers)
}

// merge specific and common middleware
func (group *RouterGroup) mergeHandlers(handlers HandlerChain) HandlerChain {
	finalSize := len(group.Handlers) + len(handlers)
	mergedHandlers := make(HandlerChain, finalSize)
	copy(mergedHandlers, group.Handlers)
	copy(mergedHandlers[len(group.Handlers):], handlers)
	return mergedHandlers
}
