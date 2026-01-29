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

func (r *PsuRepository) GetPsus(ctx context.Context, params dto.PsuParams) ([]orm.Product, error) {
	var psus []orm.Product

	q := getDefaultProduct(r.db, ctx, params.DefaultParams)

	q.Joins("JOIN psus on psus.product_id = products.id").Preload("Psu")

	if params.FormFactor != "" {
		q = q.Where("psus.form_factor = ?", params.FormFactor)
	}

	if params.Efficiency != "" {
		q = q.Where("psus.efficiency ILIKE ?", params.Efficiency)
	}

	if params.Power != 0 {
		q = q.Where("psus.power = ?", params.Power)
	}

	if err := q.Find(&psus).Error ; err != nil {
		return nil, err
	}

	return psus, nil
}