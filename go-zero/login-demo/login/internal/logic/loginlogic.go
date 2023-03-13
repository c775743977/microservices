package logic

import (
	"context"

	"login/internal/svc"
	"login/internal/types/user"
	service "login/internal/types/loginservice"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"fmt"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *user.LRequest) (resp *user.LResponse, err error) {
	// todo: add your logic here and delete this line
	userRpc := service.NewLoginService(zrpc.MustNewClient(l.svcCtx.Config.UserRpc))
	var request = &user.Request{
		UserName : req.UserName,
		Password : req.Password,
	}
	res, err := userRpc.Login(context.Background(), request)
	resp = &user.LResponse{
		Message : res.Res,
	}
	if err != nil {
		fmt.Println("grpc calls login error:", err)
		return resp, err
	}
	return resp, nil
}

// func (l *LoginLogic) Login(req *user.LRequest) (*user.LResponse, error) {
// 	var resp = &user.LResponse{
// 		Message : "login ok",
// 	}
// 	return resp, nil
// }