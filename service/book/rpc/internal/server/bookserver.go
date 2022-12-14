// Code generated by goctl. DO NOT EDIT!
// Source: book.proto

package server

import (
	"context"

	"book/service/book/rpc/book"
	"book/service/book/rpc/internal/logic"
	"book/service/book/rpc/internal/svc"
)

type BookServer struct {
	svcCtx *svc.ServiceContext
	book.UnimplementedBookServer
}

func NewBookServer(svcCtx *svc.ServiceContext) *BookServer {
	return &BookServer{
		svcCtx: svcCtx,
	}
}

func (s *BookServer) FindByName(ctx context.Context, in *book.FindByNameReq) (*book.FindByNameResp, error) {
	l := logic.NewFindByNameLogic(ctx, s.svcCtx)
	return l.FindByName(in)
}
