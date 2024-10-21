// models/course.go
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Course struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	BannerURL   string             `bson:"banner_url" json:"banner_url"`
	Rating      float32            `bson:"rating" json:"rating"`
	Units       []Unit             `bson:"units" json:"units"`
}

type Unit struct {
	Title   string  `bson:"title" json:"title"`
	Classes []Class `bson:"classes" json:"classes"`
}

type Class struct {
	Title       string `bson:"title" json:"title"`
	Description string `bson:"description" json:"description"`
	VideoURL    string `bson:"video_url" json:"video_url"`
}
