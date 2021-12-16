package main

import (
	cache "example/micro/cache/proto"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func main()  {
	got := &cache.GetResponse{
		Value:      "aa",
		Expiration: "",
	}
	src := &cache.GetResponse{
		Value:      "",
		Expiration: "bb",
	}
	proto.Merge(src, got) // The first param is destination
	fmt.Println("merge", got)
	fmt.Println("merge", src)
}