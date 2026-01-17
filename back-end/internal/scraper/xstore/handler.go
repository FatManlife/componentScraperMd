package xstore

import (
	"strings"
	"sync"
	"time"

	"github.com/FatManlife/component-finder/back-end/internal/collector"
	"github.com/FatManlife/component-finder/back-end/internal/models"
	"github.com/FatManlife/component-finder/back-end/internal/utils"
	"github.com/gocolly/colly"
)

//Generalize the categories to send
var components map[string]string = map[string]string{
	"procesoare": "cpu",
	"plăci de bază": "motherboard",
	"plăci video": "gpu",
	"memorie operativă (ram)": "ram",
	"unități de stocare date": "storage",
	"carcase": "case",
	"surse de alimentare (psu)": "psu",
	"sisteme de racire": "cooling systems",
	"m.2 nvme": "ssd",
	"ssd": "ssd",
	"hdd": "hdd",
	"ssd externe": "ssd",
	"hdd externe": "hdd",
	"sisteme de racire cu apă": "cooler",
	"coolere procesoare": "cooler",
	"ventilatoare pc": "fan",
	"all-in-one pc": "aio",
	"mini pc": "pc_mini",

}

func requestBodyProduct(categoryColly *colly.Collector, pageColly *colly.Collector, productColly *colly.Collector, productLinks *chan models.Link, semaphor *chan struct{}){
	//Seting up filter for urls
	var sentUrls sync.Map

	//Scrape the link for categories
	categoryColly.OnHTML("div.xabs_header div.layer1.side-menu ul[data-xcat] > li ", func(e *colly.HTMLElement) {	
		//Getting categories names
		category := strings.ToLower(e.DOM.Find("a").First().Text())

		//Filterin categories
		if strings.Contains(category,"accesorii") || (!strings.Contains(category, "laptopuri") && !strings.Contains(category, "macbook") && !strings.Contains(category, "calculatoare") && !strings.Contains(category, "pc")){
			return
		}

		//Scraping each category
		e.ForEach("ul li", func(_ int, el *colly.HTMLElement){
			subCategory := strings.ToLower(el.ChildText("a"))

			//Filtering for unwanted categories
			if strings.Contains(subCategory,"software") || strings.Contains(subCategory, "setup"){
				return 
			}
			
			link := el.ChildAttr("a", "href")

			//Filtering existing links
			if _, exists := sentUrls.LoadOrStore(link, true); exists{
				return 
			}

			// sending the correct category 
			if category != "componente pc"{
				if strings.Contains(category, "laptopuri") || strings.Contains(category, "macbook"){
					subCategory = "laptop"			
				} else if  category != "all-in-one pc" && category != "mini pc" && (strings.Contains(category, "pc") || strings.Contains(category, "calculatoare")){ 
					subCategory = "pc"
				} else {
					subCategory = components[category]
				} 
			} else {		
				subCategory = components[subCategory]

				if subCategory == ""{
					return
				}
			}	

			//visiting link and seting up a semaphor
			utils.SafeVisit(pageColly, link, collector.NewContext("category",subCategory))
			//Setting up a semaphor to block the for loop
			<- *semaphor 

		})	
	})

	//Setting up a channel to throttler the requests
	for i := 0; i < 1; i++ {
		go func (){	
			for link := range *productLinks {
				utils.SafeVisit(productColly, link.Url, collector.NewContext("category",link.Category))
				time.Sleep(1500 * time.Millisecond)
			}
		}()
	}

	//Specific category Page (cooler/storage)
	pageColly.OnHTML("div.category-baner-item",func(e *colly.HTMLElement){
		category := e.Request.Ctx.Get("category")

		category = "storage"

		if category != "storage" && category != "cooling systems" {
			return 
		}

		category = components[strings.ToLower(e.ChildText("a"))]

		if category == ""{
			return
		}

		utils.SafeVisit(pageColly,e.ChildAttr("a","href"),collector.NewContext("category",category))		
	})
	
	//Category page
	pageColly.OnHTML("div.category-prods.xlists figure.card-product", func(e *colly.HTMLElement){
		category := e.Request.Ctx.Get("category")

		if category == "storage" || category == "cooling systems" {
			return 
		}

		*productLinks <- models.Link{Category: category, Url: e.ChildAttr("a.img-wrap", "href")}
	})

	//Unblocking the sempahor 
	pageColly.OnScraped(func(r *colly.Response) {
		if r.Ctx.Get("category") != "" {
			*semaphor <- struct{}{}
		}
	})

	//going to the next page
	pageColly.OnHTML("a[aria-label=\"Următor\"]", func(e *colly.HTMLElement){
		category := e.Request.Ctx.Get("category")	
		link := e.Attr("href")

		if category != "storage" && category != "cooling systems" && link != "#" {
			utils.SafeVisit(pageColly, link, collector.NewContext("category", category))
		}

	})

	//Scraping the details of the product
	productColly.OnHTML("div.container.page_product",func(e *colly.HTMLElement){
		category := e.Request.Ctx.Get("category")

		switch category {
			case "cpu": cpuHandler(e)
			case "motherboard": motherboardHandler(e)
			case "gpu": gpuHandler(e)
			case "ram": ramHandler(e)
			case "ssd": ssdHandler(e)
			case "hdd": hddHandler(e)
			case "fan": fanHandler(e)
			case "case": caseHandler(e)
			case "psu": psuHandler(e)
			case "cooler": coolerHandler(e)
			case "laptop": laptopHandler(e)
			case "pc": pcHandler(e)
			case "aio": aioHandler(e)
			case "pc_mini": pcMiniHandler(e)
		}
	})	
}



