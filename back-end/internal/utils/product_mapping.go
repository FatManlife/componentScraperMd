package utils

import (
	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"github.com/FatManlife/component-finder/back-end/internal/models/orm"
)

func CoolerMapping(product orm.Product) dto.CoolerResponse {
	var compatibilities []string
	
	for _, v := range product.Cooler.Cpus {
		compatibilities = append(compatibilities, v.Compatibility.Cpu)
	}

	return dto.CoolerResponse{
		Product: ProductMapping(product),
		Type: product.Cooler.Type,
		FanRPM: product.Cooler.FanRPM,
		Noise: product.Cooler.Noise,
		Size: product.Cooler.Size,
		Compatibility: compatibilities,
	}
}

func ProductMapping(product orm.Product) dto.ProductResponse{
	return dto.ProductResponse{
			ID: product.ID,
			Name: product.Name,
			ImageURL: product.ImageURL,
			Brand: product.Brand,
			Price: product.Price,
			Url: product.URL,
			Category: product.Category,
			Website_id: product.WebsiteID,
		}	
}

func HddMapping(product orm.Product) dto.HddResponse {
	return dto.HddResponse{
		Product: ProductMapping(product),
		Capacity: product.Hdd.Capacity,
		RotationSpeed: product.Hdd.RotationSpeed,
		FormFactor: product.Hdd.FormFactor,
	}
}	

func CpuMapping(product orm.Product) dto.CpuResponse{
	return dto.CpuResponse{
		Product: ProductMapping(product),
		Cores: product.Cpu.Cores,
		Threads: product.Cpu.Threads,
		BaseClock: product.Cpu.BaseClock,
		BoostClock: product.Cpu.BoostClock,
		Tdp: product.Cpu.Tdp,
		Socket: product.Cpu.Socket,
	}
}

func AioMapping(product orm.Product) dto.AioResponse{
	return dto.AioResponse{
		Product: ProductMapping(product),
		Diagonal: product.Aio.Diagonal,
		Gpu: product.Aio.Gpu,
		Cpu: product.Aio.Cpu,
		Ram: product.Aio.Ram,
		Storage: product.Aio.Storage,
	}
}

func CaseMapping(product orm.Product) dto.CaseResponse{
	return dto.CaseResponse{
		Product: ProductMapping(product),
		Format: product.Case.Format,
		MotherboardFormFactor: product.Case.MotherboardFormFactor,
	}
}

func GpuMapping(product orm.Product) dto.GpuResponse {
	return dto.GpuResponse{
		Product:      ProductMapping(product),
		Chipset:      product.Gpu.Chipset,
		Vram:         product.Gpu.Vram,
		GpuFrequency: product.Gpu.GpuFrequency,
		VramFrequency: product.Gpu.VramFrequency,
	}
}

func FanMapping(product orm.Product) dto.FanResponse{
	return dto.FanResponse{
		Product: ProductMapping(product),
		FanRPM: product.Fan.FanRPM,
		Size: product.Fan.Size,
		Noise: product.Fan.Noise,
	}
}

func MbMapping(product orm.Product) dto.MotherboardResponse{
	return dto.MotherboardResponse{
		Product: ProductMapping(product),
		Chipset: product.Motherboard.Chipset,
		Socket: product.Motherboard.Socket,
		FormFactor: product.Motherboard.FormFactor,
		RamSupport: product.Motherboard.RamSupport,
		FormFactorRam: product.Motherboard.FormFactorRam,
	}
}

func LaptopMapping(product orm.Product) dto.LaptopResponse {
	return dto.LaptopResponse{
		Product: ProductMapping(product),
		Cpu: product.Laptop.Cpu,
		Gpu: product.Laptop.Gpu,
		Ram: product.Laptop.Ram,
		Storage: product.Laptop.Storage,
		Diagonal: product.Laptop.Diagonal,
		Battery: product.Laptop.Battery,
	}
}

func PcMiniMapping(product orm.Product) dto.PcMiniResponse{
	return dto.PcMiniResponse{
		Product: ProductMapping(product),
		Cpu: product.PcMini.Cpu,
		Gpu: product.PcMini.Gpu,
		Ram: product.PcMini.Ram,
		Storage: product.PcMini.Storage,
	}
}

func PcMapping(product orm.Product) dto.PcResponse{
	return dto.PcResponse{
		Product: ProductMapping(product),
		Case: product.Pc.PcCase,
		Cpu: product.Pc.Cpu,
		Gpu: product.Pc.Gpu,
		Motherboard: product.Pc.Motherboard,
		Psu: product.Pc.Psu,
		Ram: product.Pc.Ram,
		Storage: product.Pc.Storage,
	}
}

func PsuMapping(product orm.Product) dto.PsuResponse {
	return dto.PsuResponse{
		Product:    ProductMapping(product),
		FormFactor: product.Psu.FormFactor,
		Efficiency: product.Psu.Efficiency,
		Power:      product.Psu.Power,
	}
}

func SsdMapping(product orm.Product) dto.SsdResponse {
	return dto.SsdResponse{
		Product: ProductMapping(product),
		Capacity: product.Ssd.Capacity,
		ReadingSpeed: product.Ssd.ReadingSpeed,
		WritingSpeed: product.Ssd.WritingSpeed,
		FormFactor: product.Ssd.FormFactor,
	}
}

func RamMapping(product orm.Product) dto.RamResponse {
	return dto.RamResponse{
		Product: ProductMapping(product),
		Capacity: product.Ram.Capacity,
		Speed: product.Ram.Speed,
		Type: product.Ram.Type,
		Compatibility: product.Ram.Compatibility,
		Configuration: product.Ram.Configuration,
	}
}
