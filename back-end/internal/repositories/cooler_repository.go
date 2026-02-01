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

func (r *CoolerRepository) GetCoolers(ctx context.Context, params dto.CoolerParams) ([]orm.Product, error) {
	var coolers []orm.Product

	q := getDefaultProduct(r.db,ctx,params.DefaultParams)

	q.Group("products.id").Joins("JOIN coolers ON coolers.product_id = products.id JOIN cooler_cpus ON cooler_cpus.cooler_id = coolers.id JOIN cooler_compatibilities ON cooler_compatibilities.id = cooler_cpus.compatibility_id")

	if params.Type != "" {
		q = q.Where("coolers.type = ?", params.Type)
	}
	if params.FanRPM != 0 {
		q = q.Where("coolers.fan_rpm = ?", params.FanRPM)
	}
	if params.Noise != 0 {
		q = q.Where("coolers.noise = ?", params.Noise)
	}
	if params.Size != "" {
		q = q.Where("coolers.size = ?", params.Size)
	}
	if len(params.Compatibility) > 0 {
		q = q.Where("cooler_compatibilities.cpu IN ?", params.Compatibility)
	}

	if err := q.Find(&coolers).Error; err != nil {
		return nil, err
	}

	return coolers, nil
}
