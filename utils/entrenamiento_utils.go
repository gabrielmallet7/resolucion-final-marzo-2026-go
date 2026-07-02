package utils

import (
	"final_marzo_2026/dto"
	"final_marzo_2026/models"
	"time"
)

// Convierte el request recibido en un model base.
// Los campos calculados se completan después en el service.
func ConvertEntrenamientoRequestToModel(request dto.EntrenamientoRequest) models.Entrenamiento {
	return models.Entrenamiento{
		Atleta:      request.Atleta,
		Velocidades: request.Velocidades,
		Fecha:       time.Now(),
	}
}

// Convierte el model ya calculado en response.
func ConvertEntrenamientoModelToResponse(entrenamiento models.Entrenamiento) dto.EntrenamientoResponse {
	return dto.EntrenamientoResponse{
		Atleta:            entrenamiento.Atleta,
		PromedioVelocidad: entrenamiento.PromedioVelocidad,
		CantidadRegistros: entrenamiento.CantidadRegistros,
		Fecha:             entrenamiento.Fecha,
	}
}
