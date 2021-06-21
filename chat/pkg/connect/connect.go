package connect

import (
	"context"
	"net"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

//IConnection 定义连接接口
type IConnection interface {
	Start()                   //启动连接，让当前连接开始工作
	Stop()                    //停止连接，结束当前连接状态M
	Context() context.Context //返回ctx，用于用户自定义的go程获取连接退出状态
	Auth(userID uint32)       //连接鉴权

	GetConnID() uint32    //获取当前连接ID
	GetUID() uint32       //获取当前连接绑定的用户id
	RemoteAddr() net.Addr //获取远程客户端地址信息

	SendMsg(msgID uint32, data []byte) error     //直接将Message数据发送至远程客户端
	SendBuffMsg(msgID uint32, data []byte) error //发送至缓冲区

	GetHTTPRequest() *http.Request    //获取http请求对象
	GetTCPConnection() *net.TCPConn   //从当前连接获取原始的socket
	GetWsConnection() *websocket.Conn //获取websocket连接对象

	SetProperty(key string, value interface{})   //设置连接属性
	GetProperty(key string) (interface{}, error) //获取连接属性
	RemoveProperty(key string)                   //移除连接属性
}

//Connection 连接
type Connection struct {
	//当前连接的ID 也可以称作为SessionID，ID全局唯一
	ConnID uint32
	//用户ID
	UID uint32
	//告知该链接已经退出/停止的channel
	Ctx    context.Context
	Cancel context.CancelFunc
	//无缓冲管道，用于读、写两个goroutine之间的消息通信
	MsgChan chan []byte
	//有缓冲管道，用于读、写两个goroutine之间的消息通信
	MsgBuffChan chan []byte
	//释放ip限流
	FreeLimit func()

	sync.RWMutex
	//链接属性
	property map[string]interface{}
	//保护当前property的锁
	propertyLock sync.Mutex
}

//Context 返回ctx，用于用户自定义的go程获取连接退出状态
func (c *Connection) Context() context.Context {
	return c.Ctx
}

//GetConnID 获取当前连接ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

//GetUID 获取当前连接绑定的用户id
func (c *Connection) GetUID() uint32 {
	return c.UID
}

// GetWsConnection 获取websocket连接
func (c *Connection) GetWsConnection() *websocket.Conn { return nil }

//GetHTTPRequest 获取连接的请求对象
func (c *Connection) GetHTTPRequest() *http.Request { return nil }

//GetTCPConnection 从当前连接获取原始的socket TCPConn
func (c *Connection) GetTCPConnection() *net.TCPConn { return nil }

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
