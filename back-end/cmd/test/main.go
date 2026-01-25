package main

import (
	"fmt"
	"strings"

	"github.com/FatManlife/component-finder/back-end/internal/config"
	ormsql "github.com/FatManlife/component-finder/back-end/internal/db/orm_sql"
	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"github.com/FatManlife/component-finder/back-end/internal/utils"
	"github.com/gocolly/colly"
)

func main() {
	db := config.ConnDb()

	links, err := ormsql.GetAllProductLinks(db)
	if err != nil {
		fmt.Println("Error fetching product links:", err)
	} 

	utils.PopulateVisitedUrls(links)

	utils.PrintVisitedUrls()

}

func setBaseAttrs(e *colly.HTMLElement, product *dto.BaseProduct, category_id int){
	product.Name = strings.TrimSpace(e.ChildText("div.top-title h1"))
	product.ImageURL = strings.TrimSpace(e.ChildAttr("div.row.prod_page img", "src"))
	product.Price = utils.CastFloat64(e.ChildText("div.xp-price"))
	product.Website_id = 1
	product.Url = e.Request.URL.String()
	product.Category_id = category_id 
}

