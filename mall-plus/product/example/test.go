package main

import (
	"context"
	"fmt"
	"pkg/redis"
	"reflect"
)

func main()  {
	redis.InitTestRedis()
	var str interface{}
	str = "床垫"
	inter(&str)
}

func inter(val interface{})  {
	fmt.Printf("val:%#v\n", val)
	fmt.Printf("val:%#v\n", reflect.ValueOf(val).Elem().String())
	redis.Client.Set(context.Background(), "test", val, 0)
}