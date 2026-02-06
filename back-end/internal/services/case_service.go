package service

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	repo "github.com/FatManlife/component-finder/back-end/internal/repositories"
)

type CaseService struct {
	repo *repo.CaseRepository
}

func NewCaseService(repo *repo.CaseRepository) *CaseService {
	return &CaseService{repo: repo}
}

func (s *CaseService) GetSpecs(ctx context.Context) (dto.CaseSpecs, error) {
	var specs dto.CaseSpecs

	formats, err := s.repo.GetFormats(ctx)
	if err != nil {
		return specs, err
	}

	motherboardFormFactors, err := s.repo.GetMotherboardFormFactors(ctx)
	if err != nil {
		return specs, err
	}

	specs.Format = formats
	specs.MotherboardFormFactor = motherboardFormFactors

	return specs, nil

}