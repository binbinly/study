package tcp

// 用于触发编译期的接口的合理性检查机制
var _ IRequest = (*Request)(nil)

/*
	IRequest 接口：
	实际上是把客户端请求的链接信息 和 请求的数据 包装到了 Request里
*/
type IRequest interface {
	GetConnection() *Connection //获取请求连接信息
	GetData() []byte            //获取请求消息的数据
	GetMsgID() uint32           //获取请求的消息ID
}

//Request 请求
type Request struct {
	conn *Connection //已经和客户端建立好的 链接
	msg  IMessage    //tcp客户端请求的数据
}

// 请求
func NewRequest(conn *Connection, msg IMessage) *Request {
	return &Request{
		conn: conn,
		msg:  msg,
	}
}

// GetConnection 获取请求连接信息
func (r *Request) GetConnection() *Connection {
	return r.conn
}

// GetData 获取请求消息的数据
func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

// GetMsgID 获取请求的消息的ID
func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMsgID()
}
