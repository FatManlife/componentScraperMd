package pcprime

import "github.com/FatManlife/component-finder/back-end/internal/models"

func Run(){
	category := categoryColly()
	page := pageColly()
	product := productColly()
	productLinks := make(chan models.Link)

	requestBodyProducts(category,page,product,&productLinks)

	category.Visit("https://prime-pc.md/")	
	page.Wait()
	product.Wait()

	close(productLinks)
}
