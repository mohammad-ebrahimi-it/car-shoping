package services

import (
	"context"
	"github.com/mohammad-ebrahimi-it/car-shoping/api/dto"
	"github.com/mohammad-ebrahimi-it/car-shoping/config"
	"github.com/mohammad-ebrahimi-it/car-shoping/data/db"
	"github.com/mohammad-ebrahimi-it/car-shoping/data/models"
	"github.com/mohammad-ebrahimi-it/car-shoping/pkg/logging"
)

type CountryService struct {
	base *BaseService[
		models.Country,
		dto.CreateUpdateCountryRequest,
		dto.CreateUpdateCountryRequest,
		dto.CountryResponse,
	]
}

func NewCountryService(cfg *config.Config) *CountryService {
	return &CountryService{
		base: &BaseService[
			models.Country,
			dto.CreateUpdateCountryRequest,
			dto.CreateUpdateCountryRequest,
			dto.CountryResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
		},
	}
}

func (s *CountryService) Create(ctx context.Context, req *dto.CreateUpdateCountryRequest) (*dto.CountryResponse, error) {
	return s.base.Create(ctx, req)

}

func (s *CountryService) Update(ctx context.Context, id int, request *dto.CreateUpdateCountryRequest) (*dto.CountryResponse, error) {
	return s.base.Update(ctx, id, request)
}

func (s *CountryService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

func (s *CountryService) GetByID(ctx context.Context, id int) (*dto.CountryResponse, error) {
	return s.base.GetById(ctx, id)
}
