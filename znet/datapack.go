package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"zinx-learning/utils"
	"zinx-learning/ziface"
)

type DataPack struct {
}

func (dp *DataPack) GetHeadLen() uint32 {
	// MsgID(4) + MsgDataLen(4)
	return 8
}

func (dp *DataPack) Unpack(data []byte) (ziface.IMessage, error) {
	dataBuff := bytes.NewReader(data)

	msg := &Message{}

	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}

	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.ID); err != nil {
		return nil, err
	}

	if utils.GlobalObj.MaxPacketSize > 0 && utils.GlobalObj.MaxPacketSize < msg.DataLen {
		fmt.Println("size:", utils.GlobalObj.MaxPacketSize, msg.DataLen)
		return nil, errors.New("Too large msg data received.")
	}

	// 只读出包头，数据根据DataLen后面读出。
	return msg, nil
}

func (dp *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	dataBuff := bytes.NewBuffer([]byte{})

	// 1． Little endian：将低序字节存储在起始地址（低位字节存储在内存中低位地址）符合思维逻辑
	// 2． Big endian：将高序字节存储在起始地址（高位字节存储在内存中低位地址）看起来直观
	// x86架构的cpu不管操作系统是NT还是unix系的，都是小字节序。而PowerPC 、SPARC和Motorola处理器则很多是大字节序。
	// 网络字节顺序采用big endian排序方式。
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}

	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgID()); err != nil {
		return nil, err
	}

	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}

	return dataBuff.Bytes(), nil
}
