package service

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"github.com/FatManlife/component-finder/back-end/internal/models/orm"
	repo "github.com/FatManlife/component-finder/back-end/internal/repositories"
)

type CpuService struct {
	repo *repo.CpuRepository
}

func NewCpuService(repo *repo.CpuRepository) *CpuService {
	return &CpuService{repo: repo}
}

func cpuMapping(product orm.Product) dto.CpuResponse{
	return dto.CpuResponse{
		Product: productMapping(product),
		Cores: product.Cpu.Cores,
		Threads: product.Cpu.Threads,
		BaseClock: product.Cpu.BaseClock,
		BoostClock: product.Cpu.BoostClock,
		Tdp: product.Cpu.Tdp,
		Socket: product.Cpu.Socket,
	}
}

func (s *CpuService) GetCpus(ctx context.Context, params dto.CpuParams) ([]dto.CpuResponse, error) {
	checkDefaultParams(&params.DefaultParams)

	cpus, err := s.repo.GetCpus(ctx, params)

	if err != nil {
		return nil, err
	}

	var cpusResponse []dto.CpuResponse

	for _, cpu := range cpus {
		cpusResponse = append(cpusResponse, cpuMapping(cpu))
	}
	
	return cpusResponse, nil
}