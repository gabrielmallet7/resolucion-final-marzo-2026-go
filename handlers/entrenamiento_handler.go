package handlers

import (
	"final_marzo_2026/dto"
	"final_marzo_2026/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EntrenamientoHandler struct {
	service services.EntrenamientoServiceInterface
}

func NewEntrenamientoHandler(service services.EntrenamientoServiceInterface) *EntrenamientoHandler {
	return &EntrenamientoHandler{service: service}
}

func (handler *EntrenamientoHandler) CalcularPromedioVelocidad(c *gin.Context) {
	var request dto.EntrenamientoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := handler.service.CalcularPromedioVelocidad(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}
