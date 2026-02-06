package repo

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"github.com/FatManlife/component-finder/back-end/internal/models/orm"
	"gorm.io/gorm"
)

type CaseRepository struct {
	db *gorm.DB
}

func NewCaseRepository(db *gorm.DB) *CaseRepository {
	return &CaseRepository{db: db}
}

func (r *CaseRepository) GetCases(ctx context.Context, params dto.CaseParams) ([]orm.Product, int64, error){
	var cases []orm.Product
	var count int64

	q := getDefaultProduct(r.db, ctx, params.DefaultParams)

	q.Joins("JOIN pc_cases ON pc_cases.product_id = products.id")

	if len(params.Format) > 0 {
		q = q.Where("pc_cases.format IN ?", params.Format)
	}

	if len(params.MotherboardFormFactor) > 0 {
		q = q.Where("pc_cases.motherboard_form_factor IN ?", params.MotherboardFormFactor)
	}

	if err := q.Count(&count).Error; err != nil {
		return nil, 0, err
	}	

	setLimits(q, params.DefaultParams.Offset)

	if err := q.Find(&cases).Error; err != nil {
		return nil, 0, err
	}

	return cases, count, nil
}

func (r *CaseRepository) GetFormats(ctx context.Context) ([]string, error) {
	var formats []string
	if err := r.db.WithContext(ctx).Model(&orm.PcCase{}).Distinct().Pluck("format", &formats).Error; err != nil {
		return nil, err
	}

	return formats, nil
}

func (r *CaseRepository) GetMotherboardFormFactors(ctx context.Context) ([]string, error) {
	var formFactors []string
	if err := r.db.WithContext(ctx).Model(&orm.PcCase{}).Distinct().Pluck("motherboard_form_factor", &formFactors).Error; err != nil {
		return nil, err
	}

	return formFactors, nil
}