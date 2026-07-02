package dto

import "time"

type ReporteResponse struct {
	Tipo      string    `json:"tipo"`
	Atleta    string    `json:"atleta"`
	Resultado float64   `json:"resultado"`
	Fecha     time.Time `json:"fecha"`
}
