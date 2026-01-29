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

func (r *RamRepository) GetRams(ctx context.Context, ramParams dto.RamParams, ) ([]orm.Product, error) {
	var ram []orm.Product

	q := getDefaultProduct(r.db, ctx, ramParams.DefaultParams)

	q.Joins("JOIN rams on rams.product_id = products.id").Preload("Ram")

	if ramParams.Capacity != 0{
		q = q.Where("rams.capacity = ?", ramParams.Capacity)
	}

	if ramParams.Speed != 0{
		q = q.Where("rams.speed = ?", ramParams.Speed)
	}

	if ramParams.Type != ""{
		q = q.Where("rams.type = ?", ramParams.Type)
	}

	if ramParams.Compatibility != ""{
		q = q.Where("rams.compatibility = ?", ramParams.Compatibility)
	}

	if ramParams.Configuration != ""{
		q = q.Where("rams.configuration = ?", ramParams.Configuration)
	}

	err := q.Find(&ram).Error

	if err != nil {
		return nil, err
	}

	return ram, nil

}