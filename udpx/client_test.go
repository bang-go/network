package udpx_test

import (
	"github.com/bang-go/network/udpx"
	"log"
	"testing"
	"time"
)

func TestClientStart(t *testing.T) {
	client := udpx.NewClient(&udpx.ClientConfig{Addr: "127.0.0.1:8083", Timeout: time.Second * 5})
	connect, err := client.Dail()
	if err != nil {
		log.Fatal(err)
	}
	defer connect.Close()
	for {
		log.Println("start-send")
		err = connect.Send([]byte("testsend"))
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(5 * time.Second)
		buf := make([]byte, 1024)
		log.Println("start-receive")
		err = connect.Receive(buf)
		log.Println(string(buf))
		if err != nil {
			log.Fatal(err)
		}
	}

}
