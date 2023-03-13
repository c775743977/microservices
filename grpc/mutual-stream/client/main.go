package main

import (
	"grpc/mutual-stream/client/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"fmt"
	"context"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	conn, err := grpc.Dial(":30020", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("dial 30020 error:", err)
		return
	}
	client := service.NewHelloServiceClient(conn)
	stream, err := client.SayHello(context.Background())
	if err != nil {
		fmt.Println("client.SayHello error:", err)
		return
	}
	wg.Add(1)
	go SendAndRevc(stream, &wg)
	wg.Wait()
}

func SendAndRevc(stream service.HelloService_SayHelloClient, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		req := &service.HelloRequest{
			Req : "CDL" + fmt.Sprint(i),
		}
		err := stream.Send(req)
		if err != nil {
			fmt.Println("stream.Send(req) error:", err)
			return
		}
		time.Sleep(time.Second*2)
		res, err := stream.Recv()
		if err != nil {
			fmt.Println("stream.Recv() error:", err)
			return
		}
		fmt.Println("从服务器收到的数据:", res.Res)
		time.Sleep(time.Second*2)
	}
}