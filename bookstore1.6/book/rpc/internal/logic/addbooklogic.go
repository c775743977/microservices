package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/types/book"
	"rpc/db"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddBookLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddBookLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddBookLogic {
	return &AddBookLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddBookLogic) AddBook(in *book.BookRes) (*book.Res, error) {
	// todo: add your logic here and delete this line
	res := db.DB.Create(in)
	if res.Error != nil {
		return nil, res.Error
	} else {
		return &book.Res{
			Result : "",
		}, nil
	}
}