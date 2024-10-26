package controllers

import (
	"net/http"

	"go-API/models"
	"go-API/services"

	"github.com/gin-gonic/gin"
)

// ClaseControlador gestiona las rutas relacionadas con las clases.
type ClaseControlador struct {
	servicio *services.ClaseService
}

// NewClaseControlador crea un nuevo controlador para las clases.
func NewClaseControlador(servicio *services.ClaseService) *ClaseControlador {
	return &ClaseControlador{servicio: servicio}
}

// ObtenerClasesPorUnidad obtiene todas las clases de una unidad.
func (cc *ClaseControlador) ObtenerClasesPorUnidad(c *gin.Context) {
	id := c.Param("id") // ID de la unidad

	clases, err := cc.servicio.ObtenerClasesPorUnidad(id)
	if err != nil {
		if err.Error() == "unidad no encontrada" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, clases)
}

// CrearClaseParaUnidad crea una nueva clase asociada a una unidad.
func (cc *ClaseControlador) CrearClaseParaUnidad(c *gin.Context) {
	unidadID := c.Param("id") // ID de la unidad

	var clase models.Clase
	if err := c.ShouldBindJSON(&clase); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	// Llamar al servicio para crear la clase
	result, err := cc.servicio.CrearClaseParaUnidad(unidadID, &clase)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"inserted_id": result.InsertedID})
}
