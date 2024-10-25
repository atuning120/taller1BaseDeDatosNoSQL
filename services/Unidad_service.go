package services

import (
	"context"
	"errors"

	"go-API/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CursoService gestiona la lógica relacionada con los cursos.
type UnidadService struct {
	UnidadCollection *mongo.Collection
}

// NewCursoService crea un nuevo servicio para los cursos.
func NewUnidadService(db *mongo.Database) *UnidadService {
	return &UnidadService{
		UnidadCollection: db.Collection("unidades"),
	}
}

// ObtenerCursos obtiene todos los cursos de la base de datos.
func (s *UnidadService) ObtenerUnidades() ([]models.Unidad, error) {
	var unidades []models.Unidad
	cursor, err := s.UnidadCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &unidades); err != nil {
		return nil, err
	}
	return unidades, nil
}

// ObtenerClasesDeUnidad obtiene las clases de una unidad por su ID.
func (s *UnidadService) ObtenerClasesDeUnidad(id string) (*models.Unidad, error) {
	objectID, err := primitive.ObjectIDFromHex(id) // Convertir el ID a ObjectID
	if err != nil {
		return nil, errors.New("ID inválido")
	}

	var unidad models.Unidad
	err = s.UnidadCollection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&unidad)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("unidad no encontrada")
		}
		return nil, err
	}

	return &unidad, nil
}
