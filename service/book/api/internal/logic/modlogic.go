package logic

import (
	"book/service/book/api/internal/svc"
	"book/service/book/api/internal/types"
	"book/service/book/model"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewModLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModLogic {
	return &ModLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ModLogic) Mod(req *types.ModReq) error {
	b := &model.Book{
		Id:    req.Id,
		Name:  req.Name,
		Type:  req.Type,
		Count: int64(req.Count),
	}

	return l.svcCtx.BookModel.Update(l.ctx, b)
}
