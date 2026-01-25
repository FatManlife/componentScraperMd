package neocomputer

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/FatManlife/component-finder/back-end/internal/constants"
	rawsql "github.com/FatManlife/component-finder/back-end/internal/db/raw_sql"
	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"github.com/FatManlife/component-finder/back-end/internal/utils"
	"github.com/gocolly/colly"
	"gorm.io/gorm"
)

func setBaseAttrs(e *colly.HTMLElement, product *dto.BaseProduct, category string){
	product.Name = strings.TrimSpace(e.ChildText("div.product_container_wrap.container.p-lg-50 h1.section-title.mb-15"))
	product.ImageURL = strings.TrimSpace(e.ChildAttr("div.product_container_wrap.container.p-lg-50 img", "src"))
	product.Price = utils.CastFloat64(e.ChildText("div.product_container_wrap.container.p-lg-50 div.price__head.mb-12 span.price__current > span.value"))
	product.Website_id = constants.WebIdMap["neocomputer"]
	product.Url = e.Request.URL.String()
	product.Category_id = constants.CategoryIdMap[category]
}

func fanHandler(e *colly.HTMLElement, db *gorm.DB, category string){
	var fan dto.Fan

	setBaseAttrs(e, &fan.BaseAttrs, category)	
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

	if err := rawsql.InsertFan(db, &fan); err != nil {
		fmt.Println("Error inserting Fan:", err)
		return
	}		

	data,_:= json.MarshalIndent(fan,""," ")	
	fmt.Println(string(data))
}

func coolerHandler(e *colly.HTMLElement, db *gorm.DB, category string){
	var cooler dto.Cooler

	setBaseAttrs(e, &cooler.BaseAttrs, category)	
	tempName := strings.TrimSpace(e.ChildText("div.product_container_wrap.container.p-lg-50 h1.section-title.mb-15"))

	if strings.Contains(tempName, "Mounting Kit"){
		return
	}

	e.ForEach("div.spec__group ul.spec__list li.spec",func (_ int, el *colly.HTMLElement) {	 
		category := strings.TrimSpace(el.ChildText("span.spec__name"))

		switch category {
		case "Viteza ventilatorului": 
			if m := regexp.MustCompile(`(\d+)\D*$`).FindStringSubmatch(el.ChildText("span.spec__value")); len(m) > 1 {
				cooler.FanRPM = utils.CastInt(m[1])
			}
		case "Viteza de rotatie":
			if m := regexp.MustCompile(`(\d+)\D*$`).FindStringSubmatch(el.ChildText("span.spec__value")); len(m) > 1 {
				cooler.FanRPM = utils.CastInt(m[1])
			}
		case "Producator": cooler.BaseAttrs.Brand = strings.TrimSpace(el.ChildText("span.spec__value"))
		case "Dimensiunile ventilatorului": cooler.Size = strings.TrimSpace(el.ChildText("span.spec__value"))
		case "Dimensiune": cooler.Size = strings.TrimSpace(el.ChildText("span.spec__value"))
		case "Tip racire": cooler.Type = strings.TrimSpace(el.ChildText("span.spec__value"))
		case "Nivelul zgomotului": 
			if m := regexp.MustCompile(`\d+(?:\.\d+)?`).FindAllString(el.ChildText("span.spec__value"), -1); len(m) > 0 {
				cooler.Noise = utils.CastFloat64(m[len(m)-1])
			}
		case "Nivel zgomot":
			if m := regexp.MustCompile(`\d+(?:\.\d+)?`).FindAllString(el.ChildText("span.spec__value"), -1); len(m) > 0 {
				cooler.Noise = utils.CastFloat64(m[len(m)-1])
			}
		case "Soket INTEL": cooler.Compatibility = append(cooler.Compatibility, strings.Split(el.ChildText("span.spec__value"),"/")...)
		case "Soket AMD": cooler.Compatibility = append(cooler.Compatibility, strings.Split(el.ChildText("span.spec__value"),"/")...)
		case "Compatibilitate": cooler.Compatibility = append(cooler.Compatibility, strings.Split(el.ChildText("span.spec__value"),"/")...)
		}
	})

	cooler_id, err := rawsql.InsertCooler(db, &cooler)

	if err != nil {
		fmt.Println("Error inserting Cooler:", err)
		return
	}

	if err := rawsql.InsertCoolerCompatibility(db, cooler_id, cooler.Compatibility); err != nil {
		fmt.Println("Error inserting Cooler Compatibility:", err)
		return
	}

	data,_:= json.MarshalIndent(cooler,""," ")	
	fmt.Println(string(data))
}

