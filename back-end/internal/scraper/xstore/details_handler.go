package xstore

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/FatManlife/component-finder/back-end/internal/constants"
	rawsql "github.com/FatManlife/component-finder/back-end/internal/db/raw_sql"
	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"github.com/FatManlife/component-finder/back-end/internal/utils"
	"github.com/gocolly/colly"
)

type handler struct {
	storage *rawsql.Storage
}

func newDetailsHandler(s *rawsql.Storage) *handler {
	return &handler{storage: s}
}

func setBaseAttrs(e *colly.HTMLElement, product *dto.BaseProduct, category string){
	product.Name = strings.TrimSpace(e.ChildText("div.top-title h1"))
	product.ImageURL = strings.TrimSpace(e.ChildAttr("div.row.prod_page img", "src"))
	product.Price = utils.CastFloat64(e.ChildText("div.xp-price"))
	product.Website_id = constants.WebIdMap["xstore"]
	product.Url = e.Request.URL.String()
	product.Category = category 
}

func (h *handler) ssdHandler(e *colly.HTMLElement){
	var ssd dto.Ssd

	setBaseAttrs(e, &ssd.BaseAttrs, "ssd")

	e.ForEach("div.tab-content div.chars-item p", func(_ int, el *colly.HTMLElement){
		spec := el.ChildText("span:nth-child(1)") 

		switch spec { 
		case "Producator": ssd.BaseAttrs.Brand = strings.ToLower(strings.TrimSpace(el.ChildText("span:nth-child(2)")))
		case "Capacitatea totală a memoriei":
			capacity := el.ChildText("span:nth-child(2)")

			if strings.Contains(capacity, "TB"){
				ssd.Capacity= utils.CastInt(strings.TrimSpace(capacity)) * 1000
				return 
			}

			ssd.Capacity= utils.CastInt(strings.TrimSpace(capacity))
		case "Viteza maximă de scriere": ssd.WritingSpeed = utils.CastInt(strings.TrimSpace(el.ChildText("span:nth-child(2)")))
		case "Viteza maximă de citire": ssd.ReadingSpeed = utils.CastInt(strings.TrimSpace(el.ChildText("span:nth-child(2)")))
		case "Form Factor": ssd.FormFactor= strings.TrimSpace(strings.ReplaceAll(el.ChildText("span:nth-child(2)"), "\"", ""))
		}
	})

	if err := h.storage.InsertSsd(&ssd); err != nil {
		fmt.Println("Error inserting SSD:", err)
		return
	}

	data, _ := json.MarshalIndent(ssd, "", "  ")
	fmt.Println(string(data))
}

func (h *handler) hddHandler(e *colly.HTMLElement){
	var hdd dto.Hdd

	setBaseAttrs(e, &hdd.BaseAttrs, "hdd")

	e.ForEach("div.tab-content div.chars-item p", func(_ int, el *colly.HTMLElement){
		spec := el.ChildText("span:nth-child(1)") 

		switch spec { 
		case "Producator": hdd.BaseAttrs.Brand = strings.ToLower(strings.TrimSpace(el.ChildText("span:nth-child(2)")))
		case "Capacitatea totală a memoriei":
			capacity := strings.TrimSpace(el.ChildText("span:nth-child(2)"))

			if strings.Contains(capacity, "TB"){
				hdd.Capacity = utils.CastInt(capacity) * 1000
			} else {
				hdd.Capacity = utils.CastInt(strings.TrimSpace(capacity))
			}
		case "Viteza de rotație": hdd.RotationSpeed = utils.CastInt(strings.TrimSpace(el.ChildText("span:nth-child(2)")))
		case "Form Factor": hdd.FormFactor = strings.TrimSpace(strings.Split(el.ChildText("span:nth-child(2)"), "\"")[0])
		}
	})

	if err := h.storage.InsertHdd(&hdd); err != nil {
		fmt.Println("Error inserting HDD:", err)
		return
	}

	data, _ := json.MarshalIndent(hdd, "", "  ")
	fmt.Println(string(data))
}

