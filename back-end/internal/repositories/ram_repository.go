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

	if params.MinCapacity > 0{
		q = q.Where("rams.capacity >= ?", params.MinCapacity)
	}

	if params.MaxCapacity > 0{
		q = q.Where("rams.capacity <= ?", params.MaxCapacity)
	}

	if params.MinSpeed > 0{
		q = q.Where("rams.speed >= ?", params.MinSpeed)
	}

	if params.MaxSpeed > 0{
		q = q.Where("rams.speed <= ?", params.MaxSpeed)
	}

	if len(params.Type) > 0{ 
		q = q.Where("rams.type IN ?", params.Type)
	}

	if len(params.Compatibility) > 0{
		q = q.Where("rams.compatibility IN ?", params.Compatibility)
	}

	if len(params.Configuration) > 0{ 
		q = q.Where("rams.configuration IN ?", params.Configuration)
	}

	if err := q.Find(&ram).Error; err != nil {
		return nil, err
	}

	return ram, nil

}