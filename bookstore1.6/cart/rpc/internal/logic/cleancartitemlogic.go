package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/types/cart"
	"rpc/db"

	"github.com/zeromicro/go-zero/core/logx"
)

type CleanCartItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCleanCartItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CleanCartItemLogic {
	return &CleanCartItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CleanCartItemLogic) CleanCartItem(in *cart.CartItem) (*cart.Result, error) {
	// todo: add your logic here and delete this line
	if in.BookID == 0 {
		res := db.DB.Where("cart_id = ?", in.CartID).Delete(in)
		if res.Error != nil {
			return nil, res.Error
		}
	} else {
		res := db.DB.Where("cart_id = ? and book_id = ?", in.CartID, in.BookID).Delete(in)
		if res.Error != nil {
			return nil, res.Error
		}
	}
	return &cart.Result{
		Res : "ok",
	}, nil
}
