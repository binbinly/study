package tcp

import (
	"context"
	"io"
	"net"
	"time"

	"github.com/pkg/errors"

	"chat/pkg/connect"
	"chat/pkg/log"
)

var _ connect.IConnection = (*Connection)(nil)

//Connection 连接
type Connection struct {
	*connect.Connection
	//当前Conn属于哪个Server
	TCPServer *Server
	//当前连接的socket TCP套接字
	Conn *net.TCPConn
	//当前连接的状态
	status status
	//连接初始化定时器，未鉴权则关闭
	timer *time.Timer
}

// NewConnect 创建连接的方法
func NewConnect(s *Server, conn *net.TCPConn, connID uint32) connect.IConnection {
	//初始化Conn属性
	return &Connection{
		Connection: &connect.Connection{
			ConnID:      connID,
			MsgChan:     make(chan []byte),
			MsgBuffChan: make(chan []byte, s.c.MaxMsgChanLen),
		},
		TCPServer: s,
		Conn:      conn,
		status:    statusInit,
	}
}

//startWriter 写消息Goroutine， 用户将数据发送给客户端
func (c *Connection) startWriter() {
	log.Debug("[tcp.write] Writer Goroutine is running")
	defer log.Debugf("[tcp.write] %v conn Writer exit!", c.RemoteAddr().String())

	for {
		select {
		case data := <-c.MsgChan:
			log.Debugf("[tcp.write] data:%v", string(data))
			// 有数据要写入客户端
			if _, err := c.Conn.Write(data); err != nil {
				log.Warnf("[tcp.write] Send Data err:%v", err)
				return
			}
		case data, ok := <-c.MsgBuffChan:
			log.Debugf("[tcp.write] buff data:%v", string(data))
			if ok {
				// 有数据要写客户端
				if _, err := c.Conn.Write(data); err != nil {
					log.Warnf("[tcp.write] Send Buff Data err:%v", err)
					return
				}
			} else {
				log.Info("[tcp.write] msgBuffChan is Closed")
				break
			}
		case <-c.Ctx.Done():
			return
		}
	}
}

//startReader 读消息Goroutine，用于从客户端中读取数据
func (c *Connection) startReader() {
	log.Debug("[tcp.read] Reader Goroutine is running")
	defer c.Stop()

	for {
		select {
		case <-c.Ctx.Done():
			return
		default:
			// 创建拆包解包对象
			dp := NewDataPack()

			//读取客户端的msg head
			headData := make([]byte, dp.GetHeadLen())
			if _, err := io.ReadFull(c.Conn, headData); err != nil {
				if err != io.EOF && err != io.ErrUnexpectedEOF {
					log.Warnf("[tcp.read] read msg head err:%v", err)
				}
				return
			}

			//拆包，得到msgID 和 dataLen 放入msg中
			msg, err := dp.Unpack(headData)
			if err != nil {
				log.Warnf("[tcp.read] unpack err:%v", err)
				return
			}

			if c.status == statusInit { //连接未鉴权只允许鉴权消息通过,否则关闭连接
				if msg.GetMsgID() == MsgIDAuth {
					//如果用户注册了该连接的鉴权回调业务，那么在此调用
					if suc := c.TCPServer.CallOnConnAuth(c, msg.GetData()); !suc { //鉴权失败，关闭连接
						return
					}
				} else {
					log.Debugf("[tcp.read] head data msgId:%v not auth to close", msg.GetMsgID())
					return
				}
			}

			// 根据dataLen 读取data，放入msg.data中
			var data []byte
			if msg.GetDataLen() > 0 {
				data = make([]byte, msg.GetDataLen())
				if _, err = io.ReadFull(c.Conn, data); err != nil {
					log.Warnf("[tcp.read] read msg data err:%v", err)
					return
				}
			}
			msg.SetData(data)
			log.Debugf("[tcp.read] msgId:%v, data:%v", msg.GetMsgID(), string(data))

			// 得到当前客户端请求的request数据
			req := connect.NewRequest(c, connect.WithBinaryMsg(msg))

			if c.TCPServer.c.WorkerPoolSize > 0 {
				// 已经启动工作池机制，将消息交给Worker处理
				c.TCPServer.MsgHandler.SendMsgToTaskQueue(req)
			} else {
				go c.TCPServer.MsgHandler.DoMsgHandler(req)
			}
		}
	}
}

