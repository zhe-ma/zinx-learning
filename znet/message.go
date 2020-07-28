package znet

type Message struct {
	ID      uint32
	DataLen uint32
	Data    []byte
}

func (msg *Message) GetMsgID() uint32 {
	return msg.ID
}

func (msg *Message) SetMsgID(id uint32) {
	msg.ID = id
}

func (msg *Message) GetDataLen() uint32 {
	return msg.DataLen
}

func (msg *Message) SetDataLen(dataLen uint32) {
	msg.DataLen = dataLen
}

func (msg *Message) GetData() []byte {
	return msg.Data
}

func (msg *Message) SetData(data []byte) {
	msg.Data = data
}

func NewMessage(id uint32, data []byte) *Message {
	return &Message{
		ID:      id,
		DataLen: uint32(len(data)),
		Data:    data,
	}
}
