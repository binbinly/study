package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

var (
	receiveData = make([]byte, 1024)
	sendData    = make([]byte, 1024)
	reader      = bufio.NewReader(os.Stdin)
)

func main()  {
	conn, err := net.Dial("tcp", "localhost:8888")
	defer conn.Close()

	if err != nil {
		log.Panic("Server not found")
	}

	fmt.Println("Connection is OK")
	fmt.Print("Please enter your name")
	fmt.Scanf("%s", &sendData)
	_, err = conn.Write(sendData)
	if err != nil {
		log.Fatalf("Error when write to server:%s \n", err.Error())
	}
	fmt.Println("Now you can talk...")

	go read(conn)

	for {
		fmt.Scanf("%s", &sendData)

		if string(sendData) == "quit" {
			fmt.Println("quit the client.......")
			os.Exit(-1)
		}

		_, err := conn.Write(sendData)
		if err != nil {
			fmt.Printf("Error when send server,err:%s \n", err.Error())
		}
	}
}

// 读取从server发来的信息
func read(conn net.Conn)  {
	for {
		length, err := conn.Read(receiveData)
		if err != nil {
			log.Fatalf("Error when receive from server:%s", err)
		}
		data := string(receiveData[:length])

		fmt.Println(data)
	}
}
