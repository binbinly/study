package main

import (
	"fmt"
	"io"
	"net"
)

func main()  {
	listener, err := net.Listen("tcp", ":8082")
	if err != nil {
		panic(err)
	}

	fmt.Println("tcp server is running...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go process(conn)
	}
}

func process(conn net.Conn)  {
	defer conn.Close()

	for {
		var data [1024]byte
		n, err := conn.Read(data[:])
		if err != nil && err != io.EOF {
			fmt.Printf("failed to read msg client, err: %v \n", err)
			break
		}

		str := string(data[:n])
		if str == "exit" {
			fmt.Println("client exit ...")
			break
		}
		fmt.Printf("read msg: %s \n", str)

		conn.Write([]byte(fmt.Sprintf("%s OK!", str)))
	}
}