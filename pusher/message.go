// Copyright (C) goomp. 2025-present.
//
// Created at 2025-02-20, by liasica

package pusher

import (
	"encoding/json"
	"fmt"
	"io"

	"resty.dev/v3"
)

type PostMessageDataList []PostMessageData

type PostMessageData struct {
	Uid              string      `json:"uid"`
	TopicId          interface{} `json:"topicId"`
	MessageId        int         `json:"messageId"`
	MessageContentId int         `json:"messageContentId"`
	SendRecordId     int         `json:"sendRecordId"`
	Code             int         `json:"code"`
	Status           string      `json:"status"`
}

type PostMessageRequest struct {
	AppToken      string   `json:"appToken"`
	Content       string   `json:"content"`
	Summary       string   `json:"summary"`
	ContentType   int      `json:"contentType"`
	TopicIds      []int    `json:"topicIds"`
	Uids          []string `json:"uids"`
	Url           string   `json:"url"`
	VerifyPay     bool     `json:"verifyPay"`
	VerifyPayType int      `json:"verifyPayType"`
}

func (p *Pusher) PostMessage(req PostMessageRequest) {
	req.AppToken = p.token

	var res Response[PostMessageDataList]
	r, err := resty.New().R().
		SetBody(req).
		Post(`https://wxpusher.zjiecode.com/api/send/message`)

	if err != nil {
		fmt.Printf("post message error: %v\n", err)
		return
	}

	b, _ := io.ReadAll(r.RawResponse.Body)
	fmt.Printf("post message response: %s\n", b)

	err = json.Unmarshal(b, &res)
	if err != nil {
		fmt.Printf("post message unmarshal error: %v\n", err)
		return
	}

	if res.Code != 1000 {
		fmt.Printf("post message error: %v\n", res.Msg)
		return
	}

	for _, item := range res.Data {
		fmt.Printf("post message record id: %d\n", item.SendRecordId)
	}
}
