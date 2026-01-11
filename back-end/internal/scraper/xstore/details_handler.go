package xstore

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/FatManlife/component-finder/back-end/internal/models"
	"github.com/FatManlife/component-finder/back-end/internal/utils"
	"github.com/gocolly/colly"
)

func laptopHandler(e *colly.HTMLElement){
	var laptop models.Laptop	

	laptop.Name = strings.TrimSpace(e.ChildText("div.top-title h1"))
	laptop.ImageURL = strings.TrimSpace(e.ChildAttr("div.row.prod_page img", "src"))
	price := utils.CastFloat64(e.ChildText("div.xp-price"))
	laptop.Price = price

	e.ForEach("div.tab-content div.chars-item p", func(_ int, el *colly.HTMLElement){
		spec := el.ChildText("span:nth-child(1)") 

		switch spec {
		case "Diagonală":
			laptop.Diagonal = strings.TrimSpace(el.ChildText("span:nth-child(2)"))
		case "Producator":
			laptop.Brand = strings.TrimSpace(el.ChildText("span:nth-child(2)"))
		case "Procesor":
			laptop.Cpu = strings.TrimSpace(el.ChildText("span:nth-child(2)"))
		case "Placă video":
			laptop.Gpu = strings.TrimSpace(el.ChildText("span:nth-child(2)"))	
		case "Capacitatea RAM":
			laptop.Ram = strings.TrimSpace(el.ChildText("span:nth-child(2)"))	
		case "Unitate de stocare":
			laptop.Storage = strings.TrimSpace(el.ChildText("span:nth-child(2)"))	
		case "Capacitate baterie":
			laptop.Battery= utils.CastFloat64(el.ChildText("span:nth-child(2)"))
		}
	})

	data,_ := json.MarshalIndent(laptop, "", " ")
	fmt.Println(string(data))
}

