package services

import (
	"github.com/mohammad-ebrahimi-it/car-shoping/api/dto"
	"github.com/mohammad-ebrahimi-it/car-shoping/common"
	"github.com/mohammad-ebrahimi-it/car-shoping/config"
	"github.com/mohammad-ebrahimi-it/car-shoping/constans"
	"github.com/mohammad-ebrahimi-it/car-shoping/data/db"
	"github.com/mohammad-ebrahimi-it/car-shoping/data/models"
	"github.com/mohammad-ebrahimi-it/car-shoping/pkg/logging"
	"github.com/mohammad-ebrahimi-it/car-shoping/pkg/service_errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	logger       logging.Logger
	cfg          *config.Config
	OtpService   *OtpService
	database     *gorm.DB
	tokenService *TokenService
}

func NewUserService(cfg *config.Config) *UserService {
	database := db.GetDb()
	logger := logging.NewLogger(cfg)

	return &UserService{logger: logger,
		cfg:          cfg,
		database:     database,
		OtpService:   NewOtpService(cfg),
		tokenService: NewTokenService(cfg),
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

func (s *UserService) RegisterByUsername(request *dto.RegisterUserByUsernameRequest) error {
	u := models.User{
		Username:  request.Username,
		Email:     request.Email,
		FirstName: request.FirstName,
		LastName:  request.LastName,
	}

	exists, err := s.existsByEmail(u.Email)

	if err != nil {
		return err
	}

	if exists {
		return &service_errors.ServiceError{
			EndUserMessage: service_errors.EmailExists,
		}
	}

	exists, err = s.existsByUsername(u.Username)

	if err != nil {
		return err
	}

	if exists {
		return &service_errors.ServiceError{
			EndUserMessage: service_errors.UsernameExists,
		}
	}

	byte_password := []byte(request.Password)

	hash_password, err := bcrypt.GenerateFromPassword(byte_password, bcrypt.DefaultCost)

	if err != nil {
		s.logger.Error(logging.General, logging.HashPassword, err.Error(), nil)
		return err
	}
	u.Password = string(hash_password)

	roleId, err := s.getDefaultRole()

	if err != nil {
		s.logger.Error(logging.General, logging.DefaultRoleNotFound, err.Error(), nil)
		return err
	}

	tx := s.database.Begin()

	err = tx.Create(&u).Error

	if err != nil {
		s.logger.Error(logging.General, logging.Rollback, err.Error(), nil)
		tx.Rollback()
		return err
	}

	err = tx.Create(&models.UserRole{RoleId: roleId, UserId: u.Id}).Error

	if err != nil {
		s.logger.Error(logging.General, logging.Rollback, err.Error(), nil)
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil

}

func (s *UserService) RegisterByMobileNumber(req *dto.RegisterLoginByMobileNumber) (*dto.TokenDetail, error) {
	err := s.OtpService.ValidateOtp(req.MobileNumber, req.Otp)

	if err != nil {
		return nil, err
	}

	exists, err := s.existsByPhoneNumber(req.MobileNumber)

	if err != nil {
		return nil, err
	}

	u := &models.User{Mobile: req.MobileNumber, Username: req.MobileNumber}

	if exists {
		var user models.User

		err = s.database.
			Model(&models.User{}).
			Where("mobile = ?", u.Mobile).
			Preload("UserRoles", func(tx *gorm.DB) *gorm.DB {
				return tx.Preload("Role")
			}).
			Find(&user).Error

		if err != nil {
			return nil, err
		}

		tdot := tokenDto{
			UserId:       user.Id,
			FirstName:    user.FirstName,
			LastName:     user.LastName,
			Email:        user.Mobile,
			MobileNumber: user.Mobile,
		}

		if len(*user.UserRoles) > 0 {
			for _, ur := range *user.UserRoles {
				tdot.Roles = append(tdot.Roles, ur.Role.Name)
			}
		}

		token, err := s.tokenService.GenerateToken(&tdot)

		if err != nil {
			return nil, err
		}

		return token, nil
	}
	byte_password := []byte(common.GeneratePassword())

	hash_password, err := bcrypt.GenerateFromPassword(byte_password, bcrypt.DefaultCost)

	if err != nil {
		s.logger.Error(logging.General, logging.HashPassword, err.Error(), nil)
		return nil, err
	}
	u.Password = string(hash_password)

	roleId, err := s.getDefaultRole()

	if err != nil {
		s.logger.Error(logging.General, logging.DefaultRoleNotFound, err.Error(), nil)
		return nil, err
	}

	tx := s.database.Begin()

	err = tx.Create(&u).Error

	if err != nil {
		s.logger.Error(logging.General, logging.Rollback, err.Error(), nil)
		tx.Rollback()
		return nil, err
	}

	err = tx.Create(&models.UserRole{RoleId: roleId, UserId: u.Id}).Error

	if err != nil {
		s.logger.Error(logging.General, logging.Rollback, err.Error(), nil)
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	var user models.User
	err = s.database.
		Model(&models.User{}).
		Where("username = ?", u.Mobile).
		Preload("UserRoles", func(tx *gorm.DB) *gorm.DB {
			return tx.Preload("Role")
		}).
		Find(&user).Error

	if err != nil {
		return nil, err
	}

	tdot := tokenDto{
		UserId:       user.Id,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Email:        user.Mobile,
		MobileNumber: user.Mobile,
	}

	if len(*user.UserRoles) > 0 {
		for _, ur := range *user.UserRoles {
			tdot.Roles = append(tdot.Roles, ur.Role.Name)
		}
	}

	token, err := s.tokenService.GenerateToken(&tdot)

	if err != nil {
		return nil, err
	}

	return token, nil

}

func (s *UserService) LoginByUsername(req *dto.LoginByUsernameRequest) (*dto.TokenDetail, error) {
	var user models.User
	err := s.database.
		Model(&models.User{}).
		Where("username = ?", req.Username).
		Preload("UserRoles", func(tx *gorm.DB) *gorm.DB {
			return tx.Preload("Role")
		}).
		Find(&user).Error

	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	if err != nil {
		return nil, err
	}

	tdot := tokenDto{
		UserId:       user.Id,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Email:        user.Mobile,
		MobileNumber: user.Mobile,
	}

	if len(*user.UserRoles) > 0 {
		for _, ur := range *user.UserRoles {
			tdot.Roles = append(tdot.Roles, ur.Role.Name)
		}
	}

	token, err := s.tokenService.GenerateToken(&tdot)

	if err != nil {
		return nil, err
	}

	return token, nil
}

func (s *UserService) existsByEmail(email string) (bool, error) {
	var exists bool
	if err := s.database.Model(&models.User{}).
		Select("count(*) > 0").
		Where("email = ?", email).
		Find(&exists).
		Error; err != nil {
		return false, err
	}
	return exists, nil
}

func (s *UserService) existsByUsername(username string) (bool, error) {
	var exists bool
	if err := s.database.Model(&models.User{}).
		Select("count(*) > 0").
		Where("username = ?", username).
		Find(&exists).
		Error; err != nil {
		return false, err
	}
	return exists, nil
}

func (s *UserService) existsByPhoneNumber(mobileNumber string) (bool, error) {
	var exists bool
	if err := s.database.Model(&models.User{}).
		Select("count(*) > 0").
		Where("mobile = ?", mobileNumber).
		Find(&exists).
		Error; err != nil {
		return false, err
	}
	return exists, nil
}

func (s *UserService) getDefaultRole() (roleId int, err error) {
	if err = s.database.Model(&models.Role{}).
		Select("id").
		Where("name = ?", constans.DefaultRoleName).
		First(&roleId).
		Error; err != nil {
		return 0, nil
	}

	return roleId, nil
}
