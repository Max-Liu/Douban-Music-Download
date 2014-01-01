package main

import (
	// spew "github.com/davecgh/go-spew/spew"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

func main() {
	data, _ := ioutil.ReadFile("data.html")
	patter := "<a title=\".{0,} href=\"http://site.douban.com/kindergartenkiller/widget/playlist/.{0,}/download\\?song_id=.{0,}"
	pattern, err := regexp.Compile(patter)
	if err != nil {
		log.Println(err)
	}
	find := pattern.FindAll(data, -1)

	for k, v := range find {
		if k < 17 {
			continue
		}

		log.Print("downloading " + strconv.Itoa(k) + "of " + strconv.Itoa(len(find)))
		patternTitle := "下载.{0,} "
		patternTitlePoint, err := regexp.Compile(patternTitle)
		if err != nil {
			log.Println(err)
		}
		findTitle := patternTitlePoint.Find(v)[7 : len(patternTitlePoint.Find(v))-2]

		patternLink := "href=\".{0,}"
		patternLinkPoint, err := regexp.Compile(patternLink)
		if err != nil {
			log.Println(err)
		}
		findLink := patternLinkPoint.Find(v)[6:]
		log.Println(string(findTitle))

		client := &http.Client{
		// CheckRedirect: redirectPolicyFunc,
		}
		req, err := http.NewRequest("GET", string(findLink), nil)
		// ...
		req.Header.Add("Accept", `text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8`)
		req.Header.Add("Accept-Encoding", `gzip,deflate,sdch"`)
		req.Header.Add("Accept-Language", `en-US,en;q=0.8,ja;q=0.6,zh-CN;q=0.4,zh-TW;q=0.2`)
		req.Header.Add("Connection", `keep-alive"`)
		req.Header.Add("Cookie", ``)
		req.Header.Add("IHost", `site.douban.com`)
		req.Header.Add("Referer", `http://site.douban.com/kindergartenkiller/`)
		req.Header.Add("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.63 Safari/537.36`)
		resp, err := client.Do(req)
		// resp.ti

		file, _ := os.Create(string(findTitle) + ".mp3")
		io.Copy(file, resp.Body)
		log.Println("done")

	}
}
