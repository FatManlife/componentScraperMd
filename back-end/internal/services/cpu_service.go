package service

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	repo "github.com/FatManlife/component-finder/back-end/internal/repositories"
)

type CpuService struct {
	cpuRepo *repo.CpuRepository
}

func NewCpuService(cpuRepo *repo.CpuRepository) *CpuService {
	return &CpuService{cpuRepo: cpuRepo}
}

func (s *CpuService) GetSpecs(ctx context.Context) (dto.CpuSpecs, error)	{
	var specs dto.CpuSpecs

	cores, err := s.cpuRepo.GetCores(ctx)
	if err != nil {
		return specs, err
	}
	specs.Cores = cores

	threads, err := s.cpuRepo.GetThreads(ctx)
	if err != nil {
		return specs, err
	}
	specs.Threads = threads

	baseClocks, err := s.cpuRepo.GetBaseClocks(ctx)
	if err != nil {
		return specs, err
	}
	specs.BaseClock = baseClocks

	boostClocks, err := s.cpuRepo.GetBoostClocks(ctx)
	if err != nil {
		return specs, err
	}
	specs.BoostClock = boostClocks
	
	sockets, err := s.cpuRepo.GetSockets(ctx)	
	if err != nil {
		return specs, err
	}
	specs.Socket = sockets

	return specs, nil
}