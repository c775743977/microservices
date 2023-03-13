package main

import (
	"grpc/test2/server/service"
	"google.golang.org/grpc"
	"net"
	"fmt"
)

func main() {
	server := grpc.NewServer()
	service.RegisterLoginServiceServer(server, service.LS)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("listen to 8080 error:", err)
		return
	}
	server.Serve(listener)
}