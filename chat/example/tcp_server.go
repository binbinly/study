package main

import (
	"chat/app/connect/conf"
	"chat/pkg/app"
	logger "chat/pkg/log"
	"chat/pkg/server"
	"chat/pkg/server/tcp"
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
func DoConnectionBegin(conn server.IConnection) bool {
	conn.SetProperty("Name", "test")
	log.Println("DoConnectionBegin is Called ... ")
	return true
}

//连接断开的时候执行
func DoConnectionLost(conn server.IConnection) {
	if name, err := conn.GetProperty("Name"); err == nil {
		log.Println("Conn Property Name = ", name)
	}
	log.Println("DoConnectionLost is Called ... ")
	log.Println("user offline ")
}

func DoConnectionFinish(conn server.IConnection) {
	log.Println("DoConnectionFinish is Called ... ")
	log.Println("user online")
}

func main() {
	//创建一个server句柄
	engine := tcp.NewEngine()
	engine.Use(func(c *tcp.Context) {
		log.Println("start exec router middleware")
		c.Next()
		log.Println("end exec router middleware")
	})
	engine.AddRoute(tcp.MsgIdAuth, Auth)
	engine.AddRoute(1, Ping)

	s := tcp.NewServer(&conf.Conf.Tcp, engine)

	//注册链接hook回调函数
	s.SetOnConnStart(DoConnectionBegin)
	s.SetOnConnStop(DoConnectionLost)
	s.SetOnConnFinish(DoConnectionFinish)

	//开启服务
	s.Serve()
	//阻塞主进程退出
	select {}
}

//Ping Handle
func Ping(c *tcp.Context) {

	log.Println("Call PingRouter Handle")
	//先读取客户端的数据，再回写ping...ping...ping
	log.Println("recv from client : msgId=", c.Req.GetMsgID(), ", data=", string(c.Req.GetData()))

	err := c.Req.GetConnection().SendBuffMsg(0, []byte("ping...ping...ping"))
	if err != nil {
		log.Println("err:", err)
	}
	log.Println("send success!!!")
}

func Auth(c *tcp.Context) {
	log.Println("Call AuthRouter Handle")
	p, err := app.Parse(string(c.Req.GetData()), conf.Conf.App.JwtSecret)
	if err != nil {
		log.Println("token parse", err)
		c.Req.GetConnection().Stop()
		return
	}
	c.Req.GetConnection().Auth(uint32(p.UserId))
}
