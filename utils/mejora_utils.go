package utils

import "final_marzo_2026/dto"

func ConvertMejoraRequestToResponse(request dto.MejoraRequest) dto.MejoraResponse {
	return dto.MejoraResponse{
		RendimientoInicial: request.RendimientoActual,
		TasaMejoraSemanal:  request.TasaMejora,
		Proyeccion:         []dto.SemanaProyectada{},
	}
}

func ObtenerSemanaProyectada(rendimientoActual, tasaMejora float64, semana int) dto.SemanaProyectada {
	return dto.SemanaProyectada{
		Semana:      semana,
		Rendimiento: rendimientoActual * (1 + tasaMejora/100),
	}
}
