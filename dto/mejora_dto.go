package dto

type MejoraRequest struct {
	RendimientoActual float64 `json:"rendimiento_actual"`
	TasaMejora        float64 `json:"tasa_mejora"`
	Semanas           int     `json:"semanas"`
}

type SemanaProyectada struct {
	Semana      int     `json:"semana"`
	Rendimiento float64 `json:"rendimiento"`
}

type MejoraResponse struct {
	RendimientoInicial float64            `json:"rendimiento_inicial"`
	TasaMejoraSemanal  float64            `json:"tasa_mejora_semanal"`
	Proyeccion         []SemanaProyectada `json:"proyeccion"`
}
