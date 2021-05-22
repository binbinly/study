package server

import (
	"errors"
)

var (
	// 数据包超过定义最大值
	ErrLargeReceived = errors.New("too large msg data received")
	// 连接未完成，不可以发送消息
	ErrConnectNotFinish = errors.New("connection not finish when send msg")
	// 连接属性未找到
	ErrPropertyNotFound = errors.New("no property found")
	// 连接未找到
	ErrConnNotFound = errors.New("connection not found")
)
