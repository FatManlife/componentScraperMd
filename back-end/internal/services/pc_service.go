package service

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"github.com/FatManlife/component-finder/back-end/internal/models/orm"
	repo "github.com/FatManlife/component-finder/back-end/internal/repositories"
)

type PcService struct {
	repo *repo.PcRepository
}

func NewPcService(repo *repo.PcRepository) *PcService {
	return &PcService{repo: repo}
}

func PcMapping(product orm.Product) *dto.PcResponse{
	return &dto.PcResponse{
		Product: productMapping(product),
		Case: product.Pc.PcCase,
		Cpu: product.Pc.Cpu,
		Gpu: product.Pc.Gpu,
		Motherboard: product.Pc.Motherboard,
		Psu: product.Pc.Psu,
		Ram: product.Pc.Ram,
		Storage: product.Pc.Storage,
	}
}

func (s *PcService) GetPcs(ctx context.Context, params dto.PcParams) ([]dto.PcResponse, error) {
	checkDefaultParams(&params.DefaultParams)

	pcs, err := s.repo.GetPcs(ctx, params)

	if err != nil {
		return nil,  err
	}

	var pcResponses []dto.PcResponse
	
	for _, pc := range pcs {
		pcResponses = append(pcResponses, *PcMapping(pc))
	}

	return pcResponses, nil
}