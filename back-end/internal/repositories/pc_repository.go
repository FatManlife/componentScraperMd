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
	
	if len(params.Cpu) > 0 {
		q = q.Where("pcs.cpu IN ?", params.Cpu)
	}

	if len(params.Gpu) > 0 {
		q = q.Where("pcs.gpu IN ?", params.Gpu)
	}
	
	if len(params.Ram) > 0 {
		q = q.Where("pcs.ram IN ?", params.Ram)
	}

	if len(params.Storage) > 0 {
		q = q.Where("pcs.storage IN ?", params.Storage)
	}
	
	if err := q.Find(&pcs).Error; err != nil {
		return nil, err
	}

	return pcs, nil	
}