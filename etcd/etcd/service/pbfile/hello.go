package pbfile

import (
	"context"
)

var HS = &HelloService{}

type HelloService struct {

}

func(this *HelloService) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	return &HelloResponse{
		Res : "Hello " + req.Req,
	}, nil
}