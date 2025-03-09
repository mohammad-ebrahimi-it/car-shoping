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

func InitServer() {
	config := config.GetConfig()
	r := gin.New()
	val, ok := binding.Validator.Engine().(*validator.Validate)

	if ok {
		val.RegisterValidation("mobile", validations.IranianMobileNumberValidator, true)
	}

	r.Use(gin.Logger(), gin.Recovery())

	v1 := r.Group("/api/v1")

	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, "hi mohammad")
			return
		})

		health := v1.Group("/health")

		routers.Health(health)
	}

	r.Run(fmt.Sprintf(":%s", config.Server.Port))
}
