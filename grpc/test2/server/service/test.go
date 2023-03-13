package service

import (
	"context"
)

var LS = &LoginService{}
type LoginService struct {

}

func(this *LoginService) CheckLogin(ctx context.Context, user *User) (*Login, error) {
	if user.Name == "cdl" && user.Password == "chilang16" {
		return &Login{Res : "登录成功",}, nil
	} else {
		return &Login{Res : "登录失败",}, nil
	}
}