package main

import (
	"grpc/test/server/service"
	"google.golang.org/grpc"
	"net"
	"fmt"
)

func main() {
	rpcServer := grpc.NewServer()
	service.RegisterProductServiceServer(rpcServer, service.PS)
	listener, err := net.Listen("tcp", ":8002")
	if err != nil {
		fmt.Println("listen 8002 error:", err)
		return
	}
	rpcServer.Serve(listener)
}