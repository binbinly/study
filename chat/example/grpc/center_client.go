package main

import (
	"chat/app/center/server"
	logger "chat/pkg/log"
	"chat/pkg/net/grpc"
	"chat/pkg/registry"
	"chat/pkg/registry/consul"
	pb "chat/proto/center"
	"context"
	"fmt"
	"log"
	"time"
)

var grpcClient pb.CenterClient

func init() {
	logger.InitLog(logger.NewConfig())
	registry.RegisterPlugin(consul.NewConsul())
	// init registry
	_, err := registry.InitRegistry(context.Background(), "consul",
		registry.WithAddr([]string{"192.168.8.76:8500"}),
	)
	if err != nil {
		log.Fatalf("failed to init register: %v", err)
	}
}

func main() {
	consul.Init()
	target := "consul://192.168.8.76:8500/center"
	target = "192.168.8.2:20001"
	conn := grpc.NewRPCClientConn(&grpc.ClientConfig{
		ServiceName:      "center",
		QPSLimit:         100,
		Timeout:          15 * time.Second,
		KeepAliveTime:    5 * time.Second,
		KeepAliveTimeout: 1 * time.Second,
	}, target)
	defer conn.Close()
	grpcClient = pb.NewCenterClient(conn)
	ctx := context.Background()
	reg(ctx, "test", "123456", 13333333333)
}

func reg(ctx context.Context, username, password string, phone int64) {
	res, err := grpcClient.UserRegister(ctx, &pb.RegisterReq{
		Username: username,
		Password: password,
		Phone:    phone,
	})
	if err != nil {
		fmt.Printf("err:%v\n", err)
		err = server.HandleError(err)
		fmt.Printf("after err: %v\n", err)
	}
	fmt.Printf("res:%v\n", res)
}

func userinfo(ctx context.Context, id uint32) {
	res, err := grpcClient.UserInfo(ctx, &pb.UIDReq{Id: 1})
	if err != nil {
		fmt.Printf("err:%v\n", err)
		err = server.HandleError(err)
		fmt.Printf("after err: %v\n", err)
	}
	fmt.Println("res", res)
}
