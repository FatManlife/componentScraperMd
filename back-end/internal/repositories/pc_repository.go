package repo

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"github.com/FatManlife/component-finder/back-end/internal/models/orm"
	"gorm.io/gorm"
)

type PcRepository struct{
	db *gorm.DB
}

func NewPcRepository(db *gorm.DB) *PcRepository {
	return &PcRepository{db: db}
}	

func (r *PcRepository) GetPcs(ctx context.Context, params dto.PcParams) ([]orm.Product, error) {
	var pcs []orm.Product

	q := getDefaultProduct(r.db,ctx,params.DefaultParams)

	q.Joins("JOIN pcs on pcs.product_id = products.id").Preload("Pc")

	if params.Case != "" {
		q = q.Where("pcs.case = ?", params.Case)
	}
	if params.Cpu != "" {
		q = q.Where("pcs.cpu = ?", params.Cpu)
	}
	if params.Gpu != "" {
		q = q.Where("pcs.gpu = ?", params.Gpu)
	}
	if params.Motherboard != "" {
		q = q.Where("pcs.motherboard = ?", params.Motherboard)
	}
	if params.Psu != "" {
		q = q.Where("pcs.psu = ?", params.Psu)
	}
	if params.Ram != "" {
		q = q.Where("pcs.ram = ?", params.Ram)
	}
	if params.Storage != "" {
		q = q.Where("pcs.storage = ?", params.Storage)
	}
	if params.Motherboard != "" {
		q = q.Where("pcs.motherboard = ?", params.Motherboard)
	}
	
	if err := q.Find(&pcs).Error; err != nil {
		return nil, err
	}

	return pcs, nil	
}