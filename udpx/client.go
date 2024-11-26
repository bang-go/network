package udpx

import (
	"net"
	"time"
)

type Client interface {
	Dail() (Connect, error)
}

type ClientConfig struct {
	Addr    string
	Timeout time.Duration
}

type clientEntity struct {
	*ClientConfig
}

func NewClient(conf *ClientConfig) Client {
	return &clientEntity{ClientConfig: conf}
}

func (s *clientEntity) Dail() (Connect, error) {
	updAddr, _ := net.ResolveUDPAddr("udp", s.Addr)
	conn, err := net.DialUDP("udp", nil, updAddr)
	return NewConnect(conn, WithConnectTimeout(s.Timeout)), err
}
