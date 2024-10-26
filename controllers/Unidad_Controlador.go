package controllers

import (
	"net/http"

	"go-API/models"
	"go-API/services"

	"github.com/gin-gonic/gin"
)

// UnidadControlador maneja las rutas relacionadas con las unidades.
type UnidadControlador struct {
	servicio *services.UnidadService
}

// NewUnidadControlador crea un nuevo controlador para las unidades.
func NewUnidadControlador(servicio *services.UnidadService) *UnidadControlador {
	return &UnidadControlador{servicio: servicio}
}

// ObtenerUnidadesPorCurso obtiene las unidades de un curso.
func (ctrl *UnidadControlador) ObtenerUnidadesPorCurso(c *gin.Context) {
	id := c.Param("id")
	unidades, err := ctrl.servicio.ObtenerUnidadesPorCurso(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, unidades)
}

// CrearUnidad crea una nueva unidad y la agrega a un curso.
func (ctrl *UnidadControlador) CrearUnidad(c *gin.Context) {
	id := c.Param("id") // ID del curso
	var unidad models.Unidad

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
