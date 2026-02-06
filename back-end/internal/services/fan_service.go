package service

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	repo "github.com/FatManlife/component-finder/back-end/internal/repositories"
)

type fanService struct {
	repo *repo.FanRepository
}

func NewFanService(repo *repo.FanRepository) *fanService {
	return &fanService{repo: repo}
}

func (s *fanService) GetSpecs(ctx context.Context) (dto.FanSpecs, error) {
	var specs dto.FanSpecs
	
	fanRPMs, err := s.repo.GetFanRPMs(ctx)
	if err != nil {
		return specs, err
	}
	specs.FanRPM = fanRPMs

	noiseLevels, err := s.repo.GetNoiseLevels(ctx)
	if err != nil {
		return specs, err
	}
	specs.Noise = noiseLevels

	return specs, nil
}	