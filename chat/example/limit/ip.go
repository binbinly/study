package main

import (
	"chat/app/logic/conf"
	"chat/pkg/redis"
	"context"
	"fmt"
	redis2 "github.com/go-redis/redis/v8"
	"log"
	"os"
	"time"
)

func Init()  {
	dir, _ := os.Getwd()
	conf.Init(dir + "/config/logic.local.yaml")
	redis.Init(&conf.Conf.Redis)
}

func main()  {
	Init()
	for i := 0; i < 10; i++ {
		IpRateLimiter(5, time.Second * 2)
	}
}

var num = 0

func IpRateLimiter(limit int, slidingWindow time.Duration) {
	num++

	now := time.Now().UnixNano()
	userCntKey := fmt.Sprint("ip-limit:", "127.0.0.1")

	c := context.Background()
	pipe := redis.Client.Pipeline()
	pipe.ZRemRangeByScore(c, userCntKey,
		"0",
		fmt.Sprint(now-(slidingWindow.Nanoseconds())))

	reqs, _ := pipe.ZRange(c, userCntKey, 0, -1).Result()
	pipe.Exec(c)

	if len(reqs) >= limit {
		log.Println("too many request", num)
		return
	}
	log.Println("success", num)

	redis.Client.ZAddNX(c, userCntKey, &redis2.Z{Score: float64(now), Member: float64(now)})
	redis.Client.Expire(c, userCntKey, slidingWindow)

}