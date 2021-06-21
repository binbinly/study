package connect

import (
	"errors"
)

const (
	//ServerWs websocket服务器
	ServerWs  = 1
	//ServerTCP tcp服务器
	ServerTCP = 2
)

var (
	//MaxPacketSize 数据包大小限制/字节
	MaxPacketSize int64 = 4096
)

var (
	//ErrLargeReceived 数据包超过定义最大值
	ErrLargeReceived = errors.New("too large msg data received")
	//ErrConnectNotFinish 连接未完成，不可以发送消息
	ErrConnectNotFinish = errors.New("connection not finish when send msg")
	//ErrPropertyNotFound 连接属性未找到
	ErrPropertyNotFound = errors.New("no property found")
	//ErrConnNotFound 连接未找到
	ErrConnNotFound = errors.New("connection not found")
)
