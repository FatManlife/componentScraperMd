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

	if params.Cpu != "" {
		q = q.Where("laptops.cpu = ?", params.Cpu)
	}

	if params.Gpu != "" {
		q = q.Where("laptops.gpu = ?", params.Gpu)
	}

	if params.Ram != "" {
		q = q.Where("laptops.ram = ?", params.Ram)
	}

	if params.Storage != "" {
		q = q.Where("laptops.storage = ?", params.Storage)
	}

	if params.Diagonal != "" {
		q = q.Where("laptops.diagonal = ?", params.Diagonal)
	}

	if params.Battery != 0 {
		q = q.Where("laptops.battery = ?", params.Battery)
	}

	if err := q.Find(&laptops).Error; err != nil {
		return nil, err
	}

	return laptops, nil
}