package repo

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"github.com/FatManlife/component-finder/back-end/internal/models/orm"
	"gorm.io/gorm"
)

type CoolerRepository struct {
	db *gorm.DB
}

func NewCoolerRepository(db *gorm.DB) *CoolerRepository {
	return &CoolerRepository{db: db}
}

func (r *CoolerRepository) GetCoolers(ctx context.Context, params dto.CoolerParams) ([]orm.Product, int64, error) {
	var coolers []orm.Product
	var count int64

	q := getDefaultProduct(r.db,ctx,params.DefaultParams)

	q.Group("products.id").Joins("JOIN coolers ON coolers.product_id = products.id JOIN cooler_cpus ON cooler_cpus.cooler_id = coolers.id JOIN cooler_compatibilities ON cooler_compatibilities.id = cooler_cpus.compatibility_id")

	if len(params.Type) > 0 {
		q = q.Where("coolers.type IN ?", params.Type)
	}
	if len(params.FanRPM) > 0 {
		q = q.Where("coolers.fan_rpm IN ?", params.FanRPM)
	}
	if len(params.Noise) > 0 {
		q = q.Where("coolers.noise IN ?", params.Noise)
	}
	if len(params.Compatibility) > 0 {
		q = q.Where("cooler_compatibilities.cpu IN ?", params.Compatibility)
	}

	if err := q.Count(&count).Error; err != nil {
		return nil, 0, err
	}
	
	setLimits(q, params.DefaultParams.Offset)

	if err := q.Find(&coolers).Error; err != nil {
		return nil, 0, err
	}

	return coolers, count, nil
}

func (r *CoolerRepository) GetCompatibility(ctx context.Context) ([]string, error) {
	var compatibilities []string
	if err := r.db.WithContext(ctx).Model(&orm.CoolerCompatibility{}).Distinct().Order("cpu").Pluck("cpu", &compatibilities).Error; err != nil {
		return nil, err
	}

	return compatibilities, nil
}

func (r *CoolerRepository) GetTypes(ctx context.Context) ([]string, error) {
	var types []string
	if err := r.db.WithContext(ctx).Model(&orm.Cooler{}).Distinct().Order("type").Pluck("type", &types).Error; err != nil {
		return nil, err
	}

	return types, nil
}

func (r *CoolerRepository) GetFanRPMs(ctx context.Context) ([]int, error) {
	var fanRPMs []int
	if err := r.db.WithContext(ctx).Model(&orm.Cooler{}).Distinct().Order("fan_rpm").Pluck("fan_rpm", &fanRPMs).Error; err != nil {
		return nil, err
	}

	return fanRPMs, nil
}

func (r *CoolerRepository) GetNoises(ctx context.Context) ([]float64, error) {
	var noises []float64
	if err := r.db.WithContext(ctx).Model(&orm.Cooler{}).Distinct().Order("noise").Pluck("noise", &noises).Error; err != nil {
		return nil, err
	}

	return noises, nil
}