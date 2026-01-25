package pcprime

import (
	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"gorm.io/gorm"
)

func Run(db *gorm.DB){
	category := categoryColly()
	page := pageColly()
	product := productColly()
	productLinks := make(chan dto.Link)

	requestBodyProducts(category,page,product,&productLinks, db)

	category.Visit("https://prime-pc.md/")	
	page.Wait()
	product.Wait()

	close(productLinks)
}
