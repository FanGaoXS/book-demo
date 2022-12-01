package logic

import (
	"book/service/book/api/internal/svc"
	"book/service/book/api/internal/types"
	"book/service/book/model"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req *types.AddReq) (resp *types.AddResp, err error) {
	b := &model.Book{
		Name:  req.Name,
		Type:  req.Type,
		Count: int64(req.Count),
	}
	result, err := l.svcCtx.BookModel.Insert(l.ctx, b)
	if err != nil {
		return nil, err
	}

	id, _ := result.LastInsertId()
	return &types.AddResp{
		Id: id,
	}, nil
}
