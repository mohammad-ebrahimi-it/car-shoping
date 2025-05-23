package services

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/mohammad-ebrahimi-it/car-shoping/api/dto"
	"github.com/mohammad-ebrahimi-it/car-shoping/config"
	"github.com/mohammad-ebrahimi-it/car-shoping/constans"
	"github.com/mohammad-ebrahimi-it/car-shoping/pkg/logging"
	"github.com/mohammad-ebrahimi-it/car-shoping/pkg/service_errors"
	"time"
)

type TokenService struct {
	logger logging.Logger
	cfg    *config.Config
}

type tokenDto struct {
	UserId       int
	FirstName    string
	LastName     string
	Username     string
	Email        string
	MobileNumber string
	Roles        []string
}

func NewTokenService(cfg *config.Config) *TokenService {
	logger := logging.NewLogger(cfg)
	return &TokenService{
		logger: logger,
		cfg:    cfg,
	}
}

func (s *TokenService) GenerateToken(token *tokenDto) (*dto.TokenDetail, error) {
	accessToken := &dto.TokenDetail{}

	accessToken.AccessTokenExpireTime = time.Now().Add(s.cfg.JWT.AccessTokenExpireDuration * time.Minute).Unix()
	accessToken.RefreshTokenExpireTime = time.Now().Add(s.cfg.JWT.RefreshTokenExpireDuration * time.Minute).Unix()

	accessTokenClaims := jwt.MapClaims{}

	accessTokenClaims[constans.UserIdKey] = token.UserId
	accessTokenClaims[constans.FirstNameKey] = token.FirstName
	accessTokenClaims[constans.LastNameKey] = token.LastName
	accessTokenClaims[constans.UsernameKey] = token.Username
	accessTokenClaims[constans.EmailKey] = token.Email
	accessTokenClaims[constans.RolesKey] = token.Roles
	accessTokenClaims[constans.MobileNumberKey] = token.MobileNumber
	accessTokenClaims[constans.ExpireTimeKey] = accessToken.AccessTokenExpireTime

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)

	var err error

	accessToken.AccessToken, err = at.SignedString([]byte(s.cfg.JWT.Secret))

	if err != nil {
		return nil, err
	}

	refreshTokenClaims := jwt.MapClaims{}
	refreshTokenClaims["user_id"] = token.UserId
	refreshTokenClaims["exp"] = accessToken.RefreshTokenExpireTime

	rf := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	accessToken.RefreshToken, err = rf.SignedString([]byte(s.cfg.JWT.RefreshSecret))

	if err != nil {
		return nil, err
	}

	return accessToken, nil
}

func (s *TokenService) VerifyToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, &service_errors.ServiceError{
				EndUserMessage: service_errors.UnExpectedError,
			}
		}

		return []byte(s.cfg.JWT.Secret), nil
	})
}

func (s *TokenService) GetClaims(token string) (claimMap map[string]interface{}, err error) {
	claimMap = map[string]interface{}{}

	verifyToken, err := s.VerifyToken(token)

	if err != nil {
		return nil, err
	}

	claims, ok := verifyToken.Claims.(jwt.MapClaims)

	if ok && verifyToken.Valid {
		for key, value := range claims {
			claimMap[key] = value
		}
		return claimMap, nil
	}

	return nil, &service_errors.ServiceError{
		EndUserMessage: service_errors.ClaimNotFound,
	}
}
