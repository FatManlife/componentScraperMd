package service

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"github.com/FatManlife/component-finder/back-end/internal/models/orm"
	repo "github.com/FatManlife/component-finder/back-end/internal/repositories"
)

type RamService struct {
	repo *repo.RamRepository
}

func NewRamService(repo *repo.RamRepository) *RamService {
	return &RamService{repo: repo}
}

func ramMapping(product orm.Product) dto.RamResponse {
	return dto.RamResponse{
		Product: productMapping(product),
		Capacity: product.Ram.Capacity,
		Speed: product.Ram.Speed,
		Type: product.Ram.Type,
		Compatibility: product.Ram.Compatibility,
		Configuration: product.Ram.Configuration,
	}
}

func (s *RamService) GetRams(ctx context.Context, ramParams dto.RamParams) ([]dto.RamResponse, error) {
	checkDefaultParams(&ramParams.DefaultParams)

	rams, err := s.repo.GetRams(ctx, ramParams)

	if err != nil {
		return nil, err
	}

	var ramsResponses []dto.RamResponse

	for _, ram := range rams {
		ramsResponses = append(ramsResponses, ramMapping(ram))
	}

	return ramsResponses, nil
}