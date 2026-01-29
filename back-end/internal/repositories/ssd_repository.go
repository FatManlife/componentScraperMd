package repo

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"github.com/FatManlife/component-finder/back-end/internal/models/orm"
	"gorm.io/gorm"
)

type SSDRepository struct {
	db *gorm.DB
}

func NewSSDRepository(db *gorm.DB) *SSDRepository {
	return &SSDRepository{db: db}
}

func (r *SSDRepository) GetSsds(ctx context.Context, params dto.SsdParams) ([]orm.Product, error) {
	var ssds []orm.Product

	q := getDefaultProduct(r.db, ctx, params.DefaultParams) 

	q.Joins("JOIN ssds on ssds.product_id = products.id").Preload("Ssd")

	if params.FormFactor != "" {
		q = q.Where("ssds.form_factor = ?", params.FormFactor)
	}

	if params.Capacity != 0 {
		q = q.Where("ssds.capacity = ?", params.Capacity)
	}

	if params.ReadingSpeed != 0 {
		q = q.Where("ssds.reading_speed = ?", params.ReadingSpeed)
	}

	if params.WritingSpeed != 0 {
		q = q.Where("ssds.writing_speed = ?", params.WritingSpeed)
	}

	if err := q.Find(&ssds).Error; err != nil {
		return nil, err
	}

	return ssds, nil
}