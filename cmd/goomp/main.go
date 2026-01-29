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
	cached            = make(map[int]string)
	directory         = "./runtime"
	articlesCacheFile = filepath.Join(directory, "articles.json")
)

func currentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

var (
	lark   *pusher.Lark
	gotify *pusher.Gotify
)

func main() {
	flag.StringVar(&directory, "dir", "./runtime", "runtime directory")

	flag.Parse()

	fmt.Printf("runtime directory: %s\n", directory)

	err := os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		fmt.Println("create runtime directory failed: ", err)
		return
	}

	if _, err = os.Stat(articlesCacheFile); os.IsNotExist(err) {
		_, err = os.Create(articlesCacheFile)
		if err != nil {
			fmt.Println("create articles cache file failed: ", err)
			return
		}
	}

	b, _ := os.ReadFile(articlesCacheFile)
	err = json.Unmarshal(b, &cached)
	if err != nil {
		fmt.Println("load articles cache file failed: ", err)
		return
	}

	ticker := time.NewTicker(5 * time.Minute)

	// build gotify pusher
	gotify = pusher.NewGotify(os.Getenv("GOTIFY_URL"))

	// build lark pusher
	larkAppId := os.Getenv("LARK_APP_ID")
	larkAppSecret := os.Getenv("LARK_APP_SECRET")
	larkUserId := os.Getenv("LARK_USER_ID")
	if larkAppId != "" && larkAppSecret != "" && larkUserId != "" {
		lark = pusher.NewLark(larkAppId, larkAppSecret, larkUserId)
	} else {
		fmt.Println("LARK_APP_ID, LARK_APP_SECRET or LARK_USER_ID is not set, skip lark pusher")
	}

	fmt.Println("environment variables:")
	for _, s := range os.Environ() {
		fmt.Println("  - ", s)
	}

	fmt.Println("starting goomp article watcher...")

	for ; true; <-ticker.C {
		articles := topic.QueryPosts()

		fmt.Printf("%s: got %d articles\n", currentTime(), len(articles))

		for _, article := range articles {
			if _, ok := cached[article.ContentId]; !ok {
				pushMessage(article)
			}
		}
	}
}

func pushMessage(article *topic.Article) {
	defer func() {
		// save to cache
		fmt.Printf("%d: %s\n", article.ContentId, article.Title)
		cached[article.ContentId] = article.Title
		b, _ := json.MarshalIndent(cached, "", "  ")
		_ = os.WriteFile(articlesCacheFile, b, 0644)
	}()

	// send gotify notification
	var image *string
	if len(article.ImageContent) > 0 {
		image = &article.ImageContent[0]
	}
	msg := &pusher.Message{
		Id:        article.ContentId,
		Title:     article.Title,
		Body:      article.TextContent,
		Image:     image,
		Author:    article.CreatorName,
		CreatTime: article.CreateTime,
	}

	// send lark message
	if lark != nil {
		err := lark.Push(msg)
		if err != nil {
			fmt.Println("push lark message failed: ", err)
		}
	}

	if gotify != nil {
		err := gotify.Push(msg)
		if err != nil {
			fmt.Println("push gotify message failed: ", err)
		}
	}
}
