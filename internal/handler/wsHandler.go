package handler

import (
	"net/http"

	"websocket_demo/internal/logic"
	"websocket_demo/internal/svc"
	"websocket_demo/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func wsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewWsLogic(r.Context(), svcCtx)
		_ = l.Ws(&req, w, r)
	}
}
