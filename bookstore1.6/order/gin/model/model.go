package model

import "gin/rpc/order"

type OrderData struct {
	UserName string
	ID string
}

type Orders struct {
	UserName string
	Order []*order.OrderRes
}