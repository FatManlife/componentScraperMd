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

func (r *CpuRepository) GetCpus(ctx context.Context, params dto.CpuParams) ([]orm.Product, int64, error){
	var cpus []orm.Product
	var count int64

	q := getDefaultProduct(r.db, ctx, params.DefaultParams)

	q.Joins("JOIN cpus ON cpus.product_id = products.id").Preload("Cpu")

	if len(params.Cores) > 0 {
		q = q.Where("cpus.cores IN ?", params.Cores)
	}

	if len(params.Threads) > 0 {
		q = q.Where("cpus.threads IN ?", params.Threads)
	}

	if len(params.BaseClock) > 0 {
		q = q.Where("cpus.base_clock IN ?", params.BaseClock)
	}

	if len(params.BoostClock) > 0 {
		q = q.Where("cpus.boost_clock IN ?", params.BoostClock)
	}

	if len(params.Socket) > 0 {
		q = q.Where("cpus.socket IN ?", params.Socket)
	}

	if err := q.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	setLimits(q, params.DefaultParams.Offset)

	if err := q.Find(&cpus).Error; err != nil {
		return nil, 0, err
	}

	return cpus, count, nil
}

func (r *CpuRepository) GetSockets(ctx context.Context) ([]string, error) {
	var sockets []string
	if err := r.db.WithContext(ctx).Model(&orm.Cpu{}).Distinct().Order("socket").Pluck("socket", &sockets).Error; err != nil {
		return nil, err
	}

	return sockets, nil
}

func (r *CpuRepository) GetBaseClocks(ctx context.Context) ([]float64, error) {
	var baseClocks []float64
	if err := r.db.WithContext(ctx).Model(&orm.Cpu{}).Distinct().Order("base_clock").Pluck("base_clock", &baseClocks).Error; err != nil {
		return nil, err
	}

	return baseClocks, nil
}

func (r *CpuRepository) GetBoostClocks(ctx context.Context) ([]float64, error) {
	var boostClocks []float64
	if err := r.db.WithContext(ctx).Model(&orm.Cpu{}).Distinct().Order("boost_clock").Pluck("boost_clock", &boostClocks).Error; err != nil {
		return nil, err
	}

	return boostClocks, nil
}

func (r *CpuRepository) GetCores(ctx context.Context) ([]int, error) {
	var cores []int
	if err := r.db.WithContext(ctx).Model(&orm.Cpu{}).Distinct().Order("cores").Pluck("cores", &cores).Error; err != nil {
		return nil, err
	}

	return cores, nil
}

func (r *CpuRepository) GetThreads(ctx context.Context) ([]int, error) {
	var threads []int
	if err := r.db.WithContext(ctx).Model(&orm.Cpu{}).Distinct().Order("threads").Pluck("threads", &threads).Error; err != nil {
		return nil, err
	}

	return threads, nil
}