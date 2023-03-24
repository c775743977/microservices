package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/types/cart"
	"rpc/db"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCartByIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCartByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCartByIDLogic {
	return &GetCartByIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCartByIDLogic) GetCartByID(in *cart.CartReqByID) (*cart.Cart, error) {
	// todo: add your logic here and delete this line
	var data cart.Cart
	res := db.DB.Where("id = ?", in.CartID).Find(&data)
	if res.Error != nil {
		return nil, res.Error
	}
	var items []*cart.CartItem
	res = db.DB.Where("cart_id = ?", data.ID).Find(&items)
	if res.Error != nil {
		return nil, res.Error
	}
	data.Items = items
	return &data, nil
}
