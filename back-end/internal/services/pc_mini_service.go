package service

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"github.com/FatManlife/component-finder/back-end/internal/models/orm"
	repo "github.com/FatManlife/component-finder/back-end/internal/repositories"
)

type PcMiniService struct {
	repo *repo.PcMiniRepository
}

func NewPcMiniService(repo *repo.PcMiniRepository) *PcMiniService {
	return &PcMiniService{repo: repo}
}

func PcMiniMapping(product orm.Product) *dto.PcMiniResponse{
	return &dto.PcMiniResponse{
		Product: productMapping(product),
		Cpu: product.PcMini.Cpu,
		Gpu: product.PcMini.Gpu,
		Ram: product.PcMini.Ram,
		Storage: product.PcMini.Storage,
	}
}

func (s *PcMiniService) GetPcMinis(ctx context.Context, params dto.PcMiniParams) ([]dto.PcMiniResponse, error) {
	checkDefaultParams(&params.DefaultParams)

	pcMinis, err := s.repo.GetPcMinis(ctx, params)

	if err != nil {
		return nil,  err
	}

	var pcMinisResponses []dto.PcMiniResponse

	for _, pcMini := range pcMinis {
		pcMinisResponses = append(pcMinisResponses, *PcMiniMapping(pcMini))
	}

	return pcMinisResponses, nil
}