package main

import (
	"fmt"
	"zinx-learning/ziface"
	"zinx-learning/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

func (pr *PingRouter) PreHandle(req ziface.IRequest) {
	fmt.Println("[PingRouter] PreHandle.")
}

func (pr *PingRouter) Handle(req ziface.IRequest) {
	fmt.Println("[PingRouter] Handle.")

	fmt.Println("Request: MsgID=", req.GetMsgID(), "MsgData:", string(req.GetData()))

	if err := req.GetConnection().SendMsg(1, []byte("Pong pong...")); err != nil {
		fmt.Println("Failed to write data:", err)
	}
}

func (pr *PingRouter) PostHandle(req ziface.IRequest) {
	fmt.Println("[PingRouter] PostHandle.")
}

func main() {
	server := znet.NewServer()
	server.AddRouter(&PingRouter{})
	server.Serve()
}
