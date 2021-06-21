package main

import (
	"chat/app/logic/conf"
	"chat/pkg/connect"
	"chat/pkg/connect/tcp"
	logger "chat/pkg/log"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func init()  {
	dir, _ := os.Getwd()
	conf.Init(dir + "/config/config.local.yaml")
	// init log
	logger.InitLog(&conf.Conf.Logger)
}

/*
	模拟客户端
*/
func main() {
	conn, err := net.Dial("tcp", "192.168.8.2:9060")
	if err != nil {
		fmt.Println("client start err, exit!", err)
		return
	}

	for {
		//发封包message消息
		dp := tcp.NewDataPack()
		msg, _ := dp.Pack(connect.NewMsgPackage(tcp.MsgIdAuth, []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTk0MjQwNjAsImlhdCI6MTYxOTMzNzY2MCwibmJmIjoxNjE5MzM3NjYwLCJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6InRlc3QifQ.rHtuRf4tDSh6SprvkDwJ9FIMZ-iM0hGMeovF9cR_JKc")))
		_, err = conn.Write(msg)
		if err != nil {
			fmt.Println("write error err ", err)
			return
		}
		time.Sleep(time.Second)

		//发封包message消息
		dp = tcp.NewDataPack()
		msg, _ = dp.Pack(connect.NewMsgPackage(1, []byte("Zinx client Demo Test MsgID=0, [Ping]")))
		_, err = conn.Write(msg)
		if err != nil {
			fmt.Println("write error err ", err)
			return
		}

		//先读出流中的head部分
		headData := make([]byte, dp.GetHeadLen())
		_, err = io.ReadFull(conn, headData) //ReadFull 会把msg填充满为止
		if err != nil {
			fmt.Println("read head error", err)
			break
		}
		//将headData字节流 拆包到msg中
		msgHead, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("server unpack err:", err)
			return
		}

		if msgHead.GetDataLen() > 0 {
			//msg 是有data数据的，需要再次读取data数据
			msg := msgHead.(*connect.Message)
			msg.Data = make([]byte, msg.GetDataLen())

			//根据dataLen从io中读取字节流
			_, err := io.ReadFull(conn, msg.Data)
			if err != nil {
				fmt.Println("server unpack data err:", err)
				return
			}

			fmt.Println("==> Test Router:[Ping] Recv Msg: ID=", msg.ID, ", len=", msg.DataLen, ", data=", string(msg.Data))
		}
		time.Sleep(5 * time.Second)
	}
}
