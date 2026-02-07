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

func (r *AioRepository) GetAios (ctx context.Context, params dto.AioParams) ([]orm.Product, int64, error) {
	var aios []orm.Product
	var count int64

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

	if err := q.Count(&count).Error; err != nil {
		return nil, 0, err
	}
	
	setLimits(q,params.DefaultParams.Offset)

	if err := q.Find(&aios).Error ; err != nil {
		return nil, 0, err
	}

	return aios, count, nil
}

func (r *AioRepository) GetDiagonlas(ctx context.Context) ([]float64, error) {
	var diagonals []float64
	if err := r.db.WithContext(ctx).Model(&orm.Aio{}).Distinct().Pluck("diagonal", &diagonals).Error ; err != nil {
		return nil, err
	}

	return diagonals, nil
}

func (r *AioRepository) GetRams(ctx context.Context) ([]int, error) {
	var rams []int
	if err := r.db.WithContext(ctx).Model(&orm.Aio{}).Distinct().Pluck("ram", &rams).Error ; err != nil {
		return nil, err
	}

	return rams, nil
}

func (r *AioRepository) GetStorages(ctx context.Context) ([]int, error) {
	var storages []int
	if err := r.db.WithContext(ctx).Model(&orm.Aio{}).Distinct().Pluck("storage", &storages).Error ; err != nil {
		return nil, err
	}

	return storages, nil
}

func (r *AioRepository) GetCpus(ctx context.Context) ([]string, error) {
	var cpus []string
	if err := r.db.WithContext(ctx).Model(&orm.Aio{}).Distinct().Pluck("cpu", &cpus).Error ; err != nil {
		return nil, err
	}

	return cpus, nil
}

func (r *AioRepository) GetGpus(ctx context.Context) ([]string, error) {
	var gpus []string
	if err := r.db.WithContext(ctx).Model(&orm.Aio{}).Distinct().Pluck("gpu", &gpus).Error ; err != nil {
		return nil, err
	}

	return gpus, nil
}