package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Curso representa un curso en la base de datos.
type Curso struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Nombre      string               `bson:"nombre" json:"nombre"`
	Descripcion string               `bson:"descripcion" json:"descripcion"`
	Imagen      string               `bson:"imagen_url" json:"imagen_url"`
	Valoracion  float32              `bson:"valoracion" json:"valoracion"`
	Unidades    []primitive.ObjectID `bson:"unidades" json:"unidades"`
	Usuarios    int                  `bson:"cant_usuarios" json:"cant_usuarios"`
	Comentarios []primitive.ObjectID `bson:"comentarios" json:"comentarios"`
}
