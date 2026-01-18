package pcprime

import (
	"strings"
	"time"

	"github.com/FatManlife/component-finder/back-end/internal/collector"
	"github.com/FatManlife/component-finder/back-end/internal/models"
	"github.com/FatManlife/component-finder/back-end/internal/utils"
	"github.com/gocolly/colly"
)

var categoryMap map[string]string = map[string]string{
	"ноутбуки и аксессуары": "laptops and accessories",
	"комплектующие для компьютера": "computers components",
	"компьютеры и периферия": "computers and peripherals",
	"ноутбуки": "laptop", 
	"память для ноутбука *": "ram",
	"hdd накопители для ноутбуков *": "hdd",
	"ssd накопители для ноутбуков *": "ssd",
	"процессоры intel" : "cpu",
	"процессоры amd" : "cpu",
	"кулеры для процессора": "cooler",
	"вентиляторы": "fan",
	"материнские платы под intel": "motherboard", 
	"материнские платы под amd": "motherboard",
	"оперативная память": "ram",
	"ssd": "ssd",
	"внутренние hdd": "hdd",
	"внешние hdd, ssd": "ssd/hdd",
	"видеокарты": "gpu",
	"корпуса": "case",
	"блоки питания": "psu",
	"компьютеры prime-pc": "pc",
	"компьютеры": "pc",
	"моноблоки": "aio",
	"мини пк": "mini_pc",
}

func requestBodyProducts(categoryColly *colly.Collector, pageColly *colly.Collector, productColly *colly.Collector, productLinks *chan models.Link){
	preflink := "https://prime-pc.md/" 

	// Category extraction
	categoryColly.OnHTML("ul.main_list.dropdown-menu > li",func(h *colly.HTMLElement){
		category := categoryMap[strings.ToLower(strings.TrimSpace(h.DOM.Find("a").First().Text()))]
		
		if category == ""{
			return 	
		}

		h.ForEach("div.submenu ul > li",func(i int, el *colly.HTMLElement) {
			category = categoryMap[strings.TrimSpace(strings.ToLower(el.ChildText("a")))]

			if category == "" {
				return
			}

			link := preflink + el.ChildAttr("a","href") + "&in_stock=1"

			if category == "ssd/hdd"{
				category = "ssd"
				tempLink := link + "&261[]=SSD"
				utils.SafeVisit(pageColly,tempLink,collector.NewContext("category",category))

				category = "hdd"
				tempLink = link + "&261[]=HDD"
				utils.SafeVisit(pageColly,tempLink,collector.NewContext("category",category))
				return
			}

			utils.SafeVisit(pageColly,link,collector.NewContext("category",category))
		})
	})

	for i := 0; i < 1; i++ {
		go func(){
			for product := range *productLinks {
				utils.SafeVisit(productColly,product.Url,collector.NewContext("category",product.Category))				
				time.Sleep(1 * time.Second)
			}
		}()
	}

	// Product links extraction
	pageColly.OnHTML("div#catalogue div.catalog_tab div.product", func (h *colly.HTMLElement)  {
		category := h.Request.Ctx.Get("category") 
		link := preflink + h.ChildAttr("a","href")

		chLink := models.Link{Category: category, Url: link}
		*productLinks <- chLink	
	})

	// next page extraction
	 pageColly.OnHTML("i.pagination_next",func(h *colly.HTMLElement) {	
		next := h.ChildAttr("a","href")

		if next == "" {
			return
		}

		link := preflink + next
		categoryCtx := collector.NewContext("category" ,h.Request.Ctx.Get("category"))
		utils.SafeVisit(pageColly,link,categoryCtx)
	})

	// Product extraction
	productColly.OnHTML("div.main_product.container", func (h *colly.HTMLElement)  {
		category := h.Request.Ctx.Get("category")

	 	switch category{
		case "aio": aioHandler(h) 
		case "cooler": coolerHandler(h)
		case "cpu": cpuHandler(h)
		case "fan": fanHandler(h)
		case "case": caseHandler(h)
		case "gpu": gpuHandler(h)
		case "hdd": hddHandler(h)
		case "laptop": laptopHandler(h)
		case "motherboard": motherBoardHandler(h)
		case "mini_pc": pcMiniHandler(h)
		case "pc": pcHandler(h)
		case "psu": psuHandler(h)
		case "ram": ramHandler(h)
		case "ssd": ssdHandler(h)
		}
	})	
}
