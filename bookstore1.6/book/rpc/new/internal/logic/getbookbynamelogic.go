package logic

import (
	"context"

	"book/rpc/new/internal/svc"
	"book/rpc/new/types/book"

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

func (l *GetBookByNameLogic) GetBookByName(in *book.BookRes) (*book.BookRes, error) {
	// todo: add your logic here and delete this line

	return &book.BookRes{}, nil
}
