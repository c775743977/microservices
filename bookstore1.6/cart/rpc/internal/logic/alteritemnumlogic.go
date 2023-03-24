package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/types/cart"
	"rpc/db"

	"github.com/zeromicro/go-zero/core/logx"
)

type AlterItemNumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAlterItemNumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlterItemNumLogic {
	return &AlterItemNumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AlterItemNumLogic) AlterItemNum(in *cart.CartItem) (*cart.Result, error) {
	// todo: add your logic here and delete this line
	res := db.DB.Model(&cart.CartItem{}).Where("cart_id = ? and book_id = ?", in.CartID, in.BookID).Updates(map[string]interface{}{"num": in.Num, "amount": float64(in.Num) * in.Amount})
	if res.Error != nil {
		return nil, res.Error
	}
	return &cart.Result{
		Res : "ok",
	}, nil
}
