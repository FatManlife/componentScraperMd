package xstore

import (
	"github.com/FatManlife/component-finder/back-end/internal/collector"
	"github.com/gocolly/colly"
)

func newProductCollector() *colly.Collector{
	c := collector.New("xstore.md", true)
	collector.ApplyDefaultLimits(c, 1, 1000, 500)
	collector.AttachRetryErr429(c)
	
	return c
}

func newPageCollector() *colly.Collector{	
	c := collector.New("xstore.md", true)
	collector.ApplyDefaultLimits(c, 1, 1000, 500)
	collector.AttachRetryErr429(c)
	
	return c
}

func newCategoryCollector() *colly.Collector{	
	c := collector.New("xstore.md", true)
	collector.ApplyDefaultLimits(c, 1, 1000, 500)
	collector.AttachRetryErr429(c)
	
	return c
}
