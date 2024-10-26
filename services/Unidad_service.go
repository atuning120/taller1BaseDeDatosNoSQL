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
	ClaseCollection  *mongo.Collection
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

// CrearClaseDeUnidad crea una clase y agrega su ID a la lista de clases de la unidad.
func (s *UnidadService) CrearClaseDeUnidad(unidadID string, clase *models.Clase) error {
	// Convertir el ID de la unidad a ObjectID
	objectID, err := primitive.ObjectIDFromHex(unidadID)
	if err != nil {
		return errors.New("ID inválido")
	}

	// Inicializar el ObjectID para la nueva clase
	clase.ID = primitive.NewObjectID()
	clase.UnidadID = objectID

	// Insertar la clase en la colección de clases
	_, err = s.ClaseCollection.InsertOne(context.TODO(), clase)
	if err != nil {
		return err
	}

	// Verificar si la lista de clases está inicializada (en MongoDB)
	_, err = s.UnidadCollection.UpdateOne(
		context.TODO(),
		bson.M{"_id": objectID, "clases": bson.M{"$exists": false}}, // Verificar si no existe
		bson.M{"$set": bson.M{"clases": []primitive.ObjectID{}}},    // Inicializar si es necesario
	)
	if err != nil {
		return err
	}

	// Agregar el ID de la clase a la lista de clases de la unidad
	_, err = s.UnidadCollection.UpdateOne(
		context.TODO(),
		bson.M{"_id": objectID},
		bson.M{"$push": bson.M{"clases": clase.ID}},
	)
	if err != nil {
		return err
	}

	return nil
}
