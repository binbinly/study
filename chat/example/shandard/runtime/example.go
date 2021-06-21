package main

import (
	"fmt"
	"runtime"
	"strings"
)

func main()  {
	example()

	// 显示调用过程
	exampleFrames()
}

func example()  {

	fmt.Println(runtime.GOROOT())

	fmt.Println(runtime.Version())

	fmt.Println(runtime.NumCPU())

	fmt.Println(runtime.NumCgoCall())

	fmt.Println(runtime.NumGoroutine())

	runtime.GC()

	runtime.GOMAXPROCS(2)

	go func() {
		fmt.Println("Start Goroutine1")

		runtime.Goexit()

		fmt.Println("end Goroutine1")


	}()

	go func() {
		runtime.Gosched()

		fmt.Println("Start Goroutine2")
		fmt.Println("End Goroutine2")
	}()

	go func() {
		runtime.LockOSThread()
		fmt.Println("Start Goroutine2")
		fmt.Println("End Goroutine3")

		runtime.UnlockOSThread()
	}()
}

func exampleFrames()  {

	c := func() {
		pc := make([]uintptr, 10)

		n := runtime.Callers(0, pc)
		if n == 0 {
			return
		}

		pc = pc[:n]

		frames := runtime.CallersFrames(pc)

		for {
			frame, more := frames.Next()

			if !strings.Contains(frame.File, "runtime/"){
				break
			}

			fmt.Printf("- more:%v | %s\n", more, frame.Function)
			if !more {
				break
			}
		}
	}
	b := func() {
		c()
	}
	a := func() {
		b()
	}
	a()

}