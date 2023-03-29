package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"helloword/internal/logic"
	"helloword/internal/svc"
	"helloword/internal/types"
)

func HellowordHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewHellowordLogic(r.Context(), svcCtx)
		resp, err := l.Helloword(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
