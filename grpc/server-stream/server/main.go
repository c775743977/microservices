package main

import (
	"google.golang.org/grpc"
	"grpc/server-stream/server/service"
	"fmt"
	"net"
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