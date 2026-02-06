package service

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	repo "github.com/FatManlife/component-finder/back-end/internal/repositories"
)

type hddService struct { 
	hddRepo *repo.HddRepository
}

func NewHddService(hddRepo *repo.HddRepository) *hddService {
	return &hddService{hddRepo: hddRepo}
}

func (s *hddService) GetSpecs(ctx context.Context) (dto.HddSpecs, error) {
	var specs dto.HddSpecs

	capacities, err := s.hddRepo.GetCapacities(ctx)
	if err != nil {
		return specs, err
	}
	specs.Capacity = capacities

	rpms, err := s.hddRepo.GetRotationSpeeds(ctx)
	if err != nil {
		return specs, err
	}
	specs.RotationSpeed = rpms

	formFactors, err := s.hddRepo.GetFormFactors(ctx)
	if err != nil {
		return specs, err
	}
	specs.FormFactor = formFactors

	return specs, nil
}