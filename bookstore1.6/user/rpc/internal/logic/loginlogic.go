package logic

import (
	"context"
	"fmt"

	"rpc/internal/svc"
	"rpc/db"
	"rpc/model"
	"rpc/utils"
	"rpc/types/user"

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
	var data model.User
	db.DB.Where("name = ?", in.UserName).Find(&data)
	if data.Name == "" {
		return &user.LoginRes{
			Result : "该用户不存在",
		}, nil
	}
	if data.Password != in.UserPassword {
		return &user.LoginRes{
			Result : "密码错误",
		}, nil
	}
	cookie := CreateSession(data.Name, data.ID)
	return &user.LoginRes{
		Result : "登录成功",
		Cookie : cookie,
	}, nil
}

func CreateSession(username string, userid int64) string {
	// todo: add your logic here and delete this line
	cookie := utils.CreateUUID()
	err := db.RDB.HMSet(context.Background() , cookie, "username", username, "userID", userid).Err()
	if err != nil {
		fmt.Println("AddSession to redis error:", err)
		return ""
	}
	return cookie
}