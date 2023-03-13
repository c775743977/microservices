package logic

import (
	"context"
	"order/internal/types"
	"order/internal/svc"
	"order/internal/types/user"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderLogic {
	return &OrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	// todo: add your logic here and delete this line
	var user = &user.RegisterReq{
		Id : req.Id,
		Name : req.Name,
		Password : req.Password,
		Gender : req.Gender,
	}
	result, err := l.svcCtx.UserRpc.Register(context.Background(), user)
	if err != nil {
		fmt.Println("call register in logic error:", err)
		return nil, err
	}
	resp = &types.RegisterResponse{
		Reply : result.Res,
	}
	return resp, nil
}
