package logic

import (
	"context"
	"errors"
	"strings"

	"book/service/book/rpc/book"
	"book/service/book/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindByNameLogic {
	return &FindByNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindByNameLogic) FindByName(in *book.FindByNameReq) (*book.FindByNameResp, error) {
	if len(strings.TrimSpace(in.Name)) == 0 {
		return nil, errors.New("参数错误")
	}
	res, err := l.svcCtx.BookModel.FindOneByName(l.ctx, in.Name)
	if err != nil {
		return nil, err
	}

	return &book.FindByNameResp{
		Name:  res.Name,
		Type:  res.Type,
		Count: res.Count,
	}, nil
}
