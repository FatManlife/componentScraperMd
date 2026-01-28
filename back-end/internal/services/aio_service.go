package service

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/orm"
	repo "github.com/FatManlife/component-finder/back-end/internal/repositories"
)

type AioService struct {
	repo *repo.AioRepository
}

func NewAioService(repo *repo.AioRepository) *AioService {
	return &AioService{repo: repo}
}

func (s *AioService) GetAios(ctx context.Context, limit int, website string, after int, brand string, min float64, max float64, order string) ([]orm.Product, error) {
	if min > max {
		min = 0
		max = 0
	}

	aios, err := s.repo.GetAios(ctx, limit, website, after, brand, min, max, order)

	if err != nil {
		return nil,  err
	}

	return aios, nil
}