// Copyright (C) goomp. 2025-present.
//
// Created at 2025-02-20, by liasica

package topic

import (
	"fmt"
	"net/http"
	"time"

	"resty.dev/v3"
)

const (
	pageStart = 1
)

type Article struct {
	MetaId             string      `json:"metaId"`
	ContentId          int         `json:"contentId"`
	ParentContentId    interface{} `json:"parentContentId"`
	IsGrey             int         `json:"isGrey"`
	CreatorAvatar      string      `json:"creatorAvatar"`
	CreatorName        string      `json:"creatorName"`
	LikeCnt            int         `json:"likeCnt"`
	ReplyCnt           int         `json:"replyCnt"`
	Type               int         `json:"type"`
	ArticleContentType int         `json:"articleContentType"`
	Status             int         `json:"status"`
	OrderNum           interface{} `json:"orderNum"`
	Title              string      `json:"title"`
	SubTitle           string      `json:"subTitle"`
	TextContent        string      `json:"textContent"`
	ImageContent       []string    `json:"imageContent"`
	VideoUrl           interface{} `json:"videoUrl"`
	VideoCoverUrl      interface{} `json:"videoCoverUrl"`
	VideoSize          interface{} `json:"videoSize"`
	VideoTime          interface{} `json:"videoTime"`
	VideoResolution    interface{} `json:"videoResolution"`
	CreateTime         time.Time   `json:"createTime"`
	IsLike             interface{} `json:"isLike"`
	TopicList          []struct {
		Id             int         `json:"id"`
		Name           string      `json:"name"`
		SubName        string      `json:"subName"`
		Cover          string      `json:"cover"`
		BigTopicCover  string      `json:"bigTopicCover"`
		ParticipateCnt int         `json:"participateCnt"`
		OrderNum       interface{} `json:"orderNum"`
		IsRecommend    bool        `json:"isRecommend"`
		HotFlag        bool        `json:"hotFlag"`
		ActivityLabel  string      `json:"activityLabel"`
		ReleaseTime    time.Time   `json:"releaseTime"`
		ShareTile      string      `json:"shareTile"`
		ShareDesc      string      `json:"shareDesc"`
		ShareImageUrl  string      `json:"shareImageUrl"`
		ShareHtmlUrl   string      `json:"shareHtmlUrl"`
		IsPublic       bool        `json:"isPublic"`
		TopicType      interface{} `json:"topicType"`
	} `json:"topicList"`
	UserId             string        `json:"userId"`
	UserType           interface{}   `json:"userType"`
	LabelPath          interface{}   `json:"labelPath"`
	FileContent        string        `json:"fileContent"`
	AccessCount        int           `json:"accessCount"`
	HeatCount          string        `json:"heatCount"`
	ProvinceId         interface{}   `json:"provinceId"`
	ProvinceName       interface{}   `json:"provinceName"`
	CityId             interface{}   `json:"cityId"`
	CityName           interface{}   `json:"cityName"`
	Location           interface{}   `json:"location"`
	Longtitude         interface{}   `json:"longtitude"`
	Latitude           interface{}   `json:"latitude"`
	ReleaseTime        time.Time     `json:"releaseTime"`
	FirstReleaseTime   interface{}   `json:"firstReleaseTime"`
	UserTitle          interface{}   `json:"userTitle"`
	UserMark           []string      `json:"userMark"`
	OriginImage        interface{}   `json:"originImage"`
	Share              interface{}   `json:"share"`
	FavouriteCount     int           `json:"favouriteCount"`
	ImgContentPlus     string        `json:"imgContentPlus"`
	FileContentPlus    string        `json:"fileContentPlus"`
	PostNum            string        `json:"postNum"`
	ActivityPageType   interface{}   `json:"activityPageType"`
	NeedSignChannel    interface{}   `json:"needSignChannel"`
	NeedActivityDetail interface{}   `json:"needActivityDetail"`
	Reminds            []interface{} `json:"reminds"`
	PublisherUser      interface{}   `json:"publisherUser"`
	ForwardInfo        interface{}   `json:"forwardInfo"`
	IpSource           interface{}   `json:"ipSource"`
	StampType          interface{}   `json:"stampType"`
}

