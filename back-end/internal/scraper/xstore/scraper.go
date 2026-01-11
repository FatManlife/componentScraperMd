package xstore

import (
	"github.com/FatManlife/component-finder/back-end/internal/utils"
)


func Run () error {
	category := newCategoryCollector() 
	page := newPageCollector()
	product := newProductCollector() 

	productsLinks := make(chan utils.Link)
	semaphor := make(chan struct{})

	requestBodyProduct(category, page, product, &productsLinks, &semaphor)

	category.Visit("https://xstore.md/")
	category.Wait()
	page.Wait()
	product.Wait()

	close(productsLinks)
	close(semaphor)

	return nil
}