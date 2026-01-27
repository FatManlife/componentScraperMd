package ormsql

import (
	"github.com/FatManlife/component-finder/back-end/internal/models/orm"
	"gorm.io/gorm"
)

func GetAllProductLinks(db *gorm.DB) ([]string, error) {
    var links []string
    
    err := db.Model(&orm.Product{}).Select("url").Pluck("url", &links).Error

    if err != nil {
        return nil, err
    }

    return links, nil
}

func GetAllProducts(db *gorm.DB, nr int) ([]orm.Product, error) {
    var products []orm.Product

    err := db.Limit(nr).Find(&products).Error

    if err != nil {
        return nil, err
    }

    return products, nil
}
