package repositories

import (
	"context"
	"final_marzo_2026/database"
	"final_marzo_2026/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type EntrenamientoRepositoryInterface interface {
	InsertarEntrenamiento(entrenamiento models.Entrenamiento) (*mongo.InsertOneResult, error)
	ObtenerEntrenamientos() ([]models.Entrenamiento, error)
}

type EntrenamientoRepository struct {
	db database.DB
}

func NewEntrenamientoRepository(db database.DB) *EntrenamientoRepository {
	return &EntrenamientoRepository{
		db: db,
	}
}

func (repo *EntrenamientoRepository) ObtenerEntrenamientos() ([]models.Entrenamiento, error) {
	collection := repo.db.GetClient().Database("velocistas_db").Collection("entrenamientos")

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	var entrenamientos []models.Entrenamiento
	for cursor.Next(context.Background()) {
		var entrenamiento models.Entrenamiento
		err := cursor.Decode(&entrenamiento)
		if err != nil {
			return nil, err
		}
		entrenamientos = append(entrenamientos, entrenamiento)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return entrenamientos, nil
}

func (repo *EntrenamientoRepository) InsertarEntrenamiento(entrenamiento models.Entrenamiento) (*mongo.InsertOneResult, error) {
	collection := repo.db.GetClient().Database("velocistas_db").Collection("entrenamientos")
	result, err := collection.InsertOne(context.Background(), entrenamiento)
	if err != nil {
		return nil, err
	}
	return result, nil
}