type Response struct {
	Code       string     `json:"code"`
	ResultCode string     `json:"resultCode"`
	Msg        string     `json:"msg"`
	PostList   []*Article `json:"postList"`
	Page       struct {
		PageNum    int         `json:"pageNum"`
		PageSize   int         `json:"pageSize"`
		TotalCount int         `json:"totalCount"`
		IsLive     interface{} `json:"isLive"`
	} `json:"page"`
}

// QueryPosts fetches articles from the remote server.
//
// https://omp.uopes.cn/static/webapp/share/article_details.html?contentId=807503&fid=0004&pkgName=app.huawei.motor&EC=&userName=hid55765798
//
//	curl 'https://omp.uopes.cn/xcar/omp/xbs/cc/queryPostByTopic?topicId=729&type=1&isShowVideo=true&pageSize=10&pageNum=1' \
//	  -H 'Accept: */*' \
//	  -H 'Accept-Language: zh-CN,zh;q=0.9' \
//	  -H 'Cache-Control: no-cache' \
//	  -H 'Connection: keep-alive' \
//	  -H 'Content-type: application/json;charset=UTF-8' \
//	  -b 'HW_id_xcar_omp_uopes_cn=819625124f0aa1042074ec20da88bbff; HW_idts_xcar_omp_uopes_cn=1721036593235; HW_refts_xcar_omp_uopes_cn=1724479386480; HW_idvc_xcar_omp_uopes_cn=4; HW_viewts_xcar_omp_uopes_cn=1724479400731; cookieSign=1.2.4.300; HWWAFSESID=3b02f30dc536e9253a57; HWWAFSESTIME=1739528073170; pkgName=app.huawei.motor; abTest=A; wapClientId=92754201; EC=' \
//	  -H 'Pragma: no-cache' \
//	  -H 'Referer: https://omp.uopes.cn/static/webapp/share/special_topic.html?quickAppSwitch=0&topicId=729&pkgName=app.huawei.motor&userName=hid55765798&isFold=false&fid=0004&stype=0' \
//	  -H 'Sec-Fetch-Dest: empty' \
//	  -H 'Sec-Fetch-Mode: cors' \
//	  -H 'Sec-Fetch-Site: same-origin' \
//	  -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/133.0.0.0 Safari/537.36' \
//	  -H 'sec-ch-ua: "Not(A:Brand";v="99", "Google Chrome";v="133", "Chromium";v="133"' \
//	  -H 'sec-ch-ua-mobile: ?0' \
//	  -H 'sec-ch-ua-platform: "macOS"' \
//	  -H 'traceID: 00002383EAE03453'
func QueryPosts(opts ...Option) (articles []*Article) {
	options := &Options{
		page: pageStart,
	}

	for _, o := range opts {
		o.apply(options)
	}

	client := resty.New()
	defer client.Close()

	// HWWAFSESID=218a41d65e98ca5416bb; HWWAFSESTIME=1740016446605
	var res Response
	_, err := client.R().
		EnableTrace().
		SetHeader("xid", "519b2c42e46734df5ee3e56095901a21be446e20e3a77641").
		SetHeader("User-Agent", "XCar-APP-iOS/2.0.1.320 (iPhone; iOS 18.3.1; Scale/3.0)").
		SetHeader("pkgName", "app.huawei.motor").
		SetCookies([]*http.Cookie{
			{
				Name:  "HWWAFSESID",
				Value: "218a41d65e98ca5416bb",
			},
			{
				Name:  "HWWAFSESTIME",
				Value: "1740016446605",
			},
		}).
		SetResult(&res).
		Get(fmt.Sprintf("https://omp.uopes.cn/xcar/omp/xbs/cc/queryPostByTopic?topicId=729&type=1&isShowVideo=true&pageSize=10&pageNum=%d", options.page))
	if err != nil {
		fmt.Printf("fetch error: %v\n", err)
		return
	}

	if res.Code != "0" {
		fmt.Printf("fetch error: %v\n", res.Msg)
		return
	}

	articles = res.PostList
	page := res.Page

	if page.TotalCount > page.PageSize*page.PageNum {
		articles = append(articles, QueryPosts(WithPage(options.page+1))...)
	}

	return
}
