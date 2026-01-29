package repo

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
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
func (r *ProductRepository) GetAllProducts(ctx context.Context, params dto.ProductParams) ([]orm.Product, error) {
	var products []orm.Product

	q := getDefaultProduct(r.db, ctx, params) 	
	
	if err := q.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