func psuHandler(e *colly.HTMLElement, db *gorm.DB, category string){
	var psu dto.Psu

	setBaseAttrs(e, &psu.BaseAttrs, category)	

	tempName := strings.TrimSpace(e.ChildText("div.product_container_wrap.container.p-lg-50 h1.section-title.mb-15"))

	if strings.Contains(tempName, "Cable"){
		return
	}

	e.ForEach("div.spec__group ul.spec__list li.spec",func (_ int, el *colly.HTMLElement) {	 
		category := strings.TrimSpace(el.ChildText("span.spec__name"))

		switch category {
		case "Form factor": psu.FormFactor = strings.TrimSpace(el.ChildText("span.spec__value"))
		case "Putere sursa (W)": psu.Power = utils.CastInt(strings.TrimSpace(el.ChildText("span.spec__value")))
		case "Certificat 80+": psu.Efficiency = strings.TrimSpace(el.ChildText("span.spec__value"))
		}
	})

	if err := rawsql.InsertPsu(db, &psu); err != nil {
		fmt.Println("Error inserting PSU:", err)
		return
	}

	data,_:= json.MarshalIndent(psu,""," ")	
	fmt.Println(string(data))
}

func caseHandler(e *colly.HTMLElement, db *gorm.DB, category string){
	var pcCase dto.Case

	setBaseAttrs(e, &pcCase.BaseAttrs, category)	

	e.ForEach("div.spec__group ul.spec__list li.spec",func (_ int, el *colly.HTMLElement) {	 
		category := strings.TrimSpace(el.ChildText("span.spec__name"))

		switch category {
		case "Form factor": pcCase.MotherboardFormFactor= strings.TrimSpace(el.ChildText("span.spec__value"))
		}
	})

	if err := rawsql.InsertCase(db, &pcCase); err != nil {
		fmt.Println("Error inserting Case:", err)
		return
	}

	data,_:= json.MarshalIndent(pcCase,""," ")	
	fmt.Println(string(data))
}

func gpuHandler(e *colly.HTMLElement, db *gorm.DB, category string){
	var gpu dto.Gpu

	setBaseAttrs(e, &gpu.BaseAttrs, category)	

	e.ForEach("div.spec__group ul.spec__list li.spec",func (_ int, el *colly.HTMLElement) {	 
		category := strings.TrimSpace(el.ChildText("span.spec__name"))

		switch category {
		case "Frecventa Max procesor (MHz)": gpu.GpuFrequency = utils.CastInt(strings.TrimSpace(el.ChildText("span.spec__value")))
		case "Frecvență boost GPU": gpu.GpuFrequency = utils.CastInt(strings.TrimSpace(el.ChildText("span.spec__value")))
		case "Procesor video": gpu.Chipset = strings.TrimSpace(el.ChildText("span.spec__value"))
		case "Model placa video": gpu.Chipset = strings.TrimSpace(el.ChildText("span.spec__value"))
		case "Frecventa memorie {MHz}": gpu.VramFrequency = utils.CastInt(strings.TrimSpace(el.ChildText("span.spec__value")))
		case "Frecvență memorie": gpu.VramFrequency = utils.CastInt(strings.TrimSpace(el.ChildText("span.spec__value")))
		case "Producător GPU": gpu.BaseAttrs.Brand = strings.TrimSpace(el.ChildText("span.spec__value"))
		case "Producător Chipset": gpu.BaseAttrs.Brand = strings.TrimSpace(el.ChildText("span.spec__value"))
		case "Memorie | VGA": gpu.Vram = utils.CastInt(strings.TrimSpace(el.ChildText("span.spec__value")))
		case "Capacitate memorie": gpu.Vram = utils.CastInt(strings.TrimSpace(el.ChildText("span.spec__value")))
		}
	})

	if err := rawsql.InsertGpu(db, &gpu); err != nil {
		fmt.Println("Error inserting GPU:", err)
		return
	}

	data,_:= json.MarshalIndent(gpu,""," ")	
	fmt.Println(string(data))
}

