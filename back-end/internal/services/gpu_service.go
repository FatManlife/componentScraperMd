package service

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"github.com/FatManlife/component-finder/back-end/internal/models/orm"
	repo "github.com/FatManlife/component-finder/back-end/internal/repositories"
)

type GpuService struct {
	repo *repo.GpuRepository 
}

func NewGpuService(repo *repo.GpuRepository) *GpuService {
	return &GpuService{repo: repo}
}

func gpuMapping(product orm.Product) dto.GpuResponse {
	return dto.GpuResponse{
		Product:      productMapping(product),
		Chipset:      product.Gpu.Chipset,
		Vram:         product.Gpu.Vram,
		GpuFrequency: product.Gpu.GpuFrequency,
		VramFrequency: product.Gpu.VramFrequency,
	}
}

func (s *GpuService) GetGpus(ctx context.Context, params dto.GpuParams) ([]dto.GpuResponse, error) {
	checkDefaultParams(&params.DefaultParams)

	gpus, err := s.repo.GetGpus(ctx, params)

	if err != nil {
		return nil, err
	}
	var gpuResponses []dto.GpuResponse

	for _, gpu := range gpus {
		gpuResponses = append(gpuResponses, gpuMapping(gpu))
	}

	return gpuResponses, nil
}