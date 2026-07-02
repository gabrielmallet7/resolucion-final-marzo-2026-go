package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Rendimiento struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Atleta       string             `bson:"atleta" json:"atleta"`
	Rendimientos []float64          `bson:"rendimientos" json:"rendimientos"`
	Variabilidad float64            `bson:"variabilidad" json:"variabilidad"`
	Metodo       string             `bson:"metodo" json:"metodo"`
	Maximo       float64            `bson:"maximo" json:"maximo"`
	Minimo       float64            `bson:"minimo" json:"minimo"`
	Fecha        time.Time          `bson:"fecha" json:"fecha"`
}
