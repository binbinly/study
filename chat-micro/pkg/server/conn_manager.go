package server

import (
	"sync"
)

var _ IConnManager = (*ConnManager)(nil)

//IConnManager 连接管理器
type IConnManager interface {
	Add(conn Connection)                   //添加链接
	Remove(conn Connection)                //删除连接
	Get(connID uint32) (Connection, error) //利用ConnID获取链接
	Len() int                              //获取当前连接
	Clear()                                //删除并停止所有链接
	SendAll(msgID uint32, msg []byte)      //全部连接发送消息
}

//ConnManager 连接管理模块
type ConnManager struct {
	connections map[uint32]Connection
	lock        sync.RWMutex
}

//NewConnManager 创建一个链接管理
func NewConnManager() IConnManager {
	return &ConnManager{
		connections: make(map[uint32]Connection),
	}
}

//Add 添加连接
func (c *ConnManager) Add(conn Connection) {
	// 保护共享资源map，加写锁
	c.lock.Lock()
	defer c.lock.Unlock()

	//将conn连接加入connManger
	c.connections[conn.GetUID()] = conn
}

//Remove 删除连接
func (c *ConnManager) Remove(conn Connection) {
	c.lock.Lock()
	defer c.lock.Unlock()

	//删除连接信息
	delete(c.connections, conn.GetUID())
}

//Get 利用ConnID获取链接
func (c *ConnManager) Get(connID uint32) (Connection, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	if conn, ok := c.connections[connID]; ok {
		return conn, nil
	}
	return nil, ErrConnNotFound
}

//Len 获取当前连接
func (c *ConnManager) Len() int {
	return len(c.connections)
}

//Clear 清除并停止所有连接
func (c *ConnManager) Clear() {
	if c.Len() == 0 {
		return
	}
	c.lock.Lock()
	defer c.lock.Unlock()

	//停止并删除全部连接信息
	for connID, conn := range c.connections {
		go conn.Stop()
		//删除
		delete(c.connections, connID)
	}
}

//SendAll 所有连接发送消息
func (c *ConnManager) SendAll(msgID uint32, msg []byte) {
	for _, conn := range c.connections {
		conn.SendBuffMsg(msgID, msg)
	}
}
