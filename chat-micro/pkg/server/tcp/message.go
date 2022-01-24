package tcp

import (
	"bytes"
	"encoding/binary"
)

const (
	//headLen 数据包头长度 ID uint32(4字节) +  DataLen uint32(4字节)
	headLen = 8
)

type tcpMessage struct {
	id     uint32 //消息的ID
	data   []byte //消息的内容
	length int    //消息的长度
}

//pack 封包方法(压缩数据)
func pack(msg *tcpMessage) ([]byte, error) {
	//创建一个存放bytes字节的缓冲
	buff := bytes.NewBuffer([]byte{})

	// 写dataLen
	if err := binary.Write(buff, binary.LittleEndian, msg.length); err != nil {
		return nil, err
	}

	// 写msgId
	if err := binary.Write(buff, binary.LittleEndian, msg.id); err != nil {
		return nil, err
	}

	// 写data数据
	if err := binary.Write(buff, binary.LittleEndian, msg.data); err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}

//unpack 拆包方法(解压数据)
func unpack(binaryData []byte) (*tcpMessage, error) {
	//创建一个从输入二进制数据的ioReader
	dataBuff := bytes.NewReader(binaryData)

	//只解压head的信息，得到dataLen和msgID
	msg := &tcpMessage{}

	//读dataLen
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.length); err != nil {
		return nil, err
	}

	//读msgID
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.id); err != nil {
		return nil, err
	}

	//这里只需要把head的数据拆包出来就可以了，然后再通过head的长度，再从conn读取一次数据
	return msg, nil
}
