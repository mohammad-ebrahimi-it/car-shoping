package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mohammad-ebrahimi-it/car-shoping/api/handlers"
	"github.com/mohammad-ebrahimi-it/car-shoping/config"
)

func File(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewFileHandler(cfg)

	r.POST("/", h.Create)
	r.PUT("/:id", h.UpdateFile)
	r.DELETE("/:id", h.DeleteFile)
	r.GET("/:id", h.GetFileById)
	r.GET("/get-by-filter", h.GetByFilter)
}
