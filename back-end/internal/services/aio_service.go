package service

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	repo "github.com/FatManlife/component-finder/back-end/internal/repositories"
)

type AioService struct {
	repo *repo.AioRepository
}

func NewAioService(repo *repo.AioRepository) *AioService {
	return &AioService{repo: repo}
}

func (s *AioService) GetSpecs(ctx context.Context) (dto.AioSpecs, error){
	var filters dto.AioSpecs
	
	diagonals, err := s.repo.GetDiagonlas(ctx)
	if err != nil {
		return filters, err
	}
	filters.Diagonal = diagonals

	cpus, err := s.repo.GetCpus(ctx)
	if err != nil {
		return filters, err
	}
	filters.Cpu = cpus

	rams, err := s.repo.GetRams(ctx)
	if err != nil {
		return filters, err
	}
	filters.Ram = rams

	storages, err := s.repo.GetStorages(ctx)
	if err != nil {
		return filters, err
	}
	filters.Storage = storages

	gpus, err := s.repo.GetGpus(ctx)
	if err != nil {
		return filters, err
	}
	filters.Gpu = gpus

	return filters, nil
}