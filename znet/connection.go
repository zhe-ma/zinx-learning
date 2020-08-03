package znet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"zinx-learning/utils"
	"zinx-learning/ziface"
)

type Connection struct {
	Conn         *net.TCPConn
	ExitBuffChan chan bool
	IsClosed     bool
	ConnID       uint32
	MsgHandler   ziface.IMsgHandler
	MsgChan      chan []byte
}

func (c *Connection) startReader() {
	fmt.Println("Reader Goroutine is running.")
	defer fmt.Println(c.RemoteAddr().String(), "reader connection exit.")

	defer c.Stop()

	for {

		dataPack := DataPack{}

		headData := make([]byte, dataPack.GetHeadLen())
		if _, err := io.ReadFull(c.Conn, headData); err != nil {
			fmt.Println("Failed to read data:", err)
			break
		}

		msg, err := dataPack.Unpack(headData)
		if err != nil {
			fmt.Println("Failed to unpack data:", err)
			break
		}

		var data []byte
		if msg.GetDataLen() > 0 {
			data = make([]byte, msg.GetDataLen())
			if _, err := io.ReadFull(c.Conn, data); err != nil {
				fmt.Println("Failed to read data:", err)
				break
			}
		}

		msg.SetData(data)

		request := NewRequest(c, msg)

		if utils.GlobalObj.WorkPoolSize > 1 {
			c.MsgHandler.SendToTaskQueue(request)
		} else {
			go c.MsgHandler.HandleMsg(request)
		}
	}
}

func (c *Connection) startWriter() {
	fmt.Println("Writer Goroutine is running.")
	defer fmt.Println(c.RemoteAddr().String(), "writer connection exit.")

	for {
		select {
		case msg := <-c.MsgChan:
			if _, err := c.Conn.Write(msg); err != nil {
				fmt.Println("Failed to write data.")
				c.ExitBuffChan <- true
				return
			}

		case <-c.ExitBuffChan:
			return
		}
	}
}

func (c *Connection) Start() {
	go c.startReader()
	go c.startWriter()
}

func (c *Connection) Stop() {
	if c.IsClosed {
		return
	}

	c.IsClosed = true

	c.Conn.Close()

	c.ExitBuffChan <- true

	close(c.ExitBuffChan)
	close(c.MsgChan)
}

func (c *Connection) GetTcpConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
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

	c.MsgChan <- packedData

	return nil
}

func NewConnection(conn *net.TCPConn, connID uint32, msgHandler ziface.IMsgHandler) *Connection {
	return &Connection{
		Conn:         conn,
		ExitBuffChan: make(chan bool, 1),
		IsClosed:     false,
		ConnID:       connID,
		MsgHandler:   msgHandler,
		MsgChan:      make(chan []byte),
	}
}
