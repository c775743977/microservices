package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/types/cart"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCartByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCartByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCartByNameLogic {
	return &GetCartByNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCartByNameLogic) GetCartByName(in *cart.CartReqByName) (*cart.Cart, error) {
	// todo: add your logic here and delete this line
	return &cart.Cart{}, nil
}
