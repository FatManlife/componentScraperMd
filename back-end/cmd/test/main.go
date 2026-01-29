package main

import (
	"encoding/json"
	"fmt"

	"github.com/FatManlife/component-finder/back-end/internal/config"
	ormsql "github.com/FatManlife/component-finder/back-end/internal/db/orm_sql"
	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
)

func main() {
	db := config.ConnDb()

	productDb,err := ormsql.GetAllProducts(db,24)

	if err != nil {
		fmt.Println(err)
	}

	for _, product := range productDb {
		data := dto.BaseProduct{
			Name:       product.Name,
			ImageURL:   product.ImageURL,
			Brand:      product.Brand,
			Price:      product.Price,
			Url:        product.URL,
		}

		dataJson, err := json.MarshalIndent(data, "", "  ")

		if err != nil {
			fmt.Println("Error marshalling product:", err)
			return
		}

		fmt.Println(string(dataJson))
	}
}


