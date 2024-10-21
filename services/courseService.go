// services/courseService.go
package services

import (
	"backend/models"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var courseCollection *mongo.Collection

func InitDatabase(uri string) {
	client, _ := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	courseCollection = client.Database("training").Collection("courses")
}

func GetAllCourses(ctx context.Context) ([]models.Course, error) {
	var courses []models.Course
	cursor, err := courseCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	cursor.All(ctx, &courses)
	return courses, nil
}

func CreateCourse(ctx context.Context, course *models.Course) error {
	_, err := courseCollection.InsertOne(ctx, course)
	return err
}
