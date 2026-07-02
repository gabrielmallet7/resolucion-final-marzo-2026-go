package utils

import (
	"final_marzo_2026/dto"
	"final_marzo_2026/models"
)

// Convierte un entrenamiento guardado en un reporte.
func ConvertEntrenamientoToReporte(entrenamiento models.Entrenamiento) dto.ReporteResponse {
	return dto.ReporteResponse{
		Tipo:      "promedio_velocidad",
		Atleta:    entrenamiento.Atleta,
		Resultado: entrenamiento.PromedioVelocidad,
		Fecha:     entrenamiento.Fecha,
	}
}

// Convierte un rendimiento guardado en un reporte.
func ConvertRendimientoToReporte(rendimiento models.Rendimiento) dto.ReporteResponse {
	return dto.ReporteResponse{
		Tipo:      "variabilidad_rendimiento",
		Atleta:    rendimiento.Atleta,
		Resultado: rendimiento.Variabilidad,
		Fecha:     rendimiento.Fecha,
	}
}