func motherboardHandler(e *colly.HTMLElement, db *gorm.DB, category string){
	var motherboard dto.Motherboard

	setBaseAttrs(e, &motherboard.BaseAttrs, category)	

	e.ForEach("div.spec__group ul.spec__list li.spec",func (_ int, el *colly.HTMLElement) {	 
		category := strings.TrimSpace(el.ChildText("span.spec__name"))

		switch category {
		case "Model Chipset": motherboard.Chipset = strings.TrimSpace(el.ChildText("span.spec__value"))
		case "Cipset MB": motherboard.Chipset = strings.TrimSpace(el.ChildText("span.spec__value"))
		case "Chipset": motherboard.Chipset = strings.TrimSpace(el.ChildText("span.spec__value"))
		case "Form factor": motherboard.FormFactor = strings.TrimSpace(el.ChildText("span.spec__value"))
		case "Form-Factor": motherboard.FormFactor = strings.TrimSpace(el.ChildText("span.spec__value"))
		case "Format": motherboard.FormFactor = strings.TrimSpace(el.ChildText("span.spec__value"))
		case "Tip memorie RAM": motherboard.FormFactorRam = strings.TrimSpace(el.ChildText("span.spec__value"))
		case "Tip memorie": motherboard.FormFactorRam = strings.TrimSpace(el.ChildText("span.spec__value"))
		case "Suportul memoriei": motherboard.FormFactorRam = strings.TrimSpace(el.ChildText("span.spec__value"))
		case "Memorie maxima (GB)": motherboard.RamSupport = strings.TrimSpace(el.ChildText("span.spec__value"))
		case "Volumul maxim al memoriei operative": motherboard.RamSupport = strings.TrimSpace(el.ChildText("span.spec__value"))
		case "Capacitate maximă memorie suportată": motherboard.RamSupport = strings.TrimSpace(el.ChildText("span.spec__value"))
		case "CPU Socket": motherboard.Socket = strings.TrimSpace(el.ChildText("span.spec__value"))
		case "Socket procesor": motherboard.Socket = strings.TrimSpace(el.ChildText("span.spec__value"))
		case "Tip soket": motherboard.Socket = strings.TrimSpace(el.ChildText("span.spec__value"))
		}
	})

	if err := rawsql.InsertMotherboard(db, &motherboard); err != nil {
		fmt.Println("Error inserting Motherboard:", err)
		return
	}

	data,_:= json.MarshalIndent(motherboard,""," ")	
	fmt.Println(string(data))
}

func ramHandler(e *colly.HTMLElement, db *gorm.DB, category string){
	var ram dto.Ram

	setBaseAttrs(e, &ram.BaseAttrs, category)	

	//Reggex to extract configuration from name
	re := regexp.MustCompile(`Kit\s+of\s+(\d+)`)
	m := re.FindStringSubmatch(strings.TrimSpace(e.ChildText("div.product_container_wrap.container.p-lg-50 h1.section-title.mb-15")))

	if len(m) > 1 {
		ram.Configuration = m[1]
	} else {
		ram.Configuration = "1"
	}

	e.ForEach("div.spec__group ul.spec__list li.spec",func (_ int, el *colly.HTMLElement) {	 
		category := strings.TrimSpace(el.ChildText("span.spec__name"))

		switch category {
		case "Frecventa memorie RAM": ram.Speed = utils.CastInt(strings.TrimSpace(el.ChildText("span.spec__value")))
		case "Capacitate memorie RAM": ram.Capacity = utils.CastInt(strings.TrimSpace(el.ChildText("span.spec__value")))
		case "Compatibilitate RAM": ram.Compatibility = strings.TrimSpace(el.ChildText("span.spec__value"))
		case "Tip memorie RAM": ram.Type= strings.TrimSpace(el.ChildText("span.spec__value"))
		}
	})

	if err := rawsql.InsertRam(db, &ram); err != nil {
		fmt.Println("Error inserting RAM:", err)
		return
	}

	data,_:= json.MarshalIndent(ram,""," ")	
	fmt.Println(string(data))
}

func ssdHandler(e *colly.HTMLElement, db *gorm.DB, category string){
	var ssd dto.Ssd

	setBaseAttrs(e, &ssd.BaseAttrs, category)	

	e.ForEach("div.spec__group ul.spec__list li.spec",func (_ int, el *colly.HTMLElement) {	 
		category := strings.TrimSpace(el.ChildText("span.spec__name"))

		switch category {
		case "Interfata unitate stocare": ssd.FormFactor = strings.TrimSpace(el.ChildText("span.spec__value"))
		case "Capacitate stocare (GB)": ssd.Capacity = utils.CastInt(strings.TrimSpace(el.ChildText("span.spec__value")))
		case "Viteza de citire (MB/s)": ssd.ReadingSpeed = utils.CastInt(strings.TrimSpace(el.ChildText("span.spec__value")))
		case "Viteza de scriere (MB/s)": ssd.WritingSpeed = utils.CastInt(strings.TrimSpace(el.ChildText("span.spec__value")))
		}
	})

	if err := rawsql.InsertSsd(db, &ssd); err != nil {
		fmt.Println("Error inserting SSD:", err)
		return
	}

	data,_:= json.MarshalIndent(ssd,""," ")	
	fmt.Println(string(data))
}

