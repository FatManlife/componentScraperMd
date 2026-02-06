package repo

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"github.com/FatManlife/component-finder/back-end/internal/models/orm"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

//implemnt Getting all products with filters
func (r *ProductRepository) GetAllProducts(ctx context.Context, params dto.ProductParams) ([]orm.Product, int64, error) {
	var products []orm.Product
	var count int64

	q := getDefaultProduct(r.db, ctx, params) 	

	if err := q.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	setLimits(q, params.Offset)
	
	if err := q.Find(&products).Error; err != nil {
		return nil, 0, err
	}

	return products, count, nil
}

func (r *ProductRepository) GetProductByID(ctx context.Context, id int) (*orm.Product, error) {
	var product orm.Product

	q := r.db.WithContext(ctx).Model(&orm.Product{}).Where("id = ?", id).Preload("Website")

	if err := q.First(&product).Error; err != nil {
		return nil, err
	}

	switch  product.Category {
		case "ssd":
			q.Preload("Ssd")	
		case "ram":
			q.Preload("Ram")	
		case "psu":
			q.Preload("Psu")	
		case "pc_mini":
			q.Preload("PcMini")	
		case "pc":
			q.Preload("Pc")	
		case "motherboard":
			q.Preload("Motherboard")	
		case "aio":
			q.Preload("Aio")
		case "laptop":
			q.Preload("Laptop")
		case "gpu":
			q.Preload("Gpu")
		case "cpu":
			q.Preload("Cpu")
		case "fan":
			q.Preload("Fan")
		case "case":
			q.Preload("Case")
		case "hdd":
			q.Preload("Hdd")
		case "cooler":
			q.Preload("Cooler"). Preload("Cooler.Cpus.Compatibility")
	}

	if err := q.First(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepository) GetAllPrices(ctx context.Context, category string) ([]float64, error) {
	var prices []float64

	q := r.db.WithContext(ctx).Model(&orm.Product{}).Select("DISTINCT price").Order("price ASC")

	if category != "" {
		q = q.Where("category = ?", category)
	}

	if err := q.Pluck("price", &prices).Error; err != nil {
		return nil, err
	}

	return prices, nil
}