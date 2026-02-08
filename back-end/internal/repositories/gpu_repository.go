package repo

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"github.com/FatManlife/component-finder/back-end/internal/models/orm"
	"gorm.io/gorm"
)

type GpuRepository struct {
	db *gorm.DB	
}

func NewGpuRepository(db *gorm.DB) *GpuRepository {
	return &GpuRepository{db: db}
}

func (r *GpuRepository) GetGpus(ctx context.Context, params dto.GpuParams) ([]orm.Product, int64, error){
	var gpu []orm.Product
	var count int64

	q := getDefaultProduct(r.db, ctx, params.DefaultParams)

	q.Joins("JOIN gpus ON gpus.product_id = products.id").Preload("Gpu")

	if params.Chipset != "" {
		q = q.Where("gpus.chipset = ?", params.Chipset)
	}
	if params.MinVram > 0 {
		q = q.Where("gpus.vram >= ?", params.MinVram)
	}
	if params.MaxVram > 0 {
		q = q.Where("gpus.vram <= ?", params.MaxVram)
	}
	if params.MinGpuFrequency > 0 {
		q = q.Where("gpus.gpu_frequency >= ?", params.MinGpuFrequency)
	}
	if params.MaxGpuFrequency > 0 {
		q = q.Where("gpus.gpu_frequency <= ?", params.MaxGpuFrequency)
	}
	if params.MinVramFrequency > 0 {
		q = q.Where("gpus.vram_frequency >= ?", params.MinVramFrequency)
	}
	if params.MaxVramFrequency > 0 {
		q = q.Where("gpus.vram_frequency <= ?", params.MaxVramFrequency)
	}

	if err := q.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	setLimits(q, params.DefaultParams.Offset)

	if err := q.Find(&gpu).Error; err != nil {
		return nil, 0, err
	}

	return gpu, count, nil	
}

func (r *GpuRepository) GetChipsets(ctx context.Context) ([]string, error) {
	var chipsets []string
	if err := r.db.WithContext(ctx).Model(&orm.Gpu{}).Distinct().Order("chipset").Pluck("chipset", &chipsets).Error; err != nil {
		return nil, err
	}

	return chipsets, nil
}

func (r *GpuRepository) GetGpuFrequencies(ctx context.Context) ([]int, error) {
	var gpuFrequencies []int
	if err := r.db.WithContext(ctx).Model(&orm.Gpu{}).Distinct().Order("gpu_frequency").Pluck("gpu_frequency", &gpuFrequencies).Error; err != nil {
		return nil, err
	}

	return gpuFrequencies, nil
}

func (r *GpuRepository) GetVramFrequencies(ctx context.Context) ([]int, error) {
	var vramFrequencies []int
	if err := r.db.WithContext(ctx).Model(&orm.Gpu{}).Distinct().Order("vram_frequency").Pluck("vram_frequency", &vramFrequencies).Error; err != nil {
		return nil, err
	}

	return vramFrequencies, nil
}

func (r *GpuRepository) GetVrams(ctx context.Context) ([]int, error) {
	var vrams []int
	if err := r.db.WithContext(ctx).Model(&orm.Gpu{}).Distinct().Order("vram").Pluck("vram", &vrams).Error; err != nil {
		return nil, err
	}
	
	return vrams, nil
}
