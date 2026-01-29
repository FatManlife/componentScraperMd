package service

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"github.com/FatManlife/component-finder/back-end/internal/models/orm"
	repo "github.com/FatManlife/component-finder/back-end/internal/repositories"
)

type SSDService struct {
	repo *repo.SSDRepository
}

func NewSSDService(repo *repo.SSDRepository) *SSDService {
	return &SSDService{repo: repo}
}

func ssdsMapping(product orm.Product) dto.SsdResponse {
	return dto.SsdResponse{
		Product: productMapping(product),
		Capacity: product.Ssd.Capacity,
		ReadingSpeed: product.Ssd.ReadingSpeed,
		WritingSpeed: product.Ssd.WritingSpeed,
		FormFactor: product.Ssd.FormFactor,
	}
}

func (s *SSDService) GetSsds(ctx context.Context, ssdParams dto.SsdParams) ([]dto.SsdResponse, error) {
	checkDefaultParams(&ssdParams.DefaultParams)
	
	ssds, err := s.repo.GetSsds(ctx, ssdParams)

	if err != nil {
		return nil, err
	}

	var ssdsResponses []dto.SsdResponse

	for _, ssd := range ssds {
		ssdsResponses = append(ssdsResponses, ssdsMapping(ssd))
	}

	return ssdsResponses, nil
}