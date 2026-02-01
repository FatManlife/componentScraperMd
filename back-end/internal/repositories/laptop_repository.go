package repo

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"github.com/FatManlife/component-finder/back-end/internal/models/orm"
	"gorm.io/gorm"
)

type LaptopRepository struct {
	db *gorm.DB
}

func NewLaptopRepository(db *gorm.DB) *LaptopRepository {
	return &LaptopRepository{db: db}
}

func (r *LaptopRepository) GetLaptops(ctx context.Context, params dto.LaptopParams) ([]orm.Product, error) {
	var laptops []orm.Product

	q := getDefaultProduct(r.db,ctx, params.DefaultParams)

	q.Joins("JOIN laptops ON laptops.product_id = products.id").Preload("Laptop")

	if len(params.Cpu) > 0 {
		q = q.Where("laptops.cpu IN ?", params.Cpu)
	}

	if len(params.Gpu) > 0 {
		q = q.Where("laptops.gpu IN ?", params.Gpu)
	}

	if len(params.Ram) > 0 {
		q = q.Where("laptops.ram IN ?", params.Ram)
	}

	if len(params.Storage) > 0 {
		q = q.Where("laptops.storage IN ?", params.Storage)
	}

	if len(params.Diagonal) > 0 {
		q = q.Where("laptops.diagonal IN ?", params.Diagonal)
	}	

	if err := q.Find(&laptops).Error; err != nil {
		return nil, err
	}

	return laptops, nil
}