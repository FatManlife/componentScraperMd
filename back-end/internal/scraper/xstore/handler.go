package xstore

import (
	"strings"
	"sync"
	"time"

	"github.com/FatManlife/component-finder/back-end/internal/collector"
	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"github.com/FatManlife/component-finder/back-end/internal/utils"
	"github.com/gocolly/colly"
)

//Generalize the categories to send
var categoryMap map[string]string = map[string]string{
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

func requestBodyProduct(categoryColly *colly.Collector, pageColly *colly.Collector, productColly *colly.Collector, productLinks *chan dto.Link, semaphor *chan struct{}){
	//Seting up filter for urls
	var sentUrls sync.Map

	//Scrape the link for categories
	categoryColly.OnHTML("div.xabs_header div.layer1.side-menu ul[data-xcat] > li ", func(h *colly.HTMLElement) {	
		//Getting categories names
		category := strings.ToLower(h.DOM.Find("a").First().Text())

		//Filterin categories
		if strings.Contains(category,"accesorii") || (!strings.Contains(category, "laptopuri") && !strings.Contains(category, "macbook") && !strings.Contains(category, "calculatoare") && !strings.Contains(category, "pc")){
			return
		}

		//Scraping each category
		h.ForEach("ul li", func(_ int, el *colly.HTMLElement){
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
					subCategory = categoryMap[category]
				} 
			} else {		
				subCategory = categoryMap[subCategory]

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
	pageColly.OnHTML("div.category-baner-item",func(h *colly.HTMLElement){
		category := h.Request.Ctx.Get("category")

		category = "storage"

		if category != "storage" && category != "cooling systems" {
			return 
		}

		category = categoryMap[strings.ToLower(h.ChildText("a"))]

		if category == ""{
			return
		}

		utils.SafeVisit(pageColly,h.ChildAttr("a","href"),collector.NewContext("category",category))		
	})
	
	//Category page
	pageColly.OnHTML("div.category-prods.xlists figure.card-product", func(h *colly.HTMLElement){
		category := h.Request.Ctx.Get("category")

		if category == "storage" || category == "cooling systems" {
			return 
		}

		*productLinks <- dto.Link{Category: category, Url: h.ChildAttr("a.img-wrap", "href")}
	})

	//Unblocking the sempahor 
	pageColly.OnScraped(func(r *colly.Response) {
		if r.Ctx.Get("category") != "" {
			*semaphor <- struct{}{}
		}
	})

	//going to the next page
	pageColly.OnHTML("a[aria-label=\"Următor\"]", func(h *colly.HTMLElement){
		category := h.Request.Ctx.Get("category")	
		link := h.Attr("href")

		if category != "storage" && category != "cooling systems" && link != "#" {
			utils.SafeVisit(pageColly, link, collector.NewContext("category", category))
		}

	})

	//Scraping the details of the product
	productColly.OnHTML("div.container.page_product",func(h *colly.HTMLElement){
		category := h.Request.Ctx.Get("category")

		switch category {
			case "cpu": cpuHandler(h)
			case "motherboard": motherboardHandler(h)
			case "gpu": gpuHandler(h)
			case "ram": ramHandler(h)
			case "ssd": ssdHandler(h)
			case "hdd": hddHandler(h)
			case "fan": fanHandler(h)
			case "case": caseHandler(h)
			case "psu": psuHandler(h)
			case "cooler": coolerHandler(h)
			case "laptop": laptopHandler(h)
			case "pc": pcHandler(h)
			case "aio": aioHandler(h)
			case "pc_mini": pcMiniHandler(h)
		}
	})	
}



