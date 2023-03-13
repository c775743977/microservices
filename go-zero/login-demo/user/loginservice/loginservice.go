// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package loginservice

import (
	"context"

	"user/types/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Request  = user.Request
	Response = user.Response

	LoginService interface {
		Login(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	}

	defaultLoginService struct {
		cli zrpc.Client
	}
)

func NewLoginService(cli zrpc.Client) LoginService {
	return &defaultLoginService{
		cli: cli,
	}
}

func (m *defaultLoginService) Login(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := user.NewLoginServiceClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}
