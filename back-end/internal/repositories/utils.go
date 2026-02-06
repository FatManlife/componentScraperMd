package repo

import (
	"context"
	"strings"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"github.com/FatManlife/component-finder/back-end/internal/models/orm"
	"gorm.io/gorm"
)

var sortMap = map[string]string{
	"price_asc":  "products.price ASC",
	"price_desc": "products.price DESC",
}

func getDefaultProduct(db *gorm.DB ,ctx context.Context, params dto.ProductParams) *gorm.DB{
	q := db.Model(&orm.Product{}).WithContext(ctx)	

	if len(params.Website) > 0 {
		q = q.Joins("JOIN websites ON websites.id = products.website_id").Where("websites.name IN ?", params.Website)
	}

	if params.Name != "" {
		q = q.Where("products.name ILIKE ?", "%"+strings.TrimSpace(params.Name)+"%")
	}

	if params.Min >  0 {
		q = q.Where("products.price >= ?", params.Min)
	}

	if params.Max > 0 {
		q = q.Where("products.price <= ?", params.Max)
	}
	
	productOrder, ok := sortMap[params.Order]
	
	if !ok {
		productOrder = "products.id ASC"
	}

	q = q.Order(productOrder)

	q = q.Preload("Website")

	return q
}

func getDefaultCount(db *gorm.DB ,ctx context.Context, params dto.ProductParams) *gorm.DB{
	q := db.WithContext(ctx).Model(&orm.Product{})

	if len(params.Website) > 0 {
		q = q.Joins("JOIN websites ON websites.id = products.website_id").Where("websites.name IN ?", params.Website)
	}

	if params.Min >  0 {
		q = q.Where("products.price >= ?", params.Min)
	}

	if params.Max > 0 {
		q = q.Where("products.price <= ?", params.Max)
	}	

	return q
}

func setLimits(q *gorm.DB, offset int) {
	q = q.Limit(24)

	if offset > 1 {
		q = q.Offset((offset - 1) * 24)
	}
}


