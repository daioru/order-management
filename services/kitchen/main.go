package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewgRPCClient(addr string) *grpc.ClientConn {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return conn
}

func main() {
	httpServer := NewHttpServer(":7000")
	err := httpServer.Run()
	if err != nil {
		log.Fatalf("http server error: %v", err)
	}
}
