package controllers

import (
	"net/http"

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
