package main

import (
	"context"
	"log"

	"go-micro.dev/v4/metadata"

	"common/constvar"
	pb "common/proto/order"
	"common/app"
)

func main() {
	client := app.NewTestClient("192.168.8.76:8500")

	svc := pb.NewOrderService(constvar.ServiceOrder, client)

	// Set arbitrary headers in context
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"User-Id": "4",
	})
	reply, err := svc.SubmitSkuOrder(ctx, &pb.SkuSubmitReq{
		SkuId:     1911300,
		AddressId: 6,
		CouponId:  0,
		Note:      "",
		Num:       2,
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("reply", reply)
}
