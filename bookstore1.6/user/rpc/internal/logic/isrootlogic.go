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

type IsRootLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsRootLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsRootLogic {
	return &IsRootLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsRootLogic) IsRoot(in *user.Cookie) (*user.RootRes, error) {
	// todo: add your logic here and delete this line
	data, err := db.RDB.HMGet(context.Background(), in.Cookie, "username", "userID").Result()
	if err != nil {
		fmt.Println("RDB.HMGet error:", err)
		return nil, err
	}
	var privilege string
	res := db.DB.Model(&model.User{}).Select("privilege").Where("id = ?", data[1].(string)).Find(&privilege)
	if res.Error != nil {
		fmt.Println("IsRoot find user error:", res.Error)
		return nil, res.Error
	}
	if privilege == "Y" {
		return &user.RootRes{
			Result : true,
		}, nil
	} else {
		return &user.RootRes{
			Result : false,
		}, nil
	}
}
