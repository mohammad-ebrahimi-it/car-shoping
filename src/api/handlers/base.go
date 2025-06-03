package handlers

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mohammad-ebrahimi-it/car-shoping/api/helper"
	"github.com/mohammad-ebrahimi-it/car-shoping/config"
	"github.com/mohammad-ebrahimi-it/car-shoping/pkg/logging"
	"net/http"
	"strconv"
)

var logger = logging.NewLogger(config.GetConfig())

func Create[Ti any, To any](c *gin.Context, caller func(ctx context.Context, req *Ti) (*To, error)) {
	req := new(Ti)

	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))

		return
	}

	res, err := caller(c, req)

	if err != nil {
		c.AbortWithStatusJSON(
			helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err),
		)

		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(
		res,
		true,
		helper.Success))
	return
}

func Update[Ti any, To any](c *gin.Context, caller func(ctx context.Context, id int, req *Ti) (*To, error)) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	req := new(Ti)

	err := c.ShouldBindJSON(req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}

	res, err := caller(c, id, req)

	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err), helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, helper.Success))

	return
}

func Delete(c *gin.Context, caller func(ctx context.Context, id int) error) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	if id == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, helper.GenerateBaseResponseWithError(nil, false, helper.NotFoundError, errors.New("id not found")))

		return
	}

	err := caller(c, id)

	if err != nil {
		c.AbortWithStatusJSON(
			helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.ValidationError, err),
		)
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(nil, true, helper.Success))
	return
}

func GetById[To any](c *gin.Context, caller func(ctx context.Context, id int) (*To, error)) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	if id == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, helper.GenerateBaseResponseWithError(nil, false, helper.NotFoundError, errors.New("id not found")))

		return
	}

	res, err := caller(c, id)

	if err != nil {
		c.AbortWithStatusJSON(
			helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.ValidationError, err),
		)

		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, helper.Success))

	return
}
func GetByFilter[Ti any, To any](c *gin.Context, caller func(ctx context.Context, req *Ti) (*To, error)) {
	req := new(Ti)

	err := c.ShouldBindQuery(req)

	if err != nil {
		c.AbortWithStatusJSON(
			helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.ValidationError, err))
		return
	}

	res, err := caller(c, req)

	if err != nil {
		c.AbortWithStatusJSON(
			helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, helper.Success))
}