func hddHandler(e *colly.HTMLElement, db *gorm.DB, category string){
	var hdd dto.Hdd

	setBaseAttrs(e, &hdd.BaseAttrs, category)	

	e.ForEach("div.spec__group ul.spec__list li.spec",func (_ int, el *colly.HTMLElement) {	 
		category := strings.TrimSpace(el.ChildText("span.spec__name"))

		switch category {
		case "Interfata unitate stocare": hdd.FormFactor = strings.TrimSpace(el.ChildText("span.spec__value"))
		case "Capacitate stocare (GB)": hdd.Capacity = utils.CastInt(strings.TrimSpace(el.ChildText("span.spec__value")))
		case "Viteza de citire (MB/s)": hdd.RotationSpeed = utils.CastInt(strings.TrimSpace(el.ChildText("span.spec__value")))
		}
	})

	if err := rawsql.InsertHdd(db, &hdd); err != nil {
		fmt.Println("Error inserting HDD:", err)
		return
	}

	data,_:= json.MarshalIndent(hdd,""," ")	
	fmt.Println(string(data))
}

func cpuHandler(e *colly.HTMLElement, db *gorm.DB, category string){
	var cpu dto.Cpu

	setBaseAttrs(e, &cpu.BaseAttrs, category)	

	e.ForEach("div.spec__group ul.spec__list li.spec",func (_ int, el *colly.HTMLElement) {	 
		category := strings.TrimSpace(el.ChildText("span.spec__name"))

		switch category {
		case "Producator": cpu.BaseAttrs.Brand = strings.TrimSpace(el.ChildText("span.spec__value"))
		case "Frecvență bază": cpu.BaseClock = utils.CastFloat64(strings.TrimSpace(el.ChildText("span.spec__value")))
		case "Frecvență turbo maximă": cpu.BoostClock = utils.CastFloat64(strings.TrimSpace(el.ChildText("span.spec__value")))
		case "TDP": cpu.Tdp = utils.CastInt(strings.TrimSpace(el.ChildText("span.spec__value")))
		case "Numar thread-uri": cpu.Threads = utils.CastInt(strings.TrimSpace(el.ChildText("span.spec__value")))
		case "Numar nuclee": cpu.Cores = utils.CastInt(strings.TrimSpace(el.ChildText("span.spec__value")))
		case "Socket": cpu.Socket = strings.TrimSpace(el.ChildText("span.spec__value"))
		}
	})

	if err := rawsql.InsertCpu(db, &cpu); err != nil {
		fmt.Println("Error inserting CPU:", err)
		return
	}

	data,_:= json.MarshalIndent(cpu,""," ")	
	fmt.Println(string(data))
}

func pcMiniHandler(e *colly.HTMLElement, db *gorm.DB, category string){
	var pc dto.PcMini

	setBaseAttrs(e, &pc.BaseAttrs, category)	

	e.ForEach("div.spec__group ul.spec__list li.spec",func (_ int, el *colly.HTMLElement) {	 
		category := strings.TrimSpace(el.ChildText("span.spec__name"))

		switch category {
		case "Model procesor": pc.Cpu = strings.TrimSpace(el.ChildText("span.spec__value"))
		case "Capacitate memorie RAM": pc.Ram = strings.TrimSpace(el.ChildText("span.spec__value"))
		case "Capacitate stocare (GB)": pc.Storage = strings.TrimSpace(el.ChildText("span.spec__value")) + " GB "
		case "Tip unitate stocare": pc.Storage += strings.TrimSpace(el.ChildText("span.spec__value"))
		case "Procesor placa video": pc.Gpu = strings.TrimSpace(el.ChildText("span.spec__value"))
		}
	})

	if err := rawsql.InsertPcMini(db, &pc); err != nil {
		fmt.Println("Error inserting PC Mini:", err)
		return
	}

	data,_:= json.MarshalIndent(pc,""," ")	
	fmt.Println(string(data))
}

