package service

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	repo "github.com/FatManlife/component-finder/back-end/internal/repositories"
)

type MotherboardService struct {
	repo *repo.MotherboardRepository
}

func NewMotherboardService(repo *repo.MotherboardRepository) *MotherboardService {
	return &MotherboardService{repo: repo}
}

func (s *MotherboardService) GetSpecs(ctx context.Context) (dto.MotherboardSpecs, error) {
	var mb dto.MotherboardSpecs

	chipsets, err := s.repo.GetChipset(ctx)
	if err != nil {
		return dto.MotherboardSpecs{}, err
	}
	mb.Chipset = chipsets

	sockets, err := s.repo.GetSocket(ctx)
	if err != nil {
		return dto.MotherboardSpecs{}, err
	}
	mb.Socket = sockets

	formFactors, err := s.repo.GetFormFactor(ctx)
	if err != nil {
		return dto.MotherboardSpecs{}, err
	}
	mb.FormFactor = formFactors

	return mb, nil
}
