package service

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	repo "github.com/FatManlife/component-finder/back-end/internal/repositories"
)

type PcService struct {
	repo *repo.PcRepository
}

func NewPcService(repo *repo.PcRepository) *PcService {
	return &PcService{repo: repo}
}	

func (s *PcService) GetSpecs(context.Context)(dto.PcSpecs, error) {
	var pc dto.PcSpecs

	cpus, err := s.repo.GetCpu(context.Background())
	if err != nil {
		return dto.PcSpecs{}, err
	}
	pc.Cpu = cpus

	gpus, err := s.repo.GetGpu(context.Background())
	if err != nil {
		return dto.PcSpecs{}, err
	}
	pc.Gpu = gpus

	rams, err := s.repo.GetRam(context.Background())
	if err != nil {
		return dto.PcSpecs{}, err
	}
	pc.Ram = rams

	storages, err := s.repo.GetStorage(context.Background())
	if err != nil {
		return dto.PcSpecs{}, err
	}
	pc.Storage = storages

	return pc, nil
}

