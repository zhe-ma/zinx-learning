package ziface

type IMessage interface {
	GetMsgID() uint32
	SetMsgID(id uint32)

	GetDataLen() uint32
	SetDataLen(dataLen uint32)

	GetData() []byte
	SetData(data []byte)
}
