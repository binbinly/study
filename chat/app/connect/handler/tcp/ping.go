package tcp

import (
	"chat/pkg/server/tcp"
	"fmt"
)

//Ping
func Ping(c *tcp.Context) {
	fmt.Println(" router ping ... ")
	err := c.Req.GetConnection().SendBuffMsg(1, []byte("pong"))
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("ping success!!!")
}