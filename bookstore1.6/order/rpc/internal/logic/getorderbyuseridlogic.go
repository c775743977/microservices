package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/types/order"
	"rpc/db"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderByUserIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderByUserIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderByUserIDLogic {
	return &GetOrderByUserIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderByUserIDLogic) GetOrderByUserID(in *order.OrderReqByUserID) (*order.OrderRes, error) {
	// todo: add your logic here and delete this line
	var data order.OrderRes
	res := db.DB.Where("user_id = ?", in.UserID).Find(&data)
	if res.Error != nil {
		return nil, res.Error
	}
	return &data, nil
}
