package repo

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"github.com/FatManlife/component-finder/back-end/internal/models/orm"
	"gorm.io/gorm"
)

type HddRepository struct {
	db *gorm.DB
}

func NewHddRepository(db *gorm.DB) *HddRepository {
	return &HddRepository{db: db}
}

func (r *HddRepository) GetHdds(ctx context.Context, params dto.HddParams) ([]orm.Product, int64, error) {
	var hdds []orm.Product
	var count int64

	q := getDefaultProduct(r.db,ctx, params.DefaultParams)

	q.Joins("JOIN hdds ON hdds.product_id = products.id").Preload("Hdd")

	if params.MinCapacity > 0 {
		q = q.Where("hdds.capacity >= ?", params.MinCapacity)
	}

	if params.MaxCapacity > 0 {
		q = q.Where("hdds.capacity <= ?", params.MaxCapacity)
	}

	if params.MinRotationSpeed > 0 {
		q = q.Where("hdds.rotation_speed >= ?", params.MinRotationSpeed)
	}

	if params.MaxRotationSpeed > 0 {
		q = q.Where("hdds.rotation_speed <= ?", params.MaxRotationSpeed)
	}
	
	if len(params.FormFactor) > 0 {
		q = q.Where("hdds.form_factor IN ?", params.FormFactor)
	}

	if err:= q.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	setLimits(q, params.DefaultParams.Offset)

	if err:= q.Find(&hdds).Error; err != nil {
		return nil, 0, err
	}

	return hdds, count, nil
}

func (r *HddRepository) GetFormFactors(ctx context.Context) ([]string, error) {
	var formFactors []string
	if err := r.db.WithContext(ctx).Model(&orm.Hdd{}).Distinct().Pluck("form_factor", &formFactors).Error; err != nil {
		return nil, err
	}

	return formFactors, nil
}

func (r *HddRepository) GetRotationSpeeds(ctx context.Context) ([]int, error) {
	var rotationSpeeds []int
	if err := r.db.WithContext(ctx).Model(&orm.Hdd{}).Distinct().Pluck("rotation_speed", &rotationSpeeds).Error; err != nil {
		return nil, err
	}
	
	return rotationSpeeds, nil
}

func (r *HddRepository) GetCapacities(ctx context.Context) ([]int, error) {
	var capacities []int
	if err := r.db.WithContext(ctx).Model(&orm.Hdd{}).Distinct().Pluck("capacity", &capacities).Error; err != nil {
		return nil, err
	}
	
	return capacities, nil
}
