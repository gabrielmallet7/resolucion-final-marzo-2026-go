package utils

import (
	"final_marzo_2026/dto"
	"final_marzo_2026/models"
	"time"
)

// Convierte el request recibido en un model base.
// Los cálculos se completan después en el service.
func ConvertRendimientoRequestToModel(request dto.RendimientoRequest) models.Rendimiento {
	return models.Rendimiento{
		Atleta:       request.Atleta,
		Rendimientos: request.Rendimientos,
		Fecha:        time.Now(),
	}
}

// Convierte el model ya calculado en response.
func ConvertRendimientoModelToResponse(rendimiento models.Rendimiento) dto.RendimientoResponse {
	return dto.RendimientoResponse{
		Atleta:       rendimiento.Atleta,
		Variabilidad: rendimiento.Variabilidad,
		Metodo:       rendimiento.Metodo,
		Maximo:       rendimiento.Maximo,
		Minimo:       rendimiento.Minimo,
		Fecha:        rendimiento.Fecha,
	}
}
