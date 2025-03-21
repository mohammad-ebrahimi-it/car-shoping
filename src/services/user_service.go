package services

import (
	"github.com/mohammad-ebrahimi-it/car-shoping/api/dto"
	"github.com/mohammad-ebrahimi-it/car-shoping/common"
	"github.com/mohammad-ebrahimi-it/car-shoping/config"
	"github.com/mohammad-ebrahimi-it/car-shoping/data/db"
	"github.com/mohammad-ebrahimi-it/car-shoping/pkg/logging"
	"gorm.io/gorm"
)

type UserService struct {
	logger     logging.Logger
	cfg        *config.Config
	OtpService *OtpService
	database   *gorm.DB
}

func NewUserService(cfg *config.Config) *UserService {
	database := db.GetDb()
	logger := logging.NewLogger(cfg)

	return &UserService{logger: logger,
		cfg:        cfg,
		database:   database,
		OtpService: NewOtpService(cfg),
	}
}

func (us *UserService) SendOtp(req *dto.GetOtpRequest) error {
	otp := common.GenerateOtp()

	err := us.OtpService.SetOptService(req.MobileNumber, otp)

	if err != nil {
		return err
	}

	return nil
}
