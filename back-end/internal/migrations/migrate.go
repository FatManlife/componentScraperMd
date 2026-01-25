package migrations

import (
	"github.com/FatManlife/component-finder/back-end/internal/models/orm"
	"gorm.io/gorm"
)


func Migrate(db *gorm.DB) error{
	return db.AutoMigrate(
		&orm.Website{},
		&orm.Category{},
		&orm.Product{},
		&orm.Aio{},
		&orm.PcCase{},
		&orm.Cooler{},
		&orm.CoolerCompatibility{},
		&orm.CoolerCpu{},
		&orm.Cpu{},
		&orm.Fan{},
		&orm.Gpu{},
		&orm.Hdd{},
		&orm.Laptop{},
		&orm.Motherboard{},
		&orm.PcMini{},
		&orm.Pc{},
		&orm.Psu{},
		&orm.Ram{},
		&orm.Ssd{},
	)
}

func DelteTables(db *gorm.DB) error{
	return db.Migrator().DropTable(
		&orm.Website{},
		&orm.Product{},	
		&orm.Aio{},
		&orm.PcCase{},
		&orm.Cooler{},
		&orm.Cpu{},
		&orm.Category{},
		&orm.CoolerCompatibility{},
		&orm.CoolerCpu{},
		&orm.Fan{},
		&orm.Gpu{},
		&orm.Hdd{},
		&orm.Laptop{},
		&orm.Motherboard{},
		&orm.PcMini{},
		&orm.Pc{},
		&orm.Psu{},
		&orm.Ram{},
		&orm.Ssd{},
	)
}