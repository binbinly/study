package _defer

import (
	"fmt"
	"testing"
	"time"
)

func TestDefer(t *testing.T)  {
	doSomething()
}

func doSomething()  {
	defer countTime("doSomething")()

	time.Sleep(3 * time.Second)
	fmt.Println("done")
}

// 统计某函数的运行时间
func countTime(msg string) func() {
	start := time.Now()
	fmt.Printf("run func: %s", msg)
	return func() {
		fmt.Printf("func name: %s run time: %f s \n", msg, time.Since(start).Seconds())
	}
}