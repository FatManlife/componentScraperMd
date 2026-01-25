package neocomputer

import (
	"strings"
	"time"

	"github.com/FatManlife/component-finder/back-end/internal/collector"
	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"github.com/FatManlife/component-finder/back-end/internal/utils"
	"github.com/gocolly/colly"
	"gorm.io/gorm"
)

var categoryMap map[string]string = map[string]string {
	 "procesoare": "cpu",
	"storage hdd": "hdd",
	"storage ssd": "ssd",
	"memorii ram": "ram",
	"plăci de bază": "motherboard",
	"plăci video": "gpu",
	"carcase": "case",
	"surse de alimentare": "psu",
	"cooler procesoare": "cooler",
	"cooler carcase": "fan",
	"desktop pc": "pc",
	"gaming pc": "pc",
	"all-in-one pc": "aio",
	"mini pc": "mini_pc",
}

// requestBodyProducts scrapes product links from category and page collectors, then scrapes product details from product collector.
func requestBodyProducts(categoryColly *colly.Collector, pageColly *colly.Collector, productColly *colly.Collector, productLink *chan dto.Link, db *gorm.DB) {
	categoryColly.OnHTML("ul.dropdown-content.categories  li.nav-wrap",func(h *colly.HTMLElement) {
		category := strings.TrimSpace(strings.ToLower(h.ChildText("a.submenu")))

		if strings.Contains(category, "laptopuri"){
			link := "https://neocomputer.md/" + h.ChildAttr("a.submenu", "href")  + "/notebook"
			utils.SafeVisit(pageColly, link, collector.NewContext("category", "laptop"))
			return
		} else if !strings.Contains(category, "pc") {
			return
		}

		h.ForEach("div.subcategories ul.subcategories-list > li",func (_ int, el *colly.HTMLElement) {
			category = strings.TrimSpace(strings.ToLower(el.DOM.Find("a").First().Text()))

			if !strings.Contains(category, "componente") && !strings.Contains(category, "computere") {
				return
			}

			el.ForEach("ul > li > a.sub-title",func (_ int, e *colly.HTMLElement) {
				category := categoryMap[strings.TrimSpace(strings.ToLower(e.Text))]
				
				if category == "" {
					return
				}

				link := "https://neocomputer.md/" + e.Attr("href")
				utils.SafeVisit(pageColly, link, collector.NewContext("category", category))
			})
		})
	})

// Pagination and product links
	pageColly.OnHTML("div.row.products-list div.col-lg-4.col-6 a", func(h *colly.HTMLElement) {
		category := h.Request.Ctx.Get("category")

		*productLink <- dto.Link{Url: h.Attr("href"), Category: category}
	})

// Iterate through products
	for i := 0; i < 1; i++ {
		go func (){
			for link := range *productLink {
				utils.SafeVisit(productColly, link.Url, collector.NewContext("category", link.Category))
				time.Sleep(1000 * time.Millisecond)
			}
		}()
	}

// Next page
	pageColly.OnHTML("li.page-nav.next a",func(h *colly.HTMLElement) {	
		link := h.Attr("href")

		if link == "" {
			return
		}

		categoryCtx := collector.NewContext("category" ,h.Request.Ctx.Get("category"))
		utils.SafeVisit(pageColly,link,categoryCtx)
	})

// Product details
	productColly.OnHTML("div#product-product", func(h *colly.HTMLElement) {
		category := h.Request.Ctx.Get("category")

		switch category {
		case "cpu": cpuHandler(h, db, category)
		case "gpu": gpuHandler(h, db, category)
		case "motherboard": motherboardHandler(h, db, category)
		case "ram": ramHandler(h, db, category)
		case "hdd": hddHandler(h, db, category)
		case "ssd": ssdHandler(h, db, category)
		case "psu": psuHandler(h, db, category)
		case "case": caseHandler(h, db, category)
		case "cooler": coolerHandler(h, db, category)
		case "fan": fanHandler(h, db, category)	
		case "pc": pcHandler(h, db, category)
		case "laptop": laptopHandler(h, db, category)
		case "aio": aioHandler(h, db, category)
		case "mini_pc": pcMiniHandler(h, db, category)
		}
	})
}