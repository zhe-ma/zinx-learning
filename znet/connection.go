package znet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"zinx-learning/ziface"
)

type Connection struct {
	Conn         *net.TCPConn
	ExitBuffChan chan bool
	IsClosed     bool
	Router       ziface.IRouter
}

func (c *Connection) startReader() {
	fmt.Println("Reader Goroutine is running.")
	defer fmt.Println(c.RemoteAddr().String(), " connetion exit.")

	defer c.Stop()

	for {

		dataPack := DataPack{}

		headData := make([]byte, dataPack.GetHeadLen())
		if _, err := io.ReadFull(c.Conn, headData); err != nil {
			fmt.Println("Failed to read data:", err)
			c.ExitBuffChan <- true
			continue
		}

		msg, err := dataPack.Unpack(headData)
		if err != nil {
			fmt.Println("Failed to unpack data:", err)
			c.ExitBuffChan <- true
			continue
		}

		var data []byte
		if msg.GetDataLen() > 0 {
			data = make([]byte, msg.GetDataLen())
			if _, err := io.ReadFull(c.Conn, data); err != nil {
				fmt.Println("Failed to read data:", err)
				c.ExitBuffChan <- true
				continue
			}
		}

		msg.SetData(data)

		request := NewRequest(c, msg)
		c.Router.PreHandle(request)
		c.Router.Handle(request)
		c.Router.PostHandle(request)
	}
}

func (c *Connection) Start() {
	go c.startReader()

	for {
		select {
		case <-c.ExitBuffChan:
			return
		}
	}
}

func (c *Connection) Stop() {
	if c.IsClosed {
		return
	}

	c.IsClosed = true

	c.Conn.Close()

	c.ExitBuffChan <- true

	close(c.ExitBuffChan)
}

func (c *Connection) GetTcpConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) SendMsg(msgID uint32, data []byte) error {
	if c.IsClosed {
		return errors.New("Connection closed when send data")
	}

	dataPack := DataPack{}
	packedData, err := dataPack.Pack(NewMessage(msgID, data))
	if err != nil {
		fmt.Println("Failed to pack data:", err, "MsgId:", msgID)
		return errors.New("Failed to pack data")
	}

	_, err = c.Conn.Write(packedData)
	if err != nil {
		fmt.Println("Failed to write data:", err, "MsgId:", msgID)
		c.ExitBuffChan <- true
		return errors.New("Failed to wriete data")
	}

	return nil
}

func NewConnection(conn *net.TCPConn, router ziface.IRouter) *Connection {
	return &Connection{
		Conn:         conn,
		ExitBuffChan: make(chan bool, 1),
		IsClosed:     false,
		Router:       router,
	}
}
