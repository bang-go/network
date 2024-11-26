package tcpx

import (
	"fmt"
	"net"
	"time"
)

type Server interface {
	Start(ServerRegisterFunc) error
	Shutdown()
}

type ServerConfig struct {
	Addr    string
	Timeout time.Duration
}

type ServerRegisterFunc func(Connect)
type serverEntity struct {
	*ServerConfig
	listen *net.TCPListener
}

func NewServer(conf *ServerConfig) Server {
	return &serverEntity{ServerConfig: conf}
}

func (s *serverEntity) Start(register ServerRegisterFunc) (err error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", s.Addr)
	if err != nil {
		return
	}
	s.listen, err = net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return
	}
	for {
		conn, err1 := s.listen.Accept()
		if err1 != nil {
			fmt.Println(err1)
			continue
		}
		go register(NewConnect(conn, WithConnectTimeout(s.Timeout)))
	}
}

func (s *serverEntity) Shutdown() {
	//todo:
}
