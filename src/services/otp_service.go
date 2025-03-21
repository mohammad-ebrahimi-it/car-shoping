package services

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/mohammad-ebrahimi-it/car-shoping/config"
	"github.com/mohammad-ebrahimi-it/car-shoping/constans"
	"github.com/mohammad-ebrahimi-it/car-shoping/data/cache"
	"github.com/mohammad-ebrahimi-it/car-shoping/pkg/logging"
	"github.com/mohammad-ebrahimi-it/car-shoping/pkg/service_errors"
	"time"
)

type OtpService struct {
	logger logging.Logger
	cfg    *config.Config
	redis  *redis.Client
}

type OtpDto struct {
	Value string
	Used  bool
}

func NewOtpService(cfg *config.Config) *OtpService {
	logger := logging.NewLogger(cfg)
	redis := cache.GetRedis()
	return &OtpService{logger: logger, cfg: cfg, redis: redis}
}

func (s *OtpService) SetOptService(mobileNumber, otp string) error {
	key := fmt.Sprintf("%s:%s", constans.RedisOtpDefaultKey, mobileNumber)

	val := &OtpDto{
		Value: otp,
		Used:  false,
	}

	res, err := cache.Get[OtpDto](s.redis, key)

	if err == nil && !res.Used {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpExists}
	} else if err == nil && res.Used {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpUsed}
	}

	err = cache.Set(s.redis, key, val, s.cfg.Otp.ExpireTime*time.Second)

	if err != nil {
		return err
	}

	return nil
}

func (s *OtpService) ValidateOtp(mobileNumber, otp string) error {
	key := fmt.Sprintf("%s:%s", constans.RedisOtpDefaultKey, mobileNumber)

	res, err := cache.Get[OtpDto](s.redis, key)

	if err != nil {
		return err
	} else if res.Used {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpUsed}
	} else if res.Value != otp {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpInValid}
	} else if res.Value == otp && !res.Used {
		res.Used = true
		err = cache.Set(s.redis, key, res, s.cfg.Otp.ExpireTime*time.Second)

		return err
	}

	return nil
}
