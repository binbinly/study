package connect

import (
	"context"

	"chat/pkg/app"
)

// 用于触发编译期的接口的合理性检查机制
var _ IRequest = (*Request)(nil)

const (
	//MsgTypeBinary 二进制包消息
	MsgTypeBinary = iota + 1
	//MsgTypeJSON json消息
	MsgTypeJSON
)

//IRequest 实际上是把客户端请求的链接信息 和 请求的数据 包装到了 Request里
type IRequest interface {
	GetConnection() IConnection //获取请求连接信息
	GetBinaryMsg() IMessage     //获取请求消息的数据
	GetJSONMsg() *app.Message   //获取请求的消息ID
	GetType() int8              //消息类型
}

//Request 请求
type Request struct {
	ctx       context.Context //上下文对象
	conn      IConnection     //已经和客户端建立好的连接对象
	binaryMsg IMessage        //tcp客户端请求的数据，二进制数据格式
	jsonMsg   *app.Message    //ws客户端消息 json格式
	t         int8            //消息类型
}

//Options 选项回调
type Options func(r *Request)

//WithBinaryMsg 设置二进制消息
func WithBinaryMsg(msg IMessage) Options {
	return func(r *Request) {
		r.binaryMsg = msg
		r.t = MsgTypeBinary
	}
}

//WithJSONMsg 设置json消息
func WithJSONMsg(msg *app.Message) Options {
	return func(r *Request) {
		r.jsonMsg = msg
		r.t = MsgTypeJSON
	}
}

//NewRequest 请求对象
func NewRequest(conn IConnection, opts ...Options) IRequest {
	r := &Request{
		ctx:  context.Background(),
		conn: conn,
	}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

//Context 上下文
func (r *Request) Context() context.Context {
	return r.ctx
}

// GetConnection 获取请求连接信息
func (r *Request) GetConnection() IConnection {
	return r.conn
}

// GetBinaryMsg 获取请求消息的数据
func (r *Request) GetBinaryMsg() IMessage {
	return r.binaryMsg
}

// GetJSONMsg 获取请求的消息的ID
func (r *Request) GetJSONMsg() *app.Message {
	return r.jsonMsg
}

// GetType 获取消息类型
func (r *Request) GetType() int8 {
	return r.t
}
