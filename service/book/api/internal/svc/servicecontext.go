package svc

import (
	"book/service/book/api/internal/config"
	"book/service/book/api/internal/middleware"
	"book/service/book/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config    config.Config
	BookModel model.BookModel
	Logger    rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		BookModel: model.NewBookModel(conn, c.CacheRedis),
		Logger:    middleware.NewLoggerMiddleware().Handle,
	}
}
