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

	categoryColly.OnHTML("ul.main_list.dropdown-menu > li",func(e *colly.HTMLElement){
		category := categoryMap[strings.ToLower(strings.TrimSpace(e.DOM.Find("a").First().Text()))]
		
		if category == ""{
			return 	
		}

		e.ForEach("div.submenu ul > li",func(i int, el *colly.HTMLElement) {
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

	pageColly.OnHTML("div#catalogue div.catalog_tab div.product", func (e *colly.HTMLElement)  {
		category := e.Request.Ctx.Get("category") 
		link := preflink + e.ChildAttr("a","href")

		chLink := models.Link{Category: category, Url: link}
		*productLinks <- chLink	
	})

	 pageColly.OnHTML("i.pagination_next",func(e *colly.HTMLElement) {	
		next := e.ChildAttr("a","href")

		if next == "" {
			return
		}

		link := preflink + next
		categoryCtx := collector.NewContext("category" ,e.Request.Ctx.Get("category"))
		utils.SafeVisit(pageColly,link,categoryCtx)
	})

	productColly.OnHTML("div.main_product.container", func (e *colly.HTMLElement)  {
		category := e.Request.Ctx.Get("category")

	 	switch category{
		case "aio": aioHandler(e) 
		case "cooler": coolerHandler(e)
		case "cpu": cpuHandler(e)
		case "fan": fanHandler(e)
		case "case": caseHandler(e)
		case "gpu": gpuHandler(e)
		case "hdd": hddHandler(e)
		case "laptop": laptopHandler(e)
		case "motherboard": motherBoardHandler(e)
		case "mini_pc": pcMiniHandler(e)
		case "pc": pcHandler(e)
		case "psu": psuHandler(e)
		case "ram": ramHandler(e)
		case "ssd": ssdHandler(e)
		}
	})	
}
