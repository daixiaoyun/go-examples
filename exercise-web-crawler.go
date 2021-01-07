package main

import (
	"fmt"
	"sync"
)

type Cache struct {
	cache map[string]bool
	mutex sync.Mutex
}

var cache Cache = Cache{cache: make(map[string]bool)}

func (cache Cache) add(url string) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	cache.cache[url] = true
}

func (cache Cache) isExist(url string) bool {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	_, ok := cache.cache[url]
	//技术不存在也添加到缓存中
	if !ok {
		cache.cache[url] = true
	}
	return ok
}

type Fetcher interface {
	// Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher, end chan bool) {
	if depth <= 0 {
		end <- true
		return
	}

	if cache.isExist(url) {
		//fmt.Println("Already Exist")
		end <- true
		return
	}
	cache.add(url)

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		end <- true
		return
	}

	fmt.Printf("found: %s %q\n", url, body)
	subEnd := make(chan bool)
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher, subEnd)
	}

	for i := 0; i < len(urls); i++ {
		<-subEnd
	}

	end <- true
}

func main() {
	end := make(chan bool)
	go Crawl("https://golang.org/", 4, fetcher, end)
	for {
		if <-end {
			return
		}
	}
	return
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
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
