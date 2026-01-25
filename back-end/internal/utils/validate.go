package utils

import (
	"fmt"
	"sync"

	"github.com/gocolly/colly"
)

var visitedUrls sync.Map

func SafeVisit(c *colly.Collector, url string, ctx *colly.Context) {
	if url == "" {
		return
	}

	if _, exists := visitedUrls.LoadOrStore(url, true); exists {
		return 
	}

	if ctx != nil {
		c.Request("GET",url,nil,ctx,nil)
	} else {
		c.Visit(url)
	}
}

func PopulateVisitedUrls(urls []string) {
	for _, url := range urls {
		visitedUrls.Store(url, true)
	}
}

func PrintVisitedUrls() {
    fmt.Println("Visited URLs:")
    visitedUrls.Range(func(key, value any) bool {
        fmt.Println(key.(string))
        return true 
    })
}