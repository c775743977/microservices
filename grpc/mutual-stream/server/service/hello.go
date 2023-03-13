package service

import (
	"fmt"
	"time"
)

var HS = &HelloService{}

type HelloService struct {

}

func (this *HelloService) SayHello(stream HelloService_SayHelloServer) error {
	for i := 0; i < 10; i++ {
		req, err := stream.Recv()
		if err != nil {
			fmt.Println("stream.Recv() error:", err)
			return err
		}
		fmt.Println("从客户端收到的数据:", req.Req)
		res := &HelloResponse{
			Res : "helloA " + req.Req,
		}
		err = stream.Send(res)
		if err != nil {
			fmt.Println("stream.Send(res) error:", err)
			return err
		}
		time.Sleep(time.Second*2)
	}
	return nil
}