func (h *handler) fanHandler(e *colly.HTMLElement){
	var fan dto.Fan

	setBaseAttrs(e, &fan.BaseAttrs, "fan")

	e.ForEach("div.tab-content div.chars-item p", func(_ int, el *colly.HTMLElement){
		spec := el.ChildText("span:nth-child(1)") 

		switch spec { 
		case "Producator": fan.BaseAttrs.Brand = strings.ToLower(strings.TrimSpace(el.ChildText("span:nth-child(2)")))
		case "Viteza maximă de rotație": fan.FanRPM = utils.CastInt(el.ChildText("span:nth-child(2)")) 
		case "Dimensiuni ventilator": fan.Size = strings.TrimSpace(el.ChildText("span:nth-child(2)"))
		case "Nivel zgomot": fan.Noise = utils.CastFloat64(el.ChildText("span:nth-child(2)")) 
		}
	})

	if err := h.storage.InsertFan(&fan); err != nil {
		fmt.Println("Error inserting Fan:", err)
		return
	}

	data, _ := json.MarshalIndent(fan, "", "  ")
	fmt.Println(string(data))
}

func (h *handler) pcMiniHandler(e *colly.HTMLElement){
	var pc dto.PcMini

	setBaseAttrs(e, &pc.BaseAttrs, "pc_mini")	

	e.ForEach("div.tab-content div.chars-item p", func(_ int, el *colly.HTMLElement){
		spec := el.ChildText("span:nth-child(1)") 

		switch spec {
		case "Procesor": pc.Cpu = strings.TrimSpace(el.ChildText("span:nth-child(2)"))
		case "Producator": pc.BaseAttrs.Brand = strings.ToLower(strings.TrimSpace(el.ChildText("span:nth-child(2)")))
		case "Model placă video": pc.Gpu = strings.TrimSpace(el.ChildText("span:nth-child(2)"))	
		case "Capacitatea RAM": pc.Ram = utils.CastInt(strings.TrimSpace(el.ChildText("span:nth-child(2)")))	
		case "Unitate de stocare": pc.Storage = utils.CastInt(strings.TrimSpace(el.ChildText("span:nth-child(2)")))		
		}
	})

	if err := h.storage.InsertPcMini(&pc); err != nil {
		fmt.Println("Error inserting Pc Mini:", err)
		return
	}

	data, _ := json.MarshalIndent(pc, "", "  ")
	fmt.Println(string(data))
}

func (h *handler) laptopHandler(e *colly.HTMLElement){
	var laptop dto.Laptop	
	
	setBaseAttrs(e, &laptop.BaseAttrs, "laptop")	

	e.ForEach("div.tab-content div.chars-item p", func(_ int, el *colly.HTMLElement){
		spec := el.ChildText("span:nth-child(1)") 

		switch spec {
		case "Diagonală": laptop.Diagonal = utils.CastFloat64(el.ChildText("span:nth-child(2)"))
		case "Producator": laptop.BaseAttrs.Brand = strings.ToLower(strings.TrimSpace(el.ChildText("span:nth-child(2)")))
		case "Procesor": laptop.Cpu = strings.TrimSpace(el.ChildText("span:nth-child(2)"))
		case "Placă video": laptop.Gpu = strings.TrimSpace(el.ChildText("span:nth-child(2)"))	
		case "Procesor Video": laptop.Gpu = strings.TrimSpace(el.ChildText("span:nth-child(2)"))	
		case "Tip placă video": laptop.Gpu = strings.TrimSpace(el.ChildText("span:nth-child(2)"))	
		case "Capacitatea RAM": laptop.Ram = utils.CastInt(el.ChildText("span:nth-child(2)"))	
		case "Unitate de stocare": 
			storage := strings.TrimSpace(el.ChildText("span:nth-child(2)"))
			if strings.Contains(storage, "TB") {
				laptop.Storage = utils.CastInt(storage) * 1000
			} else {
				laptop.Storage = utils.CastInt(storage)
			}
		case "Capacitate baterie": laptop.Battery= utils.CastFloat64(el.ChildText("span:nth-child(2)"))
		}
	})

	if err := h.storage.InsertLaptop(&laptop); err != nil {
		fmt.Println("Error inserting Laptop:", err)
		return
	}

	data,_ := json.MarshalIndent(laptop, "", " ")
	fmt.Println(string(data))
}

