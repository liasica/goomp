// Copyright (C) goomp. 2025-present.
//
// Created at 2025-02-20, by liasica

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/liasica/goomp/pusher"
	"github.com/liasica/goomp/topic"
)

var (
	cached    = make(map[int]string)
	directory = "./runtime"
)

func currentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func main() {
	flag.StringVar(&directory, "dir", "./runtime", "runtime directory")

	flag.Parse()

	fmt.Printf("runtime directory: %s\n", directory)

	os.MkdirAll(directory, os.ModePerm)

	var articlesCacheFile = filepath.Join(directory, "articles.json")

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
		fmt.Printf("%s: got %d articles\n", currentTime(), len(articles))

		for _, article := range articles {
			if _, ok := cached[article.ContentId]; !ok {
				fmt.Printf("%d: %s\n", article.ContentId, article.Title)
				cached[article.ContentId] = article.Title
				b, _ = json.MarshalIndent(cached, "", "  ")
				_ = os.WriteFile(articlesCacheFile, b, 0644)

				// send notification
				p.PostMessage(pusher.PostMessageRequest{
					Content:     article.SubTitle,
					Summary:     pusher.CutMessage(article.Title, 100),
					ContentType: 1,
					TopicIds:    []int{37764},
					Url:         fmt.Sprintf(`https://omp.uopes.cn/static/webapp/share/article_details.html?contentId=%d&fid=0004&pkgName=app.huawei.motor&EC=&userName=hid55765798`, article.ContentId),
				})
			}
		}
	}
}
