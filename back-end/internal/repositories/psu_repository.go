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

func (r *PsuRepository) GetPsus(ctx context.Context, params dto.PsuParams) ([]orm.Product, int64, error) {
	var psus []orm.Product
	var count int64

	q := getDefaultProduct(r.db, ctx, params.DefaultParams)

	q.Joins("JOIN psus on psus.product_id = products.id").Preload("Psu")

	if len(params.FormFactor) > 0 {
		q = q.Where("psus.form_factor IN ?", params.FormFactor)
	}

	if len(params.Efficiency) > 0 {
		q = q.Where("psus.efficiency IN ?", params.Efficiency)
	}

	if params.MinPower != 0 {
		q = q.Where("psus.power >= ?", params.MinPower)
	}

	if params.MaxPower != 0 {
		q = q.Where("psus.power <= ?", params.MaxPower)
	}

	if err := q.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	setLimits(q, params.DefaultParams.Offset)

	if err := q.Find(&psus).Error ; err != nil {
		return nil, 0, err
	}

	return psus, count, nil
}

func (r *PsuRepository) GetPower(ctx context.Context) ([]int, error){
	var powers []int

	if err := r.db.Model(&orm.Psu{}).Distinct().Pluck("power", &powers).Error; err != nil {
		return nil, err
	}

	return powers, nil
}

func (r *PsuRepository) GetEfficiency(ctx context.Context) ([]string, error){
	var efficiencies []string	

	if err := r.db.Model(&orm.Psu{}).Distinct().Pluck("efficiency", &efficiencies).Error; err != nil {
		return nil, err
	}

	return efficiencies, nil
}

func (r *PsuRepository) GetFormFactor(ctx context.Context) ([]string, error){
	var formFactors []string

	if err := r.db.Model(&orm.Psu{}).Distinct().Pluck("form_factor", &formFactors).Error; err != nil {
		return nil, err
	}

	return formFactors, nil	
}