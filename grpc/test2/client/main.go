package main

import (
	"grpc/test2/client/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"fmt"
	"context"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Dial to 8080 error:", err)
		return
	}
	defer conn.Close()
	client := service.NewLoginServiceClient(conn)
	var user = &service.User{
		Name : "cdl",
		Password : "chilang16",
	}

	res, err := client.CheckLogin(context.Background(), user)
	if err != nil {
		fmt.Println("client.CheckLogin error:", err)
		return
	}
	fmt.Println(res.Res)
}