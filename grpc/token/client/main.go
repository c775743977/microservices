package main

import (
	"google.golang.org/grpc"
	"grpc/single-auth/service"
	"google.golang.org/grpc/credentials/insecure"
	"context"
	"fmt"
)

type Authentication struct {
    User     string
    Password string
}

func (a *Authentication) GetRequestMetadata(context.Context, ...string) (
    map[string]string, error,
) {
    return map[string]string{"user": a.User, "password": a.Password}, nil
}

func (a *Authentication) RequireTransportSecurity() bool {
    return false
}

func main() {
	user := &Authentication{
        User: "cdl",
        Password: "chilang16",
    }
	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithPerRPCCredentials(user))
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