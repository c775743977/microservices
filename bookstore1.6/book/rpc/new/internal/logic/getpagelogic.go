package logic

import (
	"context"

	"book/rpc/new/internal/svc"
	"book/rpc/new/types/book"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPageLogic {
	return &GetPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPageLogic) GetPage(in *book.PageReq) (*book.PageRes, error) {
	// todo: add your logic here and delete this line

	return &book.PageRes{}, nil
}
