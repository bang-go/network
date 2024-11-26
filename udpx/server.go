package udpx

import (
	"net"
	"time"
)

type Server interface {
	Start(ServerRegisterFunc) error
}
type ServerConfig struct {
	Addr    string
	Timeout time.Duration
}
type ServerRegisterFunc func(Connect)

type serverEntity struct {
	*ServerConfig
	listen net.Listener
}

func NewServer(conf *ServerConfig) Server {
	return &serverEntity{ServerConfig: conf}
}

func (s *serverEntity) Start(register ServerRegisterFunc) (err error) {
	updAddr, err := net.ResolveUDPAddr("udp", s.Addr)
	if err != nil {
		return
	}
	conn, err := net.ListenUDP("udp", updAddr)
	if err != nil {
		return
	}
	register(NewConnect(conn, WithConnectTimeout(s.Timeout)))
	return
}