func (h *handler) pcHandler(e *colly.HTMLElement){
	var pc dto.Pc

	if strings.Contains(strings.ToLower(strings.TrimSpace(e.ChildText("div.top-title h1"))),"setup"){
		return
	}

	setBaseAttrs(e, &pc.BaseAttrs, "pc")	

	e.ForEach("div.tab-content div.chars-item p", func(_ int, el *colly.HTMLElement){
		spec := el.ChildText("span:nth-child(1)") 

		switch spec {
		case "Model placă de bază": pc.Motherboard = strings.TrimSpace(el.ChildText("span:nth-child(2)"))
		case "Model carcasă": pc.Case = strings.TrimSpace(el.ChildText("span:nth-child(2)"))
		case "Procesor": pc.Cpu = strings.TrimSpace(el.ChildText("span:nth-child(2)"))
		case "Model placă video": pc.Gpu = strings.TrimSpace(el.ChildText("span:nth-child(2)"))	
		case "Capacitatea RAM": pc.Ram = utils.CastInt(el.ChildText("span:nth-child(2)"))	
		case "Unitate de stocare": 
			storage := el.ChildText("span:nth-child(2)")	
			if strings.Contains(storage, "TB") {
				pc.Storage = utils.CastInt(storage) * 1000
			} else {
				pc.Storage = utils.CastInt(storage)
			}
		case "Model sursa de alimentare": pc.Psu = strings.TrimSpace(el.ChildText("span:nth-child(2)"))
		}
	})

	if err := h.storage.InsertPc(&pc); err != nil {
		fmt.Println("Error inserting Pc:", err)
		return
	}

	data, _ := json.MarshalIndent(pc, "", "  ")
	fmt.Println(string(data))
}

func (h *handler) aioHandler(e *colly.HTMLElement){
	var aio dto.Aio

	setBaseAttrs(e, &aio.BaseAttrs, "aio")	

	e.ForEach("div.tab-content div.chars-item p", func(_ int, el *colly.HTMLElement){
		spec := el.ChildText("span:nth-child(1)") 

		switch spec { 
		case "Producator": aio.BaseAttrs.Brand = strings.ToLower(strings.TrimSpace(el.ChildText("span:nth-child(2)")))
		case "Diagonală": aio.Diagonal= strings.TrimSpace(strings.Split(el.ChildText("span:nth-child(2)"),"\"")[0])
		case "Procesor": aio.Cpu = strings.TrimSpace(el.ChildText("span:nth-child(2)"))
		case "Capacitatea RAM": aio.Ram = utils.CastInt(el.ChildText("span:nth-child(2)"))	
			case "Unitate de stocare":
				var storage int		
				if strings.Contains(el.ChildText("span:nth-child(2)"), "TB") {
					storage = utils.CastInt(el.ChildText("span:nth-child(2)")) * 1000
				} else {
					storage = utils.CastInt(el.ChildText("span:nth-child(2)"))
				}
				aio.Storage = storage	
		case "Placă video": aio.Gpu= strings.TrimSpace(el.ChildText("span:nth-child(2)"))
		}
	})

	if err := h.storage.InsertAio(&aio); err != nil {
		fmt.Println("Error inserting Aio:", err)
		return
	}

	data, _ := json.MarshalIndent(aio, "", "  ")
	fmt.Println(string(data))
}

