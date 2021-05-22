package server

import (
	"context"
	"net"
	"net/http"
	"sync"
)

var _ IConnection = (*Connection)(nil)

// 定义连接接口
type IConnection interface {
	Start()                    //启动连接，让当前连接开始工作
	Stop()                     //停止连接，结束当前连接状态M
	Context() context.Context  //返回ctx，用于用户自定义的go程获取连接退出状态
	Auth(userId uint32) //连接鉴权

	GetConnID() uint32         //获取当前连接ID
	GetUid() uint32         //获取当前连接绑定的用户id
	RemoteAddr() net.Addr      //获取远程客户端地址信息
	GetRequest() *http.Request //获取连接请求对象
	GetContext() context.Context //获取上下文

	SetProperty(key string, value interface{})   //设置链接属性
	GetProperty(key string) (interface{}, error) //获取链接属性
	RemoveProperty(key string)                   //移除链接属性
}

//Connection 连接
type Connection struct {
	//当前连接的ID 也可以称作为SessionID，ID全局唯一
	ConnID uint32
	//用户ID
	Uid uint32
	//告知该链接已经退出/停止的channel
	Ctx    context.Context
	Cancel context.CancelFunc
	//无缓冲管道，用于读、写两个goroutine之间的消息通信
	MsgChan chan []byte
	//有缓冲管道，用于读、写两个goroutine之间的消息通信
	MsgBuffChan chan []byte

	sync.RWMutex
	//链接属性
	property map[string]interface{}
	//保护当前property的锁
	propertyLock sync.Mutex
}

//Start 启动连接，让当前连接开始工作
func (c *Connection) Start() {}

//Stop 停止连接，结束当前连接状态M
func (c *Connection) Stop() {}

//Auth 连接鉴权
func (c *Connection) Auth(userId uint32) {}

//返回ctx，用于用户自定义的go程获取连接退出状态
func (c *Connection) Context() context.Context {
	return c.Ctx
}

//GetConnID 获取当前连接ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

//GetUid 获取当前连接绑定的用户id
func (c *Connection) GetUid() uint32 {
	return c.Uid
}

//GetContext 获取上下文
func (c *Connection) GetContext() context.Context {
	return c.Ctx
}

//RemoteAddr 获取远程客户端地址信息
func (c *Connection) RemoteAddr() net.Addr { return nil }

//GetRequest 获取连接的请求对象
func (c *Connection) GetRequest() *http.Request { return nil }

//SetProperty 设置链接属性
func (c *Connection) SetProperty(key string, value interface{}) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()
	if c.property == nil {
		c.property = make(map[string]interface{})
	}
	c.property[key] = value
}

//GetProperty 获取链接属性
func (c *Connection) GetProperty(key string) (interface{}, error) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	if value, ok := c.property[key]; ok {
		return value, nil
	}
	return nil, ErrPropertyNotFound
}

//RemoveProperty 移除链接属性
func (c *Connection) RemoveProperty(key string) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	delete(c.property, key)
}
