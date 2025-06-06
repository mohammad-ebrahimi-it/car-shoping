package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/mohammad-ebrahimi-it/car-shoping/api/middlewares"
	"github.com/mohammad-ebrahimi-it/car-shoping/api/routers"
	"github.com/mohammad-ebrahimi-it/car-shoping/api/validations"
	"github.com/mohammad-ebrahimi-it/car-shoping/config"
	"github.com/mohammad-ebrahimi-it/car-shoping/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitServer(cfg *config.Config) {
	r := gin.New()
	RegisterValidators()

	r.Use(gin.Logger(), gin.CustomRecovery(middlewares.ErrorHandler))
	r.Use(middlewares.DefaultStructuredLogger(cfg))

	RegisterRoutes(r, cfg)
	RegisterSwagger(r, cfg)

	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}

func RegisterRoutes(r *gin.Engine, cfg *config.Config) {

	v1 := r.Group("/api/v1")

	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, "hi mohammad")
			return
		})

		health := v1.Group("/health")

		routers.Health(health)

		users := v1.Group("/users")
		routers.User(users, cfg)

		countries := v1.Group("/countries", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		routers.Country(countries, cfg)

		cities := v1.Group("/cities", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		routers.City(cities, cfg)

		files := v1.Group("/files", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		routers.File(files, cfg)
	}
}

func RegisterValidators() {
	val, ok := binding.Validator.Engine().(*validator.Validate)

	if ok {
		val.RegisterValidation("mobile", validations.IranianMobileNumberValidator, true)
		val.RegisterValidation("password", validations.PasswordValidator, true)
	}
}

func RegisterSwagger(r *gin.Engine, cfg *config.Config) {
	docs.SwaggerInfo.Title = "golang web api"
	docs.SwaggerInfo.Description = "golang web api"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", cfg.Server.Port)
	docs.SwaggerInfo.Schemes = []string{"http"}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
