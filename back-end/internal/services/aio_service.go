package service

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"github.com/FatManlife/component-finder/back-end/internal/models/orm"
	repo "github.com/FatManlife/component-finder/back-end/internal/repositories"
)

type AioService struct {
	repo *repo.AioRepository
}

func NewAioService(repo *repo.AioRepository) *AioService {
	return &AioService{repo: repo}
}

func aiosMapping(product orm.Product) dto.AioResponse{
	return dto.AioResponse{
		Product: productMapping(product),
		Diagonal: product.Aio.Diagonal,
		Gpu: product.Aio.Gpu,
		Cpu: product.Aio.Cpu,
		Ram: product.Aio.Ram,
		Storage: product.Aio.Storage,
	}
}

func (s *AioService) GetAios(ctx context.Context, aioParams dto.AioParams) ([]dto.AioResponse, error) {
	checkDefaultParams(&aioParams.DefaultParams)
	
	aios, err := s.repo.GetAios(ctx, aioParams)

	if err != nil {
		return nil,  err
	}

	var aiosResponses []dto.AioResponse

	for _, aio := range aios {
		aiosResponses = append(aiosResponses, aiosMapping(aio))
	}

	return aiosResponses, nil
}