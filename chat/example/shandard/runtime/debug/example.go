package main

import (
	"chat/pkg/log"
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"time"
)

func main()  {

	debug.FreeOSMemory()

	if info, ok := debug.ReadBuildInfo();ok {
		fmt.Println(info.Path)

		fmt.Println(info.Main)

		fmt.Println(info.Deps)
	}

	// 设置垃圾回收百分比
	exampleGCPercent()

	// 设置被单个go协程调用栈可使用内存值
	exampleSetMaxStack()

	// 设置go程序可以使用的最大操作系统线程数
	exampleSetMaxThreads()

	// 设置程序请求运行是只触发panic,而不崩溃
	exampleSetPanic()

	// 垃圾收集信息
	exampleStats()

	// 将内存分配堆和其中对象的描述写入文件中
	exampleHeapDump()

	// 获取go协程调用栈踪迹
	exampleStack()
}

func exampleGCPercent()  {

	fmt.Println(debug.SetGCPercent(1))

	var dic = make([]byte, 100, 100)

	runtime.SetFinalizer(&dic, func(dic *[]byte) {
		fmt.Println("内存回收")
	})

	runtime.GC()

	var s = make([]byte, 100, 100)
	runtime.SetFinalizer(&s, func(dic *[]byte) {
		fmt.Println("内存回收2")
	})

	d := make([]byte, 300, 300)

	for index := range d {
		d[index] = 'a'
	}
	fmt.Println(d)

	time.Sleep(time.Second)

}

func exampleSetMaxStack()  {

	debug.SetMaxStack(102400)

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(1)
		}()

	}
	time.Sleep(time.Second)

}

func exampleSetMaxThreads()  {

	debug.SetMaxThreads(10)

	go func() {
		fmt.Println("Hello World")
	}()
	time.Sleep(time.Second)

}

func exampleSetPanic()  {

	go func() {
		defer func() {
			recover()
		}()

		fmt.Println(debug.SetPanicOnFault(true))
		var s *int = nil
		*s = 34
	}()
	time.Sleep(time.Second)

	fmt.Println("ddd")

}

func exampleStats()  {

	data := make([]byte, 1000, 1000)
	println(data)

	runtime.GC()

	var stats debug.GCStats

	debug.ReadGCStats(&stats)

	fmt.Println(stats.NumGC)

	fmt.Println(stats.LastGC)

	fmt.Println(stats.Pause)

	fmt.Println(stats.PauseTotal)

	fmt.Println(stats.PauseEnd)

}

func exampleHeapDump()  {

	f, _ := os.OpenFile("example/testdata/debug_heapDump.txt",  os.O_RDWR|os.O_CREATE, 0666)

	fd := f.Fd()

	debug.WriteHeapDump(fd)

	data := make([]byte, 10, 10)
	println(data)
	runtime.GC()

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

}

func exampleStack()  {

	go func() {
		fmt.Println(string(debug.Stack()))
	}()

	time.Sleep(time.Second)

}