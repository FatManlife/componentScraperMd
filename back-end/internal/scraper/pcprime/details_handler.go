package pcprime

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/FatManlife/component-finder/back-end/internal/models"
	"github.com/FatManlife/component-finder/back-end/internal/utils"
	"github.com/gocolly/colly"
)

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


func setBaseAttrs(e *colly.HTMLElement, product models.BaseProduct){
	product.Name = e.ChildText("ol.breadcrumb li:last-child")
	product.ImageURL = e.ChildAttr("img","src")
	product.Price = utils.CastFloat64(e.ChildAttr("div.productPrice b","data-price"))
	product.Brand = e.ChildText("ol.breadcrumb li:nth-last-child(2)")
	product.Website_id = 2 
	product.Url = e.Request.URL.String()
}

func aioHandler(e *colly.HTMLElement){
	var aio models.Aio

	setBaseAttrs(e, aio.BaseAttrs)

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

func pcMiniHandler(e *colly.HTMLElement){
	var pc models.Pc
	
	setBaseAttrs(e, pc.BaseAttrs)

	e.ForEach(`div[id="fullDesc"] div.table_row`,func(_ int, el *colly.HTMLElement){
		category := el.ChildText("div.table_cell:nth-child(1)")

		switch category{
		case "Процессор":
			pc.Cpu = el.ChildText("div.table_cell:nth-child(2)")
		case "Модель видеокарты":
			pc.Gpu = el.ChildText("div.table_cell:nth-child(2) ")
		case "Оперативная память":
			pc.Ram = strings.Replace(el.ChildText("div.table_cell:nth-child(2)"),"ГБ","GB",1)
		case "Объем SSD":
			pc.Storage= strings.Replace(el.ChildText("div.table_cell:nth-child(2)"),"ГБ","GB",1)
		}
	})

	data,_ := json.MarshalIndent(pc,""," ")
	fmt.Println(string(data))
}

func pcHandler(e *colly.HTMLElement){
	var pc models.Pc
	
	setBaseAttrs(e, pc.BaseAttrs)

	e.ForEach(`div[id="fullDesc"] div.table_row`,func(_ int, el *colly.HTMLElement){
		category := el.ChildText("div.table_cell:nth-child(1)")

		switch category{
		case "Название процессора и частота":
			pc.Cpu = strings.Replace(el.ChildText("div.table_cell:nth-child(2)"),"ГГц","GHz",1)
		case "Модель видеокарты":
			pc.Gpu = el.ChildText("div.table_cell:nth-child(2) ")
		case "Объем оперативной памяти":
			pc.Ram = strings.Replace(el.ChildText("div.table_cell:nth-child(2)"),"ГБ","GB",1)
		case "Объем SSD":
			pc.Storage= strings.Replace(el.ChildText("div.table_cell:nth-child(2)"),"ГБ","GB",1)
		}
	})

	data,_ := json.MarshalIndent(pc,""," ")
	fmt.Println(string(data))
}

func caseHandler(e *colly.HTMLElement){
	var pcCase models.Case
	
	setBaseAttrs(e, pcCase.BaseAttrs)

	e.ForEach(`div[id="fullDesc"] div.table_row`,func(_ int, el *colly.HTMLElement){
		category := el.ChildText("div.table_cell:nth-child(1)")

		switch category{
		case "Вид":
			pcCase.Format = el.ChildText("div.table_cell:nth-child(2)")
		case "Форм-фактор":
			pcCase.MotherboardFormFactor= el.ChildText("div.table_cell:nth-child(2)")
		}
	})

	data,_ := json.MarshalIndent(pcCase,""," ")
	fmt.Println(string(data))
}

func pcPrimeHandler(e *colly.HTMLElement){
	var pc models.Pc
	
	setBaseAttrs(e, pc.BaseAttrs)

	e.ForEach(`div[id="fullDesc"] div.table_row`,func(_ int, el *colly.HTMLElement){
		category := el.ChildText("div.table_cell:nth-child(1)")

		switch category{
		case "Процессор":
			pc.Cpu= el.ChildText("div.table_cell:nth-child(2) a")
		case "Видеокарта":
			pc.Gpu= el.ChildText("div.table_cell:nth-child(2) a")
		case "Оперативная память":
			pc.Ram= el.ChildText("div.table_cell:nth-child(2) a")
		case "SSD накопитель":
			pc.Storage= el.ChildText("div.table_cell:nth-child(2) a")
		case "Материнская плата":
			pc.Motherboard= el.ChildText("div.table_cell:nth-child(2) a")
		case "Блок питания":
			pc.Psu= el.ChildText("div.table_cell:nth-child(2) a")
		case "Корпус":
			pc.Case= el.ChildText("div.table_cell:nth-child(2) a")
		}
	})

	data,_ := json.MarshalIndent(pc,""," ")
	fmt.Println(string(data))
}

func psuHandler(e *colly.HTMLElement){
	var psu models.Psu

	setBaseAttrs(e, psu.BaseAttrs)
	
	e.ForEach(`div[id="fullDesc"] div.table_row`,func(_ int, el *colly.HTMLElement){
		category := el.ChildText("div.table_cell:nth-child(1)")

		switch category{
		case "Мощность БП":
			psu.Power= utils.CastInt(el.ChildText("div.table_cell:nth-child(2)"))
		case "Сертификат БП":
			psu.Efficiency= el.ChildText("div.table_cell:nth-child(2)")
		case "Форм-фактор":
			psu.FormFactor = el.ChildText("div.table_cell:nth-child(2)")
		}
	})

	data,_ := json.MarshalIndent(psu,""," ")
	fmt.Println(string(data))
}

func gpuHandler(e *colly.HTMLElement){
	var gpu models.Gpu

	setBaseAttrs(e, gpu.BaseAttrs)

	e.ForEach(`div[id="fullDesc"] div.table_row`,func(_ int, el *colly.HTMLElement){
		category := el.ChildText("div.table_cell:nth-child(1)")

		switch category{
		case "Объем видеопамяти":
			gpu.Vram = utils.CastInt(el.ChildText("div.table_cell:nth-child(2)"))
		case "Частота памяти":
			gpu.VramFrequency = utils.CastInt(el.ChildText("div.table_cell:nth-child(2)"))
		case "Частота ядра":
			gpu.GpuFrequency = utils.CastInt(el.ChildText("div.table_cell:nth-child(2)"))
		case "Графический чип":
			gpu.Chipset = el.ChildText("div.table_cell:nth-child(2)")
		}
	})

	data,_ := json.MarshalIndent(gpu,""," ")
	fmt.Println(string(data))
}

func ramHandler(e *colly.HTMLElement){
	var ram models.Ram

	setBaseAttrs(e, ram.BaseAttrs)
	
	e.ForEach(`div[id="fullDesc"] div.table_row`,func(_ int, el *colly.HTMLElement){
		category := el.ChildText("div.table_cell:nth-child(1)")

		switch category{
		case "Объем оперативной памяти":
			ram.Capacity = utils.CastInt(el.ChildText("div.table_cell:nth-child(2)"))
		case "Частота памяти":
			ram.Speed = utils.CastInt(el.ChildText("div.table_cell:nth-child(2)"))
		case "Тип оперативной памяти":
			ram.Type = el.ChildText("div.table_cell:nth-child(2)")
		case "Вид":
			ram.Compatibility = ruEng[el.ChildText("div.table_cell:nth-child(2)")]
		case "Количество планок":
			configuration := el.ChildText("div.table_cell:nth-child(2)")
			ram.Configuration = configuration + " x " + strconv.Itoa(ram.Capacity / utils.CastInt(configuration) ) + " GB"
		}
	})

	data,_ := json.MarshalIndent(ram,""," ")
	fmt.Println(string(data))
}

func motherBoardHandler(e *colly.HTMLElement){
	var motherboard models.Motherboard

	setBaseAttrs(e, motherboard.BaseAttrs)

	e.ForEach(`div[id="fullDesc"] div.table_row`,func(_ int, el *colly.HTMLElement){
		category := el.ChildText("div.table_cell:nth-child(1)")

		switch category{
		case "Чипсет":
			motherboard.Chipset= el.ChildText("div.table_cell:nth-child(2)")
		case "Тип сокета":
			motherboard.Socket = el.ChildText("div.table_cell:nth-child(2)")
		case "Формфактор":
			motherboard.FormFactor = el.ChildText("div.table_cell:nth-child(2)")
		case "Максимальный объем памяти":
			motherboard.RamSupport= strings.Replace(el.ChildText("div.table_cell:nth-child(2)"),"ГБ","GB",1)
		case "Поддержка памяти":
			motherboard.FormFactorRam= el.ChildText("div.table_cell:nth-child(2)")
		}
	})

	data,_ := json.MarshalIndent(motherboard,""," ")
	fmt.Println(string(data))
}

func coolerHandler(e *colly.HTMLElement){
	var cooler models.Cooler

	setBaseAttrs(e, cooler.BaseAttrs)

	e.ForEach(`div[id="fullDesc"] div.table_row`,func(_ int, el *colly.HTMLElement){
		category := el.ChildText("div.table_cell:nth-child(1)")

		switch category{
		case "Тип":
			cooler.Type= ruEng[el.ChildText("div.table_cell:nth-child(2)")]
		case "Цвет подсветки":
			cooler.Ilumination = ruEng[el.ChildText("div.table_cell:nth-child(2)")]
		case "Частота вращения":
			cooler.FanRPM = CastingCastIntFan(el.ChildText("div.table_cell:nth-child(2)"))
		case "Уровень шума":
			cooler.Noise = utils.CastFloat64(el.ChildText("div.table_cell:nth-child(2)"))
		case "Размеры, мм":
			cooler.Size = el.ChildText("div.table_cell:nth-child(2)")
		case "Тип сокета процессора":	
			s := el.ChildText("div.table_cell:nth-child(2)")
			cooler.Compatibility = func(parts []string) []string{
				for i := range parts {parts[i] = strings.TrimSpace(parts[i])} 
				return parts 
			}(strings.Split(s,"/"))
		}
	})

	data,_ := json.MarshalIndent(cooler ,""," ")
	fmt.Println(string(data))
}

func cpuHandler(e *colly.HTMLElement){
	var cpu models.Cpu

	setBaseAttrs(e, cpu.BaseAttrs)

	e.ForEach(`div[id="fullDesc"] div.table_row`,func(_ int, el *colly.HTMLElement){
		category := el.ChildText("div.table_cell:nth-child(1)")

		switch category{
		case "Количество ядер":
			cpu.Cores= utils.CastInt(el.ChildText("div.table_cell:nth-child(2)"))
		case "Количество потоков":
			cpu.Threads = utils.CastInt(el.ChildText("div.table_cell:nth-child(2)"))
		case "Базовая частота":
			cpu.BaseClock = utils.CastFloat64(el.ChildText("div.table_cell:nth-child(2)"))
		case "Максимальная частота":
			cpu.BoostClock = utils.CastFloat64(el.ChildText("div.table_cell:nth-child(2)"))
		case "Объем кэш памяти":
			cpu.Cache = el.ChildText("div.table_cell:nth-child(2)")
		case "Тип сокета":
			cpu.Socket = el.ChildText("div.table_cell:nth-child(2)")
		case "Мощность TDP":
			cpu.Tdp = utils.CastInt(el.ChildText("div.table_cell:nth-child(2)"))
		}
	})

	data,_ := json.MarshalIndent(cpu,""," ")
	fmt.Println(string(data))
}

func laptopHandler(e *colly.HTMLElement){
	var laptop models.Laptop

	setBaseAttrs(e, laptop.BaseAttrs)
	
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
			laptop.Battery = castingFloat64Laptop(el.ChildText("div.table_cell:nth-child(2)"))
		}
	})

	data,_ := json.MarshalIndent(laptop,""," ")
	fmt.Println(string(data))
}

