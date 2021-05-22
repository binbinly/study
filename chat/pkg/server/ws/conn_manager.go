package ws

import (
	"sync"

	"chat/pkg/log"
	"chat/pkg/server"
)

var _ IConnManager = (*ConnManager)(nil)

// 连接管理器
type IConnManager interface {
	Add(conn *Connection)                   //添加链接
	Remove(conn *Connection)                //删除连接
	Get(connID uint32) (*Connection, error) //利用ConnID获取链接
	Len() int                               //获取当前连接
	Clear()                                 //删除并停止所有链接
}

//ConnManager 连接管理模块
type ConnManager struct {
	Connections map[uint32]*Connection
	lock        sync.RWMutex
}

//NewConnManager 创建一个链接管理
func NewConnManager() *ConnManager {
	return &ConnManager{
		Connections: make(map[uint32]*Connection),
	}
}

//Add 添加连接
func (c *ConnManager) Add(conn *Connection) {
	// 保护共享资源map，加写锁
	c.lock.Lock()
	defer c.lock.Unlock()

	//将conn连接加入connManger
	c.Connections[conn.GetConnID()] = conn
	log.Infof("[conn.manager] connection add success; conn num=%v", c.Len())
}

//Remove 删除连接
func (c *ConnManager) Remove(conn *Connection) {
	c.lock.Lock()
	defer c.lock.Unlock()

	//删除连接信息
	delete(c.Connections, conn.GetConnID())
	log.Infof("[conn.manager] connection remove success len=%v", c.Len())
}

//Get 利用ConnID获取链接
func (c *ConnManager) Get(connID uint32) (*Connection, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	if conn, ok := c.Connections[connID]; ok {
		return conn, nil
	}
	return nil, server.ErrConnNotFound
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
	for connId, conn := range c.Connections {
		go conn.Stop()
		//删除
		delete(c.Connections, connId)
	}
	log.Infof("[conn.manager] clear all connections success len=%v", c.Len())
}
