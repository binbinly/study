package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type CounterLimit struct {
	counter int64 //计数器
	limit int64 //指定时间窗口内允许的最大请求数
	intervalNano int64 //指定的时间窗口
	unixNano int64 //时间戳
}

func NewCounterLimit(interval time.Duration, limit int64) *CounterLimit {
	return &CounterLimit{
		counter:      0,
		limit:        limit,
		intervalNano: int64(interval),
		unixNano:     time.Now().UnixNano(),
	}
}

func (c *CounterLimit) Allow() bool {
	now := time.Now().UnixNano()
	if now - c.unixNano > c.intervalNano { //如果当前过了当前的时间窗口,则重新进行计数
		atomic.StoreInt64(&c.counter, 0)
		atomic.StoreInt64(&c.unixNano, now)
		return true
	}
	atomic.AddInt64(&c.counter, 1)
	return c.counter < c.limit
}

func main()  {
	limit := NewCounterLimit(time.Second, 50)
	m := make(map[int]bool)
	for i:=0; i < 1000; i++ {
		allow := limit.Allow()
		if allow {
			m[i] = true
		} else {
			m[i] = false
		}
		time.Sleep(time.Millisecond * 10)
	}
	for i := 0; i < 1000; i++ {
		fmt.Printf("i=%d allow=%v\n", i, m[i])
	}
}