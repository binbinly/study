package main

import (
	"chat/example/limit/conn"
	"context"
	"fmt"
	"net"
	"time"
)

func runLimitedServer(lim *conn.Limiter, connTime time.Duration) (string, func()) {
	l, err := net.Listen("tcp", "127.0.0.1:10000")
	if err != nil {
		fmt.Println("err", err)
	}
	serverCtx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				fmt.Println("accept err", err)
				continue
			}
			go func() {
				defer conn.Close()

				free, err := lim.Accept(conn)
				if err != nil {
					fmt.Println("acc err", err)
					return
				}
				defer free()

				ctx, cancel := context.WithCancel(serverCtx)
				defer cancel()

				_, err = conn.Write([]byte("Hello"))
				if err != nil {
					fmt.Println("write err", err)
					return
				}

				go func() {
					for {
						bs := make([]byte, 10)
						_, err := conn.Read(bs)
						if err != nil {
							fmt.Println("read err", err)
							cancel()
							return
						}
						if ctx.Err() != nil {
							return
						}
					}
				}()

				select {
				case <-ctx.Done():
					return
				case <-time.After(connTime):
					return
				}
			}()
		}
	}()

	return l.Addr().String(), func() {
		l.Close()
		cancel()
	}
}

func clientRead(serverAddr string) (net.Conn, string, error) {
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("client err", err)
	}
	conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	
	buf := make([]byte, 10)
	n, err := conn.Read(buf)
	
	return conn, string(buf[0:n]), err

}

func main()  {
	LimiterDenies()

}

func LimiterDenies()  {
	lim := conn.NewLimiter(conn.Config{MaxConnsPerClientIP: 9})

	serverAddr, serverClose := runLimitedServer(lim, 10 * time.Minute)
	defer serverClose()

	conn1, got, err := clientRead(serverAddr)
	if err != nil {
		fmt.Println("client read err ", err)
		return
	}
	fmt.Println("got", got)
	defer conn1.Close()

	conn2, got, err := clientRead(serverAddr)
	if err != nil {
		fmt.Println("client read2 err ", err)
		return
	}
	fmt.Println("got2", got)
	defer conn2.Close()

	conn3, got, err := clientRead(serverAddr)
	if err != nil {
		fmt.Println("client read3 err", err)
		return
	}
	fmt.Println("got3", got)
	defer conn3.Close()

	attempts := 0
	for {
		conn4, got, err := clientRead(serverAddr)
		if err != nil {
			fmt.Println("client read4 err", err)
			attempts++
			if attempts < 10 {
				time.Sleep(20 * time.Millisecond)
				continue
			}
		}
		fmt.Println("got4", got)
		defer conn4.Close()
		break
	}

	lim.SetConfig(conn.Config{MaxConnsPerClientIP: 3})

	conn5, got, err := clientRead(serverAddr)
	if err != nil {
		fmt.Println("client read5 err", err)
		return
	}
	fmt.Println("got5", got)
	defer conn5.Close()

	fmt.Println("num", lim.NumOpen(conn5.RemoteAddr()))

	conn6, _, err := clientRead(serverAddr)
	if err != nil {
		fmt.Println("client read6 err", err)
		return
	}
	defer conn6.Close()
}