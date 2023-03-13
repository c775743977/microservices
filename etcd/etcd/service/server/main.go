package main

import (
	"fmt"
	"google.golang.org/grpc"
	"go-zero/etcd/service/pbfile"
	"go-zero/etcd/service"
	_"strconv"
	"net"
)

func main() {
	server := grpc.NewServer()
	serverRegist(server, pbfile.HS)
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("listen to 8080 error:", err)
		return
	}
	server.Serve(listener)
}

func serverRegist(s *grpc.Server, srv pbfile.HelloServiceServer) {
	pbfile.RegisterHelloServiceServer(s, srv)
	s1 := &service.Service{
		Name : "SayHello",
		Port : "8080",
		IP : "localhost",
		Protocol : "grpc",
	}
	go service.ServiceRegist(s1)
}