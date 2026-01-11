package xstore

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/FatManlife/component-finder/back-end/internal/collector"
	"github.com/FatManlife/component-finder/back-end/internal/utils"
	"github.com/gocolly/colly"
)

func requestBodyProduct(categoryColly *colly.Collector, pageColly *colly.Collector, productColly *colly.Collector, productLinks *chan utils.Link, semaphor *chan struct{}){
	//Seting up filter for urls
	var sentUrls sync.Map

	//Scrape the link for categories
	categoryColly.OnHTML("div.xabs_header div.layer1.side-menu ul[data-xcat] > li ", func(e *colly.HTMLElement) {	
			//Getting categories names
			category := e.DOM.Find("a").First().Text()
			//Filterin categories
			if (strings.Contains(category, "Laptopuri") || strings.Contains(category, "MacBook") || strings.Contains(category, "Calculatoare") || strings.Contains(category, "PC")){
				e.ForEach("ul li", func(_ int, el *colly.HTMLElement){
					if strings.Contains(el.ChildText("a"), "software"){
						return 
					}

					//Generalize the categories to send 
					components := map[string]string{
						"Procesoare": "cpu",
						"Plăci de bază": "motherboard",
						"Plăci video": "gpu",
						"Memorie operativă (RAM)": "ram",
						"Unități de stocare date": "storage",
						"Carcase": "case",
						"Surse de alimentare (PSU)": "psu",
						"Sisteme de racire": "cooler",
					}

					link := el.ChildAttr("a", "href")

					var subCategory string

					if (category != "Componente PC"){
						if (strings.Contains(category, "Laptopuri") || strings.Contains(category, "MacBook")){
							subCategory = "laptop"
						} else if ( category != "All-in-One PC" && (strings.Contains(category, "PC") || strings.Contains(category, "Calculatoare"))){ 
							subCategory = "pc"
						} else if (category == "All-in-One PC") {
							subCategory = "aio"
						}
					} else {	
						if (el.ChildText("a") == "Cabluri și controlere") {
							return 
						}
						subCategory = components[el.ChildText("a")]
					}

					//filter existing links
					if _, exists := sentUrls.LoadOrStore(link, true); !exists {
						//visiting link and seting up a semaphor
						utils.SafeVisit(pageColly, link, collector.NewContext("category",subCategory))
						//Setting up a semaphor to block the for loop
						<- *semaphor 
					}	
				})	
			} 	
	})

	//Setting up a channel to throttler the requests
	for i := 0; i < 1; i++ {
		go func (){	
			for link := range *productLinks {
				utils.SafeVisit(productColly, link.Url, collector.NewContext("category",link.Category))
				time.Sleep(800 * time.Millisecond)
			}
		}()
	}
	
	//Category page
	pageColly.OnHTML("div.category-prods.xlists figure.card-product", func(e *colly.HTMLElement){
		*productLinks <- utils.Link{Category: e.Request.Ctx.Get("category"), Url: e.ChildAttr("a.img-wrap", "href")}
	})

	//Unblocking the sempahor 
	pageColly.OnScraped(func(r *colly.Response) {
		if r.Ctx.Get("category") != "" {
			*semaphor <- struct{}{}
		}
	})

	//going to the next page
	pageColly.OnHTML("a[aria-label=\"Următor\"]", func(e *colly.HTMLElement){
		categroy := e.Request.Ctx.Get("category")	
		link := e.Attr("href")
		if link != "#" {
			utils.SafeVisit(pageColly, link, collector.NewContext("category", categroy))
		}
	})

	//Scraping the details of the product
	productColly.OnHTML("div.container.page_product",func(e *colly.HTMLElement){
		category := e.Request.Ctx.Get("category")

		switch category {
			case "cpu":
				CpuHandler(e)
			case "motherboard":
				MotherboardHandler(e)
			case "gpu":
				GpuHandler(e)
			case "ram":
				RamHandler(e)
			case "storage":
				StorageHandler(e)
			case "case":
				CaseHandler(e)
			case "psu":
				PsuHandler(e)
			case "cooler":
				CoolerHandler(e)
			case "laptop":
				LaptopHandler(e)
			case "pc":
				PcHandler(e)
			case "aio":
				AioHandler(e)
		}

		fmt.Println()
	})	
}



