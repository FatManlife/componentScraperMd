package repo

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"github.com/FatManlife/component-finder/back-end/internal/models/orm"
	"gorm.io/gorm"
)

type PcMiniRepository struct{
	db *gorm.DB
}

func NewPcMiniRepository(db *gorm.DB) *PcMiniRepository {
	return &PcMiniRepository{db: db}
}

func (r *PcMiniRepository) GetPcMinis(ctx context.Context, params dto.PcParams) ([]orm.Product, int64, error) {
	var pcMinis []orm.Product
	var count int64

	q := getDefaultProduct(r.db,ctx,params.DefaultParams)

	q.Joins("Join pc_minis ON pc_minis.product_id = products.id").Preload("PcMini")

	if len(params.Cpu) > 0 {
		q = q.Where("pc_minis.cpu IN ?", params.Cpu)
	}
	if len(params.Gpu) > 0 {
		q = q.Where("pc_minis.gpu IN ?", params.Gpu)
	}
	if len(params.Ram) > 0 {
		q = q.Where("pc_minis.ram IN ?", params.Ram)
	}
	if len(params.Storage) > 0 {
		q = q.Where("pc_minis.storage IN ?", params.Storage)
	}
	
	if err := q.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	setLimits(q, params.DefaultParams.Offset)

	if err := q.Find(&pcMinis).Error; err != nil {
		return nil, 0, err
	}

	return pcMinis, count, nil
}

func (r *PcMiniRepository) GetCpu(ctx context.Context) ([]string, error) {
	var cpus []string

	if err := r.db.WithContext(ctx).Model(&orm.PcMini{}).Distinct().Order("cpu").Pluck("cpu", &cpus).Error; err != nil {
		return nil, err
	}

	return cpus, nil
}

func (r *PcMiniRepository) GetGpu(ctx context.Context) ([]string, error) {
	var gpus []string

	if err := r.db.WithContext(ctx).Model(&orm.PcMini{}).Distinct().Order("gpu").Pluck("gpu", &gpus).Error; err != nil {
		return nil, err
	}

	return gpus, nil
}

func (r *PcMiniRepository) GetRam(ctx context.Context) ([]int, error) {
	var rams []int

	if err := r.db.WithContext(ctx).Model(&orm.PcMini{}).Distinct().Order("ram").Pluck("ram", &rams).Error; err != nil {
		return nil, err
	}

	return rams, nil
}

func (r *PcMiniRepository) GetStorage(ctx context.Context) ([]int, error) {
	var storages []int

	if err := r.db.WithContext(ctx).Model(&orm.PcMini{}).Distinct().Order("storage").Pluck("storage", &storages).Error; err != nil {
		return nil, err
	}

	return storages, nil
}