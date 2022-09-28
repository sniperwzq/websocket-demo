package logic

import (
	"context"
	"fmt"

	"websocket_demo/internal/pkg/xwebsocket"
	"websocket_demo/internal/svc"
	"websocket_demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List() (resp *types.ListRes, err error) {
	list := types.ListRes{}
	list.List = xwebsocket.GetHubNameList()
	fmt.Printf("%#v", list)
	return &list, nil
}
