package main

import (
	"google.golang.org/grpc"
	"grpc/single-auth/service"
	"google.golang.org/grpc/credentials"
	"context"
	"fmt"
)

func main() {
	//"*.cdl.com"是openssl.cnf里设置的DNS
	creds, err := credentials.NewClientTLSFromFile("E:/golang/go_code/src/grpc/credentials/server.pem", "*.cdl.com")
	if err != nil {
		fmt.Println("client get credentials error:", err)
		return
	}
	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(creds))
	if err != nil {
		fmt.Println("connect to 8080 error:", err)
		return
	}
	defer conn.Close()
	client := service.NewHelloServiceClient(conn)
	var helloreq = &service.HelloRequest{Req : "cdl",}
	res, err := client.SayHello(context.Background(), helloreq)
	if err != nil {
		fmt.Println("client.SayHello error:", err)
		return
	}
	fmt.Println(res)
}