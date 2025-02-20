// Copyright (C) goomp. 2025-present.
//
// Created at 2025-02-20, by liasica

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"goomp/pusher"
	"goomp/topic"
)

const (
	articlesCacheFile = "./runtime/articles.json"
)

var cached = make(map[int]string)

func main() {
	os.MkdirAll("./runtime", os.ModePerm)

	if _, err := os.Stat(articlesCacheFile); os.IsNotExist(err) {
		_, err = os.Create(articlesCacheFile)
		if err != nil {
			fmt.Println("create articles cache file failed: ", err)
			return
		}
	}

	b, _ := os.ReadFile(articlesCacheFile)
	json.Unmarshal(b, &cached)

	ticker := time.NewTicker(5 * time.Minute)

	p := pusher.NewPusher(os.Getenv("APP_TOKEN"))

	for ; true; <-ticker.C {
		articles := topic.QueryPosts()

		for _, article := range articles {
			if _, ok := cached[article.ContentId]; !ok {
				fmt.Printf("%d: %s\n", article.ContentId, article.Title)
				cached[article.ContentId] = article.Title
				b, _ = json.MarshalIndent(cached, "", "  ")
				_ = os.WriteFile(articlesCacheFile, b, 0644)
				// send notification
				p.PostMessage(pusher.PostMessageRequest{
					Content:     article.Title,
					Summary:     article.SubTitle,
					ContentType: 1,
					TopicIds:    []int{37764},
					Url:         fmt.Sprintf(`https://omp.uopes.cn/static/webapp/share/article_details.html?contentId=%d&fid=0004&pkgName=app.huawei.motor&EC=&userName=hid55765798`, article.ContentId),
				})
			}
		}
	}
}
