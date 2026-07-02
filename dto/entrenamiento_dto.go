package dto

import "time"

type EntrenamientoRequest struct {
	Atleta      string    `json:"atleta"`
	Velocidades []float64 `json:"velocidades"`
}

type EntrenamientoResponse struct {
	Atleta            string    `json:"atleta"`
	PromedioVelocidad float64   `json:"promedio_velocidad"`
	CantidadRegistros int       `json:"cantidad_registros"`
	Fecha             time.Time `json:"fecha"`
}
