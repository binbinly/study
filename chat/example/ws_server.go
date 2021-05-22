package main

import (
	"chat/app/connect/conf"
	"chat/pkg/app"
	"chat/pkg/server"
	"chat/pkg/server/ws"
	logger "chat/pkg/log"
	"log"
	"os"
)

func init() {
	dir, _ := os.Getwd()
	conf.Init(dir + "/config/config.local.yaml")
	// init log
	logger.InitLog(&conf.Conf.Logger)
}

//创建连接的时候执行
func DoConnectionWsBegin(conn server.IConnection) bool {
	conn.SetProperty("Name", "test")
	log.Println("DoConnectionBegin is Called ... ")
	return true
}

//连接断开的时候执行
func DoConnectionWsLost(conn server.IConnection) {
	if name, err := conn.GetProperty("Name"); err == nil {
		log.Println("Conn Property Name = ", name)
	}
	log.Println("DoConnectionLost is Called ... ")
	log.Println("user offline")
}

func DoConnectionWsFinish(conn server.IConnection) {
	log.Println("DoConnectionFinish is Called ... ")
	log.Println("user online")
}

func main() {
	//创建一个server句柄
	engine := ws.NewEngine()
	engine.Use(func(c *ws.Context) {
		log.Println("start exec router middleware")
		c.Next()
		log.Println("end exec router middleware")
	})
	engine.AddRoute("ping", WsPing)

	s := ws.NewServer(&conf.Conf.Ws, engine)

	//注册链接hook回调函数
	s.SetOnConnStart(DoConnectionWsBegin)
	s.SetOnConnStop(DoConnectionWsLost)
	s.SetOnConnFinish(DoConnectionWsFinish)

	//开启服务
	s.Serve()
	//阻塞主进程退出
	select {}
}

//WsPing Handle
func WsPing(c *ws.Context) {

	log.Println("Call PingRouter Handle")
	msg, err := app.NewMessagePack("ping", "pong")
	err = c.Req.GetConnection().SendBuffMsg(msg)
	if err != nil {
		log.Println("err:", err)
	}
	log.Println("send success!!!")
}