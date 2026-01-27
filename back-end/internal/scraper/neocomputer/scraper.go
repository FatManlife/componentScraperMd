package neocomputer

import (
	rawsql "github.com/FatManlife/component-finder/back-end/internal/db/raw_sql"
	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
)

func Run(s *rawsql.Storage){
	category := newCategoryColector()
	page := newPageColector()
	product := newProductCollector()

	productLink := make(chan dto.Link)

	requestBodyProducts(category, page, product, &productLink, s)

	category.Visit("https://neocomputer.md/")
	page.Wait()
}