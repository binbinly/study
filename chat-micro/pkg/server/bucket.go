package server

import (
	"strconv"

	"chat-micro/pkg/crypt"
)

var _ IBucket = (*Bucket)(nil)

//IBucket 接口定义
type IBucket interface {
	GetConnMgr(id uint32) IConnManager
	Broadcast(msgID uint32, msg []byte)
	ConnTotal() int
	Insert(conn Connection)
	Remove(conn Connection)
	Clear()
}

//Bucket 连接器桶，减少锁竞争
type Bucket struct {
	//连接管理器
	connMgr []IConnManager
	//桶大小
	size int
}

//NewBucket 实例化
func NewBucket(size int) *Bucket {
	mgr := make([]IConnManager, size)
	for i := 0; i < size; i++ {
		mgr[i] = NewConnManager()
	}
	return &Bucket{
		connMgr: mgr,
		size:    size,
	}
}

// GetConnMgr 当前连接的管理器
func (b *Bucket) GetConnMgr(id uint32) IConnManager {
	if b.size == 1 {
		return b.connMgr[0]
	}
	userIDStr := strconv.Itoa(int(id))
	idx := crypt.CityHash32([]byte(userIDStr), uint32(len(userIDStr))) % uint32(b.size)
	return b.connMgr[idx]
}

// ConnTotal 当前服务器的总连接数
func (b *Bucket) ConnTotal() int {
	var c int
	for _, manager := range b.connMgr {
		c += manager.Len()
	}
	return c
}

//Insert 插入连接
func (b *Bucket) Insert(conn Connection) {
	b.GetConnMgr(conn.GetUID()).Add(conn)
}

//Remove 删除连接
func (b *Bucket) Remove(conn Connection) {
	b.GetConnMgr(conn.GetUID()).Remove(conn)
}

//Clear 清理所有连接
func (b *Bucket) Clear() {
	for _, manager := range b.connMgr {
		manager.Clear()
	}
}

//Broadcast 广播所有连接
func (b *Bucket) Broadcast(msgID uint32, msg []byte) {
	for _, manager := range b.connMgr {
		manager.SendAll(msgID, msg)
	}
}
