package xstore

import (
	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"gorm.io/gorm"
)


func Run (db *gorm.DB) error {
	category := newCategoryCollector() 
	page := newPageCollector()
	product := newProductCollector() 

	productsLinks := make(chan dto.Link)
	semaphor := make(chan struct{})

	requestBodyProduct(category, page, product, &productsLinks, &semaphor, db)

	category.Visit("https://xstore.md/componente-pc/racire")
	category.Wait()
	page.Wait()
	product.Wait()

	close(productsLinks)
	close(semaphor)

	return nil
}