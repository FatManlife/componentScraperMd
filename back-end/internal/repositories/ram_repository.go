package repo

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"github.com/FatManlife/component-finder/back-end/internal/models/orm"
	"gorm.io/gorm"
)

type RamRepository struct {
	db *gorm.DB
}

func NewRamRepository(db *gorm.DB) *RamRepository {
	return &RamRepository{db: db}
}

func (r *RamRepository) GetRams(ctx context.Context, params dto.RamParams, ) ([]orm.Product, int64, error) {
	var ram []orm.Product
	var count int64

	q := getDefaultProduct(r.db, ctx, params.DefaultParams)

	q.Joins("JOIN rams on rams.product_id = products.id").Preload("Ram")

	if params.MinCapacity > 0{
		q = q.Where("rams.capacity >= ?", params.MinCapacity)
	}

	if params.MaxCapacity > 0{
		q = q.Where("rams.capacity <= ?", params.MaxCapacity)
	}

	if params.MinSpeed > 0{
		q = q.Where("rams.speed >= ?", params.MinSpeed)
	}

	if params.MaxSpeed > 0{
		q = q.Where("rams.speed <= ?", params.MaxSpeed)
	}

	if len(params.Type) > 0{ 
		q = q.Where("rams.type IN ?", params.Type)
	}

	if len(params.Compatibility) > 0{
		q = q.Where("rams.compatibility IN ?", params.Compatibility)
	}

	if len(params.Configuration) > 0{ 
		q = q.Where("rams.configuration IN ?", params.Configuration)
	}

	if err := q.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	setLimits(q, params.DefaultParams.Offset)

	if err := q.Find(&ram).Error; err != nil {
		return nil, 0, err
	}

	return ram, count, nil
}

func (r *RamRepository) GetCapacity(ctx context.Context) ([]int, error){
	var capacities []int

	if err := r.db.Model(&orm.Ram{}).Distinct().Pluck("capacity", &capacities).Error; err != nil {
		return nil, err
	}

	return capacities, nil
}

func (r *RamRepository) GetSpeed(ctx context.Context) ([]int, error){
	var speeds []int

	if err := r.db.Model(&orm.Ram{}).Distinct().Pluck("speed", &speeds).Error; err != nil {
		return nil, err
	}

	return speeds, nil
}

func (r *RamRepository) GetType(ctx context.Context) ([]string, error){
	var types []string

	if err := r.db.Model(&orm.Ram{}).Distinct().Pluck("type", &types).Error; err != nil {
		return nil, err
	}

	return types, nil
}

func (r *RamRepository) GetCompatibility(ctx context.Context) ([]string, error){
	var compatibilities []string	

	if err := r.db.Model(&orm.Ram{}).Distinct().Pluck("compatibility", &compatibilities).Error; err != nil {
		return nil, err
	}

	return compatibilities, nil
}

func (r *RamRepository) GetConfiguration(ctx context.Context) ([]string, error){
	var configurations []string

	if err := r.db.Model(&orm.Ram{}).Distinct().Pluck("configuration", &configurations).Error; err != nil {
		return nil, err
	}

	return configurations, nil
}