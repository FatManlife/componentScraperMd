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

func (s *ProductService) GetProducts(ctx context.Context, params dto.ProductParams) ([]dto.ProductResponse, error) {
	checkDefaultParams(&params)	

	products, err := s.repo.GetAllProducts(ctx, params)

	if err != nil {
		return nil,  err
	}

	var productResponses []dto.ProductResponse

	for _, product := range products {
		productResponses = append(productResponses, productMapping(product))
	}

	return productResponses, nil
}


