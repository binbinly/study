package v1

import (
	"chat/pkg/connect"
	"fmt"

	"chat/pkg/app"
)

//Ping 心跳
func Ping(c *connect.Context) {
	fmt.Println(" router ping ... ")
	msg, err := app.NewMessagePack("ping", "pong")
	err = c.Req.GetConnection().SendBuffMsg(0, msg)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("ping success!!!")
}