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

func (r *SSDRepository) GetSsds(ctx context.Context, params dto.SsdParams) ([]orm.Product, int64, error) {
	var ssds []orm.Product
	var count int64

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

	if err := q.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	setLimits(q, params.DefaultParams.Offset)

	if err := q.Find(&ssds).Error; err != nil {
		return nil, 0, err
	}

	return ssds, count, nil
}

func (r *SSDRepository) GetCapacity(ctx context.Context) ([]int, error) {
	var capacities []int

	if err := r.db.WithContext(ctx).Model(&orm.Ssd{}).Distinct().Pluck("capacity", &capacities).Error; err != nil {
		return nil, err
	}

	return capacities, nil
}

func (r *SSDRepository) GetReadingSpeed(ctx context.Context) ([]int, error) {
	var speeds []int

	if err := r.db.WithContext(ctx).Model(&orm.Ssd{}).Distinct().Pluck("reading_speed", &speeds).Error; err != nil {
		return nil, err
	}

	return speeds, nil
}

func (r *SSDRepository) GetWritingSpeed(ctx context.Context) ([]int, error) {
	var speeds []int

	if err := r.db.WithContext(ctx).Model(&orm.Ssd{}).Distinct().Pluck("writing_speed", &speeds).Error; err != nil {
		return nil, err
	}

	return speeds, nil
}

func (r *SSDRepository) GetFormFactor(ctx context.Context) ([]string, error) {
	var formFactors []string

	if err := r.db.WithContext(ctx).Model(&orm.Ssd{}).Distinct().Pluck("form_factor", &formFactors).Error; err != nil {
		return nil, err
	}

	return formFactors, nil
}