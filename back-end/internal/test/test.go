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
		pcHandler(e)
	})
	
	c.Visit("https://xstore.md/calculatoare-pc/gaming/raptor-x07")
}

func pcHandler(e *colly.HTMLElement){
	var pc models.Pc

	if strings.Contains(strings.ToLower(strings.TrimSpace(e.ChildText("div.top-title h1"))),"setup"){
		return
	}

	setBaseAttrs(e, &pc.BaseAttrs)	

	e.ForEach("div.tab-content div.chars-item p", func(_ int, el *colly.HTMLElement){
		spec := el.ChildText("span:nth-child(1)") 

		switch spec {
		case "Model placă de bază":
			pc.Motherboard= strings.TrimSpace(el.ChildText("span:nth-child(2)"))
		case "Model carcasă":
			pc.Case = strings.TrimSpace(el.ChildText("span:nth-child(2)"))
		case "Procesor":
			pc.Cpu = strings.TrimSpace(el.ChildText("span:nth-child(2)"))
		case "Model placă video":
			pc.Gpu = strings.TrimSpace(el.ChildText("span:nth-child(2)"))	
		case "Capacitatea RAM":
			pc.Ram = strings.TrimSpace(el.ChildText("span:nth-child(2)"))	
		case "Unitate de stocare":
			pc.Storage = strings.TrimSpace(el.ChildText("span:nth-child(2)"))	
		case "Model sursa de alimentare":
			pc.Psu = strings.TrimSpace(el.ChildText("span:nth-child(2)"))
		}
	})

	data, _ := json.MarshalIndent(pc, "", "  ")
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

