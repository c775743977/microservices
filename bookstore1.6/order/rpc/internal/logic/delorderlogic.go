package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/types/order"
	"rpc/db"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelOrderLogic {
	return &DelOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelOrderLogic) DelOrder(in *order.OrderReqByID) (*order.OrderReqByID, error) {
	// todo: add your logic here and delete this line
	res := db.DB.Where("order_id = ?", in.ID).Delete(&order.OrderItem{})
	if res.Error != nil {
		return nil, res.Error
	}
	res = db.DB.Where("id = ?", in.ID).Delete(&order.OrderRes{})
	if res.Error != nil {
		return nil, res.Error
	}
	return &order.OrderReqByID{
		ID : in.ID,
	}, nil
}
