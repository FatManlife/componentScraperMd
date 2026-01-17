package pcprime

import (
	"github.com/FatManlife/component-finder/back-end/internal/collector"
	"github.com/gocolly/colly"
)

func categoryColly() *colly.Collector{
	c := collector.New("prime-pc.md", false)
	collector.AttachRetryErr429(c)

	return c
}

func productColly() *colly.Collector{
	c := collector.New("prime-pc.md",true)
	collector.ApplyDefaultLimits(c,1,500,500)
	collector.AttachRetryErr429(c)

	return c
}

func pageColly() *colly.Collector{
	c := collector.New("prime-pc.md",true)
	collector.ApplyDefaultLimits(c,1,500,500)
	collector.AttachRetryErr429(c)

	return c
}