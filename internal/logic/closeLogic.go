package logic

import (
	"context"

	"websocket_demo/internal/pkg/xwebsocket"
	"websocket_demo/internal/svc"
	"websocket_demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CloseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCloseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CloseLogic {
	return &CloseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CloseLogic) Close(req *types.CloseReq) error {
	xwebsocket.CloseHub(req.Hub)
	return nil
}
