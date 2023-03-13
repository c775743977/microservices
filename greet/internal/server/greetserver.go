// Code generated by goctl. DO NOT EDIT.
// Source: greet.proto

package server

import (
	"context"

	"go-zero/greet/greet"
	"go-zero/greet/internal/logic"
	"go-zero/greet/internal/svc"
)

type GreetServer struct {
	svcCtx *svc.ServiceContext
	greet.UnimplementedGreetServer
}

func NewGreetServer(svcCtx *svc.ServiceContext) *GreetServer {
	return &GreetServer{
		svcCtx: svcCtx,
	}
}

func (s *GreetServer) Ping(ctx context.Context, in *greet.Request) (*greet.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}
