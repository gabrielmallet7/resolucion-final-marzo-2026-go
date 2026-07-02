package handlers

import (
	"final_marzo_2026/dto"
	"final_marzo_2026/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RendimientoHandler struct {
	service services.RendimientoServiceInterface
}

func NewRendimientoHandler(service services.RendimientoServiceInterface) *RendimientoHandler {
	return &RendimientoHandler{service: service}
}

func (handler *RendimientoHandler) CalcularVariabilidadRendimiento(c *gin.Context) {
	var request dto.RendimientoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := handler.service.CalcularVariabilidadRendimiento(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}
