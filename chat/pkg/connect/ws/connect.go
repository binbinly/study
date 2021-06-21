package ws

import (
	"context"
	"encoding/json"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/websocket"

	"chat/pkg/app"
	"chat/pkg/connect"
	"chat/pkg/log"
)

var _ connect.IConnection = (*Connection)(nil)

//Connection 连接
type Connection struct {
	*connect.Connection
	//当前Conn属于哪个Server
	WsServer *Server
	//连接对象
	Conn *websocket.Conn
	//websocket响应对象
	writer http.ResponseWriter
	//websocket请求对象
	request *http.Request
	//当前连接的关闭状态
	isClosed bool
}

// NewConnect 创建连接对象
func NewConnect(s *Server, w http.ResponseWriter, r *http.Request, connID uint32) connect.IConnection {
	//初始化Conn属性
	return &Connection{
		Connection: &connect.Connection{
			ConnID:      connID,
			MsgChan:     make(chan []byte),
			MsgBuffChan: make(chan []byte, s.c.MaxMsgChanLen),
		},
		WsServer: s,
		writer:   w,
		request:  r,
	}
}

//startWriter 写消息Goroutine， 用户将数据发送给客户端
func (c *Connection) startWriter() {
	log.Debug("[ws.write] Writer Goroutine is running")
	// 心跳由客户端发送 ping 回复 pong
	ticker := time.NewTicker(c.WsServer.c.PingPeriod)
	defer func() {
		defer log.Debugf("[ws.write] %v conn Writer exit!", c.RemoteAddr().String())
		ticker.Stop()
	}()
	for {
		select {
		case data := <-c.MsgChan:
			log.Debugf("[ws.write] data:%v", string(data))
			//write data dead time , like http timeout , default 10s
			c.Conn.SetWriteDeadline(time.Now().Add(c.WsServer.c.WriteWait))
			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				log.Warnf("[ws.write] conn.NextWriter err :%v  ", err)
				return
			}
			w.Write(data)
			if err = w.Close(); err != nil {
				return
			}
		case data, ok := <-c.MsgBuffChan:
			log.Debugf("[ws.write] buff data:%v", string(data))
			c.Conn.SetWriteDeadline(time.Now().Add(c.WsServer.c.WriteWait))
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				log.Warnf("[ws.write] conn.NextWriter err :%v  ", err)
				return
			}
			w.Write(data)
			if err = w.Close(); err != nil {
				return
			}
		case <-c.Ctx.Done():
			return
		case <-ticker.C:
			//heartbeat，if ping error will exit and close current websocket conn
			c.Conn.SetWriteDeadline(time.Now().Add(c.WsServer.c.WriteWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Warnf("[ws.write] ticker ping wrr:%v", err)
				return
			}
		}
	}
}

//startReader 读消息Goroutine，用于从客户端中读取数据
func (c *Connection) startReader() {
	log.Debug("[ws.read] Reader Goroutine is running")
	defer c.Stop()

	c.Conn.SetReadLimit(c.WsServer.c.MaxPacketSize)
	c.Conn.SetReadDeadline(time.Now().Add(c.WsServer.c.PongWait))
	c.Conn.SetPongHandler(func(string) error {
		c.Conn.SetReadDeadline(time.Now().Add(c.WsServer.c.PongWait))
		return nil
	})

	for {
		select {
		case <-c.Ctx.Done():
			return
		default:
			_, message, err := c.Conn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Warnf("[ws.read] readPump ReadMessage err:%V", err)
					return
				}
			}
			if message == nil {
				return
			}
			if string(message) == "ping" {
				c.SendBuffMsg(0, []byte("pong"))
			} else {
				log.Debugf("[ws.read] reader message:%v", string(message))

				msg := &app.Message{}
				err = json.Unmarshal(message, msg)
				if err != nil {
					log.Warnf("[ws.read] json unmarshal err:%v", err)
					return
				}

				// 得到当前客户端请求的request数据
				req := connect.NewRequest(c, connect.WithJSONMsg(msg))

				if c.WsServer.c.WorkerPoolSize > 0 {
					// 已经启动工作池机制，将消息交给Worker处理
					c.WsServer.MsgHandler.SendMsgToTaskQueue(req)
				} else {
					go c.WsServer.MsgHandler.DoMsgHandler(req)
				}
			}
		}
	}
}

