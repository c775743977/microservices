package logic

import (
	"context"

	"user/rpc/internal/svc"
	"user/rpc/types/user"
	"user/rpc/model"
	"user/rpc/db"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginRes, error) {
	// todo: add your logic here and delete this line
	db := db.ConnectMysql(l.svcCtx.Config.Mysql.DataSource)
	var data = &model.User{
		Name : in.UserName,
	}
	_ = db.FindByName(context.Background(), data)
	if data.Password == "" {
		return &user.LoginRes{
			Res : "账号不存在",
		}, nil
	}
	if data.Password != in.Password {
		return &user.LoginRes{
			Res : "login failed",
		}, nil
	} else {
		return &user.LoginRes{
			Res : "login succeeded",
		}, nil
	}
}
