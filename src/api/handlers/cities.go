package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mohammad-ebrahimi-it/car-shoping/api/dto"
	"github.com/mohammad-ebrahimi-it/car-shoping/api/helper"
	"github.com/mohammad-ebrahimi-it/car-shoping/config"
	"github.com/mohammad-ebrahimi-it/car-shoping/services"
	"net/http"
	"strconv"
)

type CityHandler struct {
	service *services.CityService
}

func NewCityHandler(cfg *config.Config) *CityHandler {
	return &CityHandler{
		service: services.NewCityService(cfg),
	}
}

// CreateCity godoc
// @Summary Create a City
// @Description Create a City
// @Tags Cities
// @Accept json
// @Produce json
// @Param Request body dto.CreateUpdateCityRequest true "Create a City"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CityResponse} "City response"
// @Failure 400 {object} helper.BaseHttpResponse "Failure"
// @Router /v1/cities [post]
// @Security AuthBearer
func (h *CityHandler) CreateCity(c *gin.Context) {
	req := &dto.CreateUpdateCityRequest{}

	err := c.ShouldBindJSON(req)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, -400, err))

		return
	}

	res, err := h.service.Create(c, req)

	if err != nil {
		c.AbortWithStatusJSON(
			helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, -400, err),
		)

		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(
		res,
		true,
		0))
	return
}

// UpdateCity godoc
// @Summary Update a City
// @Description Update a City
// @Tags Cities
// @Accept json
// @Produce json
// @Param id  path int true "Id"
// @Param Request body dto.CreateUpdateCityRequest true "Update a City"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CityResponse} "City response"
// @Failure 400 {object} helper.BaseHttpResponse "Failure"
// @Router /v1/cities/{id} [put]
// @Security AuthBearer
func (h *CityHandler) UpdateCity(c *gin.Context) {

	id, _ := strconv.Atoi(c.Params.ByName("id"))

	req := &dto.CreateUpdateCityRequest{}

	err := c.ShouldBindJSON(req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, -400, err))
		return
	}

	res, err := h.service.Update(c, id, req)

	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err), helper.GenerateBaseResponseWithError(nil, false, -400, err))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, 0))

	return
}

// DeleteCity godoc
// @Summary Delete a City
// @Description Delete a City
// @Tags Cities
// @Accept json
// @Produce json
// @Param id  path int true "Id"
// @Success 201 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failure"
// @Router /v1/cities/{id} [delete]
// @Security AuthBearer
func (h *CityHandler) DeleteCity(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	if id == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, helper.GenerateBaseResponseWithError(nil, false, -404, errors.New("id not found")))

		return
	}

	err := h.service.Delete(c, id)

	if err != nil {
		c.AbortWithStatusJSON(
			helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, -400, err),
		)
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(nil, true, 0))
	return
}

// GetCityById godoc
// @Summary Get a City
// @Description Get a City
// @Tags Cities
// @Accept json
// @Produce json
// @Param id  path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CityResponse} "City response"
// @Failure 400 {object} helper.BaseHttpResponse "Failure"
// @Router /v1/cities/{id} [get]
// @Security AuthBearer
func (h *CityHandler) GetCityById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	if id == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, helper.GenerateBaseResponseWithError(nil, false, -404, errors.New("id not found")))

		return
	}

	res, err := h.service.GetByID(c, id)

	if err != nil {
		c.AbortWithStatusJSON(
			helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, -400, err),
		)

		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, 0))

	return
}

// GetByFilter godoc
// @Summary Get Cities
// @Description Get Cities
// @Tags Cities
// @Accept json
// @Produce json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PageList[dto.CityResponse]} "City response"
// @Failure 400 {object} helper.BaseHttpResponse "Failure"
// @Router /v1/cities/get-by-filter [post]
// @Security AuthBearer
func (h *CityHandler) GetByFilter(c *gin.Context) {
	req := dto.PaginationInputWithFilter{}

	err := c.ShouldBindQuery(&req)

	if err != nil {
		c.AbortWithStatusJSON(
			helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, -400, err))
		return
	}

	res, err := h.service.GetByFilter(c, &req)

	if err != nil {
		c.AbortWithStatusJSON(
			helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, -400, err))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, 0))
}
