package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/types/book"
	"rpc/db"

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
	res := db.DB.Save(in)
	if res.Error != nil {
		return nil, res.Error
	} else {
		return &book.Res{
			Result : "",
		}, nil
	}
}
