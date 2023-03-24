package logic

import (
	"context"
	"fmt"

	"rpc/internal/svc"
	"rpc/types/cart"
	"rpc/db"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddItemLogic {
	return &AddItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddItemLogic) AddItem(in *cart.CartItem) (*cart.Result, error) {
	// todo: add your logic here and delete this line
	fmt.Println("data:", in)
	res := db.DB.Create(in)
	if res.Error != nil {
		return nil, res.Error
	}
	return &cart.Result{
		Res : "ok",
	}, nil
}
