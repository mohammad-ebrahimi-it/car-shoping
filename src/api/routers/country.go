package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mohammad-ebrahimi-it/car-shoping/api/handlers"
	"github.com/mohammad-ebrahimi-it/car-shoping/config"
)

func Country(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewCountryHandler(cfg)

	r.POST("/", h.CreateCountry)
	r.PUT("/:id", h.UpdateCountry)
	r.DELETE("/:id", h.DeleteCountry)
	r.GET("/:id", h.GetCountryById)
}
