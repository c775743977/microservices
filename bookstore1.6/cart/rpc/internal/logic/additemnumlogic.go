package logic

import (
	"context"
	"fmt"

	"rpc/internal/svc"
	"rpc/types/cart"
	"rpc/db"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddItemNumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddItemNumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddItemNumLogic {
	return &AddItemNumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddItemNumLogic) AddItemNum(in *cart.CartItem) (*cart.Result, error) {
	// todo: add your logic here and delete this line
	amount := in.Amount + (in.Amount / float64(in.Num))
	fmt.Println("cartitem:", in)
	fmt.Println("amount:", amount)
	res := db.DB.Model(&cart.CartItem{}).Where("cart_id = ? and book_id = ?", in.CartID, in.BookID).Updates(map[string]interface{}{"num": in.Num + 1, "amount": in.Amount + (in.Amount / float64(in.Num))})
	if res.Error != nil {
		return nil, res.Error
	}
	return &cart.Result{
		Res : "ok",
	}, nil
}
