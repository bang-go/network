package grpcx_test

import (
	"github.com/bang-go/network/grpcx"
	"google.golang.org/grpc"
	"log"
	"testing"
)

func TestServerStart(t *testing.T) {
	server := grpcx.NewServer(&grpcx.ServerConfig{Addr: "127.0.0.1:8081"})
	err := server.Start(func(server *grpc.Server) {})
	if err != nil {
		log.Fatal()
	}
}
