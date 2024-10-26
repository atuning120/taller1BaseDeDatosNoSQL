package controllers

import (
	"net/http"

	"go-API/models"
	"go-API/services"

	"github.com/gin-gonic/gin"
)

type CursoControlador struct {
	servicio *services.CursoService
}

func NewCursoControlador(servicio *services.CursoService) *CursoControlador {
	return &CursoControlador{servicio: servicio}
}

// Ruta para obtener todos los cursos
func (ctrl *CursoControlador) ObtenerCursos(c *gin.Context) {
	cursos, err := ctrl.servicio.ObtenerCursos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cursos)
}

// Ruta para crear un curso
// CrearCurso crea un nuevo curso.
func (ctrl *CursoControlador) CrearCurso(c *gin.Context) {
	var request struct {
		Nombre      string  `json:"nombre"`
		Descripcion string  `json:"descripcion"`
		Imagen      string  `json:"imagen_url"`
		Valoracion  float32 `json:"valoracion"`
	}

	// Verificar si los datos enviados son correctos
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Crear un nuevo curso usando el constructor
	curso := models.NewCurso(request.Nombre, request.Descripcion, request.Imagen, request.Valoracion)

	result, err := ctrl.servicio.CrearCurso(curso)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"inserted_id": result.InsertedID})
}

// Ruta para obtener un curso por su ID
func (ctrl *CursoControlador) ObtenerCursoPorID(c *gin.Context) {
	id := c.Param("id")
	curso, err := ctrl.servicio.ObtenerCursoPorID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, curso)
}

// Ruta para obtener las unidades de un curso
func (ctrl *CursoControlador) ObtenerUnidadesPorCurso(c *gin.Context) {
	id := c.Param("id")
	unidades, err := ctrl.servicio.ObtenerUnidadesPorCurso(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, unidades)
}

// Ruta para crear una unidad en un curso
// CrearUnidad crea una nueva unidad y la agrega a un curso.
func (ctrl *CursoControlador) CrearUnidad(c *gin.Context) {
	id := c.Param("id") // ID del curso
	var unidad models.Unidad

	// Verificar si los datos enviados son correctos
	if err := c.ShouldBindJSON(&unidad); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := ctrl.servicio.CrearUnidad(id, unidad)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"inserted_id": result.InsertedID})
}
