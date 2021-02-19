package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"sync"
	"time"
)

var mu sync.Mutex
var pageCount, failCount int

func main() {
	fmt.Println("crawler begin")
	start := time.Now()
	ch := make(chan string)
	startUrl := "http://books.toscrape.com/"
	resp, err := http.Get(startUrl)
	if err != nil {
		log.Panicf("fetch start url error %v", err)
		os.Exit(1)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	urls := parse(string(body), `<a href="(.+?)" title=".+?"/?>`)
	for _, href := range urls {
		url := startUrl + href
		go fetch(url, ch)
	}

	for range urls {
		fmt.Println(<-ch)
	}
	fmt.Printf("crawl finished, cost %.2f seconds, fetch %d pages, %d failed", time.Since(start).Seconds(), pageCount, failCount)
}

// 根据正则参数解析网页
func parse(html string, regStr string) []string {
	reg := regexp.MustCompile(regStr)
	arr := reg.FindAllStringSubmatch(html, -1)
	var newArr []string
	for _, e := range arr {
		newArr = append(newArr, e[1])
	}
	return newArr
}

// 抓取单个网页
func fetch(url string, ch chan<- string) (content string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Printf("fetch url %s error:%v %v\n", url, err, resp.StatusCode)
		mu.Lock()
		failCount++
		mu.Unlock()
		ch <- fmt.Sprint(err)
		return
	}
	defer func() {
		secs := time.Since(start).Seconds()
		//nbytes, _ := io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
		mu.Lock()
		pageCount++
		mu.Unlock()
		ch <- fmt.Sprintf("%.2fs %7d %s ", secs, len(content), url)
	}()
	if resp.StatusCode != http.StatusOK {
		log.Printf("fetch url %s status code: %s", url, resp.StatusCode)
		return
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}
