package main

import (
	"log"

	httpServer "github.com/asim/go-micro/plugins/server/http/v4"
	"go-micro.dev/v4"

	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/server"
)

const (
	ServerName = "demo-http"
)

func main() {
	srv := httpServer.NewServer(
		server.Name(ServerName),
		server.Address(":8888"))

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())

	demo := newDemo()
	demo.InitRouter(router)

	hd := srv.NewHandler(router)
	if err := srv.Handle(hd); err != nil {
		log.Fatalln(err)
	}

	service := micro.NewService(
		micro.Server(srv),
		micro.Registry(registry.NewRegistry()))
	service.Init()
	service.Run()
}

type demo struct{}

func newDemo() *demo {
	return &demo{}
}

func (e *demo) InitRouter(router *gin.Engine) {
	router.POST("/demo", e.demo)
}

func (e *demo) demo(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "call go-micro v3 http server success"})
}
