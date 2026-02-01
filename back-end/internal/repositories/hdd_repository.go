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

func (r *HddRepository) GetHdds(ctx context.Context, params dto.HddParams) ([]orm.Product, error) {
	var hdds []orm.Product

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
		q = q.Where("hdds.interface IN ?", params.FormFactor)
	}

	if err:= q.Find(&hdds).Error; err != nil {
		return nil, err
	}

	return hdds, nil
}