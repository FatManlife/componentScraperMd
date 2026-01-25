package neocomputer

import (
	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"gorm.io/gorm"
)

func Run(db *gorm.DB){
	category := newCategoryColector()
	page := newPageColector()
	product := newProductCollector()

	productLink := make(chan dto.Link)

	requestBodyProducts(category, page, product, &productLink, db)

	category.Visit("https://neocomputer.md/")
	page.Wait()
}