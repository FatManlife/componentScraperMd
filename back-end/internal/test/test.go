package test

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/FatManlife/component-finder/back-end/internal/collector"
	"github.com/FatManlife/component-finder/back-end/internal/models"
	"github.com/FatManlife/component-finder/back-end/internal/utils"
	"github.com/gocolly/colly"
)

func TestColly(){
	c := collector.New("xstore.md",false)

	// Extracting Computer category
	c.OnHTML("div.container.page_product",func(e *colly.HTMLElement){
		ssdHandler(e)
	})
	
	c.Visit("https://xstore.md/componente-pc/stocare/samsung-pro-ultimate-128gb")
}

func ssdHandler(e *colly.HTMLElement){
	var ssd models.Ssd

	setBaseAttrs(e, &ssd.BaseAttrs)

	e.ForEach("div.tab-content div.chars-item p", func(_ int, el *colly.HTMLElement){
		spec := el.ChildText("span:nth-child(1)") 

		switch spec { 
		case "Producator":
			ssd.BaseAttrs.Brand = strings.TrimSpace(el.ChildText("span:nth-child(2)"))
		case "Capacitatea totală a memoriei":
			capacity := el.ChildText("span:nth-child(2)")
			if strings.Contains(capacity, "TB"){
				ssd.Capacity= utils.CastInt(strings.TrimSpace(capacity)) * 1000
				return 
			}
			ssd.Capacity= utils.CastInt(strings.TrimSpace(capacity))
		case "Viteza maximă de scriere":
			ssd.WritingSpeed = utils.CastInt(strings.TrimSpace(el.ChildText("span:nth-child(2)")))
		case "Viteza maximă de citire":
			ssd.ReadingSpeed = utils.CastInt(strings.TrimSpace(el.ChildText("span:nth-child(2)")))
		case "Form Factor":
			ssd.FormFactor= el.ChildText("span:nth-child(2)")
		}
	})

	data, _ := json.MarshalIndent(ssd, "", "  ")
	fmt.Println(string(data))
}

var ruEng map[string]string = map[string]string {
	"Жидкостное охлаждение":"Liquid cooling",
	"Воздушное охлаждение" : "Air cooling",
	"Вентилятор":"Air cooling",
	"ARGB": "ARGB",
	"RGB": "RGB",
	"Многоцветный":"Multicolor",
	"Для компьютера": "Pc",
	"Для ноутбука": "Laptop",
}

func setBaseAttrs(e *colly.HTMLElement, product *models.BaseProduct){
	product.Name = strings.TrimSpace(e.ChildText("div.top-title h1"))
	product.ImageURL = strings.TrimSpace(e.ChildAttr("div.row.prod_page img", "src"))
	product.Price = utils.CastFloat64(e.ChildText("div.xp-price"))
	product.Website_id = 1
	product.Url = e.Request.URL.String()
}

