// Copyright (C) goomp. 2025-present.
//
// Created at 2025-02-20, by liasica

package pusher

import "time"

type Pusher interface {
	Push(title string, contentId int) error
}

type Message struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Image     *string   `json:"image"`
	Author    string    `json:"author"`
	CreatTime time.Time `json:"creatTime"`
}

func (m *Message) CuntContent(i int) string {
	body := []rune(m.Body)
	n := len(body)
	if n <= i {
		return m.Body
	}

	return string(body[:i]) + "..."
}
