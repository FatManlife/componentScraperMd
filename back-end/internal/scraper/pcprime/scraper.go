package pcprime

import (
	rawsql "github.com/FatManlife/component-finder/back-end/internal/db/raw_sql"
	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
)

func Run(s *rawsql.Storage){
	category := categoryColly()
	page := pageColly()
	product := productColly()
	productLinks := make(chan dto.Link)

	requestBodyProducts(category,page,product,&productLinks, s)

	category.Visit("https://prime-pc.md/")	
	page.Wait()
	product.Wait()

	close(productLinks)
}
