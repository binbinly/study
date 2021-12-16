package main

import (
	"common/constvar"
	pb "common/proto/product"
	"common/app"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func main() {
	client := app.NewTestClient("192.168.8.76:8500")

	svc := pb.NewProductService(constvar.ServiceProduct, client)
	//reply, err := svc.SkuList(context.Background(), &pb.SkuListReq{Page: 0})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Printf("reply: %v", reply)

	cat, err := svc.CategoryTree(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("cat", cat)
}
