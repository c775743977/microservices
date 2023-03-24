package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/types/cart"
	"rpc/db"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelItemLogic {
	return &DelItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelItemLogic) DelItem(in *cart.CartItem) (*cart.Result, error) {
	// todo: add your logic here and delete this line
	res := db.DB.Where("cart_id = ? and book_id = ?", in.CartID, in.BookID).Delete(&cart.CartItem{})
	if res.Error != nil {
		return nil, res.Error
	}
	return &cart.Result{
		Res : "ok",
	}, nil
}
