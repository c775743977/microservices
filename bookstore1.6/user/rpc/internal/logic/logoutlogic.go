package logic

import (
	"context"
	"fmt"

	"rpc/internal/svc"
	"rpc/types/user"
	"rpc/db"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogOutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLogOutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogOutLogic {
	return &LogOutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LogOutLogic) LogOut(in *user.Cookie) (*user.Cookie, error) {
	// todo: add your logic here and delete this line
	err := db.RDB.HDel(context.Background(), in.Cookie, "username", "userID").Err()
	if err != nil {
		fmt.Println("DelSession error:", err)
		return nil, err
	}
	return &user.Cookie{
		Cookie : in.Cookie,
	}, nil
}
