package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	pb "common/proto/market"
	"common/app"
	"google.golang.org/grpc/status"
)

func main()  {

	c := pb.NewMarketService("mall.market", app.NewTestClient("192.168.8.76:8500"))
	reply, err := c.GetHomeCatData(context.Background(), &pb.CatReq{CatId: 1220})
	if err != nil {
		log.Printf("err: %v", err)
		if s, ok := status.FromError(err); ok {
			fmt.Println("s", s.Message())
		}
	}
	log.Printf("reply: %v", reply)
	str, _ := json.Marshal(reply)
	log.Println("str", string(str))
}