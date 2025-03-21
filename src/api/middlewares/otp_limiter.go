package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mohammad-ebrahimi-it/car-shoping/api/helper"
	"github.com/mohammad-ebrahimi-it/car-shoping/config"
	"github.com/mohammad-ebrahimi-it/car-shoping/pkg/limiter"
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

func OtpLimiter(cfg *config.Config) gin.HandlerFunc {
	var lm = limiter.NewIPRateLimiter(rate.Every(cfg.Otp.Limiter*time.Second), 1)
	return func(c *gin.Context) {

		limit := lm.GetLimiter(c.Request.RemoteAddr)

		if !limit.Allow() {
			c.AbortWithStatusJSON(
				http.StatusTooManyRequests,
				helper.GenerateBaseResponseWithError(nil, false, -1, errors.New("not allowed")),
			)
			c.Abort()
		}

		c.Next()
	}
}
