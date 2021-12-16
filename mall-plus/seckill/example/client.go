package main

import (
	"context"
	"log"

	"go-micro.dev/v4/metadata"

	"common/app"
	"common/constvar"
	pb "common/proto/seckill"
)

func main() {
	client := app.NewTestClient("192.168.8.76:8500")

	svc := pb.NewSeckillService(constvar.ServiceSeckill, client)

	// Set arbitrary headers in context
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"User-Id": "4",
	})
	reply, err := svc.Kill(ctx, &pb.KillReq{
		SkuId:     1300,
		AddressId: 6,
		Num:       1,
		Key:       "e59c258eeeda8a85800be822b3383e02",
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("reply", reply)
}
