package main

import (
	_ "backend/docs" // Import necesario para Swagger
	"backend/routes"
	"backend/services"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API de Cursos
// @version 1.0
// @description API REST para gestionar cursos online.
// @host localhost:8080
// @BasePath /

func main() {
	r := gin.Default()

	services.InitDatabase("mongodb://root:toor@localhost:27017/")
	routes.RegisterRoutes(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
