package service

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	repo "github.com/FatManlife/component-finder/back-end/internal/repositories"
)

type GpuService struct {
	repo *repo.GpuRepository
}

func NewGpuService(repo *repo.GpuRepository) *GpuService {
	return &GpuService{repo: repo}
}

func (s *GpuService) GetSpecs(ctx context.Context) (dto.GpuSpecs, error) {
	var specs dto.GpuSpecs
	
	chipsets, err := s.repo.GetChipsets(ctx)
	if err != nil {
		return specs, err
	}
	specs.Chipset = chipsets

	vram, err := s.repo.GetVrams(ctx)
	if err != nil {
		return specs, err
	}
	specs.Vram = vram

	gpuFrequencies, err := s.repo.GetGpuFrequencies(ctx)
	if err != nil {
		return specs, err
	}
	specs.GpuFrequency = gpuFrequencies

	vramFrequencies, err := s.repo.GetVramFrequencies(ctx)
	if err != nil {
		return specs, err
	}
	specs.VramFrequency = vramFrequencies

	return specs, nil
}