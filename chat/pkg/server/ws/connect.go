package ws

import (
	"context"
	"encoding/json"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/websocket"

	"chat/pkg/app"
	"chat/pkg/log"
	"chat/pkg/server"
)

//Connection 连接
type Connection struct {
	*server.Connection
	//当前Conn属于哪个Server
	WsServer *Server
	//连接对象
	Conn *websocket.Conn
	//websocket响应对象
	writer http.ResponseWriter
	//websocket请求对象
	request *http.Request
	//消息管理MsgID和对应处理方法的消息管理模块
	MsgHandler IMsgHandle
	//当前连接的关闭状态
	isClosed bool
}

// NewConnect 创建连接的方法
func NewConnect(s *Server, w http.ResponseWriter, r *http.Request) *Connection {
	//初始化Conn属性
	c := &Connection{
		Connection: &server.Connection{
			MsgChan:     make(chan []byte),
			MsgBuffChan: make(chan []byte, s.Config.MaxMsgChanLen),
		},
		WsServer:   s,
		writer:     w,
		request:    r,
		MsgHandler: s.msgHandler,
	}
	return c
}

//StartWriter 写消息Goroutine， 用户将数据发送给客户端
func (c *Connection) StartWriter() {
	log.Info("[ws.write] Writer Goroutine is running")
	ticker := time.NewTicker(c.WsServer.Config.PingPeriod)
	defer func() {
		defer log.Infof("[ws.write] %v conn Writer exit!", c.RemoteAddr().String())
		ticker.Stop()
	}()
	for {
		select {
		case data := <-c.MsgChan:
			log.Infof("[ws.write] data:%v", string(data))
			//write data dead time , like http timeout , default 10s
			c.Conn.SetWriteDeadline(time.Now().Add(c.WsServer.Config.WriteWait))
			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				log.Warnf("[ws.write] conn.NextWriter err :%v  ", err)
				return
			}
			w.Write(data)
			if err := w.Close(); err != nil {
				return
			}
		case data, ok := <-c.MsgBuffChan:
			log.Infof("[ws.write] buff data:%v", string(data))
			c.Conn.SetWriteDeadline(time.Now().Add(c.WsServer.Config.WriteWait))
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
			c.Conn.SetWriteDeadline(time.Now().Add(c.WsServer.Config.WriteWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Warnf("[ws.write] ticker ping wrr:%v", err)
				return
			}
		}
	}
}

//StartReader 读消息Goroutine，用于从客户端中读取数据
func (c *Connection) StartReader() {
	log.Info("[ws.read] Reader Goroutine is running")
	defer log.Infof("[ws.read] %v conn Reader exit", c.RemoteAddr().String())
	defer c.Stop()

	c.Conn.SetReadLimit(c.WsServer.Config.MaxPacketSize)
	c.Conn.SetReadDeadline(time.Now().Add(c.WsServer.Config.PongWait))
	c.Conn.SetPongHandler(func(string) error {
		c.Conn.SetReadDeadline(time.Now().Add(c.WsServer.Config.PongWait))
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
					log.Errorf("[ws.read] readPump ReadMessage err:%V", err)
					return
				}
			}
			if message == nil {
				return
			}
			if string(message) == "ping" {
				c.Conn.SetWriteDeadline(time.Now().Add(c.WsServer.Config.WriteWait))
				if err := c.Conn.WriteMessage(websocket.PongMessage, nil); err != nil {
					log.Warnf("[ws.read] pong wrr:%v", err)
					return
				}
			} else {
				log.Infof("[ws.read] reader message:%v", string(message))

				msg := &app.Message{}
				err = json.Unmarshal(message, msg)
				if err != nil {
					log.Errorf("[ws.read] json unmarshal err:%v", err)
					return
				}

				// 得到当前客户端请求的request数据
				req := NewRequest(c, msg)

				if c.WsServer.Config.WorkerPoolSize > 0 {
					// 已经启动工作池机制，将消息交给Worker处理
					c.MsgHandler.SendMsgToTaskQueue(req)
				} else {
					go c.MsgHandler.DoMsgHandler(req)
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
		ReadBufferSize:  c.WsServer.Config.ReadBufferSize,
		WriteBufferSize: c.WsServer.Config.WriteBufferSize,
	}
	//cross origin domain support
	upGrader.CheckOrigin = func(r *http.Request) bool { return true }

	conn, err := upGrader.Upgrade(c.writer, c.request, nil)
	if err != nil {
		log.Errorf("[ws.start] serverWs err:%v", err)
		return
	}

	c.Conn = conn
	// 开启用户从客户端读取数据的Goroutine
	go c.StartReader()
	// 开启用于写回客户端数据的Goroutine
	go c.StartWriter()
}

// Stop 关闭连接
func (c *Connection) Stop() {
	c.Lock()
	defer c.Unlock()

	//如果当前链接已经关闭
	if c.isClosed == true {
		return
	}
	log.Infof("[ws.stop] conn uid:%v", c.Uid)
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
	c.WsServer.GetConnMgr(c.Uid).Remove(c)
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
	return c.Uid
}

//RemoteAddr 获取远程客户端地址信息
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

//SendMsg 直接将Message数据发送数据给远程的WS客户端
func (c *Connection) SendMsg(msg []byte) error {
	c.RLock()
	defer c.RUnlock()
	if c.isClosed == true {
		return server.ErrConnectNotFinish
	}

	//写回客户端
	c.MsgChan <- msg
	return nil
}

//SendBuffMsg  发生BuffMsg
func (c *Connection) SendBuffMsg(msg []byte) error {
	c.RLock()
	defer c.RUnlock()
	if c.isClosed == true {
		return server.ErrConnectNotFinish
	}

	//写回客户端
	c.MsgBuffChan <- msg
	return nil
}

//GetRequest 获取连接的请求对象
func (c *Connection) GetRequest() *http.Request {
	return c.request
}

// Auth 当前连接鉴权
func (c *Connection) Auth(userId uint32) {
	c.Uid = userId
	//将新创建的Conn添加到链接管理中
	c.WsServer.GetConnMgr(c.Uid).Add(c)
	//如果用户注册了该连接的鉴权完成回调业务，那么在此调用
	c.WsServer.CallOnConnFinish(c)
}
