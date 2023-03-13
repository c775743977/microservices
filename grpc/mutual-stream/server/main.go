package main

import (
	"grpc/mutual-stream/server/service"
	"google.golang.org/grpc"
	"net"
	"fmt"
)

func main() {
	server := grpc.NewServer()
	service.RegisterHelloServiceServer(server, service.HS)
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("listen to 8080 error:", err)
		return
	}
	server.Serve(listener)
}