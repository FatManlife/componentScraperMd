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

func (r *GpuRepository) GetGpus(ctx context.Context, params dto.GpuParams) ([]orm.Product, error){
	var gpu []orm.Product

	q := getDefaultProduct(r.db, ctx, params.DefaultParams)

	q.Joins("JOIN gpus ON gpus.product_id = products.id").Preload("Gpu")

	if params.Chipset != "" {
		q = q.Where("gpus.chipset = ?", params.Chipset)
	}
	if params.Vram != 0 {
		q = q.Where("gpus.vram = ?", params.Vram)
	}
	if params.GpuFrequency != 0 {
		q = q.Where("gpus.gpu_frequency = ?", params.GpuFrequency)
	}
	if params.VramFrequency != 0 {
		q = q.Where("gpus.vram_frequency = ?", params.VramFrequency)
	}

	if err := q.Find(&gpu).Error; err != nil {
		return nil, err
	}

	return gpu, nil	
}