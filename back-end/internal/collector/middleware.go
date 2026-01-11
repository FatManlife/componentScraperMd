package collector

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

func AttachRetryErr429(c *colly.Collector) {
	c.OnError(func(r *colly.Response, err error){
		if r.StatusCode == 429 {
			fmt.Println("429 encountered, retrying:", r.Request.URL)
			time.Sleep(5 * time.Second)
			r.Request.Retry()
		} 
	})
}