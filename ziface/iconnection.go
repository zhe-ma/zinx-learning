package ziface

import "net"

type IConnection interface {
	Start()
	Stop()
	GetTcpConnection() *net.TCPConn
	RemoteAddr() net.Addr

	SendMsg(msgID uint32, data []byte) error
}
