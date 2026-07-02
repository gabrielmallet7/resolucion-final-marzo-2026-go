package repositories

import (
	"context"
	"final_marzo_2026/database"
	"final_marzo_2026/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RendimientoRepositoryInterface interface {
	InsertarRendimiento(rendimiento models.Rendimiento) (*mongo.InsertOneResult, error)
	ObtenerRendimientos() ([]models.Rendimiento, error)
}

type RendimientoRepository struct {
	db database.DB
}

func NewRendimientoRepository(db database.DB) *RendimientoRepository {
	return &RendimientoRepository{
		db: db,
	}
}

func (repo *RendimientoRepository) ObtenerRendimientos() ([]models.Rendimiento, error) {
	collection := repo.db.GetClient().Database("velocistas_db").Collection("rendimientos")

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	var rendimientos []models.Rendimiento
	for cursor.Next(context.Background()) {
		var rendimiento models.Rendimiento
		err := cursor.Decode(&rendimiento)
		if err != nil {
			return nil, err
		}
		rendimientos = append(rendimientos, rendimiento)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return rendimientos, nil
}

func (repo *RendimientoRepository) InsertarRendimiento(rendimiento models.Rendimiento) (*mongo.InsertOneResult, error) {
	collection := repo.db.GetClient().Database("velocistas_db").Collection("rendimientos")
	result, err := collection.InsertOne(context.Background(), rendimiento)
	if err != nil {
		return nil, err
	}
	return result, nil
}
