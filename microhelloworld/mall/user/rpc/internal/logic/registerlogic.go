package logic

import (
	"context"

	"user/rpc/internal/svc"
	"user/rpc/types/user"
	"user/rpc/model"
	"user/rpc/db"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterRes, error) {
	// todo: add your logic here and delete this line
	var data = &model.User{
		Id : in.Id,
		Name : in.Name,
		Password : in.Password,
		Gender : in.Gender,
	}
	db := db.ConnectMysql(l.svcCtx.Config.Mysql.DataSource)
	err := db.Insert(context.Background(), data)
	if err != nil {
		fmt.Println("call db insert error:", err)
		return nil, err
	}
	return &user.RegisterRes{
		Res : "register succeeded",
	}, nil
}
