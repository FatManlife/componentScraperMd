package service

import (
	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"github.com/FatManlife/component-finder/back-end/internal/models/orm"
)

func productMapping(product orm.Product) dto.ProductResponse{
	return dto.ProductResponse{
			ID: product.ID,
			Name: product.Name,
			ImageURL: product.ImageURL,
			Brand: product.Brand,
			Price: product.Price,
			Url: product.URL,
			Website_id: product.WebsiteID,
		}	
}
