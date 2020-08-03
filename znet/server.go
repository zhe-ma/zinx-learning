package znet

import (
	"fmt"
	"net"
	"time"
	"zinx-learning/utils"
	"zinx-learning/ziface"
)

type Server struct {
	Name       string
	IPVersion  string
	IP         string
	Port       string
	MsgHandler ziface.IMsgHandler
}

func (s *Server) Start() {
	fmt.Printf("[START] Server %s is starting. Listening %s:%s.\n", s.Name, s.IP, s.Port)

	go func() {
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%s", s.IP, s.Port))
		if err != nil {
			fmt.Println("Failed to resolve address:", err)
			return
		}

		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("Failed to listen address:", err)
			return
		}

		s.MsgHandler.StartWorkPool()

		var connID uint32 = 0

		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Failed to accept connection:", err)
				continue
			}

			connID++

			connPtr := NewConnection(conn, connID, s.MsgHandler)
			go connPtr.Start()
		}
	}()
}

func (s *Server) Stop() {
	fmt.Printf("[STOP] Server %s stop.\n", s.Name)
}

func (s *Server) Serve() {
	fmt.Println("[SERVE]")
	s.Start()

	for {
		time.Sleep(time.Second)
	}
}

func (s *Server) AddRouter(msgID uint32, router ziface.IRouter) error {
	return s.MsgHandler.AddRouter(msgID, router)
}

func NewServer() ziface.IServer {
	fmt.Println(utils.GlobalObj.TCPPort, utils.GlobalObj.Host, utils.GlobalObj.ServerName)

	return &Server{
		Name:       utils.GlobalObj.ServerName,
		IPVersion:  "tcp4",
		IP:         utils.GlobalObj.Host,
		Port:       utils.GlobalObj.TCPPort,
		MsgHandler: NewMsgHandler(),
	}
}
