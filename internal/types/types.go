// Code generated by goctl. DO NOT EDIT.
package types

type WsReq struct {
	Hub string `path:"hub"`
}

type ListRes struct {
	List []string `json:"list"`
}

type CloseReq struct {
	Hub string `path:"hub"`
}