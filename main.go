package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const url = "mongodb+srv://atuning120:C21387541n@cluster0.ketg2.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

var mongoClient *mongo.Client

func init() {
	if err := connect_to_mongodb(); err != nil {
		log.Fatal("No se pudo conectar a MongoDB")
	}
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "anashei",
		})
	})
	r.Run()
}

func connect_to_mongodb() error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(url).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.TODO(), nil)
	mongoClient = client
	return err
}
