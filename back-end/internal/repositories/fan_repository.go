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

func (r *FanRepository) GetFans(ctx context.Context, params dto.FanParams) ([]orm.Product, error) {
	var fans []orm.Product

	q := getDefaultProduct(r.db, ctx, params.DefaultParams)

	q.Joins("JOIN fans ON fans.product_id = products.id").Preload("Fan")

	if params.Size != "" {
		q = q.Where("fans.size = ?", params.Size)
	}

	if params.FanRPM!= 0 {
		q = q.Where("fans.fan_rpm = ?", params.FanRPM)
	}

	if params.Noise != 0 {
		q = q.Where("fans.noise = ?", params.Noise)
	}

	if err := q.Find(&fans).Error; err != nil {
		return nil, err
	}

	return fans, nil
}	