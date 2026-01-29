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

func (r *AioRepository) GetAios (ctx context.Context, params dto.AioParams) ([]orm.Product, error) {
	var aios []orm.Product

	q := getDefaultProduct(r.db, ctx, params.DefaultParams)

	q.Joins("JOIN aios on aios.product_id = products.id").Preload("Aio")

	if params.Diagonal != "" {
		q = q.Where("aios.diagonal = ?", params.Diagonal)
	}

	if params.Ram != "" {
		q = q.Where("aios.ram = ?", params.Ram)
	}

	if params.Storage != "" {
		q = q.Where("aios.storage = ?", params.Storage)
	}

	if params.Cpu != "" {
		q = q.Where("aios.cpu = ?", params.Cpu)
	}

	if params.Gpu != "" {
		q = q.Where("aios.gpu = ?", params.Gpu)
	}

	if err := q.Find(&aios).Error ; err != nil {
		return nil, err
	}

	return aios, nil
}