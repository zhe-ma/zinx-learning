package znet

import (
	"zinx-learning/ziface"
)

type Request struct {
	Conn ziface.IConnection
	Msg  ziface.IMessage
}

func (r *Request) GetConnection() ziface.IConnection {
	return r.Conn
}

func (r *Request) GetData() []byte {
	return r.Msg.GetData()
}

func (r *Request) GetMsgID() uint32 {
	return r.Msg.GetMsgID()
}

func NewRequest(conn ziface.IConnection, msg ziface.IMessage) *Request {
	return &Request{
		Conn: conn,
		Msg:  msg,
	}
}
