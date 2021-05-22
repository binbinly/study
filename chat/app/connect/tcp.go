package connect

import (
	"chat/app/connect/routers"
	"chat/pkg/log"
	"chat/pkg/server"
	"chat/pkg/server/tcp"
)

func NewTcpServer(c *tcp.Config) *tcp.Server {
	s := tcp.NewServer(c, routers.NewTcpEngine())

	//注册链接hook回调函数
	s.SetOnConnStart(DoConnectionBegin)
	s.SetOnConnStop(DoConnectionLost)
	s.SetOnConnFinish(DoConnectionFinish)

	//开启服务
	s.Serve()
	return s
}

//创建连接的时执行
func DoConnectionBegin(conn server.IConnection) bool {
	log.Info("Do Connection begin is Called ...")
	return true
}

//连接鉴权完成时执行
func DoConnectionFinish(conn server.IConnection) {
	log.Info("Do Connection finish is Called ...")
	log.Info("user online")
}

//连接断开的时候执行
func DoConnectionLost(conn server.IConnection) {
	log.Info("Do Connection lost is Called ...")
	log.Info("user offline")
}
