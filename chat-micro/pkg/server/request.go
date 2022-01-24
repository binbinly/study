package server

// IRequest is a synchronous request interface
type IRequest interface {
	//Conn 连接对象
	Conn() Connection
	// The action requested
	Method() string
	// Header of the request
	Header() map[string]string
	// Body is the initial decoded value
	Body() []byte
}

type Request struct {
	conn   Connection
	method string
	header map[string]string
	body   []byte
}

//NewRequest 实例化请求
func NewRequest(conn Connection, method string, body []byte, header map[string]string) *Request {
	return &Request{
		conn:   conn,
		method: method,
		header: header,
		body:   body,
	}
}

func (r *Request) Conn() Connection {
	return r.conn
}

func (r *Request) Method() string {
	return r.method
}

func (r *Request) Header() map[string]string {
	return r.header
}

func (r *Request) Body() []byte {
	return r.body
}
