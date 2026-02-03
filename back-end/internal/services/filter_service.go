package service

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	repo "github.com/FatManlife/component-finder/back-end/internal/repositories"
)

type FilterService struct {
	website *repo.WebsiteRepository
	product *repo.ProductRepository
}

func NewFilterService(website *repo.WebsiteRepository, product *repo.ProductRepository) *FilterService {
	return &FilterService{website: website, product: product}
}

func (s *FilterService) GetDefaultFilters(ctx context.Context,category string) (dto.DefaultFilters, error){
	websites, err := s.website.GetAllWebsites(ctx)

	if err != nil {
		return dto.DefaultFilters{}, err
	}

	prices , err := s.product.GetAllPrices(ctx, category)

	if err != nil {
		return dto.DefaultFilters{}, err
	}

	return dto.DefaultFilters{
		Websites: websites,
		Prices: prices,
		Order: []string{"price_asc", "price_desc"},
	}, nil
}