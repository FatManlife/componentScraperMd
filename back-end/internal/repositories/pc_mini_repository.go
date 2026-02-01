package repo

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"github.com/FatManlife/component-finder/back-end/internal/models/orm"
	"gorm.io/gorm"
)

type PcMiniRepository struct{
	db *gorm.DB
}

func NewPcMiniRepository(db *gorm.DB) *PcMiniRepository {
	return &PcMiniRepository{db: db}
}

func (r *PcMiniRepository) GetPcMinis(ctx context.Context, params dto.PcParams) ([]orm.Product,error) {
	var pcMinis []orm.Product

	q := getDefaultProduct(r.db,ctx,params.DefaultParams)

	q.Joins("Join pc_minis ON pc_minis.product_id = products.id").Preload("PcMini")

	if len(params.Cpu) > 0 {
		q = q.Where("pc_minis.cpu IN ?", params.Cpu)
	}
	if len(params.Gpu) > 0 {
		q = q.Where("pc_minis.gpu IN ?", params.Gpu)
	}
	if len(params.Ram) > 0 {
		q = q.Where("pc_minis.ram IN ?", params.Ram)
	}
	if len(params.Storage) > 0 {
		q = q.Where("pc_minis.storage IN ?", params.Storage)
	}
	
	if err := q.Find(&pcMinis).Error; err != nil {
		return nil, err
	}

	return pcMinis, nil
}