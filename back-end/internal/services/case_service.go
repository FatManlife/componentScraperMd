package service

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"github.com/FatManlife/component-finder/back-end/internal/models/orm"
	repo "github.com/FatManlife/component-finder/back-end/internal/repositories"
)

type CaseService struct {
	repo *repo.CaseRepository
}

func NewCaseService(repo *repo.CaseRepository) *CaseService {
	return &CaseService{repo: repo}
}

func caseMapping(product orm.Product) dto.CaseResponse{
	return dto.CaseResponse{
		Product: productMapping(product),
		Format: product.Case.Format,
		MotherboardFormFactor: product.Case.MotherboardFormFactor,
	}
}

func (s *CaseService) GetCases(ctx context.Context, params dto.CaseParams) ([]dto.CaseResponse, error) {
	checkDefaultParams(&params.DefaultParams)

	cases, err := s.repo.GetCases(ctx, params)

	if err != nil {
		return nil, err
	}

	var casesResponse []dto.CaseResponse

	for _, caseItem := range cases {
		casesResponse = append(casesResponse, caseMapping(caseItem))
	}

	return casesResponse, nil
}