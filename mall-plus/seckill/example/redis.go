package main

import (
	"context"
	"fmt"
	"pkg/redis"
)

func main()  {
	redis.InitTestRedis()

	ctx := context.Background()

	for i := 0; i < 10; i++ {
		go func() {
			a, _ := redis.Client.Decr(ctx, "test").Result()
			if a <= 0 {
				fmt.Println("aaaaa")
				return
			}
			fmt.Printf("test: %v\n", a)
		}()
	}
	select {}
}
