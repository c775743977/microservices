package logic

import (
	"context"
	"fmt"

	"rpc/internal/svc"
	"rpc/types/user"
	"rpc/db"
	"rpc/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByNameLogic {
	return &GetUserByNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByNameLogic) GetUserByName(in *user.UserName) (*user.UserRes, error) {
	// todo: add your logic here and delete this line
	var data model.User
	res := db.DB.Where("name = ?", in.Name).Find(&data)
	if res.Error != nil {
		fmt.Println("GetUser find user error:", res.Error)
		return nil, res.Error
	}
	return &user.UserRes{
		UserID : data.ID,
		UserName : data.Name,
		UserPassword : data.Password,
		UserEmail : data.Email,
		UserPrivilege : data.Privilege,
	}, nil
}
