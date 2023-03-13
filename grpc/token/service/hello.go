package service

import (
	"context"
)

var HS = &HelloService{}

type HelloService struct {

}

func(this *HelloService) SayHello(ctx context.Context, hello *HelloRequest) (*HelloResponse, error) {
	str := "hello " + hello.Req
	return &HelloResponse{Res : str,}, nil
}