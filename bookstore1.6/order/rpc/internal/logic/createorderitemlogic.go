package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/types/order"
	"rpc/db"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderItemLogic {
	return &CreateOrderItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderItemLogic) CreateOrderItem(in *order.OrderItem) (*order.OrderReqByID, error) {
	// todo: add your logic here and delete this line
	res := db.DB.Create(in)
	if res.Error != nil {
		return nil, res.Error
	}
	return &order.OrderReqByID{
		ID : in.OrderID,
	}, nil
}
