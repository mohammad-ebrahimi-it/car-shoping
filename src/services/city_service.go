package services

import (
	"context"
	"github.com/mohammad-ebrahimi-it/car-shoping/api/dto"
	"github.com/mohammad-ebrahimi-it/car-shoping/config"
	"github.com/mohammad-ebrahimi-it/car-shoping/data/db"
	"github.com/mohammad-ebrahimi-it/car-shoping/data/models"
	"github.com/mohammad-ebrahimi-it/car-shoping/pkg/logging"
)

type CityService struct {
	base *BaseService[models.City, dto.CreateUpdateCityRequest, dto.CreateUpdateCityRequest, dto.CityResponse]
}

func NewCityService(cfg *config.Config) *CityService {

	return &CityService{
		base: &BaseService[
			models.City,
			dto.CreateUpdateCityRequest,
			dto.CreateUpdateCityRequest,
			dto.CityResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{
				{string: "Country"},
			},
		},
	}
}

func (s *CityService) Create(ctx context.Context, req *dto.CreateUpdateCityRequest) (*dto.CityResponse, error) {
	return s.base.Create(ctx, req)
}

func (s *CityService) Update(ctx context.Context, id int, request *dto.CreateUpdateCityRequest) (*dto.CityResponse, error) {
	return s.base.Update(ctx, id, request)
}

func (s *CityService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

func (s *CityService) GetByID(ctx context.Context, id int) (*dto.CityResponse, error) {
	return s.base.GetById(ctx, id)
}

func (s *CityService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PageList[dto.CityResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
