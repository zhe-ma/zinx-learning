package ziface

import "net"

type HandleFunc func(conn *net.TCPConn, data []byte, count int) error

type IConnection interface {
	Start()
	Stop()
	RemoteAddr() net.Addr
}
