package services

import (
	"errors"
	"final_marzo_2026/dto"
	"final_marzo_2026/repositories"
	"final_marzo_2026/utils"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EntrenamientoServiceInterface interface {
	CalcularPromedioVelocidad(request dto.EntrenamientoRequest) (dto.EntrenamientoResponse, error)
}

type EntrenamientoService struct {
	repo repositories.EntrenamientoRepositoryInterface
}

func NewEntrenamientoService(repo repositories.EntrenamientoRepositoryInterface) *EntrenamientoService {
	return &EntrenamientoService{
		repo: repo,
	}
}

func (service *EntrenamientoService) CalcularPromedioVelocidad(request dto.EntrenamientoRequest) (dto.EntrenamientoResponse, error) {
	if strings.TrimSpace(request.Atleta) == "" {
		return dto.EntrenamientoResponse{}, errors.New("el nombre del atleta no puede estar vacío")
	}
	if len(request.Velocidades) == 0 {
		return dto.EntrenamientoResponse{}, errors.New("al menos una velocidad debe ser proporcionada")
	}

	suma := 0.0

	for _, velocidad := range request.Velocidades {
		if velocidad < 0 {
			return dto.EntrenamientoResponse{}, errors.New("la velocidad no puede ser negativa")
		}
		suma += velocidad
	}

	promedio := suma / float64(len(request.Velocidades))

	entrenamiento := utils.ConvertEntrenamientoRequestToModel(request)
	entrenamiento.PromedioVelocidad = promedio
	entrenamiento.CantidadRegistros = len(request.Velocidades)

	result, err := service.repo.InsertarEntrenamiento(entrenamiento)
	if err != nil {
		return dto.EntrenamientoResponse{}, err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		entrenamiento.ID = oid
	}

	return utils.ConvertEntrenamientoModelToResponse(entrenamiento), nil
}
