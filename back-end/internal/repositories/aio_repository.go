package repo

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"github.com/FatManlife/component-finder/back-end/internal/models/orm"
	"gorm.io/gorm"
)

type AioRepository struct {
	db *gorm.DB
}

func NewAioRepository(db *gorm.DB) *AioRepository {
	return &AioRepository{db: db}
}

func (r *AioRepository) GetAios (ctx context.Context, aioParams dto.AioParams) ([]orm.Product, error) {
	var aios []orm.Product

	q := getDefaultProduct(r.db, ctx, aioParams.DefaultParams)

	q.Joins("JOIN aios on aios.product_id = products.id").Preload("Aio")

	if aioParams.Diagonal != "" {
		q = q.Where("aios.diagonal = ?", aioParams.Diagonal)
	}

	if aioParams.Ram != "" {
		q = q.Where("aios.ram = ?", aioParams.Ram)
	}

	if aioParams.Storage != "" {
		q = q.Where("aios.storage = ?", aioParams.Storage)
	}

	if aioParams.Cpu != "" {
		q = q.Where("aios.cpu = ?", aioParams.Cpu)
	}

	if aioParams.Gpu != "" {
		q = q.Where("aios.gpu = ?", aioParams.Gpu)
	}

	err := q.Find(&aios).Error

	if err != nil {
		return nil, err
	}

	return aios, nil
}