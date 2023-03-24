package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/types/book"
	"rpc/db"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBookByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBookByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBookByNameLogic {
	return &GetBookByNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetBookByNameLogic) GetBookByName(in *book.BookRes) (*book.Books, error) {
	// todo: add your logic here and delete this line
	var data []*book.BookRes
	res := db.DB.Where("title = ?", in.Title).Find(&data)
	if res.Error != nil {
		return &book.Books{
			Books : data,
		}, res.Error
	}
	return &book.Books{
		Books : data,
	}, nil
}
