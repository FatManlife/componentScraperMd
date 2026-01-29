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

	q.Joins("JOIN cases ON cases.product_id = products.id").Preload("Case")

	if params.Format!= "" {
		q = q.Where("cases.format = ?", params.Format)
	}

	if params.MotherboardFormFactor != "" {
		q = q.Where("cases.motherboard_form_factor = ?", params.MotherboardFormFactor)
	}

	if err := q.Find(&cases).Error; err != nil {
		return nil, err
	}

	return cases, nil
}