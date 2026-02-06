package service

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	repo "github.com/FatManlife/component-finder/back-end/internal/repositories"
)

type SsdService struct {
	repo *repo.SSDRepository
}

func NewSsdService(repo *repo.SSDRepository) *SsdService {
	return &SsdService{repo: repo}
}

func (s *SsdService) GetSpecs(ctx context.Context) (dto.SsdSpecs, error) {
	var specs dto.SsdSpecs

	capacity, err := s.repo.GetCapacity(ctx)
	if err != nil {
		return specs, err
	}
	specs.Capacity = capacity

	readingSpeed, err := s.repo.GetReadingSpeed(ctx)
	if err != nil {
		return specs, err
	}
	specs.ReadingSpeed = readingSpeed

	writingSpeed, err := s.repo.GetWritingSpeed(ctx)
	if err != nil {
		return specs, err
	}
	specs.WritingSpeed = writingSpeed

	formFactor, err := s.repo.GetFormFactor(ctx)
	if err != nil {
		return specs, err
	}
	specs.FormFactor = formFactor

	return specs, nil
}