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
type CursoService struct {
	CursoCollection *mongo.Collection
}

// NewCursoService crea un nuevo servicio para los cursos.
func NewCursoService(db *mongo.Database) *CursoService {
	return &CursoService{
		CursoCollection: db.Collection("cursos"),
	}
}

// ObtenerCursos obtiene todos los cursos de la base de datos.
func (s *CursoService) ObtenerCursos() ([]models.Curso, error) {
	var cursos []models.Curso
	cursor, err := s.CursoCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &cursos); err != nil {
		return nil, err
	}
	return cursos, nil
}

// CrearCurso agrega un nuevo curso a la base de datos.
func (s *CursoService) CrearCurso(curso models.Curso) (*mongo.InsertOneResult, error) {
	return s.CursoCollection.InsertOne(context.TODO(), curso)
}

// ObtenerCursoPorID obtiene un curso por su ID.
func (s *CursoService) ObtenerCursoPorID(id string) (*models.Curso, error) {
	objectID, err := primitive.ObjectIDFromHex(id) // Convertir a ObjectID
	if err != nil {
		return nil, errors.New("ID inválido")
	}

	var curso models.Curso
	err = s.CursoCollection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&curso)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("curso no encontrado")
		}
		return nil, err
	}

	return &curso, nil
}

// ObtenerUnidadesPorCurso obtiene las unidades de un curso.
func (s *CursoService) ObtenerUnidadesPorCurso(id string) ([]models.Unidad, error) {
	objectID, err := primitive.ObjectIDFromHex(id) // Convertir a ObjectID
	if err != nil {
		return nil, errors.New("ID inválido")
	}

	var curso models.Curso
	err = s.CursoCollection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&curso)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("curso no encontrado")
		}
		return nil, err
	}

	var unidades []models.Unidad
	cursor, err := s.CursoCollection.Database().Collection("unidades").Find(context.TODO(), bson.M{"_id": bson.M{"$in": curso.Unidades}})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &unidades); err != nil {
		return nil, err
	}

	return unidades, nil
}

// CrearUnidad crea una nueva unidad en un curso.
func (s *CursoService) CrearUnidad(id string, unidad models.Unidad) (*mongo.InsertOneResult, error) {
	objectID, err := primitive.ObjectIDFromHex(id) // Convertir a ObjectID
	if err != nil {
		return nil, errors.New("ID inválido")
	}

	// Verificar si el curso existe
	var curso models.Curso
	err = s.CursoCollection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&curso)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("curso no encontrado")
		}
		return nil, err
	}

	// Asignar el ID del curso a la unidad
	unidad.IDcurso = objectID

	// Insertar la unidad
	result, err := s.CursoCollection.Database().Collection("unidades").InsertOne(context.TODO(), unidad)
	if err != nil {
		return nil, err
	}

	// Actualizar el curso con la nueva unidad
	_, err = s.CursoCollection.UpdateOne(context.TODO(), bson.M{"_id": objectID}, bson.M{"$push": bson.M{"unidades": result.InsertedID}})
	if err != nil {
		return nil, err
	}

	return result, nil
}
