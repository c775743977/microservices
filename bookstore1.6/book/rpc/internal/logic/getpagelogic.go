package logic

import (
	"context"
	"fmt"

	"rpc/internal/svc"
	"rpc/types/book"
	_"rpc/model"
	"rpc/db"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPageLogic {
	return &GetPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPageLogic) GetPage(in *book.PageReq) (*book.PageRes, error) {
	// todo: add your logic here and delete this line
	books, count := getBooks(in.MaxPrice, in.MinPrice, in.PageNo, 4)
	var totalPages int64
	if count % 4 == 0 {
		totalPages = count / 4
	} else {
		totalPages = count / 4 + 1
	}
	return &book.PageRes{
		Books : books,
		PageNo : in.PageNo,
		PageSize : 4,
		TotalPages : totalPages,
		TotalBooks : count,
		MaxPrice : in.MaxPrice,
		MinPrice : in.MinPrice,
	}, nil
}

func getBooks(max float64, min float64, pageno int64, pagesize int64) ([]*book.BookRes, int64) {
	var count int64
	var books []*book.BookRes
	if min > max {
		return nil, 0
	}
	if min == 0 && max == 0 {
		res := db.DB.Model(&book.BookRes{}).Count(&count)
		if res.Error != nil {
			fmt.Println("count book error:", res.Error)
			return nil, 0
		}
		if pageno == 0 || pagesize == 0 {
			db.DB.Find(&books)
		} else {
			db.DB.Limit(int(pagesize)).Offset((int(pageno) - 1) * int(pagesize)).Find(&books)
		}
		return books, count
	} else {
		res := db.DB.Model(&book.BookRes{}).Where("price between ? and ?", min, max).Count(&count)
		if res.Error != nil {
			fmt.Println("count book error:", res.Error)
			return nil, 0
		}
		if pageno == 0 || pagesize == 0 {
			db.DB.Where("price between ? and ?", min, max).Find(&books)
		} else {
			db.DB.Where("price between ? and ?", min, max).Limit(int(pagesize)).Offset((int(pageno) - 1) * int(pagesize)).Find(&books)
		}
		return books, count
	}
}