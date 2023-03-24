// Code generated by goctl. DO NOT EDIT.
// Source: order.proto

package server

import (
	"context"

	"rpc/internal/logic"
	"rpc/internal/svc"
	"rpc/types/order"
)

type OrderServiceServer struct {
	svcCtx *svc.ServiceContext
	order.UnimplementedOrderServiceServer
}

func NewOrderServiceServer(svcCtx *svc.ServiceContext) *OrderServiceServer {
	return &OrderServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *OrderServiceServer) CreateOrder(ctx context.Context, in *order.OrderRes) (*order.OrderReqByID, error) {
	l := logic.NewCreateOrderLogic(ctx, s.svcCtx)
	return l.CreateOrder(in)
}

func (s *OrderServiceServer) CreateOrderItem(ctx context.Context, in *order.OrderItem) (*order.OrderReqByID, error) {
	l := logic.NewCreateOrderItemLogic(ctx, s.svcCtx)
	return l.CreateOrderItem(in)
}

func (s *OrderServiceServer) GetOrderByID(ctx context.Context, in *order.OrderReqByID) (*order.OrderRes, error) {
	l := logic.NewGetOrderByIDLogic(ctx, s.svcCtx)
	return l.GetOrderByID(in)
}

func (s *OrderServiceServer) GetOrderByUserID(ctx context.Context, in *order.OrderReqByUserID) (*order.OrderRes, error) {
	l := logic.NewGetOrderByUserIDLogic(ctx, s.svcCtx)
	return l.GetOrderByUserID(in)
}

func (s *OrderServiceServer) GetOrderItem(ctx context.Context, in *order.OrderItemReq) (*order.OrderItemRes, error) {
	l := logic.NewGetOrderItemLogic(ctx, s.svcCtx)
	return l.GetOrderItem(in)
}

func (s *OrderServiceServer) AlterOrderStatus(ctx context.Context, in *order.OrderRes) (*order.OrderRes, error) {
	l := logic.NewAlterOrderStatusLogic(ctx, s.svcCtx)
	return l.AlterOrderStatus(in)
}

func (s *OrderServiceServer) GetAllOrder(ctx context.Context, in *order.OrderReqByUserID) (*order.AllOrder, error) {
	l := logic.NewGetAllOrderLogic(ctx, s.svcCtx)
	return l.GetAllOrder(in)
}

func (s *OrderServiceServer) DelOrder(ctx context.Context, in *order.OrderReqByID) (*order.OrderReqByID, error) {
	l := logic.NewDelOrderLogic(ctx, s.svcCtx)
	return l.DelOrder(in)
}