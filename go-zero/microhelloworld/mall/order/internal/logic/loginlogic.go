package logic

import (
	"context"
	"fmt"
	"order/internal/svc"
	"order/internal/types"
	"order/internal/types/user"
	"github.com/zeromicro/go-zero/core/logx"
)

type OrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderLogic {
	return &OrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderLogic) Login(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	data := &user.LoginReq{
		UserName : req.UserName,
		Password : req.Password,
	}
	result, err := l.svcCtx.UserRpc.Login(context.Background(), data)
	if err != nil {
		fmt.Println("call rpc Login error:", err)
		return nil, err
	}
	resp = &types.Response{
		Message : result.Res,
	}
	return resp, nil
}
