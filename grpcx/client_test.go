package grpcx_test

import (
	"github.com/bang-go/network/grpcx"
	"log"
	"testing"
)

func TestClientConn(t *testing.T) {
	client := grpcx.NewClient(&grpcx.ClientConfig{Addr: "127.0.0.1:8081"})
	conn, err := client.Dial()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(conn)
}
