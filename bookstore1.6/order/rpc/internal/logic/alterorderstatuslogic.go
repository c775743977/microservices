package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/types/order"
	"rpc/db"

	"github.com/zeromicro/go-zero/core/logx"
)

type AlterOrderStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAlterOrderStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlterOrderStatusLogic {
	return &AlterOrderStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AlterOrderStatusLogic) AlterOrderStatus(in *order.OrderRes) (*order.OrderRes, error) {
	// todo: add your logic here and delete this line
	res := db.DB.Model(&order.OrderRes{}).Where("id = ?", in.ID).Update("status", in.Status)
	if res.Error != nil {
		return nil, res.Error
	}
	return &order.OrderRes{
		ID : in.ID,
	}, nil
}
