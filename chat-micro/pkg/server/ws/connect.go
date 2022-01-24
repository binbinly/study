package ws

import (
	"context"
	"encoding/json"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"

	"chat-micro/pkg/logger"
	"chat-micro/pkg/server"
)

var _ server.Connection = (*wsConnection)(nil)

const (
	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
)

//wsConnection 连接
type wsConnection struct {
	//当前连接的ID 也可以称作为SessionID，ID全局唯一
	id uint32
	//用户ID
	uid uint32
	//告知该链接已经退出/停止的channel
	ctx    context.Context
	cancel context.CancelFunc
	//无缓冲管道，用于读、写两个goroutine之间的消息通信
	msgChan chan []byte
	//有缓冲管道，用于读、写两个goroutine之间的消息通信
	msgBuffChan chan []byte
	//当前Conn属于哪个Server
	server *wsServer
	//连接对象
	conn *websocket.Conn
	//websocket响应对象
	writer http.ResponseWriter
	//websocket请求对象
	request *http.Request
	//当前连接的关闭状态
	isClosed bool

	sync.RWMutex
}

// NewConnect 创建连接的方法
func NewConnect(s *wsServer, w http.ResponseWriter, r *http.Request, connID uint32) *wsConnection {
	return &wsConnection{
		server:      s,
		writer:      w,
		request:     r,
		id:          connID,
		msgChan:     make(chan []byte),
		msgBuffChan: make(chan []byte, s.Options().MaxMsgChanLen),
	}
}

//startWriter 写消息Goroutine， 用户将数据发送给客户端
func (c *wsConnection) startWriter() {
	logger.Debug("[ws.write] Writer Goroutine is running")
	// 心跳由客户端发送 ping 回复 pong
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		defer logger.Debugf("[ws.write] %v conn Writer exit!", c.RemoteAddr().String())
		ticker.Stop()
	}()

	for {
		select {
		case data := <-c.msgChan:
			logger.Debugf("[ws.write] data:%v", string(data))
			//write data dead time , like http timeout , default 10s
			c.conn.SetWriteDeadline(time.Now().Add(c.server.Options().WriteWait))
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				logger.Warnf("[ws.write] conn.NextWriter err :%v  ", err)
				return
			}
			w.Write(data)
			if err = w.Close(); err != nil {
				return
			}
		case data, ok := <-c.msgBuffChan:
			logger.Debugf("[ws.write] buff data:%v", string(data))
			c.conn.SetWriteDeadline(time.Now().Add(c.server.Options().WriteWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				logger.Warnf("[ws.write] conn.NextWriter err :%v  ", err)
				return
			}
			w.Write(data)
			if err = w.Close(); err != nil {
				return
			}
		case <-c.ctx.Done():
			return
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(c.server.Options().WriteWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				logger.Warnf("[ws.write] ticker ping wrr:%v", err)
				return
			}
		}
	}
}

//startReader 读消息Goroutine，用于从客户端中读取数据
func (c *wsConnection) startReader() {
	logger.Debug("[ws.read] Reader Goroutine is running")
	defer c.Stop()

	c.conn.SetReadLimit(int64(c.server.Options().MaxPacketSize))
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			_, message, err := c.conn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					logger.Warnf("[ws.read] readPump ReadMessage err:%V", err)
					return
				}
			}
			if message == nil {
				return
			}
			if string(message) == "ping" {
				c.SendBuffMsg(0, []byte("pong"))
			} else {
				logger.Debugf("[ws.read] reader message:%v", string(message))

				msg := &wsMessage{}
				if err = json.Unmarshal(message, msg); err != nil {
					logger.Warnf("[ws.read] json unmarshal err:%v", err)
					return
				}
				// 构建当前客户端请求的request数据
				req := server.NewRequest(c, msg.Event, msg.Data, nil)
				if c.server.Options().WorkerPoolSize > 0 {
					// 已经启动工作池机制，将消息交给Worker处理
					c.server.MsgHandler.SendMsgToTaskQueue(req)
				} else {
					go c.server.MsgHandler.DoMsgHandler(req)
				}
			}
		}
	}
}

//Start 启动连接，让当前连接开始工作
func (c *wsConnection) Start() {
	c.ctx, c.cancel = context.WithCancel(context.Background())

	c.server.Options().OnConnStarting(c)

	req := server.NewRequest(c, "auth", []byte(c.request.URL.Query().Get("token")), nil)
	//如果用户注册了该连接的鉴权回调业务，那么在此调用
	if ok, uid := c.server.Options().OnConnAuth(c, req); ok { //鉴权成功
		c.auth(uid)
	} else { // 鉴权失败
		return
	}

	var upGrader = websocket.Upgrader{
		ReadBufferSize:  c.server.Options().ReadBufferSize,
		WriteBufferSize: c.server.Options().WriteBufferSize,
	}
	//cross origin domain support
	upGrader.CheckOrigin = func(r *http.Request) bool { return true }

	wsSocket, err := upGrader.Upgrade(c.writer, c.request, nil)
	if err != nil {
		logger.Warnf("[ws.start] Upgrade err:%v", err)
		return
	}
	c.conn = wsSocket

	// 开启用户从客户端读取数据的Goroutine
	go c.startReader()
	// 开启用于写回客户端数据的Goroutine
	go c.startWriter()
	// 按照用户传递进来的创建连接时需要处理的业务，执行钩子方法
	c.server.Options().OnConnStarted(c)
}

// Stop 关闭连接
func (c *wsConnection) Stop() {
	c.Lock()
	defer c.Unlock()

	//如果当前链接已经关闭
	if c.isClosed == true {
		return
	}

	//如果用户注册了该连接的关闭回调业务，那么在此调用
	c.server.Options().OnConnStop(c)

	//关闭socket连接
	if err := c.conn.Close(); err != nil {
		logger.Warnf("[ws.stop] connection closed err:%v", err)
	}
	//关闭Writer
	c.cancel()

	// 将连接从连接管理器中删除
	c.server.Bucket.Remove(c)
	// 关闭该连接全部管道
	close(c.msgBuffChan)
	//设置标志位
	c.isClosed = true
}

func (c *wsConnection) Context() context.Context {
	return c.ctx
}

// GetID 获取连接id
func (c *wsConnection) GetID() uint32 {
	return c.id
}

// GetUID 获取连接绑定的用户id
func (c *wsConnection) GetUID() uint32 {
	return c.uid
}

//RemoteAddr 获取远程客户端地址信息
func (c *wsConnection) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

//SendMsg 直接将Message数据发送数据给远程的WS客户端
func (c *wsConnection) SendMsg(msgID uint32, msg []byte) error {
	c.RLock()
	defer c.RUnlock()
	if c.isClosed == true {
		return server.ErrConnectNotFinish
	}

	//写回客户端
	c.msgChan <- msg
	return nil
}

//SendBuffMsg 发生BuffMsg
func (c *wsConnection) SendBuffMsg(msgID uint32, msg []byte) error {
	c.RLock()
	defer c.RUnlock()
	if c.isClosed == true {
		return server.ErrConnectNotFinish
	}

	//写回客户端
	c.msgBuffChan <- msg
	return nil
}

// auth 当前连接鉴权
func (c *wsConnection) auth(userID uint32) {
	c.uid = userID
	//将新创建的Conn添加到链接管理中
	c.server.Bucket.Insert(c)
}
