package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	startUrl := "http://books.toscrape.com/" //os.Args[2]
	resp, err := http.Get(startUrl)
	if err != nil {
		fmt.Println("start url failed %s\n", resp.StatusCode)
		os.Exit(1)
	}
	fmt.Printf("%s %d\n", startUrl, resp.StatusCode)

	body, _ := ioutil.ReadAll(resp.Body)
	urlReg := regexp.MustCompile(`<a href="(.+?)" title=".+?"/?>`)
	hrefArr := urlReg.FindAllStringSubmatch(string(body), -1)

	for _, href := range hrefArr {
		url := startUrl + href[1]
		go fetch(url, ch)
	}
	for range hrefArr{
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}


func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}