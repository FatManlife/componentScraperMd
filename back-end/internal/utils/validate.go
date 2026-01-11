package utils

import (
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