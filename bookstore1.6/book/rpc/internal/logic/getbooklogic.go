package logic

import (
	"context"
	"fmt"

	"rpc/internal/svc"
	"rpc/db"
	"rpc/types/book"

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

	return getBook(in.ID)
}

func getBook(bookid int64) (*book.BookRes, error) {
	var data book.BookRes
	res := db.DB.Where("id = ?", bookid).Find(&data)
	if res.Error != nil {
		fmt.Println("find book error:", res.Error)
		return nil, res.Error
	}
	return &data, nil
}
