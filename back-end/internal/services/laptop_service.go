package service

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	repo "github.com/FatManlife/component-finder/back-end/internal/repositories"
)

type LaptopService struct {
	repo *repo.LaptopRepository
}

func NewLaptopService(repo *repo.LaptopRepository) *LaptopService {
	return &LaptopService{repo: repo}
}

func (s *LaptopService) GetSpecs(ctx context.Context) (dto.LaptopSpecs, error) {
	var specs dto.LaptopSpecs
	
	cpus, err := s.repo.GetCpus(ctx)
	if err != nil {
		return specs, err
	}
	specs.Cpu = cpus

	gpus, err := s.repo.GetGpus(ctx)
	if err != nil {
		return specs, err
	}
	specs.Gpu = gpus

	rams, err := s.repo.GetRams(ctx)
	if err != nil {
		return specs, err
	}
	specs.Ram = rams

	storages, err := s.repo.GetStorages(ctx)
	if err != nil {
		return specs, err
	}
	specs.Storage = storages

	diagonals, err := s.repo.GetDiagonals(ctx)
	if err != nil {
		return specs, err
	}
	specs.Diagonal = diagonals

	return specs, nil
}
