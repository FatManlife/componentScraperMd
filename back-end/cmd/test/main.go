package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"github.com/FatManlife/component-finder/back-end/internal/utils"
	"github.com/gocolly/colly"
)

func main() {
	test_1()	
	 test_2()
	// test_3()
}
func test_1() {
	fmt.Println("pcprime")
	c := colly.NewCollector(colly.AllowedDomains("prime-pc.md")) 

	c.OnHTML("div.main_product.container", func (e *colly.HTMLElement)  {
		var ram dto.Ram

		e.ForEach(`div[id="fullDesc"] div.table_row`,func(_ int, el *colly.HTMLElement){
			spec := el.ChildText("div.table_cell:nth-child(1)")

			switch spec {
			case "Объем оперативной памяти": ram.Capacity = utils.CastInt(el.ChildText("div.table_cell:nth-child(2)"))
			case "Частота памяти": ram.Speed = utils.CastInt(el.ChildText("div.table_cell:nth-child(2)"))
			case "Тип оперативной памяти": ram.Type = strings.TrimSpace(el.ChildText("div.table_cell:nth-child(2)"))
			case "Количество планок": ram.Configuration = utils.CastInt(el.ChildText("div.table_cell:nth-child(2)"))
			}
		})

		data, _ := json.MarshalIndent(ram, "", "  ")
		println(string(data))
	})
	
	c.Visit("https://prime-pc.md/products/kingston-fury-beast-expo-ddr5-16gb-6800mhz-kf568c34bbe-16")
}

func test_2() {
	fmt.Println("xstore")
	c := colly.NewCollector(colly.AllowedDomains("xstore.md")) 

	c.OnHTML("div.container.page_product",func(e *colly.HTMLElement){		
		var ram dto.Ram

		e.ForEach("div.tab-content div.chars-item p", func(_ int, el *colly.HTMLElement){
			spec := el.ChildText("span:nth-child(1)") 

			switch spec {
			case "Producator": ram.BaseAttrs.Brand = strings.ToLower(strings.TrimSpace(el.ChildText("span:nth-child(2)")))
			case "Capacitatea totală a memoriei": ram.Capacity = utils.CastInt(el.ChildText("span:nth-child(2)"))
			case "Frecvență memorie": ram.Speed = utils.CastInt(el.ChildText("span:nth-child(2)"))
			case "Standard memorie operativă": ram.Type = strings.TrimSpace(el.ChildText("span:nth-child(2)"))
			case "Aplicare (Utilizare)": ram.Compatibility = strings.TrimSpace(el.ChildText("span:nth-child(2)"))
			case "Numărul de plăci în set": ram.Configuration = utils.CastInt(el.ChildText("span:nth-child(2)"))
			}
		})

		data, _ := json.MarshalIndent(ram, "", "  ")
		fmt.Println(string(data))
	})
	
	c.Visit("https://xstore.md/componente-pc/memorii-ram/kingston-fury-beast-kf432c16bb1k232")
}

func test_3(){
	fmt.Println("neocomputer")

	c := colly.NewCollector(colly.AllowedDomains("neocomputer.md")) 

	c.OnHTML("div#product-product", func (e *colly.HTMLElement)  {
		var ram dto.Ram

		e.ForEach("div.spec__group ul.spec__list li.spec",func (_ int, el *colly.HTMLElement) {	 
			spec := strings.TrimSpace(el.ChildText("span.spec__name"))

			switch spec {
			case "Frecventa memorie RAM": ram.Speed = utils.CastInt(strings.TrimSpace(el.ChildText("span.spec__value")))
			case "Capacitate memorie RAM": ram.Capacity = utils.CastInt(strings.TrimSpace(el.ChildText("span.spec__value")))
			case "Compatibilitate RAM": ram.Compatibility = strings.TrimSpace(el.ChildText("span.spec__value"))
			case "Tip memorie RAM": ram.Type= strings.TrimSpace(el.ChildText("span.spec__value"))
			}
		})

		data,_:= json.MarshalIndent(ram,"","  ")
		fmt.Println(string(data))

	})
	
	c.Visit("https://neocomputer.md/gaming-bloc-de-sistem-neo-20")
}

func castingFloat64Laptop(s string) float64{
	m := regexp.MustCompile(`(?i)(\d+(\.\d+)?)\s*Вт\*ч`).FindStringSubmatch(s)
	if len(m) > 1 {
		v, _ := strconv.ParseFloat(m[1], 64)
		return v
	}
	return 0
}

func extractCapacity(text string) int {
    re := regexp.MustCompile(`(\d+)\s*ГБ`)
    match := re.FindStringSubmatch(text)
    if len(match) > 1 {
        val, _ := strconv.Atoi(match[1])
        return val
    }
    return 0
}
