package repo

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"github.com/FatManlife/component-finder/back-end/internal/models/orm"
	"gorm.io/gorm"
)

type FanRepository struct {
	db *gorm.DB
}

func NewFanRepository(db *gorm.DB) *FanRepository {
	return &FanRepository{db: db}
}

func (r *FanRepository) GetFans(ctx context.Context, params dto.FanParams) ([]orm.Product, int64, error) {
	var fans []orm.Product
	var count int64

	q := getDefaultProduct(r.db, ctx, params.DefaultParams)

	q.Joins("JOIN fans ON fans.product_id = products.id").Preload("Fan")	

	if params.MinFanRPM > 0 {
		q = q.Where("fans.fan_rpm >= ?", params.MinFanRPM)
	}

	if params.MinNoise > 0 {
		q = q.Where("fans.noise >= ?", params.MinNoise)
	}

	if params.MaxFanRPM > 0 {
		q = q.Where("fans.fan_rpm <= ?", params.MaxFanRPM)
	}

	if params.MaxNoise > 0 {
		q = q.Where("fans.noise <= ?", params.MaxNoise)
	}

	if err := q.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	setLimits(q, params.DefaultParams.Offset)

	if err := q.Find(&fans).Error; err != nil {
		return nil, 0, err
	}

	return fans, count, nil
}	

func (r *FanRepository) GetFanRPMs (ctx context.Context) ([]int, error) {
	var fanRPMs []int
	if err := r.db.WithContext(ctx).Model(&orm.Fan{}).Distinct().Order("fan_rpm").Pluck("fan_rpm", &fanRPMs).Error; err != nil {
		return nil, err
	}

	return fanRPMs, nil
}

func (r *FanRepository) GetNoiseLevels(ctx context.Context) ([]float64, error) {
	var noiseLevels []float64
	if err := r.db.WithContext(ctx).Model(&orm.Fan{}).Distinct().Order("noise").Pluck("noise", &noiseLevels).Error; err != nil {
		return nil, err
	}
	
	return noiseLevels, nil
}