//Start 启动连接，让当前连接开始工作
func (c *Connection) Start() {
	var err error
	c.FreeLimit, err = c.TCPServer.Limit.Accept(c.RemoteAddr())
	if err != nil {
		log.Warnf("[tcp.start] limiter err:%v", err)
		c.Stop()
		return
	}
	c.Ctx, c.Cancel = context.WithCancel(context.Background())
	// 连接超时设置
	c.timer = time.AfterFunc(c.TCPServer.c.HandshakeTimeout, func() {
		if c.status == statusInit {
			log.Infof("[tcp.start] handshake timeout id:%v", c.ConnID)
			c.Stop()
		}
	})

	// 开启用户从客户端读取数据的Goroutine
	go c.startReader()
	// 开启用于写回客户端数据的Goroutine
	go c.startWriter()
	// 按照用户传递进来的创建连接时需要处理的业务，执行钩子方法
	c.TCPServer.CallOnConnStart(c)
}

//Stop 停止连接，结束当前连接状态M
func (c *Connection) Stop() {
	c.Lock()
	defer c.Unlock()

	log.Debugf("[tcp.stop] conn id:%v", c.ConnID)
	c.FreeLimit()
	//如果当前连接已关闭
	if c.status == statusClosed {
		return
	}

	//如果用户注册了该连接的关闭回调业务，那么在此调用
	c.TCPServer.CallOnConnStop(c)

	//关闭socket连接
	err := c.Conn.Close()
	if err != nil {
		log.Warnf("[conn.stop] connection closed err:%v", err)
	}
	//关闭Writer
	c.Cancel()

	if c.UID > 0 {
		// 将连接从连接管理器中删除
		c.TCPServer.GetConnMgr(c.UID).Remove(c)
	}
	// 关闭该连接全部管道
	close(c.MsgBuffChan)
	//设置标志位
	c.status = statusClosed
	//清除定时器
	c.timer.Stop()
}

//GetTCPConnection 从当前连接获取原始的socket TCPConn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

//RemoteAddr 获取远程客户端地址信息
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

//SendMsg 直接将Message数据发送数据给远程的TCP客户端
func (c *Connection) SendMsg(msgID uint32, data []byte) error {
	c.RLock()
	defer c.RUnlock()

	if c.status != statusFinish {
		return connect.ErrConnectNotFinish
	}

	// 将data封包，并且发送
	dp := NewDataPack()
	msg, err := dp.Pack(connect.NewMsgPackage(msgID, data))
	if err != nil {
		return errors.Wrapf(err, "[tcp.sendMsg] pack id:%v,data:%v", msgID, data)
	}
	//写回客户端
	c.MsgChan <- msg
	return nil
}

//SendBuffMsg  发生BuffMsg
func (c *Connection) SendBuffMsg(msgID uint32, data []byte) error {
	c.RLock()
	defer c.RUnlock()

	if c.status != statusFinish {
		return connect.ErrConnectNotFinish
	}

	// 将data封包，并发送
	dp := NewDataPack()
	msg, err := dp.Pack(connect.NewMsgPackage(msgID, data))
	if err != nil {
		return errors.Wrapf(err, "[tcp.SendBuffMsg] pack id:%v,data:%v", msgID, data)
	}
	// 写回客户端
	c.MsgBuffChan <- msg
	return nil
}

// Auth 当前连接鉴权
func (c *Connection) Auth(userID uint32) {
	if c.status == statusFinish { // 无需重复鉴权
		return
	}
	c.timer.Stop()
	c.UID = userID
	c.status = statusFinish
	//将新创建的Conn添加到链接管理中
	c.TCPServer.GetConnMgr(c.UID).Add(c)
}
