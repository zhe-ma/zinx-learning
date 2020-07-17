package znet

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	server := NewServer("test")
	server.Start()

	time.Sleep(time.Second)

	client, err := net.Dial("tcp4", "127.0.0.1:9547")
	if err != nil {
		t.Error(err)
	}

	if _, err = client.Write([]byte("Hello World")); err != nil {
		t.Error(err)
	}

	buf := make([]byte, 512)
	count, err := client.Read(buf)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(buf[:count]))

	client.Close()
}