//Start 启动连接，让当前连接开始工作
func (c *Connection) Start() {
	c.Ctx, c.Cancel = context.WithCancel(context.Background())
	// 按照用户传递进来的创建连接时需要处理的业务，执行钩子方法
	if suc := c.WsServer.CallOnConnStart(c); !suc { // 初始化失败
		return
	}

	var upGrader = websocket.Upgrader{
		ReadBufferSize:  c.WsServer.c.ReadBufferSize,
		WriteBufferSize: c.WsServer.c.WriteBufferSize,
	}
	//cross origin domain support
	upGrader.CheckOrigin = func(r *http.Request) bool { return true }

	conn, err := upGrader.Upgrade(c.writer, c.request, nil)
	if err != nil {
		log.Warnf("[ws.start] serverWs err:%v", err)
		return
	}

	c.Conn = conn
	c.FreeLimit, err = c.WsServer.Limit.Accept(c.RemoteAddr())
	if err != nil {
		log.Warnf("[ws.start] limiter err:%v", err)
		c.Stop()
		return
	}
	// 开启用户从客户端读取数据的Goroutine
	go c.startReader()
	// 开启用于写回客户端数据的Goroutine
	go c.startWriter()
}

// Stop 关闭连接
func (c *Connection) Stop() {
	c.Lock()
	defer c.Unlock()

	c.FreeLimit()
	//如果当前链接已经关闭
	if c.isClosed == true {
		return
	}
	log.Debugf("[ws.stop] conn uid:%v", c.UID)
	//如果用户注册了该连接的关闭回调业务，那么在此调用
	c.WsServer.CallOnConnStop(c)

	//关闭socket连接
	err := c.Conn.Close()
	if err != nil {
		log.Warnf("[ws.stop] connection closed err:%v", err)
	}
	//关闭Writer
	c.Cancel()

	// 将连接从连接管理器中删除
	c.WsServer.GetConnMgr(c.UID).Remove(c)
	// 关闭该连接全部管道
	close(c.MsgBuffChan)
	//设置标志位
	c.isClosed = true
}

// GetWsConnection 获取websocket连接
func (c *Connection) GetWsConnection() *websocket.Conn {
	return c.Conn
}

// GetConnID 获取当前连接绑定的用户id
func (c *Connection) GetConnID() uint32 {
	return c.UID
}

//RemoteAddr 获取远程客户端地址信息
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

//SendMsg 直接将Message数据发送数据给远程的WS客户端
func (c *Connection) SendMsg(msgID uint32, msg []byte) error {
	c.RLock()
	defer c.RUnlock()
	if c.isClosed == true {
		return connect.ErrConnectNotFinish
	}

	//写回客户端
	c.MsgChan <- msg
	return nil
}

//SendBuffMsg  发生BuffMsg
func (c *Connection) SendBuffMsg(msgID uint32, msg []byte) error {
	c.RLock()
	defer c.RUnlock()
	if c.isClosed == true {
		return connect.ErrConnectNotFinish
	}

	//写回客户端
	c.MsgBuffChan <- msg
	return nil
}

//GetHTTPRequest 获取连接的请求对象
func (c *Connection) GetHTTPRequest() *http.Request {
	return c.request
}

// Auth 当前连接鉴权
func (c *Connection) Auth(userID uint32) {
	c.UID = userID
	//将新创建的Conn添加到链接管理中
	c.WsServer.GetConnMgr(c.UID).Add(c)
}
