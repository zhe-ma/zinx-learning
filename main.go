package main

import "zinx-learning/znet"

func main() {
	server := znet.NewServer("Test")
	server.Serve()
}
