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
	//test_2()
	// test_3()
}
func test_1() {
	fmt.Println("pcprime")
	c := colly.NewCollector(colly.AllowedDomains("prime-pc.md")) 

	c.OnHTML("div.main_product.container", func (e *colly.HTMLElement)  {
		var laptop dto.Laptop

		e.ForEach(`div[id="fullDesc"] div.table_row`,func(_ int, el *colly.HTMLElement){
			spec := el.ChildText("div.table_cell:nth-child(1)")

			switch spec {
			case "Процессор": laptop.Cpu = strings.TrimSpace(strings.Split(el.ChildText("div.table_cell:nth-child(2)"),",")[0])
			case "Видеокарта": 
				gpu := el.ChildText("div.table_cell:nth-child(2)")
				switch gpu {
				case "Интегрированная": laptop.Gpu = "Integrated"
				case "Дискретная": laptop.Gpu = "Dedicated"
				default:
					if strings.Contains(gpu, "Интегрированная") || strings.Contains(gpu, "Дискретная") {
						laptop.Gpu = strings.TrimSpace(strings.Split(gpu,",")[1])
					}
					laptop.Gpu = strings.TrimSpace(laptop.Gpu)
				}
			case "Оперативная память": laptop.Ram = extractCapacity(el.ChildText("div.table_cell:nth-child(2)"))
			case "Накопитель": laptop.Storage = extractCapacity(el.ChildText("div.table_cell:nth-child(2)"))
			case "Диагональ экрана": laptop.Diagonal = utils.CastFloat64(el.ChildText("div.table_cell:nth-child(2)"))
			case "Аккумулятор": laptop.Battery = castingFloat64Laptop(el.ChildText("div.table_cell:nth-child(2)"))
			}
		})

		data, _ := json.MarshalIndent(laptop, "", "  ")
		println(string(data))
	})
	
	c.Visit("https://prime-pc.md//products/acer-nitro-v-15-anv15-52-nhqz7eu007-i5-13420h-16gb-1tb-rtx5050-linux-obsidian-black")
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
			case "Frecventa (MHz)": ram.Speed = utils.CastInt(strings.TrimSpace(el.ChildText("span.spec__value")))
			case "Capacitate memorie RAM": ram.Capacity = utils.CastInt(strings.TrimSpace(el.ChildText("span.spec__value")))
			case "Capacitate (GB)": ram.Capacity = utils.CastInt(strings.TrimSpace(el.ChildText("span.spec__value")))
			case "Compatibilitate RAM": ram.Compatibility = strings.TrimSpace(el.ChildText("span.spec__value"))
			case "Compatibilitate": ram.Compatibility = strings.TrimSpace(el.ChildText("span.spec__value"))
			case "Tip memorie RAM": ram.Type= strings.TrimSpace(el.ChildText("span.spec__value"))
			case "Tip memorie": ram.Type= strings.TrimSpace(el.ChildText("span.spec__value"))
			}
		})

		data,_:= json.MarshalIndent(ram,"","  ")
		fmt.Println(string(data))

	})
	
	c.Visit("https://neocomputer.md/4gb-ddr3l-1600-sodimm-kingston-valueram-pc12800-cl11-1-35v")
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
