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
		Content:     "测试推送消息",
		Summary:     "测试摘要",
		ContentType: 1,
		TopicIds:    []int{37764},
		Url:         `https://omp.uopes.cn/static/webapp/share/article_details.html?contentId=807503&fid=0004&pkgName=app.huawei.motor&EC=&userName=hid55765798`,
	})
}
