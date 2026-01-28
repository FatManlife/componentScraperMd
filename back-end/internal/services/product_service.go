package service

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	repo "github.com/FatManlife/component-finder/back-end/internal/repositories"
)


type ProductService struct {
	repo *repo.ProductRepository
}

func NewProductService(repo *repo.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetProducts(ctx context.Context, limit int, website string, after int, brand string, min float64, max float64, order string) ([]dto.ProductResponse, error) {
	if min > max {
		min = 0
		max = 0
	}

	products, err := s.repo.GetAllProducts(ctx, limit, website, after, brand, min, max, order)

	if err != nil {
		return nil,  err
	}

	var productResponses []dto.ProductResponse

	for _, product := range products {
		productResponses = append(productResponses, productMapping(product))
	}

	return productResponses, nil
}


