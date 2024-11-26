package tcpx_test

import (
	"github.com/bang-go/network/tcpx"
	"log"
	"testing"
	"time"
)

func TestClientStart(t *testing.T) {
	client := tcpx.NewClient(&tcpx.ClientConfig{Addr: "127.0.0.1:8082"})
	connect, err := client.Dail()
	if err != nil {
		log.Fatal(err)
	}
	defer connect.Close()
	for {
		err = connect.Send([]byte("test"))
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(5 * time.Second)
		buf := make([]byte, 1024)
		err = connect.Receive(buf)
		log.Println(string(buf))
		if err != nil {
			log.Fatal(err)
		}
	}

}
