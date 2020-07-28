package ziface

type IDataPack interface {
	GetHeadLen() uint32
	Unpack(data []byte) (IMessage, error)
	Pack(msg IMessage) ([]byte, error)
}
