package main

import (
	"fmt"
	"zinx-learning/ziface"
	"zinx-learning/znet"
)

type PingRouter struct {
}

func (pr *PingRouter) PreHandle(req ziface.IRequtest) {
	fmt.Println("[PingRouter] PreHandle.")
	if _, err := req.GetConnection().GetTcpConnection().Write([]byte("Before ping...")); err != nil {
		fmt.Println("Failed to write data.")
	}
}

func (pr *PingRouter) Handle(req ziface.IRequtest) {
	fmt.Println("[PingRouter] Handle.")
	if _, err := req.GetConnection().GetTcpConnection().Write([]byte("Ping...")); err != nil {
		fmt.Println("Failed to write data.")
	}
}

func (pr *PingRouter) PostHandle(req ziface.IRequtest) {
	fmt.Println("[PingRouter] PostHandle.")
	if _, err := req.GetConnection().GetTcpConnection().Write([]byte("Post ping...")); err != nil {
		fmt.Println("Failed to write data.")
	}
}

func main() {
	server := znet.NewServer("Test")
	server.AddRouter(&PingRouter{})
	server.Serve()
}
