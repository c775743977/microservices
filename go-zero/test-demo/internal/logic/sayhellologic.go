package logic

import (
	"context"

	"go-zero/test-demo/internal/svc"
	"go-zero/test-demo/server/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SayHelloLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSayHelloLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SayHelloLogic {
	return &SayHelloLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SayHelloLogic) SayHello(in *rpc.HelloRequest) (*rpc.HelloResponse, error) {
	// todo: add your logic here and delete this line

	return &rpc.HelloResponse{Res : "Hello Hello " + in.Req}, nil
}
