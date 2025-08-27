// Copyright (C) goomp. 2025-present.
//
// Created at 2025-08-27, by liasica

package pusher

import (
	"fmt"

	"resty.dev/v3"
)

type Gotify struct {
	url string
}

func NewGotify(url string) *Gotify {
	return &Gotify{url: url}
}

type GotifyContentType = string

const (
	GotifyContentTypePlain GotifyContentType = "text/plain"
	GotifyContentTypeMarkd GotifyContentType = "text/markdown"
)

type GotifyClick struct {
	Url string `json:"url,omitempty"`
}

type GotifyClientDisplay struct {
	ContentType string `json:"contentType,omitempty"`
}

type GotifyClientNotification struct {
	Click       GotifyClick `json:"click,omitempty"`
	BigImageUrl string      `json:"bigImageUrl,omitempty"`
}

type GotifyExtras struct {
	ClientClick        GotifyClick              `json:"client::click,omitempty"`
	ClientDisplay      GotifyClientDisplay      `json:"client::display,omitempty"`
	ClientNotification GotifyClientNotification `json:"client::notification,omitempty"`
}

type GotifyRequest struct {
	Title    string       `json:"title,omitempty"`
	Message  string       `json:"message,omitempty"`
	Priority int          `json:"priority,omitempty"`
	Extras   GotifyExtras `json:"extras,omitempty"`
}

func NewGotifyRequest(message *Message) *GotifyRequest {
	url := fmt.Sprintf(`https://omp.uopes.cn/static/webapp/share/article_details.html?contentId=%d&fid=0004&pkgName=app.huawei.motor&EC=&userName=hid55765798`, message.Id)
	content := message.CuntContent(100)

	if message.Image != nil {
		content = fmt.Sprintf(`![%s](%s)

%s`, message.Title, *message.Image, content)
	}
	req := &GotifyRequest{
		Title: "发现OTA文章",
		Message: fmt.Sprintf(`### %s



> %s - %s


%s



[查看详情](%s)`,
			message.Title,
			message.Author,
			message.CreatTime.Format("2006-01-02 15:04:05"),
			content,
			url,
		),
		Priority: 10,
		Extras: GotifyExtras{
			ClientDisplay: GotifyClientDisplay{
				ContentType: GotifyContentTypeMarkd,
			},
			ClientNotification: GotifyClientNotification{
				Click: GotifyClick{
					Url: url,
				},
			},
			ClientClick: GotifyClick{
				Url: url,
			},
		},
	}
	return req
}

func (p *Gotify) Push(message *Message) error {
	_, err := resty.New().R().
		SetContentType("application/json").
		SetBody(NewGotifyRequest(message)).
		SetDebug(true).
		Post(p.url)
	return err
}
