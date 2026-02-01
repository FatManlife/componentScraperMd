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

	if len(params.Diagonal) > 0 {
		q = q.Where("aios.diagonal IN ?", params.Diagonal)
	}

	if len(params.Ram) > 0 { 
		q = q.Where("aios.ram IN ?", params.Ram)
	}

	if len(params.Storage) > 0 {
		q = q.Where("aios.storage IN ?", params.Storage)
	}

	if len(params.Cpu) > 0 {
		q = q.Where("aios.cpu IN ?", params.Cpu)
	}

	if len(params.Gpu) > 0 {
		q = q.Where("aios.gpu IN ?", params.Gpu)
	}

	if err := q.Find(&aios).Error ; err != nil {
		return nil, err
	}

	return aios, nil
}