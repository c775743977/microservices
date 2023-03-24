package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/types/cart"
	"rpc/db"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCartItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCartItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCartItemLogic {
	return &GetCartItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCartItemLogic) GetCartItem(in *cart.CartItem) (*cart.CartItem, error) {
	// todo: add your logic here and delete this line
	var data cart.CartItem
	res := db.DB.Where("cart_id = ? and book_id = ?", in.CartID, in.BookID).Find(&data)
	if res.Error != nil {
		return nil, res.Error
	}
	return &data, nil
}
