package service

import (
	_"context"
	_"google.golang.org/grpc"
	"fmt"
	"time"
)

var HS = &HelloService{}

type HelloService struct {

}

func (this *HelloService) SayHello(stream HelloService_SayHelloServer) error {
	count := 0
	for {
		req, err := stream.Recv()
		if err != nil {
			fmt.Println("receive request error:", err)
			return err
		}
		fmt.Println("收到的流:", req.Req)
		time.Sleep(time.Second)
		count++
		if count >= 10 {
			res := &HelloResponse{
				Res : "Hello " + req.Req,
			}
			err = stream.SendAndClose(res)
			if err != nil {
				return err
			}
			return nil
		}
	}
}