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

func (r *MotherboardRepository) GetMotherboards(ctx context.Context, params dto.MotherboardParams) ([]orm.Product, int64, error) {
	var motherboards []orm.Product
	var count int64

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
		
	if err := q.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	setLimits(q, params.DefaultParams.Offset)

	if err := q.Find(&motherboards).Error; err != nil {
		return nil, 0, err
	}

	return motherboards, count, nil	
}

func (r *MotherboardRepository) GetChipset(ctx context.Context) ([]string, error) {
	var chipsets []string

	if err := r.db.WithContext(ctx).Model(&orm.Motherboard{}).Distinct().Order("chipset").Pluck("chipset", &chipsets).Error; err != nil {
		return nil, err
	}

	return chipsets, nil
}

func (r *MotherboardRepository) GetSocket(ctx context.Context) ([]string, error) {
	var sockets []string

	if err := r.db.WithContext(ctx).Model(&orm.Motherboard{}).Distinct().Order("socket").Pluck("socket", &sockets).Error; err != nil {
		return nil, err
	}

	return sockets, nil
}

func (r *MotherboardRepository) GetFormFactor(ctx context.Context) ([]string, error) {
	var formFactors []string

	if err := r.db.WithContext(ctx).Model(&orm.Motherboard{}).Distinct().Order("form_factor").Pluck("form_factor", &formFactors).Error; err != nil {
		return nil, err
	}

	return formFactors, nil
}
