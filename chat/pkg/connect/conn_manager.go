package connect

import (
	"sync"
)

var _ IConnManager = (*ConnManager)(nil)

//IConnManager 连接管理器
type IConnManager interface {
	Add(conn IConnection)                   //添加链接
	Remove(conn IConnection)                //删除连接
	Get(connID uint32) (IConnection, error) //利用ConnID获取链接
	GetConnections() map[uint32]IConnection //获取所有连接
	Len() int                               //获取当前连接
	Clear()                                 //删除并停止所有链接
}

//ConnManager 连接管理模块
type ConnManager struct {
	Connections map[uint32]IConnection
	lock        sync.RWMutex
}

//NewConnManager 创建一个链接管理
func NewConnManager() IConnManager {
	return &ConnManager{
		Connections: make(map[uint32]IConnection),
	}
}

//Add 添加连接
func (c *ConnManager) Add(conn IConnection) {
	// 保护共享资源map，加写锁
	c.lock.Lock()
	defer c.lock.Unlock()

	//将conn连接加入connManger
	c.Connections[conn.GetConnID()] = conn
}

//Remove 删除连接
func (c *ConnManager) Remove(conn IConnection) {
	c.lock.Lock()
	defer c.lock.Unlock()

	//删除连接信息
	delete(c.Connections, conn.GetConnID())
}

//Get 利用ConnID获取链接
func (c *ConnManager) Get(connID uint32) (IConnection, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	if conn, ok := c.Connections[connID]; ok {
		return conn, nil
	}
	return nil, ErrConnNotFound
}

//GetConnections 获取所有连接
func (c *ConnManager) GetConnections() map[uint32]IConnection {
	return c.Connections
}

//Len 获取当前连接
func (c *ConnManager) Len() int {
	return len(c.Connections)
}

//Clear 清除并停止所有连接
func (c *ConnManager) Clear() {
	if c.Len() == 0 {
		return
	}
	c.lock.Lock()
	defer c.lock.Unlock()

	//停止并删除全部连接信息
	for connID, conn := range c.Connections {
		go conn.Stop()
		//删除
		delete(c.Connections, connID)
	}
}
