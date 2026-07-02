package services

import (
	"errors"
	"final_marzo_2026/dto"
	"final_marzo_2026/utils"
)

type MejoraServiceInterface interface {
	CalcularProyeccionMejora(dto.MejoraRequest) (dto.MejoraResponse, error)
}

type MejoraService struct{}

func NewMejoraService() *MejoraService {
	return &MejoraService{}
}

func (service *MejoraService) CalcularProyeccionMejora(request dto.MejoraRequest) (dto.MejoraResponse, error) {
	if request.RendimientoActual < 0 || request.RendimientoActual > 100 {
		return dto.MejoraResponse{}, errors.New("el rendimiento actual debe estar entre 0 y 100")
	}

	if request.TasaMejora < 0 {
		return dto.MejoraResponse{}, errors.New("la tasa de mejora no puede ser negativa")
	}

	if request.Semanas <= 0 {
		return dto.MejoraResponse{}, errors.New("el número de semanas debe ser mayor que cero")
	}

	var mejoraResponse dto.MejoraResponse
	mejoraResponse = utils.ConvertMejoraRequestToResponse(request)
	rendimientoAnterior := request.RendimientoActual

	for i := 1; i <= request.Semanas; i++ {
		mejoraResponse.Proyeccion = append(mejoraResponse.Proyeccion, utils.ObtenerSemanaProyectada(rendimientoAnterior, mejoraResponse.TasaMejoraSemanal, i))

		rendimientoAnterior = mejoraResponse.Proyeccion[i-1].Rendimiento
	}

	return mejoraResponse, nil
}
