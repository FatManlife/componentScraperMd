package service

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"github.com/FatManlife/component-finder/back-end/internal/models/orm"
	repo "github.com/FatManlife/component-finder/back-end/internal/repositories"
)

type HddService struct {
	repo *repo.HddRepository
}

func NewHddService(hddRepo *repo.HddRepository) *HddService {
	return &HddService{repo: hddRepo}
}

func hddMapping(product orm.Product) dto.HddResponse {
	return dto.HddResponse{
		Product: productMapping(product),
		Capacity: product.Hdd.Capacity,
		RotationSpeed: product.Hdd.RotationSpeed,
		FormFactor: product.Hdd.FormFactor,
	}
}	

func (s *HddService) GetHdds(ctx context.Context, params dto.HddParams) ([]dto.HddResponse, error) {
	checkDefaultParams(&params.DefaultParams)

	hdds, err := s.repo.GetHdds(ctx, params)

	if err != nil {
		return nil, err
	}

	var hddResponse []dto.HddResponse

	for _, hdd := range hdds {
		hddResponse = append(hddResponse, hddMapping(hdd))
	}

	return hddResponse, nil
}