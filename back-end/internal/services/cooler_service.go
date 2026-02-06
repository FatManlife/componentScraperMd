package service

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	repo "github.com/FatManlife/component-finder/back-end/internal/repositories"
)

type CoolerService struct {
	repo *repo.CoolerRepository
}

func NewCoolerService(repo *repo.CoolerRepository) *CoolerService {
	return &CoolerService{repo: repo}
}

func (s *CoolerService) GetSpecs(ctx context.Context) (dto.CoolerSpecs, error) {
	var specs dto.CoolerSpecs
	
	types, err := s.repo.GetTypes(ctx)
	if err != nil {
		return specs, err
	}
	specs.Type = types

	fanRPMs, err := s.repo.GetFanRPMs(ctx)
	if err != nil {
		return specs, err
	}
	specs.FanRPM = fanRPMs

	noises, err := s.repo.GetNoises(ctx)
	if err != nil {
		return specs, err
	}
	specs.Noise = noises

	compatibilities, err := s.repo.GetCompatibility(ctx)
	if err != nil {
		return specs, err
	}
	specs.Compatibility = compatibilities

	return specs, nil
}