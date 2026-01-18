package neocomputer

import (
	"github.com/FatManlife/component-finder/back-end/internal/collector"
	"github.com/gocolly/colly"
)

func newCategoryColector() *colly.Collector {
	c := collector.New("neocomputer.md",false)
	collector.AttachRetryErr429(c)

	return c
} 

func newPageColector() *colly.Collector {
	c := collector.New("neocomputer.md",true)
	collector.AttachRetryErr429(c)
	collector.ApplyDefaultLimits(c,1,500,500)

	return c
} 

func newProductCollector() *colly.Collector {
	c := collector.New("neocomputer.md",true)
	collector.AttachRetryErr429(c)
	collector.ApplyDefaultLimits(c,1,500,500)

	return c
} 