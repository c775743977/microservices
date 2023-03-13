package main

import (
	"grpc/server-stream/client/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"fmt"
	"context"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("connect to 8080 error:", err)
		return
	}
	client := service.NewHelloServiceClient(conn)
	req := &service.HelloRequest{
		Req : "cdl",
	}
	stream, err := client.SayHello(context.Background(), req)
	if err != nil {
		fmt.Println("get stream error:", err)
	}
	wg.Add(1)
	go RecvRes(stream, &wg)
	wg.Wait()
}

func RecvRes(stream service.HelloService_SayHelloClient, wg *sync.WaitGroup) {
	defer wg.Done()
	count := 0
	for {
		res, err := stream.Recv()
		if err != nil {
			fmt.Println("stream.Recv error:", err)
			return
		}
		fmt.Println("收到的流数据:", res.Res)
		count++
		if count >= 10 {
			break
		}
	}
}