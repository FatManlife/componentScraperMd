package service

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"github.com/FatManlife/component-finder/back-end/internal/models/orm"
	repo "github.com/FatManlife/component-finder/back-end/internal/repositories"
)

type FanService struct {
	repo *repo.FanRepository
}

func NewFanService(repo *repo.FanRepository) *FanService {
	return &FanService{repo: repo}
}

func fanMapping(product orm.Product) dto.FanResponse{
	return dto.FanResponse{
		Product: productMapping(product),
		FanRPM: product.Fan.FanRPM,
		Size: product.Fan.Size,
		Noise: product.Fan.Noise,
	}
}

func (s *FanService) GetFans(ctx context.Context, params dto.FanParams) ([]dto.FanResponse, error) {
	checkDefaultParams(&params.DefaultParams)

	fans, err := s.repo.GetFans(ctx, params)

	if err != nil {
		return nil, err
	}

	var fansResponse []dto.FanResponse

	for _, fan := range fans {
		fansResponse = append(fansResponse, fanMapping(fan))
	}

	return fansResponse, nil
}	