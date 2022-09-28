package handler

import (
	"net/http"

	"websocket_demo/internal/logic"
	"websocket_demo/internal/svc"
	"websocket_demo/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func closeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CloseReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCloseLogic(r.Context(), svcCtx)
		err := l.Close(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
