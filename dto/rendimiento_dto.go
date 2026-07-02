package dto

import "time"

type RendimientoRequest struct {
	Atleta       string    `json:"atleta"`
	Rendimientos []float64 `json:"rendimientos"`
}

type RendimientoResponse struct {
	Atleta       string    `json:"atleta"`
	Variabilidad float64   `json:"variabilidad"`
	Metodo       string    `json:"metodo"`
	Maximo       float64   `json:"maximo"`
	Minimo       float64   `json:"minimo"`
	Fecha        time.Time `json:"fecha"`
}
