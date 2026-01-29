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

	if params.Capacity != 0 {
		q = q.Where("hdds.capacity = ?", params.Capacity)
	}

	if params.RotationSpeed != 0 {
		q = q.Where("hdds.rpm = ?", params.RotationSpeed)
	}

	if params.FormFactor != "" {
		q = q.Where("hdds.interface = ?", params.FormFactor)
	}

	if err:= q.Find(&hdds).Error; err != nil {
		return nil, err
	}

	return hdds, nil
}