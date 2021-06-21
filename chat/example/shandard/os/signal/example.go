package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main()  {

	c := make(chan os.Signal, 1)

	signal.Notify(c, syscall.SIGINT, syscall.SIGKILL)

	s := <-c
	fmt.Println("got signal:", s)

	signal.Stop(c)
}