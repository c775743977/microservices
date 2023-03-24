package logic

import (
	"context"

	"book/rpc/new/internal/svc"
	"book/rpc/new/types/book"

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

func (l *GetBookByAuthorLogic) GetBookByAuthor(in *book.BookRes) (*book.BookRes, error) {
	// todo: add your logic here and delete this line

	return &book.BookRes{}, nil
}
