package service

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"github.com/FatManlife/component-finder/back-end/internal/models/orm"
	"github.com/FatManlife/component-finder/back-end/internal/utils"
)

type ComponentService[T any, P any] struct {
	getAll func (context.Context, P) ([]orm.Product, error)
	getByID func (context.Context, int) (*orm.Product, error)
	mapper func(orm.Product) T
}

func NewComponentService[T any, P any](getAll func (ctx context.Context, params P) ([]orm.Product, error), getByID func (ctx context.Context, id int) (*orm.Product, error), mapper func(orm.Product) T) *ComponentService[T, P] {
	return &ComponentService[T, P]{getAll: getAll, getByID: getByID, mapper: mapper}
}

func (s *ComponentService[T, P]) GetComponents (ctx context.Context, params P) ([]dto.ProductsResponse, error) {
	products, err := s.getAll(ctx, params)

	if err != nil {
		return nil, err
	}

	var mappedProducts []dto.ProductsResponse

	for _, product := range products {
		mappedProducts = append(mappedProducts, utils.ProductsMapping(product))
	}

	return mappedProducts, nil
}

func (s *ComponentService[T, P]) GetComponentByID (ctx context.Context, id int) (*T, error) {
	product, err := s.getByID(ctx, id)

	if err != nil {
		return nil, err
	}

	mappedProduct := s.mapper(*product)

	return &mappedProduct, nil
}