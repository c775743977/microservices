package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/db"
	"rpc/types/cart"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddToCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddToCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddToCartLogic {
	return &AddToCartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddToCartLogic) AddToCart(in *cart.Cart) (*cart.Result, error) {
	// todo: add your logic here and delete this line
	res := db.DB.Create(in)
	if res.Error != nil {
		return nil, res.Error
	}
	return &cart.Result{
		Res : "ok",
	}, nil
}
