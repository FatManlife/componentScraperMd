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
	c := collector.New("prime-pc.md",false)

	// Extracting Computer category
	c.OnHTML("div.main_product.container",func(e *colly.HTMLElement){
		aioHandler(e)
	})
	
	c.Visit("https://prime-pc.md/products/asus-aio-a3402i3-1315u-8gb-512gb-intel-uhd-no-os-white")
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

func aioHandler(e *colly.HTMLElement){
	var aio models.Aio

	aio.Name = e.ChildText("ol.breadcrumb li:last-child")
	aio.ImageURL = e.ChildAttr("img","src")
	aio.Price = utils.CastFloat64(e.ChildAttr("div.productPrice b","data-price"))
	aio.Brand = e.ChildText("ol.breadcrumb li:nth-last-child(2)")

	e.ForEach(`div[id="fullDesc"] div.table_row`,func(_ int, el *colly.HTMLElement){
		category := el.ChildText("div.table_cell:nth-child(1)")

		switch category{
		case "Процессор":
			aio.Cpu = strings.Replace(el.ChildText("div.table_cell:nth-child(2)"),"ГГц","GHz",1)
		case "Модель видеокарты":
			aio.Gpu = el.ChildText("div.table_cell:nth-child(2) ")
		case "Объем оперативной памяти":
			aio.Ram = strings.Replace(el.ChildText("div.table_cell:nth-child(2)"),"ГБ","GB",1)
		case "Объем SSD":
			aio.Storage= strings.Replace(el.ChildText("div.table_cell:nth-child(2)"),"ГБ","GB",1)
		case "Экран":
			aio.Diagonal = strings.Split(el.ChildText("div.table_cell:nth-child(2)"),",")[0]
		}
	})

	data,_ := json.MarshalIndent(aio,""," ")
	fmt.Println(string(data))
}