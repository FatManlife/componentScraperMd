package migrations

import (
	"github.com/FatManlife/component-finder/back-end/internal/models/orm"
	"gorm.io/gorm"
)


func Migrate(db *gorm.DB) error{
	return db.AutoMigrate(
		&orm.Aio{},
		&orm.Case{},
		&orm.Cooler{},
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
		&orm.Website{},
	)
}