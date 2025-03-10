package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/mohammad-ebrahimi-it/car-shoping/api/routers"
	"github.com/mohammad-ebrahimi-it/car-shoping/api/validations"
	"github.com/mohammad-ebrahimi-it/car-shoping/config"
)

func InitServer(cfg *config.Config) {
	r := gin.New()
	RegisterValidators()

	r.Use(gin.Logger(), gin.Recovery())

	RegisterRoutes(r)

	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}

func RegisterRoutes(r *gin.Engine) {

	v1 := r.Group("/api/v1")

	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, "hi mohammad")
			return
		})

		health := v1.Group("/health")

		routers.Health(health)
	}
}

func RegisterValidators() {
	val, ok := binding.Validator.Engine().(*validator.Validate)

	if ok {
		val.RegisterValidation("mobile", validations.IranianMobileNumberValidator, true)
	}
}
