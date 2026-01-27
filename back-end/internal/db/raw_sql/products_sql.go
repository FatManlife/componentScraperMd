package rawsql

import (
	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"gorm.io/gorm"
)

type Storage struct{
	Db *gorm.DB
}

func NewStorage(db *gorm.DB) *Storage {
	return &Storage{Db: db}
}

func (s *Storage) insertProduct(tx *gorm.DB, product *dto.BaseProduct) (int64, error) {
	var lastId int64

	err := tx.Raw("INSERT INTO products (name, image_url, brand, price, url, website_id, category_id) VALUES (?,?,?,?,?,?,?) RETURNING id",
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

func (s *Storage) InsertFan(fan *dto.Fan) error {
	return s.Db.Transaction(func(tx *gorm.DB) error {
		productId, err := s.insertProduct(tx, &fan.BaseAttrs)
		if err != nil {
			return err
		}

		return tx.Exec(
			"INSERT INTO fans (product_id, fan_rpm, noise, size) VALUES (?,?,?,?)",
			productId,
			fan.FanRPM,
			fan.Noise,
			fan.Size,
		).Error
	})
}

func (s *Storage) InsertAio(aio *dto.Aio) error {
	return s.Db.Transaction(func(tx *gorm.DB) error{
		productId, err := s.insertProduct(tx, &aio.BaseAttrs)

		if err != nil {
			return err
		}

		return tx.Exec("INSERT INTO aios (product_id, diagonal, cpu, ram, storage, gpu) VALUES (?,?,?,?,?,?)",
			productId,
			aio.Diagonal,
			aio.Cpu,
			aio.Ram,
			aio.Storage,
			aio.Gpu,
		).Error
	})		
}

func (s *Storage) InsertCase(caseProduct *dto.Case) error {
	return s.Db.Transaction(func(tx *gorm.DB) error {
		productId, err := s.insertProduct(tx, &caseProduct.BaseAttrs)
		
		if err != nil {
			return err
		}

		return tx.Exec("INSERT INTO pc_cases (product_id, format, motherboard_form_factor) VALUES (?,?,?)",
			productId,
			caseProduct.Format,
			caseProduct.MotherboardFormFactor,
		).Error
	})
}
func (s *Storage) InsertCooler(cooler *dto.Cooler) (int64, error) {
	var lastId int64

	err := s.Db.Transaction(func(tx *gorm.DB) error {
		productId, err := s.insertProduct(tx, &cooler.BaseAttrs)

		if err != nil {
			return err
		}


		err = tx.Raw("INSERT INTO coolers (product_id, type, fan_rpm, noise, size) VALUES (?,?,?,?,?) RETURNING id",
			productId,
			cooler.Type,
			cooler.FanRPM,
			cooler.Noise,
			cooler.Size,
		).Scan(&lastId).Error

		return err
	})
	
	return lastId, err
}

func (s *Storage) InsertCoolerCompatibility(cooler_id int64, compatibility []string) error {
	for _, comp := range compatibility {
		lastId := s.GetCoolerCompatibilityIdByName(comp)

		if lastId == 0 {
			if err := s.Db.Raw("INSERT INTO cooler_compatibilities (cpu) VALUES (?) RETURNING id",
				comp,
			).Scan(&lastId).Error; err != nil {
				return err
			}
		}

		if err := s.InsertCoolerCpus(cooler_id, lastId); err != nil {
			return err
		}
	}

	return nil
}

func (s *Storage) InsertCoolerCpus(cooler_id int64, compatibility_id int64) error {
	return s.Db.Exec("INSERT INTO cooler_cpus(cooler_id, compatibility_id) VALUES (?,?)",
		cooler_id,
		compatibility_id,
	).Error
}

func (s *Storage) GetCoolerCompatibilityIdByName(name string) (int64) {
	var id int64

	s.Db.Raw("SELECT id FROM cooler_compatibilities WHERE cpu = ?",
		name,
	).Scan(&id)

	return id
}

func (s *Storage) InsertCpu(cpu *dto.Cpu) error {
	return s.Db.Transaction(func(tx *gorm.DB) error {
		productId, err := s.insertProduct(tx, &cpu.BaseAttrs)

		if err != nil {
			return err
		}

		return tx.Exec("INSERT INTO cpus (product_id, cores, threads, base_clock, boost_clock, socket, tdp) VALUES (?,?,?,?,?,?,?)",
			productId,
			cpu.Cores,
			cpu.Threads,
			cpu.BaseClock,
			cpu.BoostClock,
			cpu.Socket,
			cpu.Tdp,
		).Error
	})
}

func (s *Storage) InsertGpu(gpu *dto.Gpu) error {
	return s.Db.Transaction(func(tx *gorm.DB) error {
		productId, err := s.insertProduct(tx, &gpu.BaseAttrs)

		if err != nil {
			return err
		}

		return tx.Exec("INSERT INTO gpus (product_id, chipset, vram, gpu_frequency, vram_frequency) VALUES (?,?,?,?,?)",
			productId,
			gpu.Chipset,
			gpu.Vram,
			gpu.GpuFrequency,
			gpu.VramFrequency,
		).Error
	})
}

func (s *Storage) InsertHdd(hdd *dto.Hdd) error {
	return s.Db.Transaction(func(tx *gorm.DB) error {
		productId, err := s.insertProduct(tx, &hdd.BaseAttrs)

		if err != nil {
			return err
		}

		return tx.Exec("INSERT INTO hdds (product_id, capacity, rotation_speed, form_factor) VALUES (?,?,?,?)",
			productId,
			hdd.Capacity,
			hdd.RotationSpeed,
			hdd.FormFactor,
		).Error
	})
}

func (s *Storage) InsertLaptop(laptop *dto.Laptop) error {
	return s.Db.Transaction(func(tx *gorm.DB) error {
		productId, err := s.insertProduct(tx, &laptop.BaseAttrs)

		if err != nil {
			return err
		}

		return tx.Exec("INSERT INTO laptops (product_id, cpu, gpu, ram, storage, diagonal, battery) VALUES (?,?,?,?,?,?,?)",
			productId,
			laptop.Cpu,
			laptop.Gpu,
			laptop.Ram,
			laptop.Storage,
			laptop.Diagonal,
			laptop.Battery,
		).Error
	})
}

func (s *Storage) InsertMotherboard(motherboard *dto.Motherboard) error {
	return s.Db.Transaction(func(tx *gorm.DB) error {
		productId, err := s.insertProduct(tx, &motherboard.BaseAttrs)

		if err != nil {
			return err
		}

		return tx.Exec("INSERT INTO motherboards (product_id, chipset, socket, form_factor, ram_support, form_factor_ram) VALUES (?,?,?,?,?,?)",
			productId,
			motherboard.Chipset,
			motherboard.Socket,
			motherboard.FormFactor,
			motherboard.RamSupport,
			motherboard.FormFactorRam,
		).Error
	})
}

func (s *Storage) InsertPcMini(pcmini *dto.PcMini) error {
	return s.Db.Transaction(func(tx *gorm.DB) error{
		productId, err := s.insertProduct(tx, &pcmini.BaseAttrs)

		if err != nil {
			return err
		}

		return tx.Exec("INSERT INTO pc_minis (product_id, cpu, gpu, ram, storage) VALUES (?,?,?,?,?)",
			productId,
			pcmini.Cpu,
			pcmini.Gpu,
			pcmini.Ram,
			pcmini.Storage,
		).Error
	})
}

func (s *Storage) InsertPc(pc *dto.Pc) error {
	return s.Db.Transaction(func(tx *gorm.DB) error {
		productId, err := s.insertProduct(tx, &pc.BaseAttrs)

		if err != nil {
			return err
		}

		return tx.Exec("INSERT INTO pcs (product_id, cpu, gpu, ram, storage, motherboard, psu, pc_case) VALUES (?,?,?,?,?,?,?,?)",
			productId,
			pc.Cpu,
			pc.Gpu,
			pc.Ram,
			pc.Storage,
			pc.Motherboard,
			pc.Psu,
			pc.Case,
		).Error
	})
}
	
func (s *Storage) InsertPsu(psu *dto.Psu) error {
	return s.Db.Transaction(func(tx *gorm.DB) error {
		productId, err := s.insertProduct(tx, &psu.BaseAttrs)

		if err != nil {
			return err
		}

		return tx.Exec("INSERT INTO psus (product_id, power, efficiency, form_factor) VALUES (?,?,?,?)",
			productId,
			psu.Power,
			psu.Efficiency,
			psu.FormFactor,
		).Error
	})
}

func (s *Storage) InsertRam(ram *dto.Ram) error {
	return s.Db.Transaction(func(tx *gorm.DB) error {
		productId, err := s.insertProduct(tx, &ram.BaseAttrs)

		if err != nil {
			return err
		}

		return tx.Exec("INSERT INTO rams (product_id, capacity, speed, type, compatibility, configuration) VALUES (?,?,?,?,?,?)",
			productId,
			ram.Capacity,
			ram.Speed,
			ram.Type,
			ram.Compatibility,
			ram.Configuration,
		).Error
	})
}

func (s *Storage) InsertSsd(ssd *dto.Ssd) error {
	return s.Db.Transaction(func(tx *gorm.DB) error {
		productId, err := s.insertProduct(tx, &ssd.BaseAttrs)

		if err != nil {
			return err
		}

		return tx.Exec("INSERT INTO ssds (product_id, capacity, reading_speed, writing_speed, form_factor) VALUES (?,?,?,?,?)",
			productId,
			ssd.Capacity,
			ssd.ReadingSpeed,
			ssd.WritingSpeed,
			ssd.FormFactor,
		).Error
	})
}