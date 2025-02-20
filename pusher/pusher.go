// Copyright (C) goomp. 2025-present.
//
// Created at 2025-02-20, by liasica

package pusher

type Pusher struct {
	token string
}

type ResponseData interface {
	PostMessageDataList
}

type Response[T ResponseData] struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Data    T      `json:"data"`
	Success bool   `json:"success"`
}

func NewPusher(token string) *Pusher {
	return &Pusher{token: token}
}
