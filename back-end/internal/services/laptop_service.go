package service

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"github.com/FatManlife/component-finder/back-end/internal/models/orm"
	repo "github.com/FatManlife/component-finder/back-end/internal/repositories"
)


type LaptopService struct {
	repo *repo.LaptopRepository
}

func NewLaptopService(laptopRepository *repo.LaptopRepository) *LaptopService {
	return &LaptopService{repo: laptopRepository}
}

func laptopMapping(product orm.Product) dto.LaptopResponse {
	return dto.LaptopResponse{
		Product: productMapping(product),
		Cpu: product.Laptop.Cpu,
		Gpu: product.Laptop.Gpu,
		Ram: product.Laptop.Ram,
		Storage: product.Laptop.Storage,
		Diagonal: product.Laptop.Diagonal,
		Battery: product.Laptop.Battery,
	}
}

func (s *LaptopService) GetLaptops(ctx context.Context, params dto.LaptopParams) ([]dto.LaptopResponse, error) {
	checkDefaultParams(&params.DefaultParams)

	laptops, err := s.repo.GetLaptops(ctx, params)

	if err != nil {
		return nil, err
	}

	var laptopResponse []dto.LaptopResponse

	for _, laptop := range laptops {
		laptopResponse = append(laptopResponse, laptopMapping(laptop))
	}

	return laptopResponse, nil
}