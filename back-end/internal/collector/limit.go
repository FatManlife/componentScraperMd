package collector

import (
	"time"

	"github.com/gocolly/colly"
)

func ApplyDefaultLimits(c *colly.Collector, parallelism int, delay int, randomDelay int) {
	c.Limit(&colly.LimitRule{
		Parallelism: parallelism,
		Delay:       time.Duration(delay)* time.Millisecond,
		RandomDelay: time.Duration(randomDelay)* time.Millisecond,
	})
}