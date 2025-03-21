package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mohammad-ebrahimi-it/car-shoping/api/dto"
	"github.com/mohammad-ebrahimi-it/car-shoping/api/helper"
	"github.com/mohammad-ebrahimi-it/car-shoping/config"
	"github.com/mohammad-ebrahimi-it/car-shoping/services"
	"net/http"
)

type UsersHandler struct {
	service *services.UserService
}

func NewUsersHandler(cfg *config.Config) *UsersHandler {
	service := services.NewUserService(cfg)
	return &UsersHandler{service: service}
}

// SendOtp user
// @Summary Send Otp To User
// @Description send otp to user
// @Tags users
// @Accept json
// @Produce json
// @Param Request body dto.GetOtpRequest true "get otp request"
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failure"
// @Failure 409 {object} helper.BaseHttpResponse "Failure"
// @Router /v1/users/send-otp/ [post]
func (uh *UsersHandler) SendOtp(c *gin.Context) {
	otp := new(dto.GetOtpRequest)
	err := c.ShouldBindJSON(&otp)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, -1, err),
		)
		return
	}

	err = uh.service.SendOtp(otp)

	if err != nil {
		c.AbortWithStatusJSON(
			helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, -1, err),
		)
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(nil, true, 0))
}
