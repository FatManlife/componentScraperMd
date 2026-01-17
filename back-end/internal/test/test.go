package test

import (
	"fmt"
	"strings"

	"github.com/FatManlife/component-finder/back-end/internal/collector"
	"github.com/FatManlife/component-finder/back-end/internal/models"
	"github.com/FatManlife/component-finder/back-end/internal/utils"
	"github.com/gocolly/colly"
)

func TestColly(){
	c := collector.New("neocomputer.md",false)

	// Extracting Computer category
	c.OnHTML("main#mm-subwrapper",func(e *colly.HTMLElement){
		aioHandler(e)
	})
	
	c.Visit("https://neocomputer.md/all-in-one-asus-a5702")
}

func aioHandler(e *colly.HTMLElement){
	var aio models.Aio
	fmt.Println(aio)
	//setBaseAttrs(e, &aio.BaseAttrs)	
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

