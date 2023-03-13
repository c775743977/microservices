package main

import (
	"google.golang.org/grpc"
	"grpc/client-stream/client/service"
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
	stream, err := client.SayHello(context.Background())
	if err != nil {
		fmt.Println("client.SayHello error:", err)
		return
	}
	// res := make(chan struct{}, 1)
	wg.Add(1)
	go HelloReq(stream, &wg)

	// select {
	// case <- res:
	// 	recv, err := stream.CloseAndRecv()
    //     if err != nil {
    //         fmt.Println("stream.CloseAndRecv error:", err)
	// 		return
    //     }
    //     stock := recv.Res
    //     fmt.Println("客户端收到响应：", stock)
	// }
	wg.Wait()
	recv, err := stream.CloseAndRecv()
    if err != nil {
        fmt.Println("stream.CloseAndRecv error:", err)
		return
    }
    stock := recv.Res
    fmt.Println("客户端收到响应：", stock)
}

func HelloReq(stream service.HelloService_SayHelloClient, wg *sync.WaitGroup) {
	count := 0
	defer wg.Done()
	for {
		request := &service.HelloRequest{
			Req : "陈鼎立" + fmt.Sprint(count),
		}
		err := stream.Send(request)
		if err != nil {
			fmt.Println("stream.Send error:", err)
			return
		}
		count++
		if count > 10 {
            break
        }
	}
}