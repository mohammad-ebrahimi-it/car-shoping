package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/mohammad-ebrahimi-it/car-shoping/api/helper"
	"github.com/mohammad-ebrahimi-it/car-shoping/config"
	"github.com/mohammad-ebrahimi-it/car-shoping/constans"
	"github.com/mohammad-ebrahimi-it/car-shoping/pkg/service_errors"
	"github.com/mohammad-ebrahimi-it/car-shoping/services"
	"net/http"
	"strings"
)

func Authentication(cfg *config.Config) gin.HandlerFunc {
	tokenService := services.NewTokenService(cfg)
	return func(c *gin.Context) {
		var err error
		claimMap := map[string]interface{}{}

		auth := c.GetHeader(constans.AuthorizationHeaderKey)

		if auth == "" {
			err = &service_errors.ServiceError{
				EndUserMessage: service_errors.TokenRequired,
			}
		} else {

			token := strings.Split(auth, " ")
			claimMap, err = tokenService.GetClaims(token[1])

			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					err = &service_errors.ServiceError{
						EndUserMessage: service_errors.TokenExpired,
					}
				default:
					err = &service_errors.ServiceError{
						EndUserMessage: service_errors.TokenInvalid,
					}
				}
			}
		}

		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				helper.GenerateBaseResponseWithError(nil, false, -401, err),
			)
			return
		}

		c.Set(constans.UserIdKey, claimMap[constans.UserIdKey])
		c.Set(constans.FirstNameKey, claimMap[constans.FirstNameKey])
		c.Set(constans.LastNameKey, claimMap[constans.LastNameKey])
		c.Set(constans.EmailKey, claimMap[constans.EmailKey])
		c.Set(constans.MobileNumberKey, claimMap[constans.MobileNumberKey])
		c.Set(constans.RolesKey, claimMap[constans.RolesKey])
		c.Set(constans.ExpireTimeKey, claimMap[constans.ExpireTimeKey])
		c.Set(constans.UsernameKey, claimMap[constans.UsernameKey])

		c.Next()

	}
}

func Authorization(validRoles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.Keys) == 0 {
			c.AbortWithStatusJSON(
				http.StatusForbidden,
				helper.GenerateBaseResponse(nil, false, -403),
			)

			return
		}

		rolesVal := c.Keys[constans.RolesKey]

		fmt.Println(rolesVal)

		if rolesVal == nil {
			c.AbortWithStatusJSON(
				http.StatusForbidden,
				helper.GenerateBaseResponse(nil, false, -403),
			)

			return
		}

		roles := rolesVal.([]interface{})

		val := map[string]int{}

		for _, role := range roles {
			val[role.(string)] = 0
		}

		for _, role := range validRoles {
			if _, ok := val[role]; !ok {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(
			http.StatusForbidden,
			helper.GenerateBaseResponse(nil, false, -403),
		)
		return
	}
}
