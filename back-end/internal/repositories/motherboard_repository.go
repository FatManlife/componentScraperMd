package repo

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"github.com/FatManlife/component-finder/back-end/internal/models/orm"
	"gorm.io/gorm"
)

type MotherboardRepository struct {
	db *gorm.DB
}

func NewMotherboardRepository(db *gorm.DB) *MotherboardRepository {
	return &MotherboardRepository{db: db}
}

func (r *MotherboardRepository) GetMotherboards(ctx context.Context, params dto.MotherboardParams) ([]orm.Product, error) {
	var motherboards []orm.Product

	q := getDefaultProduct(r.db,ctx, params.DefaultParams)

	q.Joins("Join motherboards ON motherboards.product_id = products.id").Preload("Motherboard")

	if len(params.Chipset) > 0 {
		q = q.Where("motherboards.chipset IN ?", params.Chipset)
	}

	if len(params.Socket) > 0 {
		q = q.Where("motherboards.socket IN ?", params.Socket)
	}

	if len(params.FormFactor) > 0 {
		q = q.Where("motherboards.form_factor IN ?", params.FormFactor)
	}
		
	if err := q.Find(&motherboards).Error; err != nil {
		return nil, err
	}

	return motherboards, nil	
}
