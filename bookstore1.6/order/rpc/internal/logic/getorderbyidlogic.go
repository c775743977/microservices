package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/types/order"
	"rpc/db"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderByIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderByIDLogic {
	return &GetOrderByIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderByIDLogic) GetOrderByID(in *order.OrderReqByID) (*order.OrderRes, error) {
	// todo: add your logic here and delete this line
	var data order.OrderRes
	res := db.DB.Where("id = ?", in.ID).Find(&data)
	if res.Error != nil {
		return nil, res.Error
	}
	return &data, nil
}
