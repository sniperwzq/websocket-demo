package logic

import (
	"context"
	"net/http"

	"websocket_demo/internal/pkg/xwebsocket"
	"websocket_demo/internal/svc"
	"websocket_demo/internal/types"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type WsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WsLogic {
	return &WsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WsLogic) Ws(req *types.WsReq, w http.ResponseWriter, r *http.Request) error {
	hub := xwebsocket.GetOrOpenHub(req.Hub)

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		l.Logger.Error(err)
		return err
	}
	client := &xwebsocket.Client{Hub: hub, Conn: conn, Send: make(chan []byte, 256)}
	l.Logger.Infof("register:%s", req.Hub)
	client.Hub.Register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.WritePump()
	go client.ReadPump()

	return nil
}
