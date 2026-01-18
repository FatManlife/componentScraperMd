package neocomputer

import "github.com/FatManlife/component-finder/back-end/internal/models"

func Run(){
	category := newCategoryColector()
	page := newPageColector()
	product := newProductCollector()

	productLink := make(chan models.Link)

	requestBodyProducts(category, page, product, &productLink)

	category.Visit("https://neocomputer.md/")
	page.Wait()
}