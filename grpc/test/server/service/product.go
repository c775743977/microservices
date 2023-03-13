package service

import (
	"context"
)

var PS = &ProductService{}

type ProductService struct {

}

func (p *ProductService) GetStock(ctx context.Context, request *ProductRequest) (*ProductResponse, error) {
	stock := p.GetStockByID(request.ProdId)
	return &ProductResponse{ProdStock : stock,}, nil
}

func (p *ProductService) GetStockByID(id int32) int32 {
	return id
}

func (p *ProductService) CheckStock(ctx context.Context, request *ProductResponse) (*StockResponse, error) {
	res := p.CheckStockByID(request.ProdStock)
	return &StockResponse{Res : res,}, nil
}

func (p *ProductService) CheckStockByID(stock int32) bool {
	if stock > 5000 {
		return true
	} else {
		return false
	}
}