// Copyright (C) goomp. 2025-present.
//
// Created at 2025-02-20, by liasica

package pusher

import (
	"os"
	"testing"
)

func TestPostMessage(t *testing.T) {
	p := NewPusher(os.Getenv("APP_TOKEN"))
	p.PostMessage(PostMessageRequest{
		Content:     "涉及车型：智界R7 纯电 Ultra | R7 纯电 Max | R7 纯电 Pro | R7 增程 Max | R7 增程 Pro | S7 Ultra | S7 Max RS | S7 Max+ | S7 Max | S7 Pro",
		Summary:     CutMessage("【常用常新】鸿蒙智行OTA 智界4月功能说明", 100),
		ContentType: 1,
		TopicIds:    []int{37764},
		Url:         `https://omp.uopes.cn/static/webapp/share/article_details.html?contentId=807503&fid=0004&pkgName=app.huawei.motor&EC=&userName=hid55765798`,
	})
}

func TestLen(t *testing.T) {
	str := "涉及车型：智界R7 纯电 Ultra | R7 纯电 Max | R7 纯电 Pro | R7 增程 Max | R7 增程 Pro | S7 Ultra | S7 Max RS | S7 Max+ | S7 Max | S7 Pro"
	t.Log(CutMessage(str, 100))
	x := "【常用常新】鸿蒙智行OTA 智界4月功能说明"
	t.Log(len(x))
}
