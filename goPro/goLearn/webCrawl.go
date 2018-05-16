// goPro project webCrawl.go
package goLearn

import (
	"fmt"
)

// http://tour.studygolang.com/concurrency/9

type Fetcher interface {
	// Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
	Fetch(url string) (body string, urls []string, err error)
}

type MapHasCrawled map[string]string

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher, hasFind MapHasCrawled, b chan string) {
	// TODO: 并行的抓取 URL。
	// TODO: 不重复抓取页面。
	// 下面并没有实现上面两种情况：
	if depth <= 0 {
		b <- ""
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		b <- ""
		return
	}
	hasFind[url] = body

	var bb chan string
	bb = make(chan string)
	for _, u := range urls {
		_, ok := hasFind[u]
		if !ok {
			continue
		}
		go Crawl(u, depth-1, fetcher, hasFind, bb)
	}

	for n := 0; n < cap(urls); n++ {
		select {
		case <-bb:
			continue
		}
	}
	b <- body
	return
}

func main() {
	m := make(MapHasCrawled)
	var bb chan string
	bb = make(chan string)
	go Crawl("http://golang.org/", 4, fetcher, m, bb)
	for {
		select {
		case <-bb:
			fmt.Printf("found: http://golang.org/ %q\n", bb)
		}
	}
}

// fakeFetcher 是返回若干结果的 Fetcher。
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher 是填充后的 fakeFetcher。
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
