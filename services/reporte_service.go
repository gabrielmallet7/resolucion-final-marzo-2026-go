package services

import (
	"final_marzo_2026/dto"
	"final_marzo_2026/repositories"
	"final_marzo_2026/utils"
	"sort"
)

type ReporteServiceInterface interface {
	ObtenerReportes() ([]dto.ReporteResponse, error)
}

type ReporteService struct {
	repoEntrenamiento repositories.EntrenamientoRepositoryInterface
	repoRendimiento   repositories.RendimientoRepositoryInterface
}

func NewReporteService(repoEntrenamiento repositories.EntrenamientoRepositoryInterface, repoRendimiento repositories.RendimientoRepositoryInterface) *ReporteService {
	return &ReporteService{
		repoEntrenamiento: repoEntrenamiento,
		repoRendimiento:   repoRendimiento,
	}
}

func (service *ReporteService) ObtenerReportes() ([]dto.ReporteResponse, error) {
	entrenamientos, err := service.repoEntrenamiento.ObtenerEntrenamientos()
	if err != nil {
		return nil, err
	}

	rendimientos, err := service.repoRendimiento.ObtenerRendimientos()
	if err != nil {
		return nil, err
	}

	var reportes []dto.ReporteResponse

	for _, entrenamiento := range entrenamientos {
		reportes = append(reportes, utils.ConvertEntrenamientoToReporte(entrenamiento))
	}

	for _, rendimiento := range rendimientos {
		reportes = append(reportes, utils.ConvertRendimientoToReporte(rendimiento))
	}

	sort.Slice(reportes, func(i, j int) bool {
		return reportes[i].Fecha.After(reportes[j].Fecha)
	})

	return reportes, nil
}
