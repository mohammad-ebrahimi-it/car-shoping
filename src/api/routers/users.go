package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mohammad-ebrahimi-it/car-shoping/api/handlers"
	"github.com/mohammad-ebrahimi-it/car-shoping/api/middlewares"
	"github.com/mohammad-ebrahimi-it/car-shoping/config"
)

func User(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewUsersHandler(cfg)

	router.POST("/send-otp", middlewares.OtpLimiter(cfg), h.SendOtp)
}
