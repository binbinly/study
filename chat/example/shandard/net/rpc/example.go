package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"strings"
)

func main()  {

	example()

	example2()

}

type Hello struct {

}

func (h *Hello) Say(args *[]string, reply *string) error {
	*reply = strings.Join(*args, " ")
	return nil
}

func example()  {
	hello := new(Hello)

	server := rpc.NewServer()

	server.Register(hello)

	server.RegisterName("Hello", hello)

	l, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatalf("net.listen tcp :0:%v", err)
	}

	go server.Accept(l)

	server.HandleHTTP("/hello", "/debug")

	address, err := net.ResolveTCPAddr("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("ResolveTCPAddr error:", err)
	}

	conn, _ := net.DialTCP("tcp", nil, address)
	defer conn.Close()

	client := rpc.NewClient(conn)
	defer client.Close()

	args := &[]string{"Hello", "World"}
	reply := new(string)
	err = client.Call("Hello.Say", args, reply)
	if err != nil {
		log.Fatal("Hello error:", err)
	}
	log.Println(*reply)

}

func example2()  {
	hello := new(Hello)
	rpc.Register(hello)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":12345")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)

	client, err := rpc.DialHTTP("tcp", "127.0.0.1:12345")
	if err != nil {
		log.Fatal(err)
	}

	args := &[]string{"Hello", "Gopher"}
	reply := new(string)
	err = client.Call("Hello.Say", args, reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*reply)
}