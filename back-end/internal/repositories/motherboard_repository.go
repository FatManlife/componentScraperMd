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

	if params.Chipset != "" {
		q = q.Where("motherboards.chipset = ?", params.Chipset)
	}
	if params.Socket != "" {
		q = q.Where("motherboards.socket = ?", params.Socket)
	}
	if params.FormFactor != "" {
		q = q.Where("motherboards.form_factor = ?", params.FormFactor)
	}
	if params.RamSupport != "" {
		q = q.Where("motherboards.ram_support = ?", params.RamSupport)
	}
	if params.FormFactorRam != "" {
		q = q.Where("motherboards.form_factor_ram = ?", params.FormFactorRam)
	}
	
	if err := q.Find(&motherboards).Error; err != nil {
		return nil, err
	}

	return motherboards, nil	
}
