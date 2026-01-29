package service

import (
	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"github.com/FatManlife/component-finder/back-end/internal/models/orm"
)

func checkDefaultParams(prams *dto.ProductParams) {
	//Cheking min and max values
	if prams.Min > prams.Max {
		prams.Min = 0
		prams.Max = 0
	}

	if prams.Min < 0 {
		prams.Min = 0
	}

	if prams.Max < 0 {
		prams.Max = 0
	}

	// limit
	if prams.Limit <= 0 {
		prams.Limit = 20
	}

	// after
	if prams.After < 0 {
		prams.After = 0
	}
}

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
