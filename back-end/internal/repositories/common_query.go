package repo

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"github.com/FatManlife/component-finder/back-end/internal/models/orm"
	"gorm.io/gorm"
)

func applyCommonFilters(db *gorm.DB ,ctx context.Context, params dto.ProductParams) *gorm.DB{
	q := db.Model(&orm.Product{}).WithContext(ctx)

	if params.After > 0 {
		q = q.Where("products.id > ?", params.After)
	}	

	if params.Website != "" {
		q = q.Joins("JOIN websites ON websites.id = products.website_id").Where("websites.name ILIKE ?", params.Website)
	}

	if params.Brand != "" {
		q = q.Where("products.brand ILIKE ?", params.Brand)
	}

	if params.Min != 0 {
		q = q.Where("products.price >= ?", params.Min)
	}

	if params.Max != 0 {
		q = q.Where("products.price <= ?", params.Max)
	}

	sortMap := map[string]string{
		"price_asc":  "products.price ASC",
		"price_desc": "products.price DESC",
	}

	productOrder, ok := sortMap[params.Order]
	
	if !ok {
		productOrder = "products.id ASC"
	}

	q = q.Order(productOrder)

	if params.Limit > 0 {
		q = q.Limit(params.Limit)
	}
	
	return q
}