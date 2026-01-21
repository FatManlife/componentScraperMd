package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/FatManlife/component-finder/back-end/internal/collector"
	"github.com/FatManlife/component-finder/back-end/internal/config"
	rawsql "github.com/FatManlife/component-finder/back-end/internal/db/raw_sql"
	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"github.com/FatManlife/component-finder/back-end/internal/utils"
	"github.com/gocolly/colly"
)

func main() {
	c := collector.New("neocomputer.md",false)

	// Extracting Computer category
	c.OnHTML("div#product-product",func(e *colly.HTMLElement){
		fanHandler(e)
	})
	
	c.Visit("https://neocomputer.md/xilence-performance-a-series-xpf120gargbpwm-white")
}

func fanHandler(e *colly.HTMLElement){
	var fan dto.Fan

	setBaseAttrs(e, &fan.BaseAttrs)	
	tempName := strings.TrimSpace(e.ChildText("div.product_container_wrap.container.p-lg-50 h1.section-title.mb-15"))

	if strings.Contains(tempName, "Fan Hub"){
		return
	}

	e.ForEach("div.spec__group ul.spec__list li.spec",func (_ int, el *colly.HTMLElement) {	 
		category := strings.TrimSpace(el.ChildText("span.spec__name"))

		switch category {
		case "Nivel zgomot": 
			if m := regexp.MustCompile(`\d+(?:\.\d+)?`).FindAllString(strings.TrimSpace(el.ChildText("span.spec__value")), -1); len(m) > 0 {
				fan.Noise = utils.CastFloat64(m[len(m)-1])
			}
		case "Nivelul zgomotului": 
			if m := regexp.MustCompile(`\d+(?:\.\d+)?`).FindAllString(strings.TrimSpace(el.ChildText("span.spec__value")), -1); len(m) > 0 {
				fan.Noise = utils.CastFloat64(m[len(m)-1])
			}
		case "Viteza de rotatie": 
			if m := regexp.MustCompile(`(\d+)\D*$`).FindStringSubmatch(strings.TrimSpace(el.ChildText("span.spec__value"))); len(m) > 1 {
				fan.FanRPM = utils.CastInt(m[1])
			}
		case "Viteza maximă a ventilatorului {rpm}":
			if m := regexp.MustCompile(`(\d+)\D*$`).FindStringSubmatch(strings.TrimSpace(el.ChildText("span.spec__value"))); len(m) > 1 {
				fan.FanRPM = utils.CastInt(m[1])
			}
		case "Viteza ventilatorului": 
			if m := regexp.MustCompile(`(\d+)\D*$`).FindStringSubmatch(strings.TrimSpace(el.ChildText("span.spec__value"))); len(m) > 1 {
				fan.FanRPM = utils.CastInt(m[1])
			}
		case "Dimensiune": fan.Size = strings.TrimSpace(el.ChildText("span.spec__value"))
		case "Dimensiune ventilatoare incluse": fan.Size = strings.TrimSpace(el.ChildText("span.spec__value"))
		case "Dimensiunile ventilatorului": fan.Size = strings.TrimSpace(el.ChildText("span.spec__value"))
		case "Dimensiuni": fan.Size = strings.TrimSpace(el.ChildText("span.spec__value"))
		}
	})

	data,_:= json.MarshalIndent(fan,""," ")	
	fmt.Println(string(data))

	db := config.ConnDb()

	lastId, err := rawsql.InsertProduct(db, &fan.BaseAttrs)

	if err != nil {
		fmt.Println("Error inserting product:", err)
		return
	}

	if err := rawsql.InsertFan(db, &fan, lastId);  err != nil {
		fmt.Println("Error inserting fan:", err)
		return
	}
}

func setBaseAttrs(e *colly.HTMLElement, product *dto.BaseProduct){
	product.Name = strings.TrimSpace(e.ChildText("div.product_container_wrap.container.p-lg-50 h1.section-title.mb-15"))
	product.ImageURL = strings.TrimSpace(e.ChildAttr("div.product_container_wrap.container.p-lg-50 img", "src"))
	product.Price = utils.CastFloat64(e.ChildText("div.product_container_wrap.container.p-lg-50 div.price__head.mb-12 span.price__current > span.value"))
	product.Website_id = 3
	product.Url = e.Request.URL.String()
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

