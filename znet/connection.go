package znet

import (
	"fmt"
	"net"
	"zinx-learning/ziface"
)

type Connection struct {
	conn         *net.TCPConn
	exitBuffChan chan bool
	isClosed     bool
	handler      ziface.HandleFunc
}

func (c *Connection) startReader() {
	fmt.Println("Reader Goroutine is running.")
	defer fmt.Println(c.RemoteAddr().String(), " connetion exit.")

	defer c.Stop()

	for {
		buf := make([]byte, 512)
		count, err := c.conn.Read(buf)
		if err != nil {
			fmt.Println("Failed to read data:", err)
			c.exitBuffChan <- true
			continue
		}

		if err := c.handler(c.conn, buf, count); err != nil {
			fmt.Println("Failed to call handler:", err)
			c.exitBuffChan <- true
			return
		}
	}
}

func (c *Connection) Start() {
	go c.startReader()

	for {
		select {
		case <-c.exitBuffChan:
			return
		}
	}
}

func (c *Connection) Stop() {
	if c.isClosed {
		return
	}

	c.isClosed = true

	c.conn.Close()

	c.exitBuffChan <- true

	close(c.exitBuffChan)
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func NewConnection(conn *net.TCPConn, handler ziface.HandleFunc) *Connection {
	return &Connection{
		conn:         conn,
		exitBuffChan: make(chan bool, 1),
		isClosed:     false,
		handler:      handler,
	}
}
