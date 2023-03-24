package logic

import (
	"context"

	"book/rpc/new/internal/svc"
	"book/rpc/new/types/book"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelBookLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelBookLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelBookLogic {
	return &DelBookLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelBookLogic) DelBook(in *book.BookReq) (*book.Res, error) {
	// todo: add your logic here and delete this line

	return &book.Res{}, nil
}
