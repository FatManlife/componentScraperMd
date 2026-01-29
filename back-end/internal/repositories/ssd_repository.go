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

func (r *SSDRepository) GetSsds(ctx context.Context, ssdParams dto.SsdParams) ([]orm.Product, error) {
	var ssds []orm.Product

	q := getDefaultProduct(r.db, ctx, ssdParams.DefaultParams) 

	q.Joins("JOIN ssds on ssds.product_id = products.id").Preload("Ssd")

	if ssdParams.FormFactor != "" {
		q = q.Where("ssds.form_factor = ?", ssdParams.FormFactor)
	}

	if ssdParams.Capacity != 0 {
		q = q.Where("ssds.capacity = ?", ssdParams.Capacity)
	}

	if ssdParams.ReadingSpeed != 0 {
		q = q.Where("ssds.reading_speed = ?", ssdParams.ReadingSpeed)
	}

	if ssdParams.WritingSpeed != 0 {
		q = q.Where("ssds.writing_speed = ?", ssdParams.WritingSpeed)
	}

	err := q.Find(&ssds).Error

	if err != nil {
		return nil, err
	}

	return ssds, nil
}