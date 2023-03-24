package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/types/cart"
	"rpc/db"

	"github.com/zeromicro/go-zero/core/logx"
)

type CleanCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCleanCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CleanCartLogic {
	return &CleanCartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CleanCartLogic) CleanCart(in *cart.CartReqByUserID) (*cart.Result, error) {
	// todo: add your logic here and delete this line
	res := db.DB.Model(&cart.Cart{}).Where("user_id = ?", in.UserID).Updates(map[string]interface{}{"num": 0, "amount": 0.00})
	if res.Error != nil {
		return nil, res.Error
	}
	return &cart.Result{
		Res : "ok",
	}, nil
}
