package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"websocket_demo/internal/logic"
	"websocket_demo/internal/svc"
)

func listHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewListLogic(r.Context(), svcCtx)
		resp, err := l.List()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
