package repo

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"github.com/FatManlife/component-finder/back-end/internal/models/orm"
	"gorm.io/gorm"
)

type PsuRepository struct {
	db *gorm.DB
}

func NewPsuRepository(db *gorm.DB) *PsuRepository {
	return &PsuRepository{db: db}
}

func (r *PsuRepository) GetPsus(ctx context.Context, psuParams *dto.PsuParams) ([]orm.Product, error) {
	var psus []orm.Product

	q := getDefaultProduct(r.db, ctx, psuParams.DefaultParams)

	q.Joins("JOIN psus on psus.product_id = products.id").Preload("Psu")

	if psuParams.FormFactor != "" {
		q = q.Where("psus.form_factor = ?", psuParams.FormFactor)
	}

	if psuParams.Efficiency != "" {
		q = q.Where("psus.efficiency = ?", psuParams.Efficiency)
	}

	if psuParams.Power != 0 {
		q = q.Where("psus.power = ?", psuParams.Power)
	}

	err := q.Find(&psus).Error

	if err != nil {
		return nil, err
	}

	return psus, nil
}