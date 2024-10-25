package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Unidad representa las unidades dentro de un curso
type Unidad struct {
	ID      primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	IDcurso primitive.ObjectID   `bson:"idcurso" json:"idcurso"`
	Nombre  string               `bson:"nombre" json:"nombre"`
	Clases  []primitive.ObjectID `bson:"clases" json:"clases"`
}
