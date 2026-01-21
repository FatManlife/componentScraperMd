package rawsql

import (
	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"gorm.io/gorm"
)

func InsertProduct(db *gorm.DB, product *dto.BaseProduct) (int64, error) {
	var lastId int64

	err := db.Raw("INSERT INTO products (name, image_url, brand, price, url, website_id) VALUES (?,?,?,?,?,?) RETURNING id",
		product.Name,
		product.ImageURL,
		product.Brand,
		product.Price,
		product.Url,
		product.Website_id,
	).Scan(&lastId).Error 

	return lastId, err
}

func InsertFan(db *gorm.DB, fan *dto.Fan, productId int64) error {
	return db.Exec("INSERT INTO fans (product_id, fan_rpm, noise, size) VALUES (?,?,?,?)",
		productId,
		fan.FanRPM,
		fan.Noise,
		fan.Size,
	).Error
}

func InsertAio(db *gorm.DB, aio *dto.Aio, productId int64) error {
	return db.Exec("INSERT INTO aios (product_id, diagonal, cpu, ram, storage, gpu) VALUES (?,?,?,?,?,?)",
		productId,
		aio.Diagonal,
		aio.Cpu,
		aio.Ram,
		aio.Storage,
		aio.Gpu,
	).Error
}

func InsertCase(db *gorm.DB, caseProduct *dto.Case, productId int64) error {
	return db.Exec("INSERT INTO cases (product_id, format, motherboard_form_factor) VALUES (?,?,?,?,?)",
		productId,
		caseProduct.Format,
		caseProduct.MotherboardFormFactor,
	).Error
}

func InsertCooler(db *gorm.DB, cooler *dto.Cooler, productId int64) error {
	return db.Exec("INSERT INTO coolers (product_id, type, fan_rpm, noise, size) VALUES (?,?,?,?,?)",
		productId,
		cooler.Type,
		cooler.FanRPM,
		cooler.Noise,
		cooler.Size,
	).Error
}

func InsertCoolerCompatibility(db *gorm.DB, coolerId int64, compatibility []string) error {
	for _, comp := range compatibility {
		if err := db.Exec("INSERT INTO cooler_compatibility (cooler_id, compatibility) VALUES (?,?)",
			coolerId,
			comp,
		).Error; err != nil {
			return err
		}
	}
	return nil
}

func InsertCpu(db *gorm.DB, cpu *dto.Cpu, productId int64) error {
	return db.Exec("INSERT INTO cpus (product_id, cores, threads, base_clock, boost_clock, socket, tdp) VALUES (?,?,?,?,?,?,?)",
		productId,
		cpu.Cores,
		cpu.Threads,
		cpu.BaseClock,
		cpu.BoostClock,
		cpu.Socket,
		cpu.Tdp,
	).Error
}

func InsertGpu(db *gorm.DB, gpu *dto.Gpu, productId int64) error {
	return db.Exec("INSERT INTO gpus (product_id, chipset, vram, gpu_frequency, vram_frequency) VALUES (?,?,?,?,?,?)",
		productId,
		gpu.Chipset,
		gpu.Vram,
		gpu.GpuFrequency,
		gpu.VramFrequency,
	).Error
}

func InsertHdd(db *gorm.DB, hdd *dto.Hdd, productId int64) error {
	return db.Exec("INSERT INTO hdds (product_id, capacity, rotation_speed, form_factor) VALUES (?,?,?,?)",
		productId,
		hdd.Capacity,
		hdd.RotationSpeed,
		hdd.FormFactor,
	).Error
}

func InsertLaptop(db *gorm.DB, laptop *dto.Laptop, productId int64) error {
	return db.Exec("INSERT INTO laptops (product_id, cpu, gpu, ram, storage, diagonal, battery) VALUES (?,?,?,?,?,?,?)",
		productId,
		laptop.Cpu,
		laptop.Gpu,
		laptop.Ram,
		laptop.Storage,
		laptop.Diagonal,
		laptop.Battery,
	).Error
}

func InsertMotherboard(db *gorm.DB, motherboard *dto.Motherboard, productId int64) error {
	return db.Exec("INSERT INTO motherboards (product_id, chipset, socket, form_factor, ram_support, form_factor_ram) VALUES (?,?,?,?,?,?)",
		productId,
		motherboard.Chipset,
		motherboard.Socket,
		motherboard.FormFactor,
		motherboard.RamSupport,
		motherboard.FormFactorRam,
	).Error
}

func InsertPcMini(db *gorm.DB, pcmini *dto.PcMini, productId int64) error {
	return db.Exec("INSERT INTO pc_minis (product_id, cpu, gpu, ram, storage) VALUES (?,?,?,?,?)",
		productId,
		pcmini.Cpu,
		pcmini.Gpu,
		pcmini.Ram,
		pcmini.Storage,
	).Error
}

func InsertPc(db *gorm.DB, pc *dto.Pc, productId int64) error {
	return db.Exec("INSERT INTO pcs (product_id, cpu, gpu, ram, storage, motherboard, psu, case) VALUES (?,?,?,?,?,?,?,?)",
		productId,
		pc.Cpu,
		pc.Gpu,
		pc.Ram,
		pc.Storage,
		pc.Motherboard,
		pc.Psu,
		pc.Case,
	).Error
}

func InsertPsu(db *gorm.DB, psu *dto.Psu, productId int64) error {
	return db.Exec("INSERT INTO psus (product_id, power, efficiency, form_factor) VALUES (?,?,?,?)",
		productId,
		psu.Power,
		psu.Efficiency,
		psu.FormFactor,
	).Error
}

func InsertRam(db *gorm.DB, ram *dto.Ram, productId int64) error {
	return db.Exec("INSERT INTO rams (product_id, capacity, speed, type, compatibility, configuration) VALUES (?,?,?,?,?,?)",
		productId,
		ram.Capacity,
		ram.Speed,
		ram.Type,
		ram.Compatibility,
		ram.Configuration,
	).Error
}

func InsertSsd(db *gorm.DB, ssd *dto.Ssd, productId int64) error {
	return db.Exec("INSERT INTO ssds (product_id, capacity, reading_speed, writing_speed, form_factor) VALUES (?,?,?,?,?)",
		productId,
		ssd.Capacity,
		ssd.ReadingSpeed,
		ssd.WritingSpeed,
		ssd.FormFactor,
	).Error
}