package grpc

import (
	"errors"
	"sync"
	"time"

	"chat/proto/logic"
)

//路由处理方法
type HandlerFunc func(c *Context) (*logic.ReceiveReply, error)

// implement http.Handler
type Engine struct {
	tree map[string]HandlerFunc
	pool sync.Pool
}

type Context struct {
	Req     *logic.ReceiveReq
	handler HandlerFunc
}

// get a new engine(tcp.Handler)
func NewEngine() *Engine {
	engine := &Engine{
		tree: make(map[string]HandlerFunc),
	}
	engine.pool.New = func() interface{} {
		return engine.allocateContext()
	}
	return engine
}

// Deadline returns the time when work done on behalf of this context
// should be canceled. Deadline returns ok==false when no deadline is
// set. Successive calls to Deadline return the same results.
func (c *Context) Deadline() (deadline time.Time, ok bool) {
	return
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
	return c.Req
}

func (e *Engine) allocateContext() *Context {
	return &Context{}
}

//执行入口
func (e *Engine) Start(req *logic.ReceiveReq) (*logic.ReceiveReply, error) {
	c := e.pool.Get().(*Context)
	c.Req = req
	handler := e.getHandler(c.Req.GetEvent())
	if handler != nil {
		return handler(c)
	}
	return nil, errors.New("event not found")
}

func (e *Engine) AddRoute(event string, handler HandlerFunc) {
	e.tree[event] = handler
}

func (e *Engine) getHandler(event string) (handler HandlerFunc) {
	handler, ok := e.tree[event]
	if !ok {
		return nil
	}
	return
}
