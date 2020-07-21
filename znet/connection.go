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
	router       ziface.IRouter
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

		request := NewRequest(c, buf[:count])
		c.router.PreHandle(request)
		c.router.Handle(request)
		c.router.PostHandle(request)
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

func (c *Connection) GetTcpConnection() *net.TCPConn {
	return c.conn
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func NewConnection(conn *net.TCPConn, router ziface.IRouter) *Connection {
	return &Connection{
		conn:         conn,
		exitBuffChan: make(chan bool, 1),
		isClosed:     false,
		router:       router,
	}
}
