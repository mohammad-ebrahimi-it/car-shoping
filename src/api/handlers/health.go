package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mohammad-ebrahimi-it/car-shoping/api/helper"
	"net/http"
)

type HealthHandler struct{}

type Person struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone" binding:"required,mobile"`
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// Health user
// @Summary register
// @Description regiser user
// @Tags register_user
// @Accept json
// @Produce json
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failure"
// @Router /v1/health/ [get]
func (h *HealthHandler) Health(c *gin.Context) {
	c.JSON(200, helper.GenerateBaseResponse("ebi", true, 0))
	return
}

// HealthPost user
// @Summary HealthPost
// @Description HealthPost
// @Tags HealthPost
// @Accept json
// @Produce json
// @Param person body Person true "person data"
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failure"
// @Router /v1/health/ [post]
func (h *HealthHandler) HealthPost(c *gin.Context) {

	var person Person
	err := c.ShouldBind(&person)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, 1, err))
		return
	}
	c.JSON(200, gin.H{
		"status": true,
		"data":   person,
	})

	return
}
