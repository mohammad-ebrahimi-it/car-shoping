package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mohammad-ebrahimi-it/car-shoping/api/handlers"
	"github.com/mohammad-ebrahimi-it/car-shoping/config"
)

func City(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewCityHandler(cfg)

	r.POST("/", h.CreateCity)
	r.PUT("/:id", h.UpdateCity)
	r.DELETE("/:id", h.DeleteCity)
	r.GET("/:id", h.GetCityById)
	r.POST("/get-by-filter", h.GetByFilter)
}
