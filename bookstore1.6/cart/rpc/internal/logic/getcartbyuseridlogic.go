package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/types/cart"
	"rpc/db"
	_"rpc/internal/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCartByUserIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCartByUserIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCartByUserIDLogic {
	return &GetCartByUserIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCartByUserIDLogic) GetCartByUserID(in *cart.CartReqByUserID) (*cart.Cart, error) {
	// todo: add your logic here and delete this line
	var data cart.Cart
	res := db.DB.Where("user_id = ?", in.UserID).Find(&data)
	if res.Error != nil {
		return nil, res.Error
	}
	var items []*cart.CartItem
	res = db.DB.Where("cart_id = ?", data.ID).Find(&items)
	if res.Error != nil {
		return nil, res.Error
	}
	total_amount := 0.00
	var total_num int64 = 0
	for _, k := range items {
		total_amount += k.Amount
		total_num += k.Num
	}
	data.Amount = total_amount
	data.Num = total_num
	data.Items = items
	res = db.DB.Save(&data)
	return &data, nil
}
