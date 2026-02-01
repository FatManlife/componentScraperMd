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
func (r *ProductRepository) GetAllProducts(ctx context.Context, params dto.ProductParams) ([]orm.Product, error) {
	var products []orm.Product

	q := getDefaultProduct(r.db, ctx, params) 	
	
	if err := q.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) GetProductByID(ctx context.Context, id int) (*orm.Product, error) {
	var product orm.Product

	q := r.db.WithContext(ctx).Model(&orm.Product{}).Where("id = ?", id)

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

