package service

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"github.com/FatManlife/component-finder/back-end/internal/models/orm"
	repo "github.com/FatManlife/component-finder/back-end/internal/repositories"
)

type MotherboardService struct {
	repo *repo.MotherboardRepository
}

func NewMotherboardService(mbRepo *repo.MotherboardRepository) *MotherboardService {
	return &MotherboardService{repo: mbRepo}
}

func mbMapping(product orm.Product) dto.MotherboardResponse{
	return dto.MotherboardResponse{
		Product: productMapping(product),
		Chipset: product.Motherboard.Chipset,
		Socket: product.Motherboard.Socket,
		FormFactor: product.Motherboard.FormFactor,
		RamSupport: product.Motherboard.RamSupport,
		FormFactorRam: product.Motherboard.FormFactorRam,
	}
}

func (s *MotherboardService) GetMotherboards(ctx context.Context, params dto.MotherboardParams) ([]dto.MotherboardResponse, error) {
	checkDefaultParams(&params.DefaultParams)

	motherboards, err := s.repo.GetMotherboards(ctx, params)

	if err != nil {
		return nil, err
	}

	var mbResponse []dto.MotherboardResponse

	for _, mb := range motherboards {
		mbResponse = append(mbResponse, mbMapping(mb))
	}

	return mbResponse, nil
}