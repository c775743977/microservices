package logic

import (
	"context"

	"user/rpc/internal/svc"
	"user/rpc/types/user"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"github.com/zeromicro/go-zero/core/logx"
	"fmt"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

type User struct {
	Name string
	Password string
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.Request) (*user.Response, error) {
	// todo: add your logic here and delete this line
	if CheckInfo(in) {
		return &user.Response{
			Res : "login succeeded",
		}, nil
	} else {
		return &user.Response{
			Res : "login failed",
		}, nil
	}
}

func CheckInfo(in *user.Request) bool {
	db, err := gorm.Open(mysql.Open("root:Chen@123@tcp(localhost)/testdb"), &gorm.Config{})
	if err != nil {
		fmt.Println("connect to mysql error", err)
		return false
	}
	var user User
	db.Where("name = ?", in.UserName).Find(&user)
	if in.Password == user.Password {
		return true
	} else {
		return false
	}
}