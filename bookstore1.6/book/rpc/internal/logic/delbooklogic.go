package logic

import (
	"context"
	"fmt"

	"rpc/internal/svc"
	"rpc/types/book"
	"rpc/db"
	"rpc/model"

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
	var data model.Book
	res := db.DB.Where("id = ?", in.ID).Delete(&data)
	fmt.Println("delete error:", res.Error)
	if res.Error != nil {
		return nil, res.Error
	} else {
		return &book.Res{
			Result : "ok",
		}, nil
	}
}
