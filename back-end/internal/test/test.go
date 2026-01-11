package test

import (
	"github.com/FatManlife/component-finder/back-end/internal/collector"
	"github.com/gocolly/colly"
)

func TestColly(){
	c := collector.New("xstore.md",false)

	// Extracting Computer category
	c.OnHTML("div.container.page_product",func(e *colly.HTMLElement){
	})
	
	c.Visit("https://xstore.md/laptopuri/gaming/lenovo-legion-pro-7-16iax10h-16-ultra-9-275hx-32gb-ram-2tb-ssd-rtx5070-ti")
}

// price := utils.CastFloat64(e.ChildText("div.xp-price"))



