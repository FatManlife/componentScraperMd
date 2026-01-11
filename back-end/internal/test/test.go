package test

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/FatManlife/component-finder/back-end/internal/collector"
	"github.com/FatManlife/component-finder/back-end/internal/models"
	"github.com/FatManlife/component-finder/back-end/internal/scraper/pcprime"
	"github.com/FatManlife/component-finder/back-end/internal/utils"
	"github.com/FatManlife/component-finder/back-end/scraper/pcprime"
	"github.com/gocolly/colly"
)

func TestColly(){
	c := collector.New("prime-pc.md",false)

	// Extracting Computer category
	c.OnHTML("div.main_product.container",func(e *colly.HTMLElement){
		LaptopHandler(e)
	})
	
	c.Visit("https://prime-pc.md/products/lenovo-v17-g4-iru-i5-13420h-8gb-512gb-intel-uhd-no-os-iron-grey")
}

func LaptopHandler(e *colly.HTMLElement){
	var laptop models.Laptop

	laptop.Name = e.ChildText("ol.breadcrumb li:last-child")
	laptop.ImageURL = e.ChildAttr("img","src")
	laptop.Brand= e.ChildText("ol.breadcrumb li:nth-last-child(2)")	
	laptop.Price = utils.CastFloat64(e.ChildAttr("div.productPrice b","data-price"))

	e.ForEach(`div[id="fullDesc"] div.table_row`,func(_ int, el *colly.HTMLElement){
		category := el.ChildText("div.table_cell:nth-child(1)")

		switch category{
		case "Процессор":
			laptop.Cpu = strings.Replace(el.ChildText("div.table_cell:nth-child(2)"),"ГГц","GHz",1)
		case "Видеокарта":
			laptop.Gpu = strings.Replace(el.ChildText("div.table_cell:nth-child(2)"),"Интегрированная","Integrated",1)
		case "Оперативная память":
			laptop.Ram = strings.Replace(el.ChildText("div.table_cell:nth-child(2)"),"ГБ","GB",1)
		case "Накопитель":
			laptop.Storage = strings.Replace(el.ChildText("div.table_cell:nth-child(2)"),"ГБ","GB",1)
		case "Диагональ экрана":
			laptop.Diagonal = el.ChildText("div.table_cell:nth-child(2)")
		case "Аккумулятор":
			laptop.Battery = pcprime.CastingFloat64Laptop(el.ChildText("div.table_cell:nth-child(2)"))
		}

	})

	data,_ := json.MarshalIndent(laptop,""," ")
	fmt.Println(string(data))
}




