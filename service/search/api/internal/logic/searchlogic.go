package logic

import (
	"book/service/book/rpc/bookclient"
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"book/service/search/api/internal/svc"
	"book/service/search/api/internal/types"
	"book/service/user/rpc/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req *types.SearchReq) (resp *types.SearchResp, err error) {
	userIdStr := fmt.Sprintf("%v", l.ctx.Value("userId"))
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		return nil, err
	}
	logx.Infof("userId: %d", userId)

	// 如果user不存在，则无法使用search方法
	_, err = l.svcCtx.UserRpc.GetUserById(l.ctx, &userclient.GetUserByIdReq{Id: userId})
	if err != nil {
		return nil, err
	}

	if len(strings.TrimSpace(req.Name)) == 0 {
		return nil, errors.New("参数错误")
	}
	book, err := l.svcCtx.BookRpc.FindByName(l.ctx, &bookclient.FindByNameReq{Name: req.Name})
	if err != nil {
		return nil, err
	}

	return &types.SearchResp{
		Name:  book.Name,
		Type:  book.Type,
		Count: int(book.Count),
	}, nil
}
