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

func (r *CaseRepository) GetCases(ctx context.Context, params dto.CaseParams) ([]orm.Product, error){
	var cases []orm.Product

	q := getDefaultProduct(r.db, ctx, params.DefaultParams)

	q.Joins("JOIN pc_cases ON pc_cases.product_id = products.id")

	if len(params.Format) > 0 {
		q = q.Where("cases.format IN ?", params.Format)
	}

	if len(params.MotherboardFormFactor) > 0 {
		q = q.Where("cases.motherboard_form_factor IN ?", params.MotherboardFormFactor)
	}

	if err := q.Find(&cases).Error; err != nil {
		return nil, err
	}

	return cases, nil
}