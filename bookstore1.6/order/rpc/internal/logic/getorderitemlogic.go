package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/types/order"
	"rpc/db"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderItemLogic {
	return &GetOrderItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderItemLogic) GetOrderItem(in *order.OrderItemReq) (*order.OrderItemRes, error) {
	// todo: add your logic here and delete this line
	var data []*order.OrderItem
	res := db.DB.Where("order_id = ?", in.OrderID).Find(&data)
	if res.Error != nil {
		return nil, res.Error
	}
	return &order.OrderItemRes{
		Items : data,
	}, nil
}
