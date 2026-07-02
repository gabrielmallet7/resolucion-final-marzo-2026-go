package handlers

import (
	"final_marzo_2026/dto"
	"final_marzo_2026/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MejoraHandler struct {
	service services.MejoraServiceInterface
}

func NewMejoraHandler(service services.MejoraServiceInterface) *MejoraHandler {
	return &MejoraHandler{service: service}
}

func (handler *MejoraHandler) CalcularProyeccionMejora(c *gin.Context) {
	var request dto.MejoraRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := handler.service.CalcularProyeccionMejora(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}
