package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/types/cart"
	"rpc/db"

	"github.com/zeromicro/go-zero/core/logx"
)

type AlterItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAlterItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlterItemLogic {
	return &AlterItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AlterItemLogic) AlterItem(in *cart.CartItem) (*cart.Result, error) {
	// todo: add your logic here and delete this line
	res := db.DB.Where("cart_id = ? and book_id = ?", in.CartID, in.BookID).Save(in)
	if res.Error != nil {
		return nil, res.Error
	}
	return &cart.Result{
		Res : "ok",
	}, nil
}
