package xstore

import (
	rawsql "github.com/FatManlife/component-finder/back-end/internal/db/raw_sql"
	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
)


func Run (s *rawsql.Storage) error {
	category := newCategoryCollector() 
	page := newPageCollector()
	product := newProductCollector() 

	productsLinks := make(chan dto.Link)
	semaphor := make(chan struct{})

	requestBodyProduct(category, page, product, &productsLinks, &semaphor, s)

	category.Visit("https://xstore.md/componente-pc/racire")
	category.Wait()
	page.Wait()
	product.Wait()

	close(productsLinks)
	close(semaphor)

	return nil
}