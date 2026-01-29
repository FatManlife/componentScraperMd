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

func (r *PcMiniRepository) GetPcMinis(ctx context.Context, params dto.PcMiniParams) ([]orm.Product,error) {
	var pcMinis []orm.Product

	q := getDefaultProduct(r.db,ctx,params.DefaultParams)

	q.Joins("Join pc_minis ON pc_minis.product_id = products.id").Preload("PcMini")

	if params.Cpu != "" {
		q = q.Where("pc_minis.cpu = ?", params.Cpu)
	}
	if params.Gpu != "" {
		q = q.Where("pc_minis.gpu = ?", params.Gpu)
	}
	if params.Ram != "" {
		q = q.Where("pc_minis.ram = ?", params.Ram)
	}
	if params.Storage != "" {
		q = q.Where("pc_minis.storage = ?", params.Storage)
	}
	
	if err := q.Find(&pcMinis).Error; err != nil {
		return nil, err
	}

	return pcMinis, nil
}