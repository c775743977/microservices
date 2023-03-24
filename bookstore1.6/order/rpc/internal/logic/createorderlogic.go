package logic

import (
	"context"
	"fmt"

	"rpc/internal/svc"
	"rpc/types/order"
	"rpc/db"
	"rpc/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderLogic) CreateOrder(in *order.OrderRes) (*order.OrderReqByID, error) {
	// todo: add your logic here and delete this line
	orderid := utils.CreateOrderID()
	res := db.DB.Create(&order.OrderRes{
		ID : orderid,
		CreateTime : utils.GetTime(),
		TotalCount : in.TotalCount,
		TotalAmount : in.TotalAmount,
		UserID : in.UserID,
		Status : 0,
	})
	if res.Error != nil {
		return nil, res.Error
	}
	fmt.Println("data:", in)
	return &order.OrderReqByID{
		ID : orderid,
	}, nil
}