func pcHandler(e *colly.HTMLElement){
	var pc models.Pc

	pc.Name = strings.TrimSpace(e.ChildText("div.top-title h1"))
	pc.ImageURL = strings.TrimSpace(e.ChildAttr("div.row.prod_page img", "src"))
	price := utils.CastFloat64(e.ChildText("div.xp-price"))
	pc.Price = price

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

func aioHandler(e *colly.HTMLElement){
	var aio models.Aio

	aio.Name = strings.TrimSpace(e.ChildText("div.top-title h1"))
	aio.ImageURL = strings.TrimSpace(e.ChildAttr("div.row.prod_page img", "src"))
	price := utils.CastFloat64(e.ChildText("div.xp-price"))
	aio.Price = price

	e.ForEach("div.tab-content div.chars-item p", func(_ int, el *colly.HTMLElement){
		spec := el.ChildText("span:nth-child(1)") 

		switch spec { 
		case "Producator":
			aio.Brand = strings.TrimSpace(el.ChildText("span:nth-child(2)"))
		case "Diagonală":
			aio.Diagonal= strings.TrimSpace(el.ChildText("span:nth-child(2)"))
		case "Procesor":
			aio.Cpu = strings.TrimSpace(el.ChildText("span:nth-child(2)"))
		case "Capacitatea RAM":
			aio.Ram = strings.TrimSpace(el.ChildText("span:nth-child(2)"))	
		case "Unitate de stocare":
			aio.Storage = strings.TrimSpace(el.ChildText("span:nth-child(2)"))	
		case "Placă video":
			aio.Gpu= strings.TrimSpace(el.ChildText("span:nth-child(2)"))
		}
	})

	data, _ := json.MarshalIndent(aio, "", "  ")
	fmt.Println(string(data))
}

func cpuHandler(e *colly.HTMLElement){
	var cpu models.Cpu

	cpu.Name = strings.TrimSpace(e.ChildText("div.top-title h1"))
	cpu.ImageURL = strings.TrimSpace(e.ChildAttr("div.row.prod_page img", "src"))
	price := utils.CastFloat64(e.ChildText("div.xp-price"))
	cpu.Price = price

	e.ForEach("div.tab-content div.chars-item p", func(_ int, el *colly.HTMLElement){
		spec := el.ChildText("span:nth-child(1)") 

		switch spec { 
		case "Producator":
			cpu.Brand = strings.TrimSpace(el.ChildText("span:nth-child(2)"))
		case "Numărul de nuclee":
			cpu.Cores = utils.CastInt(el.ChildText("span:nth-child(2)"))
		case "Numărul threads":
			cpu.Threads = utils.CastInt(el.ChildText("span:nth-child(2)"))
		case "Frecvență turbo":
			cpu.BoostClock = utils.CastFloat64(el.ChildText("span:nth-child(2)"))
		case "Frecvența de bază":
			cpu.BaseClock = utils.CastFloat64(el.ChildText("span:nth-child(2)"))
		case "Cache L3":
			cpu.Cache = strings.TrimSpace(el.ChildText("span:nth-child(2)"))
		case "Degajarea de căldură (TDP)":
			cpu.Tdp = utils.CastInt(el.ChildText("span:nth-child(2)"))
		case "Socket procesor":
			cpu.Socket = el.ChildText("span:nth-child(2)")
		}	
	})

	data, _ := json.MarshalIndent(cpu, "", "  ")
	fmt.Println(string(data))
}

func motherboardHandler(e *colly.HTMLElement){
	var motherboard models.Motherboard

	motherboard.Name = strings.TrimSpace(e.ChildText("div.top-title h1"))
	motherboard.ImageURL = strings.TrimSpace(e.ChildAttr("div.row.prod_page img", "src"))
	price :=  utils.CastFloat64(e.ChildText("div.xp-price"))
	motherboard.Price = price

	e.ForEach("div.tab-content div.chars-item p", func(_ int, el *colly.HTMLElement){
		spec := el.ChildText("span:nth-child(1)") 

		switch spec { 
		case "Producator":
			motherboard.Brand = strings.TrimSpace(el.ChildText("span:nth-child(2)"))
		case "Chipset":
			motherboard.Chipset = strings.TrimSpace(el.ChildText("span:nth-child(2)"))
		case "Socket procesor":
			motherboard.Socket = strings.TrimSpace(el.ChildText("span:nth-child(2)"))	
		case "Form-factor memorie operativă":
			motherboard.FormFactorRam = strings.TrimSpace(el.ChildText("span:nth-child(2)"))	
		case "Memorie maximă suportată":
			motherboard.RamSupport = strings.TrimSpace(el.ChildText("span:nth-child(2)"))
		case "Viteza adaptorului de rețea":
			motherboard.NetWork = strings.TrimSpace(el.ChildText("span:nth-child(2)"))
		case "Bluetooth":
			motherboard.Blueetooth = strings.TrimSpace(el.ChildText("span:nth-child(2)"))
		case "Form Factor placă de bază":
			motherboard.FormFactor = strings.TrimSpace(el.ChildText("span:nth-child(2)"))
		}	
	})

	data, _ := json.MarshalIndent(motherboard, "", "  ")
	fmt.Println(string(data))
}

func ramHandler(e *colly.HTMLElement){
	var ram models.Ram

	ram.Name = strings.TrimSpace(e.ChildText("div.top-title h1"))
	ram.ImageURL = strings.TrimSpace(e.ChildAttr("div.row.prod_page img", "src"))
	price := utils.CastFloat64(e.ChildText("div.xp-price"))
	ram.Price = price

	e.ForEach("div.tab-content div.chars-item p", func(_ int, el *colly.HTMLElement){
		spec := el.ChildText("span:nth-child(1)") 

		switch spec { 
		case "Producator":
			ram.Brand = strings.TrimSpace(el.ChildText("span:nth-child(2)"))
		case "Capacitatea totală a memoriei":
			ram.Capacity = utils.CastInt(el.ChildText("span:nth-child(2)"))
		case "Frecvență memorie":
			ram.Speed = utils.CastInt(el.ChildText("span:nth-child(2)"))
		case "Standard memorie operativă":
			ram.Type = strings.TrimSpace(el.ChildText("span:nth-child(2)"))
		case "Form-factor memorie operativă":
			ram.Compatibility = "Pc"
		case "Numărul de plăci în set":
			ram.Configuration = strings.TrimSpace(el.ChildText("span:nth-child(2)"))	
		}
	})

	data, _ := json.MarshalIndent(ram, "", "  ")
	fmt.Println(string(data))
}

func gpuHandler(e *colly.HTMLElement){
	var gpu models.Gpu

	gpu.Name = strings.TrimSpace(e.ChildText("div.top-title h1"))
	gpu.ImageURL = strings.TrimSpace(e.ChildAttr("div.row.prod_page img", "src"))
	price :=  utils.CastFloat64(e.ChildText("div.xp-price"))
	gpu.Price = price

	e.ForEach("div.tab-content div.chars-item p", func(_ int, el *colly.HTMLElement){
		spec := el.ChildText("span:nth-child(1)") 

		switch spec { 
		case "Producator":
			gpu.Brand = strings.TrimSpace(el.ChildText("span:nth-child(2)"))
		case "Frecvența maximă a GPU-ului":
			gpu.GpuFrequency = utils.CastInt(el.ChildText("span:nth-child(2)"))
		case "Memorie video":
			gpu.Vram = utils.CastInt(el.ChildText("span:nth-child(2)"))
		case "Procesor Video":
			gpu.Chipset = strings.TrimSpace(el.ChildText("span:nth-child(2)"))
		case "Frecvența memoriei video":
			gpu.VramFrequency = utils.CastInt(el.ChildText("span:nth-child(2)")) 
		}
	})

	data, _ := json.MarshalIndent(gpu, "", "  ")
	fmt.Println(string(data))
}

func storageHandler(e *colly.HTMLElement){
	var storage models.Storage

	storage.Name = strings.TrimSpace(e.ChildText("div.top-title h1"))
	storage.ImageURL = strings.TrimSpace(e.ChildAttr("div.row.prod_page img", "src"))
	price := utils.CastFloat64(e.ChildText("div.xp-price"))
	storage.Price = price

	e.ForEach("div.tab-content div.chars-item p", func(_ int, el *colly.HTMLElement){
		spec := el.ChildText("span:nth-child(1)") 

		switch spec { 
		case "Producator":
			storage.Brand = strings.TrimSpace(el.ChildText("span:nth-child(2)"))
		case "Capacitatea totală a memoriei":
			unit := strings.Split(strings.TrimSpace(el.ChildText("span:nth-child(2)"))," ")[1]
			if (unit == "TB"){
				storage.Capacity = utils.CastInt(el.ChildText("span:nth-child(2)")) * 1000
			} else {
				storage.Capacity = utils.CastInt(el.ChildText("span:nth-child(2)")) 
			}
		case "Form Factor":
			storage.FormFactor = strings.TrimSpace(el.ChildText("span:nth-child(2)"))
		}
	})

	data, _ := json.MarshalIndent(storage, "", "  ")
	fmt.Println(string(data))
}

func caseHandler(e *colly.HTMLElement){
	var pcCase models.Case

	pcCase.Name = strings.TrimSpace(e.ChildText("div.top-title h1"))
	pcCase.ImageURL = strings.TrimSpace(e.ChildAttr("div.row.prod_page img", "src"))
	price := utils.CastFloat64(e.ChildText("div.xp-price"))
	pcCase.Price = price

	e.ForEach("div.tab-content div.chars-item p", func(_ int, el *colly.HTMLElement){
		spec := el.ChildText("span:nth-child(1)") 
		switch spec { 
		case "Producator":
			pcCase.Brand = strings.TrimSpace(el.ChildText("span:nth-child(2)"))
		case "Format":
			pcCase.Format = strings.TrimSpace(el.ChildText("span:nth-child(2)"))
		case "Compatibilitate form factor placă de bază":
			pcCase.MotherboardFormFactor = strings.TrimSpace(el.ChildText("span:nth-child(2)"))	
		}
	})

	data, _ := json.MarshalIndent(pcCase, "", "  ")
	fmt.Println(string(data))
}

func psuHandler(e *colly.HTMLElement){
	var psu models.Psu

	psu.Name = strings.TrimSpace(e.ChildText("div.top-title h1"))
	psu.ImageURL = strings.TrimSpace(e.ChildAttr("div.row.prod_page img", "src"))
	price := utils.CastFloat64(e.ChildText("div.xp-price"))
	psu.Price = price

	e.ForEach("div.tab-content div.chars-item p", func(_ int, el *colly.HTMLElement){
		spec := el.ChildText("span:nth-child(1)") 

		switch spec { 
		case "Producator":
			psu.Brand = strings.TrimSpace(el.ChildText("span:nth-child(2)"))
		case "Puterea":
			psu.Power = utils.CastInt(el.ChildText("span:nth-child(2)"))
		case "Certificat 80+":
			psu.Efficiency = strings.TrimSpace(el.ChildText("span:nth-child(2)"))
		case "Cabluri detașabile":
			psu.Modularity = strings.TrimSpace(el.ChildText("span:nth-child(2)"))
		case "Form Factor":
			psu.FormFactor = strings.TrimSpace(el.ChildText("span:nth-child(2)"))	
		}
	})

	data, _ := json.MarshalIndent(psu, "", "  ")
	fmt.Println(string(data))
}

func coolerHandler(e *colly.HTMLElement){
	var cooler models.Cooler

	cooler.Name = strings.TrimSpace(e.ChildText("div.top-title h1"))
	cooler.ImageURL = strings.TrimSpace(e.ChildAttr("div.row.prod_page img", "src"))
	price := utils.CastFloat64(e.ChildText("div.xp-price"))
	cooler.Price = price

	e.ForEach("div.tab-content div.chars-item p", func(_ int, el *colly.HTMLElement){
		spec := el.ChildText("span:nth-child(1)") 

		switch spec { 
		case "Producator":
			cooler.Brand = strings.TrimSpace(el.ChildText("span:nth-child(2)"))
		case "Tip răcire":
			cooler.Type = strings.TrimSpace(el.ChildText("span:nth-child(2)"))
		case "Iluminare":
			cooler.Ilumination = strings.TrimSpace(el.ChildText("span:nth-child(2)"))
		case "Viteza maximă de rotație":
			cooler.FanRPM = utils.CastInt(el.ChildText("span:nth-child(2)")) 
		case "Dimensiuni ventilator":
			cooler.Size = strings.TrimSpace(el.ChildText("span:nth-child(2)"))
		case "Nivel zgomot":
			cooler.Noise = utils.CastFloat64(el.ChildText("span:nth-child(2)")) 
		}
	})

	e.ForEach("div.tab-content div.chars-item", func(_ int, el *colly.HTMLElement){
		if el.ChildText("div.chr-title") != "Compatibilitate Socket" {
			return 
		}

		el.ForEach("p", func(_ int, element *colly.HTMLElement){
			if (element.ChildText("span:nth-child(2)") == "Da"){
				cooler.Compatibility = append(cooler.Compatibility, element.ChildText("span:nth-child(1)"))
			}
		})
	})

	data, _ := json.MarshalIndent(cooler, "", "  ")
	fmt.Println(string(data))
}
