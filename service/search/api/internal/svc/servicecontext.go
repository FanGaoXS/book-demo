package svc

import (
	"book/service/book/rpc/bookclient"
	"book/service/search/api/internal/config"
	"book/service/user/rpc/userclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userclient.User
	BookRpc bookclient.Book
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		BookRpc: bookclient.NewBook(zrpc.MustNewClient(c.BookRpc)),
	}
}
