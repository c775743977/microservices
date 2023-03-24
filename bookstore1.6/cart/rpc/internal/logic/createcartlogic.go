package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/types/cart"
	"rpc/db"

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

func (l *CreateCartLogic) CreateCart(in *cart.CartReqByUserID) (*cart.Result, error) {
	// todo: add your logic here and delete this line
	res := db.DB.Create(&cart.Cart{
		UserID : in.UserID,
	})
	if res.Error != nil {
		return nil, res.Error
	}
	return &cart.Result{
		Res : "ok",
	}, nil
}
