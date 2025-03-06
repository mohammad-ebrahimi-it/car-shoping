package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mohammad-ebrahimi-it/car-shoping/api/routers"
	"github.com/mohammad-ebrahimi-it/car-shoping/config"
)

func InitServer() {
	config := config.GetConfig()
	r := gin.New()
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
