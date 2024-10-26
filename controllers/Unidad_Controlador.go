package controllers

import (
	"net/http"

	"go-API/models"
	"go-API/services"

	"github.com/gin-gonic/gin"
)

type UnidadControlador struct {
	servicio *services.UnidadService
}

func NewUnidadControlador(servicio *services.UnidadService) *UnidadControlador {
	return &UnidadControlador{servicio: servicio}
}

// Ruta para obtener todos los cursos
func (ctrl *UnidadControlador) ObtenerUnidades(c *gin.Context) {
	unidades, err := ctrl.servicio.ObtenerUnidades()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, unidades)
}

// ObtenerClasesDeUnidad obtiene las clases de una unidad por su ID.
func (ctrl *UnidadControlador) ObtenerClasesDeUnidad(c *gin.Context) {
	id := c.Param("id")
	unidad, err := ctrl.servicio.ObtenerClasesDeUnidad(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, unidad.Clases)
}

// CrearClaseDeUnidad crea una nueva clase en una unidad.
func (ctrl *UnidadControlador) CrearClaseDeUnidad(c *gin.Context) {
	id := c.Param("id") // ID de la unidad
	var clase models.Clase

	if err := c.ShouldBindJSON(&clase); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.servicio.CrearClaseDeUnidad(id, &clase)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Clase creada exitosamente", "clase_id": clase.ID.Hex()})
}
