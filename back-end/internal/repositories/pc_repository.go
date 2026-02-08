package repo

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"github.com/FatManlife/component-finder/back-end/internal/models/orm"
	"gorm.io/gorm"
)

type PcRepository struct{
	db *gorm.DB
}

func NewPcRepository(db *gorm.DB) *PcRepository {
	return &PcRepository{db: db}
}	

func (r *PcRepository) GetPcs(ctx context.Context, params dto.PcParams) ([]orm.Product, int64, error) {
	var pcs []orm.Product
	var count int64

	q := getDefaultProduct(r.db,ctx,params.DefaultParams)

	q.Joins("JOIN pcs on pcs.product_id = products.id").Preload("Pc")
	
	if len(params.Cpu) > 0 {
		q = q.Where("pcs.cpu IN ?", params.Cpu)
	}

	if len(params.Gpu) > 0 {
		q = q.Where("pcs.gpu IN ?", params.Gpu)
	}
	
	if len(params.Ram) > 0 {
		q = q.Where("pcs.ram IN ?", params.Ram)
	}

	if len(params.Storage) > 0 {
		q = q.Where("pcs.storage IN ?", params.Storage)
	}

	if err := q.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	setLimits(q, params.DefaultParams.Offset)
	
	if err := q.Find(&pcs).Error; err != nil {
		return nil, 0, err
	}

	return pcs, count, nil	
}

func (r *PcRepository) GetCpu(ctx context.Context) ([]string, error) {
	var cpus []string

	if err := r.db.WithContext(ctx).Model(&orm.Pc{}).Distinct().Order("cpu").Pluck("cpu", &cpus).Error; err != nil {
		return nil, err
	}

	return cpus, nil
}

func (r *PcRepository) GetGpu(ctx context.Context) ([]string, error) {
	var gpus []string

	if err := r.db.WithContext(ctx).Model(&orm.Pc{}).Distinct().Order("gpu").Pluck("gpu", &gpus).Error; err != nil {
		return nil, err
	}

	return gpus, nil
}	

func (r *PcRepository) GetRam(ctx context.Context) ([]int, error) {
	var rams []int

	if err := r.db.WithContext(ctx).Model(&orm.Pc{}).Distinct().Order("ram").Pluck("ram", &rams).Error; err != nil {
		return nil, err
	}

	return rams, nil
}

func (r *PcRepository) GetStorage(ctx context.Context) ([]int, error) {
	var storages []int

	if err := r.db.WithContext(ctx).Model(&orm.Pc{}).Distinct().Order("storage").Pluck("storage", &storages).Error; err != nil {
		return nil, err
	}

	return storages, nil
}	