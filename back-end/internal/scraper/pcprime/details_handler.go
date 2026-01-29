package pcprime

import (
	"encoding/json"
	"fmt"
	"strconv"
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

var specRuEng map[string]string = map[string]string {
	"Жидкостное охлаждение":"Liquid cooling",
	"Воздушное охлаждение" : "Air cooling",
	"Вентилятор":"Air cooling",
	"ARGB": "ARGB",
	"RGB": "RGB",
	"Многоцветный":"Multicolor",
	"Для компьютера": "Pc",
	"Для ноутбука": "Laptop",
}

func setBaseAttrs(e *colly.HTMLElement, product *dto.BaseProduct){
	product.Name = e.ChildText("ol.breadcrumb li:last-child")
	product.ImageURL = e.ChildAttr("img","src")
	product.Price = utils.CastFloat64(e.ChildAttr("div.productPrice b","data-price"))
	product.Brand = e.ChildText("ol.breadcrumb li:nth-last-child(2)")
	product.Website_id = constants.WebIdMap["pcprime"]
	product.Url = e.Request.URL.String()
}

func (h *handler) ssdHandler(e *colly.HTMLElement){
	var ssd dto.Ssd

	setBaseAttrs(e, &ssd.BaseAttrs)

	e.ForEach(`div[id="fullDesc"] div.table_row`,func(_ int, el *colly.HTMLElement){
		spec := el.ChildText("div.table_cell:nth-child(1)")

		switch spec{
		case "Объем SSD": ssd.Capacity = utils.CastInt(el.ChildText("div.table_cell:nth-child(2)"))
		case "Скорость записи": ssd.WritingSpeed = utils.CastInt(el.ChildText("div.table_cell:nth-child(2)"))
		case "Скорость чтения": ssd.ReadingSpeed = utils.CastInt(el.ChildText("div.table_cell:nth-child(2)"))
		case "Формфактор": ssd.FormFactor = strings.TrimSpace(el.ChildText("div.table_cell:nth-child(2)"))
		}
	})

	if err := h.storage.InsertSsd(&ssd); err != nil {
		fmt.Println("Error inserting SSD:", err)
		return
	}

	data,_ := json.MarshalIndent(ssd,""," ")
	fmt.Println(string(data))
}

func (h *handler) hddHandler(e *colly.HTMLElement){
	var hdd dto.Hdd

	setBaseAttrs(e, &hdd.BaseAttrs)

	e.ForEach(`div[id="fullDesc"] div.table_row`,func(_ int, el *colly.HTMLElement){
		spec := el.ChildText("div.table_cell:nth-child(1)")

		switch spec {
		case "Объем HDD": hdd.Capacity = utils.CastInt(el.ChildText("div.table_cell:nth-child(2)"))
		case "Скорость вращения шпинделя": hdd.RotationSpeed= utils.CastInt(el.ChildText("div.table_cell:nth-child(2)"))
		case "Формфактор": hdd.FormFactor = strings.TrimSpace(el.ChildText("div.table_cell:nth-child(2)"))
		}
	})

	if err := h.storage.InsertHdd(&hdd); err != nil {
		fmt.Println("Error inserting HDD:", err)
		return
	}

	data,_ := json.MarshalIndent(hdd,""," ")
	fmt.Println(string(data))
}

func (h *handler) fanHandler(e *colly.HTMLElement){
	var fan dto.Fan

	setBaseAttrs(e, &fan.BaseAttrs)

	e.ForEach(`div[id="fullDesc"] div.table_row`,func(_ int, el *colly.HTMLElement){
		spec := el.ChildText("div.table_cell:nth-child(1)")

		switch spec {
		case "Частота вращения": fan.FanRPM = CastingIntFan(el.ChildText("div.table_cell:nth-child(2)"))
		case "Уровень шума": fan.Noise= utils.CastFloat64(el.ChildText("div.table_cell:nth-child(2)"))
		case "Размеры, мм": fan.Size = strings.TrimSpace(el.ChildText("div.table_cell:nth-child(2)"))
		}
	})

	if err := h.storage.InsertFan(&fan); err != nil {
		fmt.Println("Error inserting Fan:", err)
		return
	}

	data,_ := json.MarshalIndent(fan,""," ")
	fmt.Println(string(data))
}

func (h *handler) aioHandler(e *colly.HTMLElement){
	var aio dto.Aio

	setBaseAttrs(e, &aio.BaseAttrs)

	e.ForEach(`div[id="fullDesc"] div.table_row`,func(_ int, el *colly.HTMLElement){
		spec := el.ChildText("div.table_cell:nth-child(1)")

		switch spec {
		case "Процессор": aio.Cpu = strings.TrimSpace(strings.Replace(el.ChildText("div.table_cell:nth-child(2)"),"ГГц","GHz",1))
		case "Модель видеокарты": aio.Gpu = strings.TrimSpace(el.ChildText("div.table_cell:nth-child(2)"))
		case "Объем оперативной памяти": aio.Ram = strings.TrimSpace(strings.Replace(el.ChildText("div.table_cell:nth-child(2)"),"ГБ","GB",1))
		case "Объем SSD": aio.Storage=strings.TrimSpace(strings.Replace(el.ChildText("div.table_cell:nth-child(2)"),"ГБ","GB",1)) 
		case "Экран": aio.Diagonal = strings.TrimSpace(strings.Split(el.ChildText("div.table_cell:nth-child(2)"),"\"")[0])
		}
	})

	if err := h.storage.InsertAio(&aio); err != nil {
		fmt.Println("Error inserting AIO:", err)
		return
	}

	data,_ := json.MarshalIndent(aio,""," ")
	fmt.Println(string(data))
}

func (h *handler) pcMiniHandler(e *colly.HTMLElement){
	var pc dto.Pc
	
	setBaseAttrs(e, &pc.BaseAttrs)

	e.ForEach(`div[id="fullDesc"] div.table_row`,func(_ int, el *colly.HTMLElement){
		spec := el.ChildText("div.table_cell:nth-child(1)")

		switch spec{
		case "Процессор": pc.Cpu = strings.TrimSpace(el.ChildText("div.table_cell:nth-child(2)"))
		case "Модель видеокарты": pc.Gpu = strings.TrimSpace(el.ChildText("div.table_cell:nth-child(2) "))
		case "Оперативная память": pc.Ram = strings.TrimSpace(strings.Replace(el.ChildText("div.table_cell:nth-child(2)"),"ГБ","GB",1))
		case "Объем SSD": pc.Storage= strings.TrimSpace(strings.Replace(el.ChildText("div.table_cell:nth-child(2)"),"ГБ","GB",1))
		}
	})

	if err := h.storage.InsertPc(&pc); err != nil {
		fmt.Println("Error inserting PC:", err)
		return
	}

	data,_ := json.MarshalIndent(pc,""," ")
	fmt.Println(string(data))
}

func (h *handler) pcHandler(e *colly.HTMLElement){
	var pc dto.Pc
	
	setBaseAttrs(e, &pc.BaseAttrs)

	e.ForEach(`div[id="fullDesc"] div.table_row`,func(_ int, el *colly.HTMLElement){
		spec := el.ChildText("div.table_cell:nth-child(1)")

		switch spec {
		case "Название процессора и частота": pc.Cpu = strings.TrimSpace(strings.Replace(el.ChildText("div.table_cell:nth-child(2)"),"ГГц","GHz",1))
		case "Модель видеокарты": pc.Gpu = strings.TrimSpace(el.ChildText("div.table_cell:nth-child(2) "))
		case "Объем оперативной памяти": pc.Ram = strings.TrimSpace(strings.Replace(el.ChildText("div.table_cell:nth-child(2)"),"ГБ","GB",1))
		case "Объем SSD": pc.Storage = strings.TrimSpace(strings.Replace(el.ChildText("div.table_cell:nth-child(2)"),"ГБ","GB",1)) 
		//Different Case
		case "Процессор": pc.Cpu = strings.TrimSpace(el.ChildText("div.table_cell:nth-child(2) a"))
		case "Видеокарта": pc.Gpu = strings.TrimSpace(el.ChildText("div.table_cell:nth-child(2) a")) 
		case "Оперативная память": pc.Ram = strings.TrimSpace(el.ChildText("div.table_cell:nth-child(2) a"))
		case "SSD накопитель": pc.Storage = strings.TrimSpace(el.ChildText("div.table_cell:nth-child(2) a"))
		case "Материнская плата": pc.Motherboard = strings.TrimSpace(el.ChildText("div.table_cell:nth-child(2) a")) 
		case "Блок питания": pc.Psu = strings.TrimSpace(el.ChildText("div.table_cell:nth-child(2) a"))
		case "Корпус": pc.Case = strings.TrimSpace(el.ChildText("div.table_cell:nth-child(2) a"))
		}
	})

	if err := h.storage.InsertPc(&pc); err != nil {
		fmt.Println("Error inserting PC:", err)
		return
	}

	data,_ := json.MarshalIndent(pc,""," ")
	fmt.Println(string(data))
}

func (h *handler) caseHandler(e *colly.HTMLElement){
	var pcCase dto.Case
	
	setBaseAttrs(e, &pcCase.BaseAttrs)

	e.ForEach(`div[id="fullDesc"] div.table_row`,func(_ int, el *colly.HTMLElement){
		spec := el.ChildText("div.table_cell:nth-child(1)")

		switch spec {
		case "Вид": pcCase.Format = strings.TrimSpace(el.ChildText("div.table_cell:nth-child(2)"))
		case "Форм-фактор": pcCase.MotherboardFormFactor = strings.TrimSpace(el.ChildText("div.table_cell:nth-child(2)"))
		}
	})

	if err := h.storage.InsertCase(&pcCase); err != nil {
		fmt.Println("Error inserting Case:", err)
		return
	}

	data,_ := json.MarshalIndent(pcCase,""," ")
	fmt.Println(string(data))
}

func (h *handler) psuHandler(e *colly.HTMLElement){
	var psu dto.Psu

	setBaseAttrs(e, &psu.BaseAttrs)
	
	e.ForEach(`div[id="fullDesc"] div.table_row`,func(_ int, el *colly.HTMLElement){
		spec := el.ChildText("div.table_cell:nth-child(1)")

		switch spec {
		case "Мощность БП": psu.Power = utils.CastInt(el.ChildText("div.table_cell:nth-child(2)"))
		case "Сертификат БП": psu.Efficiency = strings.TrimSpace(strings.ReplaceAll(el.ChildText("div.table_cell:nth-child(2)"), "+", " PLUS"))
		case "Форм-фактор": psu.FormFactor = strings.TrimSpace(el.ChildText("div.table_cell:nth-child(2)"))
		}
	})

	if psu.Efficiency == "" {
		psu.Efficiency = "Standard"
	}

	if err := h.storage.InsertPsu(&psu); err != nil {
		fmt.Println("Error inserting PSU:", err)
		return
	}

	data,_ := json.MarshalIndent(psu,""," ")
	fmt.Println(string(data))
}

func (h *handler) gpuHandler(e *colly.HTMLElement){
	var gpu dto.Gpu

	setBaseAttrs(e, &gpu.BaseAttrs)

	e.ForEach(`div[id="fullDesc"] div.table_row`,func(_ int, el *colly.HTMLElement){
		spec := el.ChildText("div.table_cell:nth-child(1)")

		switch spec {
		case "Объем видеопамяти": gpu.Vram = utils.CastInt(el.ChildText("div.table_cell:nth-child(2)"))
		case "Частота памяти": gpu.VramFrequency = utils.CastInt(el.ChildText("div.table_cell:nth-child(2)"))
		case "Частота ядра": gpu.GpuFrequency = utils.CastInt(el.ChildText("div.table_cell:nth-child(2)"))
		case "Графический чип": gpu.Chipset = strings.TrimSpace(el.ChildText("div.table_cell:nth-child(2)"))
		}
	})

	if err := h.storage.InsertGpu(&gpu); err != nil {
		fmt.Println("Error inserting GPU:", err)
		return
	}

	data,_ := json.MarshalIndent(gpu,""," ")
	fmt.Println(string(data))
}

func (h *handler) ramHandler(e *colly.HTMLElement){
	var ram dto.Ram

	setBaseAttrs(e, &ram.BaseAttrs)
	
	e.ForEach(`div[id="fullDesc"] div.table_row`,func(_ int, el *colly.HTMLElement){
		spec := el.ChildText("div.table_cell:nth-child(1)")

		switch spec {
		case "Объем оперативной памяти": ram.Capacity = utils.CastInt(el.ChildText("div.table_cell:nth-child(2)"))
		case "Частота памяти": ram.Speed = utils.CastInt(el.ChildText("div.table_cell:nth-child(2)"))
		case "Тип оперативной памяти": ram.Type = strings.TrimSpace(el.ChildText("div.table_cell:nth-child(2)"))
		case "Вид": ram.Compatibility = specRuEng[el.ChildText("div.table_cell:nth-child(2)")]
		case "Количество планок":
			configuration := strings.TrimSpace(el.ChildText("div.table_cell:nth-child(2)"))
			ram.Configuration = configuration + " x " + strconv.Itoa(ram.Capacity / utils.CastInt(configuration) ) + " GB"
		}
	})

	if err := h.storage.InsertRam(&ram); err != nil {
		fmt.Println("Error inserting RAM:", err)
		return
	}

	data,_ := json.MarshalIndent(ram,""," ")
	fmt.Println(string(data))
}

func (h *handler) motherBoardHandler(e *colly.HTMLElement){
	var motherboard dto.Motherboard

	setBaseAttrs(e, &motherboard.BaseAttrs)

	e.ForEach(`div[id="fullDesc"] div.table_row`,func(_ int, el *colly.HTMLElement){
		spec := el.ChildText("div.table_cell:nth-child(1)")

		switch spec {
		case "Чипсет":
			motherboard.Chipset= strings.TrimSpace(el.ChildText("div.table_cell:nth-child(2)"))
		case "Тип сокета":
			motherboard.Socket = strings.TrimSpace(el.ChildText("div.table_cell:nth-child(2)"))
		case "Формфактор":
			motherboard.FormFactor = strings.TrimSpace(el.ChildText("div.table_cell:nth-child(2)"))
		case "Максимальный объем памяти":
			motherboard.RamSupport= strings.TrimSpace(strings.Replace(el.ChildText("div.table_cell:nth-child(2)"),"ГБ","GB",1))
		case "Поддержка памяти":
			motherboard.FormFactorRam= strings.TrimSpace(el.ChildText("div.table_cell:nth-child(2)"))
		}
	})

	if err := h.storage.InsertMotherboard(&motherboard); err != nil {
		fmt.Println("Error inserting Motherboard:", err)
		return
	}

	data,_ := json.MarshalIndent(motherboard,""," ")
	fmt.Println(string(data))
}

func (h *handler) coolerHandler(e *colly.HTMLElement){
	var cooler dto.Cooler

	setBaseAttrs(e, &cooler.BaseAttrs)

	e.ForEach(`div[id="fullDesc"] div.table_row`,func(_ int, el *colly.HTMLElement){
		spec := el.ChildText("div.table_cell:nth-child(1)")

		switch spec{
		case "Тип": cooler.Type= specRuEng[el.ChildText("div.table_cell:nth-child(2)")]
		case "Частота вращения": cooler.FanRPM = CastingIntFan(el.ChildText("div.table_cell:nth-child(2)"))
		case "Уровень шума": cooler.Noise = utils.CastFloat64(el.ChildText("div.table_cell:nth-child(2)"))
		case "Размеры, мм": cooler.Size = strings.TrimSpace(el.ChildText("div.table_cell:nth-child(2)"))
		case "Тип сокета процессора":	
			s := el.ChildText("div.table_cell:nth-child(2)")
			cooler.Compatibility = func(parts []string) []string{
				for i := range parts {parts[i] = strings.TrimSpace(parts[i])} 
				return parts 
			}(strings.Split(s,"/"))
		}
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

	data,_ := json.MarshalIndent(cooler ,""," ")
	fmt.Println(string(data))
}

func (h *handler) cpuHandler(e *colly.HTMLElement){
	var cpu dto.Cpu

	setBaseAttrs(e, &cpu.BaseAttrs)

	e.ForEach(`div[id="fullDesc"] div.table_row`,func(_ int, el *colly.HTMLElement){
		spec := el.ChildText("div.table_cell:nth-child(1)")

		switch spec {
		case "Количество ядер": cpu.Cores= utils.CastInt(el.ChildText("div.table_cell:nth-child(2)"))
		case "Количество потоков": cpu.Threads = utils.CastInt(el.ChildText("div.table_cell:nth-child(2)"))
		case "Базовая частота": cpu.BaseClock = utils.CastFloat64(el.ChildText("div.table_cell:nth-child(2)"))
		case "Максимальная частота": cpu.BoostClock = utils.CastFloat64(el.ChildText("div.table_cell:nth-child(2)"))
		case "Тип сокета": cpu.Socket = strings.TrimSpace(el.ChildText("div.table_cell:nth-child(2)")) 
		case "Мощность TDP": cpu.Tdp = utils.CastInt(el.ChildText("div.table_cell:nth-child(2)"))
		}
	})

	if err := h.storage.InsertCpu(&cpu); err != nil {
		fmt.Println("Error inserting CPU:", err)
		return
	}

	data,_ := json.MarshalIndent(cpu,""," ")
	fmt.Println(string(data))
}

func (h *handler) laptopHandler(e *colly.HTMLElement){
	var laptop dto.Laptop

	setBaseAttrs(e, &laptop.BaseAttrs)
	
	e.ForEach(`div[id="fullDesc"] div.table_row`,func(_ int, el *colly.HTMLElement){
		spec := el.ChildText("div.table_cell:nth-child(1)")

		switch spec{
		case "Процессор": laptop.Cpu = strings.TrimSpace(strings.Replace(el.ChildText("div.table_cell:nth-child(2)"),"ГГц","GHz",1))
		case "Видеокарта": laptop.Gpu = strings.TrimSpace(strings.Replace(el.ChildText("div.table_cell:nth-child(2)"),"Интегрированная","Integrated",1))
		case "Оперативная память": laptop.Ram = strings.TrimSpace(strings.Replace(el.ChildText("div.table_cell:nth-child(2)"),"ГБ","GB",1))
		case "Накопитель": laptop.Storage = strings.TrimSpace(strings.Replace(el.ChildText("div.table_cell:nth-child(2)"),"ГБ","GB",1))
		case "Диагональ экрана": laptop.Diagonal = strings.TrimSpace(strings.Split(el.ChildText("div.table_cell:nth-child(2)"),"\"")[0])
		case "Аккумулятор": laptop.Battery = castingFloat64Laptop(el.ChildText("div.table_cell:nth-child(2)"))
		}
	})

	if err := h.storage.InsertLaptop(&laptop); err != nil {
		fmt.Println("Error inserting Laptop:", err)
		return
	}

	data,_ := json.MarshalIndent(laptop,""," ")
	fmt.Println(string(data))
}

