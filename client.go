package main

import (
	"fmt"
	"net"
)

func main() {
	client, err := net.Dial("tcp4", "0.0.0.0:9547")
	if err != nil {
		panic(err)
	}
	defer client.Close()

	for i := 0; i < 5; i++ {
		if _, err := client.Write([]byte("Hello world!")); err != nil {
			panic(err)
		}

		buf := make([]byte, 512)
		count, err := client.Read(buf)
		if err != nil {
			panic(err)
		}

		fmt.Println("Recv:", string(buf[:count]))
	}
}
