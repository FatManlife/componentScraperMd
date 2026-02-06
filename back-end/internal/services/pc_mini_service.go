package service

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	repo "github.com/FatManlife/component-finder/back-end/internal/repositories"
)

type PcMiniService struct {
	repo *repo.PcMiniRepository
}

func NewPcMiniService(repo *repo.PcMiniRepository) *PcMiniService {
	return &PcMiniService{repo: repo}
}

func (s *PcMiniService) GetSpecs(ctx context.Context) (dto.PcSpecs, error) {
	var pcMini dto.PcSpecs

	cpus, err := s.repo.GetCpu(ctx)
	if err != nil {
		return dto.PcSpecs{}, err
	}
	pcMini.Cpu = cpus

	gpus, err := s.repo.GetGpu(ctx)
	if err != nil {
		return dto.PcSpecs{}, err
	}
	pcMini.Gpu = gpus

	rams, err := s.repo.GetRam(ctx)
	if err != nil {
		return dto.PcSpecs{}, err
	}
	pcMini.Ram = rams

	storages, err := s.repo.GetStorage(ctx)
	if err != nil {
		return dto.PcSpecs{}, err
	}
	pcMini.Storage = storages

	return pcMini, nil
}