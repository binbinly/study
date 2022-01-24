package server

import (
	"context"
	"net"
)

//Connection 定义连接接口
type Connection interface {
	Start()                   //启动连接，让当前连接开始工作
	Stop()                    //停止连接，结束当前连接状态
	Context() context.Context //返回ctx，用于用户自定义的go程获取连接退出状态

	GetID() uint32        //获取当前连接ID
	GetUID() uint32       //获取当前连接绑定的用户id
	RemoteAddr() net.Addr //获取远程客户端地址信息

	SendMsg(msgID uint32, data []byte) error     //直接将Message数据发送至远程客户端
	SendBuffMsg(msgID uint32, data []byte) error //发送至缓冲区
}
