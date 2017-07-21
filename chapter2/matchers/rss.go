package matchers

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/Lovelcp/go-in-action-learning/chapter2/search"
)

type (
	// item根据item字段的标签，将定义的字段
	// 与rss文档的字段关联起来
	item struct {
		XMLName xml.Name `xml:"item"`
		PubDate string `xml:"pubDate"`
		Title string `xml:"title"`
		Description string `xml:"description"`
		Link string `xml:"link"`
		GUID string `xml:"guid"`
		GeoRssPoint string `xml:"georss:point"`
	}

	// image根据image字段的标签，将定义的字段与rss文档的字段关联起来
	image struct {
		XMLName xml.Name `xml:"image"`
		URL string `xml:"url"`
		Title string `xml:"title"`
		Link string `xml:"link"`
	}

	// channel根据channel字段的标签，将定义的字段与rss文档的字段关联起来
	channel struct {
		XMLName xml.Name `xml:"channel"`
		Title string `xml:"title"`
		Description string `xml:"description"`
		Link string `xml:"link"`
		PubDate string `mxl:"pubDate"`
		LastBuildDate string `xml:"lastBuildDate"`
		TTL string `xml:"ttl"`
		Language string `xml:"language"`
		ManagingEditor string `xml:"managingEditor"`
		WebMaster string `xml:"webMaster"`
		Image image `xml:"image"`
		Item []item `xml:"item"`
	}

	// rssDocument定义了与rss文档关联的字段
	rssDocument struct {
		XMLName xml.Name `xml:"rss"`
		Channel channel `xml:"channel"`
	}
)

// rssMatcher实现了Matcher接口
type rssMatcher struct {}

// init将匹配器注册到程序里
func init() {
	var matcher rssMatcher
	search.Register("rss", matcher)
}

// retrieve发送HTTP Get请求获取rss数据源并解码
func (m rssMatcher) retrieve(feed *search.Feed) (*rssDocument, error) {
	if feed.URI == "" {
		return nil, errors.New("No rss feed URI provided")
	}

	// 从网络获得rss数据源文档
	resp, err := http.Get(feed.URI)
	if err != nil {
		return nil, err
	}

	// 一旦从函数返回，关闭返回的响应链接
	defer resp.Body.Close()

	// 检查状态码是不是200，这样就能知道是不是收到了正确的响应
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Response Error %d\n", resp.StatusCode)
	}

	// 将rss数据源文档解码到我们定义的结构类型里
	// 不需要检查错误，调用这会做这件事
	var document rssDocument
	err = xml.NewDecoder(resp.Body).Decode(&document)
	return &document, err
}