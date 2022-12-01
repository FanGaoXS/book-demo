package handler

import (
	"net/http"

	"book/service/book/api/internal/logic"
	"book/service/book/api/internal/svc"
	"book/service/book/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func addHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAddLogic(r.Context(), svcCtx)
		resp, err := l.Add(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
