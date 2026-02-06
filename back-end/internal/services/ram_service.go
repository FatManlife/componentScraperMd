package service

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	repo "github.com/FatManlife/component-finder/back-end/internal/repositories"
)

type RamService struct {
	repo *repo.RamRepository
}

func NewRamService(repo *repo.RamRepository) *RamService {
	return &RamService{repo: repo}
}

func (s *RamService) GetSpecs(ctx context.Context) (dto.RamSpecs,error) {
	var specs dto.RamSpecs

	capacities, err := s.repo.GetCapacity(ctx)
	if err != nil {
		return dto.RamSpecs{}, err
	}
	specs.Capacity = capacities

	speeds, err := s.repo.GetSpeed(ctx)
	if err != nil {
		return dto.RamSpecs{}, err
	}
	specs.Speed = speeds

	types, err := s.repo.GetType(ctx)
	if err != nil {
		return dto.RamSpecs{}, err
	}
	specs.Type = types

	compatibilities, err := s.repo.GetCompatibility(ctx)
	if err != nil {
		return dto.RamSpecs{}, err
	}
	specs.Compatibility = compatibilities

	configurations, err := s.repo.GetConfiguration(ctx)
	if err != nil {
		return dto.RamSpecs{}, err
	}
	specs.Configuration = configurations

	return specs, nil
}