package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mohammad-ebrahimi-it/car-shoping/api/dto"
	_ "github.com/mohammad-ebrahimi-it/car-shoping/api/helper"
	"github.com/mohammad-ebrahimi-it/car-shoping/config"
	"github.com/mohammad-ebrahimi-it/car-shoping/services"
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
	Create(c, h.service.Create)
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
	Update(c, h.service.Update)
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
	Delete(c, h.service.Delete)
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
	GetById(c, h.service.GetByID)
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
	GetByFilter(c, h.service.GetByFilter)
}
