package ws

import (
	"fmt"

	"chat/pkg/app"
	"chat/pkg/server/ws"
)

//Ping
func Ping(c *ws.Context) {
	fmt.Println(" router ping ... ")
	msg, err := app.NewMessagePack("ping", "pong")
	err = c.Req.GetConnection().SendBuffMsg(msg)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("ping success!!!")
}