// Copyright (C) goomp. 2026-present.
//
// Created at 2026-01-29, by liasica

package pusher

import (
	"context"
	"encoding/json"
	"fmt"

	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

type Lark struct {
	client *lark.Client
	userId string
}

func NewLark(appId, appSecret, userId string) *Lark {
	return &Lark{
		client: lark.NewClient(appId, appSecret),
		userId: userId,
	}
}

type LarkTextMessage struct {
	Text string `json:"text"`
}

func (l *Lark) CreateMessage(text string) (id string, err error) {
	b, _ := json.Marshal(LarkTextMessage{Text: text})

	req := larkim.NewCreateMessageReqBuilder().
		ReceiveIdType("user_id").
		Body(larkim.NewCreateMessageReqBodyBuilder().
			ReceiveId(l.userId).
			MsgType("text").
			Content(string(b)).
			Build()).
		Build()

	var res *larkim.CreateMessageResp
	res, err = l.client.Im.V1.Message.Create(context.Background(), req)
	if err != nil {
		return
	}

	if !res.Success() {
		err = fmt.Errorf("lark push message failed: %s", res.Msg)
		return
	}

	if res.Data == nil || res.Data.MessageId == nil {
		err = fmt.Errorf("lark push message failed: no message id returned")
		return
	}

	id = *res.Data.MessageId
	return
}

func (l *Lark) UrgentAppMessage(id string) (err error) {
	req := larkim.NewUrgentAppMessageReqBuilder().
		MessageId(id).
		UserIdType("user_id").
		UrgentReceivers(larkim.NewUrgentReceiversBuilder().
			UserIdList([]string{l.userId}).
			Build()).
		Build()

	var res *larkim.UrgentAppMessageResp
	res, err = l.client.Im.V1.Message.UrgentApp(context.Background(), req)
	if err != nil {
		return
	}

	if !res.Success() {
		err = fmt.Errorf("lark urgent app message failed: %s", res.Msg)
	}

	return
}

// Push 发送Lark消息
func (l *Lark) Push(message *Message) error {
	id, err := l.CreateMessage(fmt.Sprintf("### %s\n\n> %s - %s\n\n%s\n\n[查看详情](https://omp.uopes.cn/static/webapp/share/article_details.html?quickAppSwitch=0&contentId=%d&spid=57b08ddd503072072e6eae111be5b924&pkgName=app.huawei.motor&fid=0004&EC=SERES&isFold=true)",
		message.Title,
		message.Author,
		message.CreatTime.Format("2006-01-02 15:04:05"),
		message.CuntContent(100),
		message.Id,
	))
	if err != nil {
		return err
	}

	return l.UrgentAppMessage(id)
}
