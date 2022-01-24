package tcp

import (
	"context"
	"io"
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/prometheus/common/log"

	"chat-micro/pkg/logger"
	"chat-micro/pkg/server"
)

const (
	statusInit   = iota + 1 //连接初始化完成
	statusFinish            //连接已完成鉴权
	statusClosed            //连接已关闭
)

var _ server.Connection = (*tcpConnection)(nil)

// 连接状态
type status int8

//tcpConnection 连接
type tcpConnection struct {
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
	server *tcpServer
	//当前连接的socket TCP套接字
	conn *net.TCPConn
	//当前连接的状态
	status status
	//连接初始化定时器，未鉴权则关闭
	timer *time.Timer

	sync.RWMutex
}

// NewConnect 创建连接的方法
func NewConnect(s *tcpServer, conn *net.TCPConn, connID uint32) *tcpConnection {
	return &tcpConnection{
		server:      s,
		conn:        conn,
		id:          connID,
		msgChan:     make(chan []byte),
		msgBuffChan: make(chan []byte, s.Options().MaxMsgChanLen),
		status:      statusInit,
	}
}

//Context 上下文
func (c *tcpConnection) Context() context.Context {
	return c.ctx
}

//GetID 连接id
func (c *tcpConnection) GetID() uint32 {
	return c.id
}

//GetUID 连接绑定的用户id
func (c *tcpConnection) GetUID() uint32 {
	return c.uid
}

//startWriter 写消息Goroutine， 用户将数据发送给客户端
func (c *tcpConnection) startWriter() {
	logger.Debug("[tcp.write] Writer Goroutine is running")
	defer logger.Debugf("[tcp.write] %v conn Writer exit!", c.RemoteAddr().String())

	for {
		select {
		case data := <-c.msgChan:
			logger.Debugf("[tcp.write] data:%v", string(data))
			//write data dead time , like http timeout , default 10s
			c.conn.SetWriteDeadline(time.Now().Add(c.server.Options().WriteWait))
			// 有数据要写入客户端
			if _, err := c.conn.Write(data); err != nil {
				logger.Warnf("[tcp.write] Send Data err:%v", err)
				return
			}
		case data, ok := <-c.msgBuffChan:
			logger.Debugf("[tcp.write] buff data:%v", string(data))
			if ok {
				// 有数据要写客户端
				c.conn.SetWriteDeadline(time.Now().Add(c.server.Options().WriteWait))
				if _, err := c.conn.Write(data); err != nil {
					logger.Warnf("[tcp.write] Send Buff Data err:%v", err)
					return
				}
			} else {
				logger.Info("[tcp.write] msgBuffChan is Closed")
				return
			}
		case <-c.ctx.Done():
			return
		}
	}
}

//startReader 读消息Goroutine，用于从客户端中读取数据
func (c *tcpConnection) startReader() {
	logger.Debug("[tcp.read] Reader Goroutine is running")
	defer c.Stop()

	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			// 创建拆包解包对象
			opts := c.server.Options()
			//读取客户端的msg head
			headData := make([]byte, headLen)
			if _, err := io.ReadFull(c.conn, headData); err != nil {
				if err != io.EOF && err != io.ErrUnexpectedEOF {
					logger.Warnf("[tcp.read] read msg head err:%v", err)
				}
				return
			}
			//拆包，得到msgID 和 dataLen 放入msg中
			msg, err := unpack(headData)
			if err != nil {
				logger.Warnf("[tcp.read] unpack err:%v", err)
				return
			}

			//判断dataLen的长度是否超出我们允许的最大包长度
			if msg.length > opts.MaxPacketSize {
				logger.Info("[tcp.read] data length limit")
				return
			} else if msg.length > 0 {
				// 根据dataLen 读取data，放入msg.data中
				data := make([]byte, msg.length)
				if _, err = io.ReadFull(c.conn, data); err != nil {
					log.Warnf("[tcp.read] read msg data err:%v", err)
					return
				}
				msg.data = data
			}

			req := server.NewRequest(c, strconv.Itoa(int(msg.id)), msg.data, nil)
			if c.status == statusInit { //连接未鉴权只允许鉴权消息通过,否则关闭连接
				//如果用户注册了该连接的鉴权回调业务，那么在此调用
				if ok, uid := opts.OnConnAuth(c, req); ok { //鉴权成功
					c.auth(uid)
				} else { // 鉴权失败，关闭连接
					return
				}
			}
			logger.Debugf("[tcp.read] msgId:%v, body:%v", msg.id, string(msg.data))

			if opts.WorkerPoolSize > 0 {
				// 已经启动工作池机制，将消息交给Worker处理
				c.server.MsgHandler.SendMsgToTaskQueue(req)
			} else {
				go c.server.MsgHandler.DoMsgHandler(req)
			}
		}
	}
}

