// Copyright (C) goomp. 2026-present.
//
// Created at 2026-01-29, by liasica

package pusher

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestLarkPushMessage(t *testing.T) {
	larkAppId := os.Getenv("LARK_APP_ID")
	larkAppSecret := os.Getenv("LARK_APP_SECRET")
	larkUserId := os.Getenv("LARK_USER_ID")

	if larkAppId == "" || larkAppSecret == "" || larkUserId == "" {
		t.Skip("LARK_APP_ID, LARK_APP_SECRET or LARK_USER_ID is not set, skip Lark push message test")
	}

	l := NewLark(larkAppId, larkAppSecret, larkUserId)
	// id, err := l.CreateMessage("https://omp.uopes.cn/static/webapp/share/article_details.html?quickAppSwitch=0&contentId=1514180&spid=57b08ddd503072072e6eae111be5b924&pkgName=app.huawei.motor&fid=0004&EC=SERES&isFold=true")
	// require.NoError(t, err)
	//
	// err = l.UrgentAppMessage(id)
	// require.NoError(t, err)

	err := l.Push(&Message{
		Id:        1514180,
		Title:     "发现OTA文章",
		Body:      "为您准备 | OTA升级指南",
		Image:     nil,
		Author:    "OTA资讯特派员",
		CreatTime: time.Now(),
	})
	require.NoError(t, err)
}
