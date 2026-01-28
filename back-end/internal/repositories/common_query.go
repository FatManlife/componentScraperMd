package repo

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/orm"
	"gorm.io/gorm"
)

func applyCommonFilters(db *gorm.DB ,ctx context.Context, limit int, website string, after int, brand string, min float64, max float64, order string) *gorm.DB{
	q := db.Model(&orm.Product{}).WithContext(ctx)

	if after > 0 {
		q = q.Where("products.id > ?", after)
	}	

	if website != "" {
		q = q.Joins("JOIN websites ON websites.id = products.website_id").Where("websites.name ILIKE ?", website)
	}

	if brand != "" {
		q = q.Where("products.brand ILIKE ?", brand)
	}

	if min != 0 {
		q = q.Where("products.price >= ?", min)
	}

	if max != 0 {
		q = q.Where("products.price <= ?", max)
	}

	sortMap := map[string]string{
		"price_asc":  "products.price ASC",
		"price_desc": "products.price DESC",
	}

	productOrder, ok := sortMap[order]
	
	if !ok {
		productOrder = "products.id ASC"
	}

	q = q.Order(productOrder)

	if limit > 0 {
		q = q.Limit(limit)
	}
	
	return q
}