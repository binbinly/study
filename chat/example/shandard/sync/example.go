package main

import (
	"fmt"
	"sync"
)

func main()  {

	// 协程同步
	exampleWaitGroup()
	// 协程中只调用单次方法
	exampleOnce()

	// 利用互斥锁进行线程操作
	exampleCond()

	examplePool()

	// 并发读写互斥锁
	exampleRWMutex()
	// 并发map
	exampleMap()
}

func exampleWaitGroup()  {

	var wg sync.WaitGroup

	for i :=0; i<10;i++{
		wg.Add(1)

		go func(i int) {
			fmt.Println(i)

			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("finish")

}

func exampleOnce()  {

	var once sync.Once

	done := make(chan bool)

	for i := 0; i< 10;i++{
		go func() {
			once.Do(func() {
				fmt.Println("Only once")
			})
			done <- true
		}()
	}
	for i := 0; i< 10; i++{
		<- done
	}
}

func exampleCond()  {

	var locker = new(sync.Mutex)

	c :=  sync.NewCond(locker)

	c.L.Lock()

	c.Broadcast()

	c.Signal()

	c.L.Unlock()
}

func examplePool()  {

	var pool = new(sync.Pool)

	pool.New = func() interface{} {
		return 5
	}

	pool.Put(10)

	item := pool.Get()
	fmt.Println(item)

}

func exampleRWMutex()  {

	var rw = new(sync.RWMutex)

	rw.Lock()

	rw.Unlock()

	rw.RLock()

	rw.RUnlock()

	rw.RLocker()

}

func exampleMap()  {

	var m sync.Map

	m.Store("a", "Hello world")
	m.Store("b", "Hello Gopher")
	m.Store("c", "Hello zc")

	fmt.Println(m.Load("c"))

	m.Delete("c")

	m.Range(func(key, value interface{}) bool {
		fmt.Println("iterate:", key, value)
		return true
	})

}