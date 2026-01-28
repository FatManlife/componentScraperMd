package repo

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/orm"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

//implemnt Getting all products with filters
func (r *ProductRepository) GetAllProducts(ctx context.Context, limit int, website string, after int, brand string, min float64, max float64, order string) ([]orm.Product, error) {
	var products []orm.Product

	q := applyCommonFilters(r.db, ctx, limit, website, after, brand, min, max, order) 	
	
	err := q.Find(&products).Error

	if err != nil {
		return nil, err
	}

	return products, nil
}

