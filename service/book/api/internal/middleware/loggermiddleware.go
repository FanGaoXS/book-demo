package middleware

import (
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoggerMiddleware struct {
}

func NewLoggerMiddleware() *LoggerMiddleware {
	return &LoggerMiddleware{}
}

func (m *LoggerMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("enabled middleware")

		// Passthrough to next handler if need
		next(w, r)
	}
}
