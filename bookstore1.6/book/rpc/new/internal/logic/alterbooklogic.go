package logic

import (
	"context"

	"book/rpc/new/internal/svc"
	"book/rpc/new/types/book"

	"github.com/zeromicro/go-zero/core/logx"
)

type AlterBookLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAlterBookLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlterBookLogic {
	return &AlterBookLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AlterBookLogic) AlterBook(in *book.BookRes) (*book.Res, error) {
	// todo: add your logic here and delete this line

	return &book.Res{}, nil
}
