package handlers

import (
	"final_marzo_2026/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ReporteHandler struct {
	service services.ReporteServiceInterface
}

func NewReporteHandler(service services.ReporteServiceInterface) *ReporteHandler {
	return &ReporteHandler{service: service}
}

func (handler *ReporteHandler) ObtenerReportes(c *gin.Context) {
	response, err := handler.service.ObtenerReportes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}
