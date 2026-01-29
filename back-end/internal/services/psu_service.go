package service

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"github.com/FatManlife/component-finder/back-end/internal/models/orm"
	repo "github.com/FatManlife/component-finder/back-end/internal/repositories"
)

type PsuService struct {
	repo *repo.PsuRepository
}

func NewPsuService(repo *repo.PsuRepository) *PsuService {
	return &PsuService{repo: repo}
}

func psuMapping(product orm.Product) dto.PsuResponse {
	return dto.PsuResponse{
		Product:    productMapping(product),
		FormFactor: product.Psu.FormFactor,
		Efficiency: product.Psu.Efficiency,
		Power:      product.Psu.Power,
	}
}

func (s *PsuService) GetPsus(ctx context.Context, psuParams *dto.PsuParams) ([]dto.PsuResponse, error) {
	checkDefaultParams(&psuParams.DefaultParams)

	psus, err := s.repo.GetPsus(ctx, psuParams)

	if err != nil {
		return nil, err
	}

	var psuResponses []dto.PsuResponse

	for _, psu := range psus {
		psuResponses = append(psuResponses, psuMapping(psu))
	}

	return psuResponses, nil
}