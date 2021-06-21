package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

func main()  {

	exampleValue()

	example()
}

func exampleValue()  {

	cfg := make(map[string]int)

	rand.Seed(time.Now().Unix())

	cfg["test"] = rand.Intn(500)

	var config atomic.Value

	config.Store(cfg)

	c := config.Load()

	fmt.Println(c)

}

func example()  {

	var i32 int32

	atomic.StoreInt32(&i32, 15)

	fmt.Println(atomic.LoadInt32(&i32))

	fmt.Println(atomic.AddInt32(&i32, 12))

}
