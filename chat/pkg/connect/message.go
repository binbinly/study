package connect

// 用于触发编译期的接口的合理性检查机制
var _ IMessage = (*Message)(nil)

//IMessage 将请求的一个消息封装到message中，定义抽象层接口
type IMessage interface {
	GetDataLen() int  //获取消息数据段长度
	GetMsgID() uint32 //获取消息ID
	GetData() []byte  //获取消息内容

	SetData(data []byte) // 设置消息内容
}

// Message 消息
type Message struct {
	DataLen int    // 消息的长度
	ID      uint32 // 消息的ID
	Data    []byte // 消息的内容
}

// NewMsgPackage 创建一个Message消息包
func NewMsgPackage(id uint32, data []byte) IMessage {
	return &Message{
		DataLen: len(data),
		ID:      id,
		Data:    data,
	}
}

// GetDataLen 获取消息数据段长度
func (m *Message) GetDataLen() int {
	return m.DataLen
}

// GetMsgID 获取消息ID
func (m *Message) GetMsgID() uint32 {
	return m.ID
}

// GetData 获取消息内容
func (m *Message) GetData() []byte {
	return m.Data
}

// SetData 设置数据
func (m *Message) SetData(data []byte) {
	m.Data = data
}
