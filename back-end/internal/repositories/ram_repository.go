package repo

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"github.com/FatManlife/component-finder/back-end/internal/models/orm"
	"gorm.io/gorm"
)

type RamRepository struct {
	db *gorm.DB
}

func NewRamRepository(db *gorm.DB) *RamRepository {
	return &RamRepository{db: db}
}

func (r *RamRepository) GetRams(ctx context.Context, params dto.RamParams, ) ([]orm.Product, error) {
	var ram []orm.Product

	q := getDefaultProduct(r.db, ctx, params.DefaultParams)

	q.Joins("JOIN rams on rams.product_id = products.id").Preload("Ram")

	if params.Capacity != 0{
		q = q.Where("rams.capacity = ?", params.Capacity)
	}

	if params.Speed != 0{
		q = q.Where("rams.speed = ?", params.Speed)
	}

	if params.Type != ""{
		q = q.Where("rams.type = ?", params.Type)
	}

	if params.Compatibility != ""{
		q = q.Where("rams.compatibility = ?", params.Compatibility)
	}

	if params.Configuration != ""{
		q = q.Where("rams.configuration = ?", params.Configuration)
	}

	if err := q.Find(&ram).Error; err != nil {
		return nil, err
	}

	return ram, nil

}