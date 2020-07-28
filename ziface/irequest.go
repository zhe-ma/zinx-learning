package ziface

type IRequest interface {
	GetConnection() IConnection
	GetMsgID() uint32
	GetData() []byte
}
