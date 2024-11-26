package tcpx

import (
	"github.com/bang-go/opt"
	"net"
	"time"
)

type Connect interface {
	Send([]byte) error
	Receive([]byte) error
	SetTimeout(time.Duration)
	Close()
}

type connectEntity struct {
	conn    net.Conn
	timeout time.Duration
}


func NewConnect(conn net.Conn, opts ...opt.Option[connectOptions]) Connect {
	options := &connectOptions{}
	opt.Each(options, opts...)
	return &connectEntity{conn: conn, timeout: options.timeout}
}

// Send 发送数据
func (c *connectEntity) Send(data []byte) (err error) {
	if c.timeout > 0 {
		if err = c.conn.SetWriteDeadline(c.calcDeadline()); err != nil {
			return
		}
	}
	_, err = c.conn.Write(data)
	return
}

// Receive 接受数据
func (c *connectEntity) Receive(data []byte) (err error) {
	if c.timeout > 0 {
		if err = c.conn.SetReadDeadline(c.calcDeadline()); err != nil {
			return
		}
	}
	_, err = c.conn.Read(data)
	return
}
func (c *connectEntity) SetTimeout(timeout time.Duration) {
	c.timeout = timeout
}
func (c *connectEntity) calcDeadline() time.Time {
	return time.Now().Add(c.timeout)
}
func (c *connectEntity) Close() {
	_ = c.conn.Close()
}
