package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mohammad-ebrahimi-it/car-shoping/api/dto"
	_ "github.com/mohammad-ebrahimi-it/car-shoping/api/helper"
	"github.com/mohammad-ebrahimi-it/car-shoping/config"
	"github.com/mohammad-ebrahimi-it/car-shoping/services"
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
	Create(c, h.service.Create)
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
	Update(c, h.service.Update)
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
	Delete(c, h.service.Delete)
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
	GetById(c, h.service.GetByID)
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
	GetByFilter(c, h.service.GetByFilter)
}
