package logic

import (
	"context"
	"errors"

	"rpc/internal/svc"
	"rpc/types/user"
	"rpc/db"
	"rpc/model"

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
	var data = model.User{
		Name : in.UserName,
		Password : in.UserPassword,
		Email : in.UserEmail,
	}
	if RegisterCheck(in.UserName) {
		return &user.RegisterRes{
			Result : "用户名已存在",
		}, nil
	}
	res := db.DB.Select("name", "password", "email").Create(&data)
	if res.RowsAffected == 0 {
		return nil, errors.New("添加新用户到数据库失败")
	}
	return &user.RegisterRes{
		Result : "注册成功",
	}, nil
}

func RegisterCheck(username string) bool {
	var data = model.User{}
	db.DB.Where("name = ?", username).Find(&data)
	if data.Password == "" {
		return false
	} else {
		return true
	}
}
