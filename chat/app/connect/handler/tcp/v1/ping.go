package v1

import (
	"chat/pkg/connect"
	"fmt"
)

//Ping 心跳
func Ping(c *connect.Context) {
	fmt.Println(" router ping ... ")
	err := c.Req.GetConnection().SendBuffMsg(2, []byte("pong"))
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("ping success!!!")
}