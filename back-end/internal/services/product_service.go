package service

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"github.com/FatManlife/component-finder/back-end/internal/models/orm"
	repo "github.com/FatManlife/component-finder/back-end/internal/repositories"
)

func productMapping(products []orm.Product) []dto.ProductResponse{
	productList := []dto.ProductResponse{}

	for _, product := range products {
		prod := dto.ProductResponse{
			ID: product.ID,
			Name: product.Name,
			ImageURL: product.ImageURL,
			Brand: product.Brand,
			Price: product.Price,
			Url: product.URL,
			Website_id: product.WebsiteID,
			Category_id: product.CategoryID,
		}	

		productList = append(productList, prod)
	}

	return productList		
}


type ProductService struct {
	repo *repo.ProductRepository
}

func NewProductService(repo *repo.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAllProducts(ctx context.Context, limit int) ([]dto.ProductResponse, error) {
	pr, err := s.repo.GetAllProducts(ctx, limit)

	if err != nil {
		return nil,  err
	}

	return productMapping(pr), nil
}