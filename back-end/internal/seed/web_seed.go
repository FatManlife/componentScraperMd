package seed

import (
	"errors"

	"github.com/FatManlife/component-finder/back-end/internal/models/orm"
	"gorm.io/gorm"
)

func SeedWeb(db *gorm.DB) error {
	var w orm.Website
	result := db.First(&w)

	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil 
	}

	webs := []orm.Website{
		{Name: "Neocomputer", URL: "https://neocomputer.md/", Image: "https://neocomputer.md/image/catalog/logo.png"},
		{Name: "Xstore", URL: "https://xstore.md/", Image: "https://xstore.md/img/logo.png"},
		{Name: "PcPrime", URL: "https://prime-pc.md/", Image: "https://prime-pc.md/design/MegaPrime/images/logo.png"},
	}

	return db.Create(&webs).Error
}

func SeedCategory(db *gorm.DB) error {
	var c orm.Category
	result := db.First(&c)

	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil 
	}

	categories := []orm.Category{
		{Name: "aio"},
		{Name: "case"},
		{Name: "cooler"},
		{Name: "cpu"},
		{Name: "fan"},
		{Name: "gpu"},
		{Name: "hdd"},
		{Name: "laptop"},
		{Name: "motherboard"},
		{Name: "pc mini"},
		{Name: "pc"},
		{Name: "psu"},
		{Name: "ram"},
		{Name: "ssd"},
	}

	return db.Create(&categories).Error
}