package services

import (
	"errors"
	"final_marzo_2026/dto"
	"final_marzo_2026/models"
	"final_marzo_2026/repositories"
	"final_marzo_2026/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"slices"
	"strings"
)

type RendimientoServiceInterface interface {
	CalcularVariabilidadRendimiento(rendimiento dto.RendimientoRequest) (dto.RendimientoResponse, error)
}

type RendimientoService struct {
	repo repositories.RendimientoRepositoryInterface
}

func NewRendimientoService(repo repositories.RendimientoRepositoryInterface) *RendimientoService {
	return &RendimientoService{
		repo: repo,
	}
}

func (service *RendimientoService) CalcularVariabilidadRendimiento(request dto.RendimientoRequest) (dto.RendimientoResponse, error) {
	if strings.TrimSpace(request.Atleta) == "" {
		return dto.RendimientoResponse{}, errors.New("el nombre del atleta no puede estar vacío")
	}

	if len(request.Rendimientos) == 0 {
		return dto.RendimientoResponse{}, errors.New("al menos un rendimiento debe ser proporcionado")
	}

	for _, rendimiento := range request.Rendimientos {
		if rendimiento < 0 {
			return dto.RendimientoResponse{}, errors.New("los rendimientos no pueden ser negativos")
		}
	}

	var rendimiento models.Rendimiento
	rendimiento = utils.ConvertRendimientoRequestToModel(request)
	rendimiento.Maximo = slices.Max(rendimiento.Rendimientos)
	rendimiento.Minimo = slices.Min(rendimiento.Rendimientos)
	rendimiento.Variabilidad = rendimiento.Maximo - rendimiento.Minimo
	rendimiento.Metodo = "maximo - minimo"

	result, err := service.repo.InsertarRendimiento(rendimiento)
	if err != nil {
		return dto.RendimientoResponse{}, err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		rendimiento.ID = oid
	}

	return utils.ConvertRendimientoModelToResponse(rendimiento), nil
}
