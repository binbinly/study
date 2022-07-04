package context

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// 带可取消的Context，调用cancel可取消所有Context执行
// 一般cancel都是放在defer里面，通过return，执行defer里面的取消函数
func contextWithCancel() {
	rand.Seed(time.Now().Unix())
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	getIP := func(ctx context.Context) <-chan int64 {
		ipChan := make(chan int64)
		go func() {
			for {
				i := rand.Intn(1000)
				select {
				case <-ctx.Done():
					return
				case <-time.After(time.Microsecond * time.Duration(i)):
					fmt.Println("模拟耗时操作...")
					ipChan <- time.Now().UnixNano()
				}
			}
		}()
		return ipChan
	}

	for ip := range getIP(ctx) {
		fmt.Printf("Get IP:%d \n", ip)
		return
	}
}

// 带Value的Context
func contextWithValue(kv map[string]interface{}) {
	ctx := context.Background()
	for key, value := range kv {
		ctx = context.WithValue(ctx, key, value)
	}

	for key := range kv {
		printKey(ctx, key)
	}
}

func printKey(ctx context.Context, key string) {
	time.Sleep(time.Microsecond * 2)
	fmt.Println(ctx.Value(key))
}

func contextWithDeadLine(deadLine time.Time) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "name", "time dead line")
	ctx, cancel := context.WithDeadline(ctx, deadLine)
	defer cancel()

	select {
	case <-time.After(time.Second):
		fmt.Println(ctx.Value("name"))
	case <-ctx.Done():
		fmt.Println("到达deadLine,结束执行")
	}
}

func contextWithTimeout(timeout time.Duration) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "name", "time dead line")
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	select {
	case <-time.After(time.Second):
		fmt.Println(ctx.Value("name"))
	case <-ctx.Done():
		fmt.Println("已超时,结束执行")
	}
}
