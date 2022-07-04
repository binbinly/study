package sync

import (
	"log"
	"os"
	"sync"
	"testing"
	"time"
)

var (
	locker   sync.Mutex
	rwLocker sync.RWMutex
)

/**
排它锁sync.Mutex
只能加一次锁，重复加锁会导致死锁
*/
func TestMutex(t *testing.T) {
	go mutex("1")
	go mutex("2")
	time.Sleep(time.Millisecond * 10)
}

func TestRWMutex(t *testing.T) {
	go rwMutex("1")
	go rwMutex("2")
	time.Sleep(time.Millisecond * 10)
}

func mutex(str string) string {
	log.Printf("first:%s", str)
	locker.Lock()
	log.Printf("this is test %s", str)
	if str == "1" {
		return "1"
	}
	time.Sleep(time.Second)
	os.Exit(0)
	locker.Unlock()
	return "2"
}

func rwMutex(str string) string {
	log.Printf("first:%s", str)
	rwLocker.RLock()
	rwLocker.RLock()

	log.Printf("this is test %s", str)
	if str == "1" {
		return "1"
	}

	time.Sleep(time.Second)
	defer rwLocker.RUnlock()
	defer rwLocker.RUnlock()

	return "2"

}