package connect

import (
	"fmt"
	"net"
	"sync"
)

//Limiter ip限流器结构
type Limiter struct {
	ips sync.Map
	maxIPCount int
}

//NewLimiter 实例化限流器
func NewLimiter(count int) *Limiter {
	return &Limiter{maxIPCount: count}
}

//Accept 记录数量
func (l *Limiter) Accept(addr net.Addr) (func(), error) {
	key := addrKey(addr)

	n := l.count(key)

	if l.maxIPCount > 0 && n >= l.maxIPCount {
		return func() {}, fmt.Errorf("client connection limit reached key:%v, num:%v", key, n)
	}
	n++
	l.ips.Store(key, n)
	return func() {
		l.free(addr)
	}, nil
}

//Num 当前数量
func (l *Limiter) Num(addr net.Addr) int {
	key := addrKey(addr)

	return l.count(key)
}

func (l *Limiter) free(addr net.Addr)  {
	key := addrKey(addr)

	n := l.count(key)

	if n > 1 {
		n--
		l.ips.Store(key, n)
		return
	}
	l.ips.Delete(key)
}

func (l *Limiter) count(key string) int {
	ipc, ok := l.ips.Load(key)
	n := 0
	if ok {
		n = ipc.(int)
	}
	return n
}

func addrKey(addr net.Addr) string {
	switch a := addr.(type) {
	case *net.TCPAddr:
		return "ip:" + a.IP.String()
	case *net.UDPAddr:
		return "ip:" + a.IP.String()
	case *net.IPAddr:
		return "ip:" + a.IP.String()
	default:
		// not sure what to do with this, just assume whole Addr is relevant?
		return addr.Network() + "/" + addr.String()
	}
}