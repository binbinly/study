package sync

import (
	"sync"
	"testing"
)

var (
	mailbox  uint8
	lock     sync.RWMutex
	sendCond = sync.NewCond(&lock)
	recvCond = sync.NewCond(lock.RLocker())
	// sign 用于传递演示完成的信号。
	sign = make(chan struct{}, 2)
	max  = 5
)

// 条件变量，主要起到一个通知的作用
func TestCond(t *testing.T) {
	go func(max int) {
		defer func() {
			sign <- struct{}{}
		}()
		for i := 0; i < max; i++ {
			lock.Lock()
			for mailbox == 1 {
				sendCond.Wait()
			}
			mailbox = 1
			t.Logf("第%d次:放置信息", i)
			lock.Unlock()
			recvCond.Signal()
		}
	}(max)

	go func(max int) {
		defer func() {
			sign <- struct{}{}
		}()
		for i := 0; i < max; i++ {
			lock.RLock()
			for mailbox == 0 {
				recvCond.Wait()
			}
			mailbox = 0
			t.Logf("第%d次:取走信息", i)
			lock.RUnlock()
			sendCond.Signal()
		}
	}(max)

	<-sign
	<-sign
}