package repo

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/orm"
	"gorm.io/gorm"
)

type AioRepository struct {
	db *gorm.DB
}

func NewAioRepository(db *gorm.DB) *AioRepository {
	return &AioRepository{db: db}
}

func (r *AioRepository) GetAios (ctx context.Context, limit int, website string, after int, brand string, min float64, max float64, order string,
	 diagonal string, ram string, storage string, cpu string, gpu string) ([]orm.Product, error) {
	var aios []orm.Product

	q := applyCommonFilters(r.db, ctx, limit, website, after, brand, min, max, order)

	q.Joins("JOIN aios on aios.product_id = products.id").Preload("Aio")

	if diagonal != "" {
		q = q.Where("aios.diagonal = ?", diagonal)
	}

	if ram != "" {
		q = q.Where("aios.ram = ?", ram)
	}

	if storage != "" {
		q = q.Where("aios.storage = ?", storage)
	}

	if cpu != "" {
		q = q.Where("aios.cpu = ?", cpu)
	}

	if gpu != "" {
		q = q.Where("aios.gpu = ?", gpu)
	}

	err := q.Find(&aios).Error

	if err != nil {
		return nil, err
	}

	return aios, nil
}