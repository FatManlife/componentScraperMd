package rawsql

import (
	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"gorm.io/gorm"
)

func insertProduct(db *gorm.DB, product *dto.BaseProduct) (int64, error) {
	var lastId int64

	err := db.Raw("INSERT INTO products (name, image_url, brand, price, url, website_id, category_id) VALUES (?,?,?,?,?,?,?) RETURNING id",
		product.Name,
		product.ImageURL,
		product.Brand,
		product.Price,
		product.Url,
		product.Website_id,
		product.Category_id,
	).Scan(&lastId).Error 
	
	return lastId, err
}

func InsertFan(db *gorm.DB, fan *dto.Fan) error {
	productId, err := insertProduct(db, &fan.BaseAttrs)

	if err != nil {
		return err
	}

	return db.Exec("INSERT INTO fans (product_id, fan_rpm, noise, size) VALUES (?,?,?,?)",
		productId,
		fan.FanRPM,
		fan.Noise,
		fan.Size,
	).Error
}

func InsertAio(db *gorm.DB, aio *dto.Aio) error {
	productId, err := insertProduct(db, &aio.BaseAttrs)

	if err != nil {
		return err
	}

	return db.Exec("INSERT INTO aios (product_id, diagonal, cpu, ram, storage, gpu) VALUES (?,?,?,?,?,?)",
		productId,
		aio.Diagonal,
		aio.Cpu,
		aio.Ram,
		aio.Storage,
		aio.Gpu,
	).Error
}

func InsertCase(db *gorm.DB, caseProduct *dto.Case) error {
	productId, err := insertProduct(db, &caseProduct.BaseAttrs)

	if err != nil {
		return err
	}

	return db.Exec("INSERT INTO pc_cases (product_id, format, motherboard_form_factor) VALUES (?,?,?)",
		productId,
		caseProduct.Format,
		caseProduct.MotherboardFormFactor,
	).Error
}

func InsertCooler(db *gorm.DB, cooler *dto.Cooler) (int64, error) {
	productId, err := insertProduct(db, &cooler.BaseAttrs)

	if err != nil {
		return 0, err
	}

	var lastId int64

	err = db.Raw("INSERT INTO coolers (product_id, type, fan_rpm, noise, size) VALUES (?,?,?,?,?) RETURNING id",
		productId,
		cooler.Type,
		cooler.FanRPM,
		cooler.Noise,
		cooler.Size,
	).Scan(&lastId).Error
	
	return lastId, err
}

func InsertCoolerCompatibility(db *gorm.DB, cooler_id int64, compatibility []string) error {
	for _, comp := range compatibility {
		lastId := GetCoolerCompatibilityIdByName(db, comp)

		if lastId == 0 {
			if err := db.Raw("INSERT INTO cooler_compatibilities (cpu) VALUES (?) RETURNING id",
				comp,
			).Scan(&lastId).Error; err != nil {
				return err
			}
		}

		if err := InsertCoolerCpus(db, cooler_id, lastId); err != nil {
			return err
		}
		
	}

	return nil
}

func InsertCoolerCpus(db *gorm.DB, cooler_id int64, compatibility_id int64) error {
	return db.Exec("INSERT INTO cooler_cpus(cooler_id, compatibility_id) VALUES (?,?)",
		cooler_id,
		compatibility_id,
	).Error
}

func GetCoolerCompatibilityIdByName(db *gorm.DB, name string) (int64) {
	var id int64

	db.Raw("SELECT id FROM cooler_compatibilities WHERE cpu = ?",
		name,
	).Scan(&id)

	return id
}

func InsertCpu(db *gorm.DB, cpu *dto.Cpu) error {
	productId, err := insertProduct(db, &cpu.BaseAttrs)

	if err != nil {
		return err
	}

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

func InsertGpu(db *gorm.DB, gpu *dto.Gpu) error {
	productId, err := insertProduct(db, &gpu.BaseAttrs)

	if err != nil {
		return err
	}

	return db.Exec("INSERT INTO gpus (product_id, chipset, vram, gpu_frequency, vram_frequency) VALUES (?,?,?,?,?)",
		productId,
		gpu.Chipset,
		gpu.Vram,
		gpu.GpuFrequency,
		gpu.VramFrequency,
	).Error
}

func InsertHdd(db *gorm.DB, hdd *dto.Hdd) error {
	productId, err := insertProduct(db, &hdd.BaseAttrs)

	if err != nil {
		return err
	}

	return db.Exec("INSERT INTO hdds (product_id, capacity, rotation_speed, form_factor) VALUES (?,?,?,?)",
		productId,
		hdd.Capacity,
		hdd.RotationSpeed,
		hdd.FormFactor,
	).Error
}

func InsertLaptop(db *gorm.DB, laptop *dto.Laptop) error {
	productId, err := insertProduct(db, &laptop.BaseAttrs)

	if err != nil {
		return err
	}

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

func InsertMotherboard(db *gorm.DB, motherboard *dto.Motherboard) error {
	productId, err := insertProduct(db, &motherboard.BaseAttrs)

	if err != nil {
		return err
	}

	return db.Exec("INSERT INTO motherboards (product_id, chipset, socket, form_factor, ram_support, form_factor_ram) VALUES (?,?,?,?,?,?)",
		productId,
		motherboard.Chipset,
		motherboard.Socket,
		motherboard.FormFactor,
		motherboard.RamSupport,
		motherboard.FormFactorRam,
	).Error
}

func InsertPcMini(db *gorm.DB, pcmini *dto.PcMini) error {
	productId, err := insertProduct(db, &pcmini.BaseAttrs)

	if err != nil {
		return err
	}

	return db.Exec("INSERT INTO pc_minis (product_id, cpu, gpu, ram, storage) VALUES (?,?,?,?,?)",
		productId,
		pcmini.Cpu,
		pcmini.Gpu,
		pcmini.Ram,
		pcmini.Storage,
	).Error
}

func InsertPc(db *gorm.DB, pc *dto.Pc) error {
	productId, err := insertProduct(db, &pc.BaseAttrs)

	if err != nil {
		return err
	}

	return db.Exec("INSERT INTO pcs (product_id, cpu, gpu, ram, storage, motherboard, psu, pc_case) VALUES (?,?,?,?,?,?,?,?)",
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

func InsertPsu(db *gorm.DB, psu *dto.Psu) error {
	productId, err := insertProduct(db, &psu.BaseAttrs)

	if err != nil {
		return err
	}

	return db.Exec("INSERT INTO psus (product_id, power, efficiency, form_factor) VALUES (?,?,?,?)",
		productId,
		psu.Power,
		psu.Efficiency,
		psu.FormFactor,
	).Error
}

func InsertRam(db *gorm.DB, ram *dto.Ram) error {
	productId, err := insertProduct(db, &ram.BaseAttrs)

	if err != nil {
		return err
	}

	return db.Exec("INSERT INTO rams (product_id, capacity, speed, type, compatibility, configuration) VALUES (?,?,?,?,?,?)",
		productId,
		ram.Capacity,
		ram.Speed,
		ram.Type,
		ram.Compatibility,
		ram.Configuration,
	).Error
}

func InsertSsd(db *gorm.DB, ssd *dto.Ssd) error {
	productId, err := insertProduct(db, &ssd.BaseAttrs)

	if err != nil {
		return err
	}

	return db.Exec("INSERT INTO ssds (product_id, capacity, reading_speed, writing_speed, form_factor) VALUES (?,?,?,?,?)",
		productId,
		ssd.Capacity,
		ssd.ReadingSpeed,
		ssd.WritingSpeed,
		ssd.FormFactor,
	).Error
}