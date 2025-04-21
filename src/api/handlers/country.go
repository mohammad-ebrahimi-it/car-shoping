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

type CountryHandler struct {
	service *services.CountryService
}

func NewCountryHandler(cfg *config.Config) *CountryHandler {

	return &CountryHandler{
		service: services.NewCountryService(cfg),
	}
}

// CreateCountry godoc
// @Summary Create a Country
// @Description Create a Country
// @Tags Countries
// @Accept json
// @Produce json
// @Param Request body dto.CreateUpdateCountryRequest true "Create a Country"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CountryResponse} "Country response"
// @Failure 400 {object} helper.BaseHttpResponse "Failure"
// @Router /v1/countries [post]
// @Security AuthBearer
func (h *CountryHandler) CreateCountry(c *gin.Context) {
	req := &dto.CreateUpdateCountryRequest{}

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

// UpdateCountry godoc
// @Summary Update a Country
// @Description Update a Country
// @Tags Countries
// @Accept json
// @Produce json
// @Param id  path int true "Id"
// @Param Request body dto.CreateUpdateCountryRequest true "Update a Country"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CountryResponse} "Country response"
// @Failure 400 {object} helper.BaseHttpResponse "Failure"
// @Router /v1/countries/{id} [put]
// @Security AuthBearer
func (h *CountryHandler) UpdateCountry(c *gin.Context) {

	id, _ := strconv.Atoi(c.Params.ByName("id"))

	req := &dto.CreateUpdateCountryRequest{}

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

// DeleteCountry godoc
// @Summary Delete a Country
// @Description Delete a Country
// @Tags Countries
// @Accept json
// @Produce json
// @Param id  path int true "Id"
// @Success 201 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failure"
// @Router /v1/countries/{id} [delete]
// @Security AuthBearer
func (h *CountryHandler) DeleteCountry(c *gin.Context) {
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

// GetCountryById godoc
// @Summary Get a Country
// @Description Get a Country
// @Tags Countries
// @Accept json
// @Produce json
// @Param id  path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CountryResponse} "Country response"
// @Failure 400 {object} helper.BaseHttpResponse "Failure"
// @Router /v1/countries/{id} [get]
// @Security AuthBearer
func (h *CountryHandler) GetCountryById(c *gin.Context) {
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
// @Summary Get Countries
// @Description Get Countries
// @Tags Countries
// @Accept json
// @Produce json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PageList[dto.CountryResponse]} "Country response"
// @Failure 400 {object} helper.BaseHttpResponse "Failure"
// @Router /v1/countries/get-by-filter [post]
// @Security AuthBearer
func (h *CountryHandler) GetByFilter(c *gin.Context) {
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