func laptopHandler(e *colly.HTMLElement, db *gorm.DB, category string){
	var laptop dto.Laptop

	setBaseAttrs(e, &laptop.BaseAttrs, category)	

	e.ForEach("div.spec__group ul.spec__list li.spec",func (_ int, el *colly.HTMLElement) {	 
		category := strings.TrimSpace(el.ChildText("span.spec__name"))

		switch category {
		case "Model procesor": laptop.Cpu = strings.TrimSpace(el.ChildText("span.spec__value"))
		case "Dimensiune ecran (inch)": laptop.Diagonal = strings.TrimSpace(el.ChildText("span.spec__value")) 
		case "Capacitate memorie RAM": laptop.Ram = strings.TrimSpace(el.ChildText("span.spec__value"))
		case "Capacitate stocare (GB)": laptop.Storage = strings.TrimSpace(el.ChildText("span.spec__value")) + " GB "
		case "Tip unitate stocare": laptop.Storage += strings.TrimSpace(el.ChildText("span.spec__value"))
		case "Procesor placa video": laptop.Gpu = strings.TrimSpace(el.ChildText("span.spec__value"))
		case "Seria laptop": laptop.BaseAttrs.Brand = strings.TrimSpace(el.ChildText("span.spec__value"))
		}
	})

	if err := rawsql.InsertLaptop(db, &laptop); err != nil {
		fmt.Println("Error inserting Laptop:", err)
		return
	}

	data,_:= json.MarshalIndent(laptop,""," ")	
	fmt.Println(string(data))
}

func aioHandler(e *colly.HTMLElement, db *gorm.DB, category string){
	var aio dto.Aio

	setBaseAttrs(e, &aio.BaseAttrs, category)	

	e.ForEach("div.spec__group ul.spec__list li.spec",func (_ int, el *colly.HTMLElement) {	 
		category := strings.TrimSpace(el.ChildText("span.spec__name"))

		switch category {
		case "Model procesor": aio.Cpu = strings.TrimSpace(el.ChildText("span.spec__value"))
		case "Dimensiune ecran (inch)": aio.Diagonal = strings.TrimSpace(el.ChildText("span.spec__value")) 
		case "Capacitate memorie RAM": aio.Ram = strings.TrimSpace(el.ChildText("span.spec__value"))
		case "Capacitate stocare (GB)": aio.Storage = strings.TrimSpace(el.ChildText("span.spec__value")) + " GB "
		case "Tip unitate stocare": aio.Storage += strings.TrimSpace(el.ChildText("span.spec__value"))
		case "Procesor placa video": aio.Gpu = strings.TrimSpace(el.ChildText("span.spec__value"))
		case "Model all-in-one PC": aio.BaseAttrs.Brand = strings.TrimSpace(el.ChildText("span.spec__value"))
		}
	})

	if err := rawsql.InsertAio(db, &aio); err != nil {
		fmt.Println("Error inserting AIO:", err)
		return
	}

	data,_:= json.MarshalIndent(aio,""," ")	
	fmt.Println(string(data))
}

func pcHandler(e *colly.HTMLElement, db *gorm.DB, category string){
	var pc dto.Pc

	setBaseAttrs(e, &pc.BaseAttrs, category)	

	e.ForEach("div.spec__group ul.spec__list li.spec",func (_ int, el *colly.HTMLElement) {	 
		category := strings.TrimSpace(el.ChildText("span.spec__name"))

		switch category {
		case "Model procesor": pc.Cpu = strings.TrimSpace(el.ChildText("span.spec__value"))
		case "Capacitate memorie RAM": pc.Ram = strings.TrimSpace(el.ChildText("span.spec__value"))
		case "Capacitate stocare (GB)": pc.Storage = strings.TrimSpace(el.ChildText("span.spec__value")) + " GB "
		case "Tip unitate stocare": pc.Storage += strings.TrimSpace(el.ChildText("span.spec__value"))
		case "Procesor placa video": pc.Gpu = strings.TrimSpace(el.ChildText("span.spec__value"))
		case "Model placa de baza": pc.Motherboard = strings.TrimSpace(el.ChildText("span.spec__value"))
		case "Model carcasa": pc.Case = strings.TrimSpace(el.ChildText("span.spec__value"))
		case "Model sursa de alimentare": pc.Psu = strings.TrimSpace(el.ChildText("span.spec__value"))
		}
	})

	if err := rawsql.InsertPc(db, &pc); err != nil {
		fmt.Println("Error inserting PC:", err)
		return
	}
	
	data,_:= json.MarshalIndent(pc,""," ")	
	fmt.Println(string(data))
}