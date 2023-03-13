package main

import (
	"grpc/single-auth/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
	"fmt"
)

func main() {
	creds, err := credentials.NewServerTLSFromFile("E:/golang/go_code/src/grpc/credentials/server.pem", "E:/golang/go_code/src/grpc/credentials/server.key")
	if err != nil {
		fmt.Println("credentials.NewServerTLSFromFile error:", err)
		return
	}
	server := grpc.NewServer(grpc.Creds(creds))
	service.RegisterHelloServiceServer(server, service.HS)
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("listen to 8080 error:", err)
		return
	}
	server.Serve(listener)
}