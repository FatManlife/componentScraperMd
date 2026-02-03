package service

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	repo "github.com/FatManlife/component-finder/back-end/internal/repositories"
	"github.com/FatManlife/component-finder/back-end/internal/utils"
)


type ProductService struct {
	repo *repo.ProductRepository
}

func NewProductService(repo *repo.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetProducts(ctx context.Context, params dto.ProductParams) ([]dto.ProductsResponse, error) {
	products, err := s.repo.GetAllProducts(ctx, params)

	if err != nil {
		return nil,  err
	}

	var productResponses []dto.ProductsResponse

	for _, product := range products {
		productResponses = append(productResponses, utils.ProductsMapping(product))
	}

	return productResponses, nil
}


func (s *ProductService) GetProductsCount(ctx context.Context, category string)(int64, error){
	return s.repo.GetProductsCount(ctx, category)
}