package services

import (
	"context"
	"database/sql"
	"github.com/mohammad-ebrahimi-it/car-shoping/api/dto"
	"github.com/mohammad-ebrahimi-it/car-shoping/config"
	"github.com/mohammad-ebrahimi-it/car-shoping/constans"
	"github.com/mohammad-ebrahimi-it/car-shoping/data/db"
	"github.com/mohammad-ebrahimi-it/car-shoping/data/models"
	"github.com/mohammad-ebrahimi-it/car-shoping/pkg/logging"
	"gorm.io/gorm"
	"time"
)

type CountryService struct {
	database *gorm.DB
	logger   logging.Logger
}

func NewCountryService(cfg *config.Config) *CountryService {
	return &CountryService{
		database: db.GetDb(),
		logger:   logging.NewLogger(cfg),
	}
}

func (s *CountryService) Create(ctx context.Context, req *dto.CreateUpdateCountryRequest) (*dto.CreateCountryResponse, error) {
	country := models.Country{
		Name: req.Name,
	}

	country.CreatedBy = int(ctx.Value(constans.UserIdKey).(float64))
	country.CreatedAt = time.Now()

	tx := s.database.WithContext(ctx).Begin()

	err := tx.Create(&country).Error

	if err != nil {
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Insert, err.Error(), nil)

		return nil, err
	}

	tx.Commit()

	dto := &dto.CreateCountryResponse{Name: country.Name, Id: country.Id}

	return dto, nil
}

func (s *CountryService) Update(ctx context.Context, id int, request *dto.CreateUpdateCountryRequest) (*dto.CreateCountryResponse, error) {
	updateMap := map[string]interface{}{
		"name":        request.Name,
		"modified_by": &sql.NullInt64{Int64: int64(ctx.Value(constans.UserIdKey).(float64)), Valid: true},
		"modified_at": &sql.NullTime{Time: time.Now().UTC(), Valid: true},
	}

	tx := s.database.WithContext(ctx).Begin()

	err := tx.Model(&models.Country{}).
		Where("id = ?", id).
		Updates(updateMap).Error

	if err != nil {
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Update, err.Error(), nil)
		return nil, err
	}

	country := &models.Country{}

	err = tx.Model(&models.Country{}).
		Where("id = ? AND deleted_by is null ", id).
		First(&country).Error

	if err != nil {
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)

		return nil, err
	}

	tx.Commit()
	dto := &dto.CreateCountryResponse{Name: country.Name, Id: country.Id}

	return dto, nil
}

func (s *CountryService) Delete(ctx context.Context, id int) error {
	tx := s.database.WithContext(ctx).Begin()

	deleteMap := map[string]interface{}{
		"deleted_by": &sql.NullInt64{Int64: int64(ctx.Value(constans.UserIdKey).(float64)), Valid: true},
		"deleted_at": &sql.NullTime{Time: time.Now().UTC(), Valid: true},
	}

	err := tx.Model(&models.Country{}).
		Where("id = ?", id).
		Updates(&deleteMap).
		Error

	if err != nil {
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Delete, err.Error(), nil)
		return err
	}

	tx.Commit()

	return nil
}

func (s *CountryService) GetByID(ctx context.Context, id int) (*dto.CountryResponse, error) {
	country := &models.Country{}

	err := s.database.
		Where("id = ? AND deleted_at is null ", id).
		First(country).Error

	if err != nil {
		s.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return nil, err
	}

	dto := &dto.CountryResponse{Name: country.Name, Id: country.Id}

	return dto, nil
}
