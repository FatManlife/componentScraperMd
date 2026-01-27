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

func (r *ProductRepository) GetAllProducts(ctx context.Context ,limit int) ([]orm.Product, error) {
	var products []orm.Product

	err := r.db.WithContext(ctx).Limit(limit).Find(&products).Error

	if err != nil {
		return nil, err
	}

	return products, nil
}
