package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/types/book"
	"rpc/db"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBookByAuthorLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBookByAuthorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBookByAuthorLogic {
	return &GetBookByAuthorLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetBookByAuthorLogic) GetBookByAuthor(in *book.BookRes) (*book.Books, error) {
	// todo: add your logic here and delete this line
	var data []*book.BookRes
	res := db.DB.Where("author = ?", in.Author).Find(&data)
	if res.Error != nil {
		return &book.Books{
			Books : data,
		}, res.Error
	}
	return &book.Books{
		Books : data,
	}, nil
}