func (h *handler) cpuHandler(e *colly.HTMLElement){
	var cpu dto.Cpu

	setBaseAttrs(e, &cpu.BaseAttrs, "cpus")

	e.ForEach("div.tab-content div.chars-item p", func(_ int, el *colly.HTMLElement){
		spec := el.ChildText("span:nth-child(1)") 

		switch spec { 
		case "Producator": cpu.BaseAttrs.Brand = strings.ToLower(strings.TrimSpace(el.ChildText("span:nth-child(2)")))
		case "Numărul de nuclee": cpu.Cores = utils.CastInt(el.ChildText("span:nth-child(2)"))
		case "Numărul threads": cpu.Threads = utils.CastInt(el.ChildText("span:nth-child(2)"))
		case "Frecvență turbo": cpu.BoostClock = utils.CastFloat64(el.ChildText("span:nth-child(2)"))
		case "Frecvența de bază": cpu.BaseClock = utils.CastFloat64(el.ChildText("span:nth-child(2)"))
		case "Degajarea de căldură (TDP)": cpu.Tdp = utils.CastInt(el.ChildText("span:nth-child(2)"))
		case "Socket procesor": cpu.Socket = el.ChildText("span:nth-child(2)")
		}	
	})

	if err := h.storage.InsertCpu(&cpu); err != nil {
		fmt.Println("Error inserting CPU:", err)
		return
	}

	data, _ := json.MarshalIndent(cpu, "", "  ")
	fmt.Println(string(data))
}

func (h *handler) motherboardHandler(e *colly.HTMLElement){
	var motherboard dto.Motherboard

	setBaseAttrs(e, &motherboard.BaseAttrs, "motherboard")

	e.ForEach("div.tab-content div.chars-item p", func(_ int, el *colly.HTMLElement){
		spec := el.ChildText("span:nth-child(1)") 

		switch spec { 
		case "Producator": motherboard.BaseAttrs.Brand = strings.ToLower(strings.TrimSpace(el.ChildText("span:nth-child(2)")))
		case "Chipset": motherboard.Chipset = strings.TrimSpace(el.ChildText("span:nth-child(2)"))
		case "Socket procesor": motherboard.Socket = strings.TrimSpace(el.ChildText("span:nth-child(2)"))	
		case "Form-factor memorie operativă": motherboard.FormFactorRam = strings.TrimSpace(el.ChildText("span:nth-child(2)"))	
		case "Memorie maximă suportată": motherboard.RamSupport = strings.TrimSpace(el.ChildText("span:nth-child(2)"))	
		case "Form Factor placă de bază": motherboard.FormFactor = strings.TrimSpace(el.ChildText("span:nth-child(2)"))
		}	
	})

	if err := h.storage.InsertMotherboard(&motherboard); err != nil {
		fmt.Println("Error inserting Motherboard:", err)
		return
	}

	data, _ := json.MarshalIndent(motherboard, "", "  ")
	fmt.Println(string(data))
}

func (h *handler) ramHandler(e *colly.HTMLElement){
	var ram dto.Ram

	setBaseAttrs(e, &ram.BaseAttrs, "ram")

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

	if err := h.storage.InsertRam(&ram); err != nil {
		fmt.Println("Error inserting Ram:", err)
		return
	}

	data, _ := json.MarshalIndent(ram, "", "  ")
	fmt.Println(string(data))
}

func (h *handler) gpuHandler(e *colly.HTMLElement){
	var gpu dto.Gpu

	setBaseAttrs(e, &gpu.BaseAttrs, "gpu")	

	e.ForEach("div.tab-content div.chars-item p", func(_ int, el *colly.HTMLElement){
		spec := el.ChildText("span:nth-child(1)") 

		switch spec { 
		case "Frecvența maximă a GPU-ului": gpu.GpuFrequency = utils.CastInt(el.ChildText("span:nth-child(2)"))
		case "Memorie video": gpu.Vram = utils.CastInt(el.ChildText("span:nth-child(2)"))
		case "Procesor Video": gpu.Chipset = strings.TrimSpace(el.ChildText("span:nth-child(2)"))
		case "Frecvența memoriei video": gpu.VramFrequency = utils.CastInt(el.ChildText("span:nth-child(2)")) 
		}
	})

	if err := h.storage.InsertGpu(&gpu); err != nil {
		fmt.Println("Error inserting GPU:", err)
		return
	}

	data, _ := json.MarshalIndent(gpu, "", "  ")
	fmt.Println(string(data))
}

