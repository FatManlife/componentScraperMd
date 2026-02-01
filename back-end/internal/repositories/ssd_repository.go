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

	q.Joins("Join ssds on ssds.product_id = products.id")

	if len(params.FormFactor) != 0 {
		q = q.Where("ssds.form_factor IN ?", params.FormFactor)
	}

	if params.MaxCapacity > 0 {
		q = q.Where("ssds.capacity <= ?", params.MaxCapacity)
	}

	if params.MinCapacity > 0 {
		q = q.Where("ssds.capacity >= ?", params.MinCapacity)
	}

	if params.MaxReadingSpeed > 0 {
		q = q.Where("ssds.reading_speed <= ?", params.MaxReadingSpeed)
	}

	if params.MinReadingSpeed > 0 {
		q = q.Where("ssds.reading_speed >= ?", params.MinReadingSpeed)
	}

	if params.MaxWritingSpeed > 0 {
		q = q.Where("ssds.writing_speed <= ?", params.MaxWritingSpeed)
	}

	if params.MinWritingSpeed > 0 {
		q = q.Where("ssds.writing_speed >= ?", params.MinWritingSpeed)
	}

	if err := q.Find(&ssds).Error; err != nil {
		return nil, err
	}

	return ssds, nil
}