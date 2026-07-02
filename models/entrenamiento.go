package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Entrenamiento struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Atleta            string             `bson:"atleta" json:"atleta"`
	Velocidades       []float64          `bson:"velocidades" json:"velocidades"`
	PromedioVelocidad float64            `bson:"promedio_velocidad" json:"promedio_velocidad"`
	CantidadRegistros int                `bson:"cantidad_registros" json:"cantidad_registros"`
	Fecha             time.Time          `bson:"fecha" json:"fecha"`
}