//Start 启动连接，让当前连接开始工作
func (c *tcpConnection) Start() {
	c.server.Options().OnConnStarting(c)

	c.ctx, c.cancel = context.WithCancel(context.Background())
	opts := c.server.Options()
	// 连接超时设置
	c.timer = time.AfterFunc(opts.Timeout, func() {
		if c.status == statusInit {
			logger.Infof("[tcp.start] handshake timeout id:%v", c.id)
			c.Stop()
		}
	})

	// 开启用户从客户端读取数据的Goroutine
	go c.startReader()
	// 开启用于写回客户端数据的Goroutine
	go c.startWriter()
	// 按照用户传递进来的创建连接时需要处理的业务，执行钩子方法
	opts.OnConnStarted(c)
}

//Stop 停止连接，结束当前连接状态M
func (c *tcpConnection) Stop() {
	c.Lock()
	defer c.Unlock()

	//如果当前连接已关闭
	if c.status == statusClosed {
		return
	}

	//如果用户注册了该连接的关闭回调业务，那么在此调用
	c.server.Options().OnConnStop(c)

	//关闭socket连接
	if err := c.conn.Close(); err != nil {
		logger.Warnf("[conn.stop] connection closed err:%v", err)
	}
	//关闭Writer
	c.cancel()

	if c.uid > 0 {
		// 将连接从连接管理器中删除
		c.server.Bucket.Remove(c)
	}
	// 关闭该连接全部管道
	close(c.msgBuffChan)
	//设置标志位
	c.status = statusClosed
	//清除定时器
	c.timer.Stop()
}

//RemoteAddr 获取远程客户端tcpConnection地址信息
func (c *tcpConnection) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

//SendMsg 直接将Message数据发送数据给远程的TCP客户端
func (c *tcpConnection) SendMsg(msgID uint32, data []byte) error {
	c.RLock()
	defer c.RUnlock()

	if c.status != statusFinish {
		return server.ErrConnectNotFinish
	}

	dp := &tcpMessage{
		id:     msgID,
		data:   data,
		length: len(data),
	}
	msg, err := pack(dp)
	if err != nil {
		return errors.Wrapf(err, "[tcp.sendMsg] pack id:%v,data:%v", msgID, data)
	}

	//写回客户端
	c.msgChan <- msg
	return nil
}

//SendBuffMsg  发生BuffMsg
func (c *tcpConnection) SendBuffMsg(msgID uint32, data []byte) error {
	c.RLock()
	defer c.RUnlock()

	if c.status != statusFinish {
		return server.ErrConnectNotFinish
	}

	// 将data封包，并发送
	dp := &tcpMessage{
		id:     msgID,
		data:   data,
		length: len(data),
	}
	msg, err := pack(dp)
	if err != nil {
		return errors.Wrapf(err, "[tcp.SendBuffMsg] pack id:%v,data:%v", msgID, data)
	}
	// 写回客户端
	c.msgBuffChan <- msg
	return nil
}

// auth 当前连接鉴权
func (c *tcpConnection) auth(userID uint32) {
	if c.status == statusFinish { // 无需重复鉴权
		return
	}
	c.timer.Stop()
	c.uid = userID
	c.status = statusFinish
	//将新创建的Conn添加到链接管理中
	c.server.Bucket.Insert(c)
}
