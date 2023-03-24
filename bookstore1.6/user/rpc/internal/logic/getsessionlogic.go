package logic

import (
	"context"
	"fmt"

	"rpc/internal/svc"
	"rpc/types/user"
	"rpc/db"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSessionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSessionLogic {
	return &GetSessionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSessionLogic) GetSession(in *user.Cookie) (*user.Session, error) {
	// todo: add your logic here and delete this line
	data, err := db.RDB.HMGet(context.Background(), in.Cookie, "username", "userID").Result()
	if err != nil {
		fmt.Println("RDB.HMGet error:", err)
		return nil, err
	}
	return &user.Session{
		UserName : data[0].(string),
		UserID : data[1].(string),
		Cookie : in.Cookie,
	}, nil
}
