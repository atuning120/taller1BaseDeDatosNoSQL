package controllers

import (
	"go-API/services"

	"github.com/gin-gonic/gin"
)

// UnidadControlador maneja las rutas relacionadas con las unidades.
type ComentarioControlador struct {
	servicio *services.ComentarioService
}

// NewUnidadControlador crea un nuevo controlador para las unidades.
func NewComentarioControlador(servicio *services.ComentarioService) *ComentarioControlador {
	return &ComentarioControlador{servicio: servicio}
}

// ObtenerComentariosPorClase obtiene todos los comentarios de una clase.
func (c *ComentarioControlador) ObtenerComentariosPorClase(ctx *gin.Context) {
	claseID := ctx.Param("id")
	comentarios, err := c.servicio.ObtenerComentariosPorClase(claseID)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, comentarios)
}

// CrearComentarioParaClase crea un nuevo comentario para una clase.
func (c *ComentarioControlador) CrearComentarioParaClase(ctx *gin.Context) {
	claseID := ctx.Param("id")
	comentario, err := c.servicio.CrearComentarioParaClase(claseID, ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(201, comentario)
}
