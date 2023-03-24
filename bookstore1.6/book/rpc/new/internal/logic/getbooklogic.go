package logic

import (
	"context"

	"book/rpc/new/internal/svc"
	"book/rpc/new/types/book"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBookLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBookLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBookLogic {
	return &GetBookLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetBookLogic) GetBook(in *book.BookReq) (*book.BookRes, error) {
	// todo: add your logic here and delete this line

	return &book.BookRes{}, nil
}
