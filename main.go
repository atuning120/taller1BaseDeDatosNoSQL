package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"go-API/controllers"
	services "go-API/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func init() {
	if err := loadEnv(); err != nil {
		log.Fatal("Error al cargar las variables de entorno:", err)
	}
	if err := connectToMongoDB(); err != nil {
		log.Fatal("No se pudo conectar a MongoDB:", err)
	}
}

func main() {
	router := gin.Default()

	// Inicializar servicio y controlador
	db := mongoClient.Database("miBaseDeDatos")
	cursoService := services.NewCursoService(db)
	cursoControlador := controllers.NewCursoControlador(cursoService)

	unidadService := services.NewUnidadService(db)
	unidadControlador := controllers.NewUnidadControlador(unidadService)

	// Definición de rutas
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Conexión exitosa"})
	})

	//Cursos
	router.GET("/api/cursos", cursoControlador.ObtenerCursos)
	router.GET("/api/cursos/:id", cursoControlador.ObtenerCursoPorID)
	router.POST("/api/cursos", cursoControlador.CrearCurso)
	router.GET("/api/cursos/:id/unidades", cursoControlador.ObtenerUnidadesPorCurso)
	router.POST("/api/cursos/:id/unidades", cursoControlador.CrearUnidad)

	//Unidades
	router.GET("/api/unidades/:id/clases", unidadControlador.ObtenerClasesDeUnidad)
	router.POST("/api/unidades/:id/clases", unidadControlador.CrearClaseDeUnidad)

	// Iniciar el servidor
	go func() {
		if err := router.Run(); err != nil {
			log.Fatal("Error al iniciar el servidor:", err)
		}
	}()

	// Manejar señal de interrupción
	gracefulShutdown()
}

func connectToMongoDB() error {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		return fmt.Errorf("la URI de MongoDB no está configurada")
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return err
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		return err
	}

	log.Println("Conectado a MongoDB")
	mongoClient = client
	return nil
}

func loadEnv() error {
	return godotenv.Load()
}

func gracefulShutdown() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
	log.Println("Cerrando la conexión con MongoDB...")
	if err := mongoClient.Disconnect(context.TODO()); err != nil {
		log.Fatal("Error al desconectar MongoDB:", err)
	}
	log.Println("Conexión con MongoDB cerrada. Apagando servidor.")
	os.Exit(0)
}
