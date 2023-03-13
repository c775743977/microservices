package service

import (
	"fmt"
	"time"
)

var HS = &HelloService{}

type HelloService struct {

}

func(this *HelloService) SayHello(req *HelloRequest, stream HelloService_SayHelloServer) error {
	count := 0
	fmt.Println("服务器收到的数据:", req.Req)
	for {
		res := &HelloResponse{
			Res : "Helloooo " + req.Req + fmt.Sprint(count),
		}
		err := stream.Send(res)
		time.Sleep(time.Second)
		if err != nil {
			fmt.Println("stream.Send error:", err)
			return err
		}
		count++
		if count >= 10 {
			break
		}
	}
	return nil
}