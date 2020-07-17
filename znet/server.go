package znet

import (
	"fmt"
	"net"
	"time"
	"zinx-learning/ziface"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      string
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

		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Failed to accept connection:", err)
				continue
			}

			go func() {
				for {
					buf := make([]byte, 512)
					count, err := conn.Read(buf)
					if err != nil {
						fmt.Println("Failed to read data:", err)
						break
					}

					if _, err = conn.Write(buf[:count]); err != nil {
						fmt.Println("Failed to write data:", err)
						continue
					}
				}
			}()
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

func NewServer(name string) ziface.IServer {
	return &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      "9547",
	}
}