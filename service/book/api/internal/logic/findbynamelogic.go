package logic

import (
	"context"
	"errors"
	"strings"

	"book/service/book/api/internal/svc"
	"book/service/book/api/internal/types"
	"book/service/book/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindByNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindByNameLogic {
	return &FindByNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindByNameLogic) FindByName(req *types.FindByNameReq) (resp *types.FindByNameResp, err error) {
	if len(strings.TrimSpace(req.Name)) == 0 {
		return nil, errors.New("参数错误")
	}

	book, err := l.svcCtx.BookModel.FindOneByName(l.ctx, req.Name)
	switch err {
	case nil:
	case model.ErrNotFound:
		return nil, errors.New("书不存在")
	default:
		return nil, err
	}

	return &types.FindByNameResp{
		Name:  book.Name,
		Type:  book.Type,
		Count: int(book.Count),
	}, nil
}
