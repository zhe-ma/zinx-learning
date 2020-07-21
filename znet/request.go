package znet

import (
	"zinx-learning/ziface"
)

type Request struct {
	conn ziface.IConnection
	data []byte
}

func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.data
}

func NewRequest(conn ziface.IConnection, data []byte) ziface.IRequtest {
	return &Request{
		conn: conn,
		data: data,
	}
}
