package service

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	repo "github.com/FatManlife/component-finder/back-end/internal/repositories"
	"github.com/FatManlife/component-finder/back-end/internal/utils"
)


type ProductService struct {
	product *repo.ProductRepository
	website *repo.WebsiteRepository
}

func NewProductService(product *repo.ProductRepository, website *repo.WebsiteRepository) *ProductService {
	return &ProductService{product: product, website: website}
}

func (s *ProductService) GetProducts(ctx context.Context, params dto.ProductParams) ([]dto.ProductsResponse, int64, error) {
	products, count, err := s.product.GetAllProducts(ctx, params)

	if err != nil {
		return nil,  0, err
	}

	var productResponses []dto.ProductsResponse

	for _, product := range products {
		productResponses = append(productResponses, utils.ProductsMapping(product))
	}

	return productResponses, count, nil
}

func (s *ProductService) GetDefaultSpecs(ctx context.Context,category string) (dto.DefaultSpecs, error){
	websites, err := s.website.GetAllWebsites(ctx)

	if err != nil {
		return dto.DefaultSpecs{}, err
	}

	prices , err := s.product.GetAllPrices(ctx, category)

	if err != nil {
		return dto.DefaultSpecs{}, err
	}

	return dto.DefaultSpecs{
		Websites: websites,
		Prices: prices,
		Order: []string{"price_asc", "price_desc"},
	}, nil
}

