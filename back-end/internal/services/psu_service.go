package service

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	repo "github.com/FatManlife/component-finder/back-end/internal/repositories"
)

type PsuService struct {
	repo *repo.PsuRepository
}

func NewPsuService(repo *repo.PsuRepository) *PsuService {
	return &PsuService{repo: repo}
}

func (s *PsuService) GetSpecs(ctx context.Context) (dto.PsuSpecs, error) {
	var specs dto.PsuSpecs

	powers, err := s.repo.GetPower(ctx)
	if err != nil {
		return dto.PsuSpecs{}, err
	}
	specs.Power = powers

	efficiencies, err := s.repo.GetEfficiency(ctx)
	if err != nil {
		return dto.PsuSpecs{}, err
	}
	specs.Efficiency = efficiencies

	formFactors, err := s.repo.GetFormFactor(ctx)
	if err != nil {
		return dto.PsuSpecs{}, err
	}
	specs.FormFactor = formFactors

	return specs, nil
}
	