package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/types/user"
	"rpc/db"
	"rpc/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCartLogic {
	return &CreateCartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateCartLogic) CreateCart(in *user.UserReq) (*user.UserReq, error) {
	// todo: add your logic here and delete this line
	var cart = model.Cart{
		UserID : in.UserID,
	}
	res := db.DB.Create(&cart)
	if res.Error != nil {
		return nil, res.Error
	}
	return in, nil
}
