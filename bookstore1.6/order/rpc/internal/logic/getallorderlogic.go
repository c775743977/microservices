package logic

import (
	"context"
	"strconv"

	"rpc/internal/svc"
	"rpc/types/order"
	"rpc/db"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllOrderLogic {
	return &GetAllOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllOrderLogic) GetAllOrder(in *order.OrderReqByUserID) (*order.AllOrder, error) {
	// todo: add your logic here and delete this line
	var orders []*order.OrderRes
	if in.UserID == "" {
		res := db.DB.Find(&orders)
		if res.Error != nil {
			return nil, res.Error
		}
	} else {
		userid, _ := strconv.ParseInt(in.UserID, 10, 64)
		res := db.DB.Where("user_id = ?", userid).Find(&orders)
		if res.Error != nil {
			return nil, res.Error
		}
	}
	return &order.AllOrder{
		Orders : orders,
	}, nil
}
