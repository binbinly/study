package ws

import (
	"chat/pkg/app"
)

//Request 请求
type Request struct {
	*app.Request
	conn *Connection  //已经和客户端建立好的 链接
}

// 请求
func NewRequest(conn *Connection, msg app.IMessage) *Request {
	return &Request{
		conn: conn,
		Request: &app.Request{Msg: msg},
	}
}

// GetConnection 获取请求连接信息
func (r *Request) GetConnection() *Connection {
	return r.conn
}