func (h *handler) caseHandler(e *colly.HTMLElement){
	var pcCase dto.Case

	setBaseAttrs(e, &pcCase.BaseAttrs, "case")

	e.ForEach("div.tab-content div.chars-item p", func(_ int, el *colly.HTMLElement){
		spec := el.ChildText("span:nth-child(1)") 
		switch spec { 
		case "Producator": pcCase.BaseAttrs.Brand = strings.ToLower(strings.TrimSpace(el.ChildText("span:nth-child(2)")))
		case "Format": pcCase.Format = strings.TrimSpace(strings.ReplaceAll(el.ChildText("span:nth-child(2)"), "-", ""))
		case "Compatibilitate form factor placă de bază": pcCase.MotherboardFormFactor = strings.TrimSpace(el.ChildText("span:nth-child(2)"))	
		}
	})

	if err := h.storage.InsertCase(&pcCase); err != nil {
		fmt.Println("Error inserting Case:", err)
		return
	}

	data, _ := json.MarshalIndent(pcCase, "", "  ")
	fmt.Println(string(data))
}

func (h *handler) psuHandler(e *colly.HTMLElement){
	var psu dto.Psu

	setBaseAttrs(e, &psu.BaseAttrs, "psu")

	e.ForEach("div.tab-content div.chars-item p", func(_ int, el *colly.HTMLElement){
		spec := el.ChildText("span:nth-child(1)") 

		switch spec { 
		case "Producator": psu.BaseAttrs.Brand = strings.ToLower(strings.TrimSpace(el.ChildText("span:nth-child(2)")))
		case "Puterea": psu.Power = utils.CastInt(el.ChildText("span:nth-child(2)"))
		case "Certificat 80+": psu.Efficiency = strings.TrimSpace(strings.ReplaceAll(el.ChildText("span:nth-child(2)"), "+", " PLUS"))
		case "Form Factor": psu.FormFactor = strings.TrimSpace(el.ChildText("span:nth-child(2)"))	
		}
	})

	if err := h.storage.InsertPsu(&psu); err != nil {
		fmt.Println("Error inserting PSU:", err)
		return
	}

	data, _ := json.MarshalIndent(psu, "", "  ")
	fmt.Println(string(data))
}

func (h *handler) coolerHandler(e *colly.HTMLElement){
	var cooler dto.Cooler

	setBaseAttrs(e, &cooler.BaseAttrs, "cooler")

	e.ForEach("div.tab-content div.chars-item p", func(_ int, el *colly.HTMLElement){
		spec := el.ChildText("span:nth-child(1)") 

		switch spec { 
		case "Producator": cooler.BaseAttrs.Brand = strings.ToLower(strings.TrimSpace(el.ChildText("span:nth-child(2)")))
		case "Tip răcire": 
			tip := strings.ToLower(el.ChildText("span:nth-child(2)"))	
			if strings.Contains(tip,"aer"){
				cooler.Type = "Air cooling"
			} else if strings.Contains(tip,"apă") {
				cooler.Type = "Liquid cooling"
			}
		case "Viteza maximă de rotație": cooler.FanRPM = utils.CastInt(el.ChildText("span:nth-child(2)")) 
		case "Dimensiuni ventilator": cooler.Size = strings.TrimSpace(el.ChildText("span:nth-child(2)"))
		case "Nivel zgomot": cooler.Noise = utils.CastFloat64(el.ChildText("span:nth-child(2)")) 
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

	cooler_id, err := h.storage.InsertCooler(&cooler)

	if err != nil {
		fmt.Println("Error inserting Cooler:", err)
		return
	}

	if err := h.storage.InsertCoolerCompatibility(cooler_id, cooler.Compatibility); err != nil {
		fmt.Println("Error inserting Cooler Compatibility:", err)
		return
	}

	data, _ := json.MarshalIndent(cooler, "", "  ")
	fmt.Println(string(data))
}