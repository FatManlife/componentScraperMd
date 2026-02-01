package repo

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"github.com/FatManlife/component-finder/back-end/internal/models/orm"
	"gorm.io/gorm"
)

type CpuRepository struct {
	db *gorm.DB
}

func NewCpuRepository(db *gorm.DB) *CpuRepository {
	return &CpuRepository{db: db}
}

func (r *CpuRepository) GetCpus(ctx context.Context, params dto.CpuParams) ([]orm.Product, error){
	var cpus []orm.Product

	q := getDefaultProduct(r.db, ctx, params.DefaultParams)

	q.Joins("JOIN cpus ON cpus.product_id = products.id").Preload("Cpu")

	if len(params.Cores) > 0 {
		q = q.Where("cpus.cores IN ?", params.Cores)
	}

	if len(params.Threads) > 0 {
		q = q.Where("cpus.threads IN ?", params.Threads)
	}

	if params.BaseClock != 0 {
		q = q.Where("cpus.base_clock = ?", params.BaseClock)
	}

	if params.BoostClock != 0 {
		q = q.Where("cpus.boost_clock = ?", params.BoostClock)
	}

	if params.Socket != "" {
		q = q.Where("cpus.socket = ?", params.Socket)
	}

	if err := q.Find(&cpus).Error; err != nil {
		return nil, err
	}

	return cpus, nil
}