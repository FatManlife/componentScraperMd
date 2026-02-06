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

func (r *LaptopRepository) GetLaptops(ctx context.Context, params dto.LaptopParams) ([]orm.Product, int64, error) {
	var laptops []orm.Product
	var count int64

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

	if err := q.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	setLimits(q, params.DefaultParams.Offset)

	if err := q.Find(&laptops).Error; err != nil {
		return nil, 0, err
	}

	return laptops, count, nil
}

func (r *LaptopRepository) GetCpus(ctx context.Context) ([]string, error) {
	var cpus []string
	if err := r.db.WithContext(ctx).Model(&orm.Laptop{}).Distinct().Pluck("cpu", &cpus).Error; err != nil {
		return nil, err
	}

	return cpus, nil
}

func (r *LaptopRepository) GetGpus(ctx context.Context) ([]string, error) {
	var gpus []string
	if err := r.db.WithContext(ctx).Model(&orm.Laptop{}).Distinct().Pluck("gpu", &gpus).Error; err != nil {
		return nil, err
	}

	return gpus, nil
}

func (r *LaptopRepository) GetRams(ctx context.Context) ([]int, error) {
	var rams []int
	if err := r.db.WithContext(ctx).Model(&orm.Laptop{}).Distinct().Pluck("ram", &rams).Error; err != nil {
		return nil, err
	}

	return rams, nil
}

func (r *LaptopRepository) GetStorages(ctx context.Context) ([]int, error) {
	var storages []int
	if err := r.db.WithContext(ctx).Model(&orm.Laptop{}).Distinct().Pluck("storage", &storages).Error; err != nil {
		return nil, err
	}

	return storages, nil
}

func (r *LaptopRepository) GetDiagonals(ctx context.Context) ([]string, error) {
	var diagonals []string
	if err := r.db.WithContext(ctx).Model(&orm.Laptop{}).Distinct().Pluck("diagonal", &diagonals).Error; err != nil {
		return nil, err
	}

	return diagonals, nil
}

func (r *LaptopRepository) GetBattery(ctx context.Context) ([]float64, error) {
	var batteries []float64
	if err := r.db.WithContext(ctx).Model(&orm.Laptop{}).Distinct().Pluck("battery", &batteries).Error; err != nil {
		return nil, err
	}

	return batteries, nil
}