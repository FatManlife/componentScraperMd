package collector

import "github.com/gocolly/colly"

func New(domain string, isAsync bool) *colly.Collector{
	return colly.NewCollector(
		colly.AllowedDomains(domain),
		colly.Async(isAsync),
	)
}