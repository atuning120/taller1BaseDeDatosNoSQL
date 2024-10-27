package controllers

import (
	"go-API/services"
	"net/http"

	"go-API/models"

	"github.com/gin-gonic/gin"
)

// UnidadControlador maneja las rutas relacionadas con las unidades.
type UsuarioControlador struct {
	servicio *services.UsuarioService
}

// NewUnidadControlador crea un nuevo controlador para las unidades.
func NewUsuarioControlador(servicio *services.UsuarioService) *UsuarioControlador {
	return &UsuarioControlador{servicio: servicio}
}

// ObtenerUnidades obtiene todos los usuarios
func (uc *UsuarioControlador) ObtenerUsuarios(c *gin.Context) {
	usuarios, err := uc.servicio.ObtenerUsuarios()
	if err != nil {
		c.JSON(
			500,
			gin.H{"error": err.Error()},
		)
		return
	}
	c.JSON(200, usuarios)
}

// CrearUsuario maneja la creación de un nuevo usuario.
func (uc *UsuarioControlador) CrearUsuario(c *gin.Context) {
	var usuario models.Usuario

	// Validar los datos enviados en la solicitud
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	// Llamar al servicio para crear el usuario
	result, err := uc.servicio.CrearUsuario(&usuario)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"inserted_id": result.InsertedID})
}

// ObtenerUsuarioPorID obtiene un usuario por su ID.
func (uc *UsuarioControlador) ObtenerUsuarioPorID(c *gin.Context) {
	id := c.Param("id")

	usuario, err := uc.servicio.ObtenerUsuarioPorID(id)
	if err != nil {
		c.JSON(
			500,
			gin.H{"error": err.Error()},
		)
		return
	}
	c.JSON(200, usuario)
}

// InscribirseACurso permite que un usuario se inscriba en un curso.
func (uc *UsuarioControlador) InscribirseACurso(c *gin.Context) {
	var inscripcion struct {
		UsuarioID string `json:"usuario_id"`
		CursoID   string `json:"curso_id"`
	}

	// Validar los datos enviados en la solicitud
	if err := c.ShouldBindJSON(&inscripcion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	// Llamar al servicio para inscribir al usuario en el curso
	err := uc.servicio.InscribirseACurso(inscripcion.UsuarioID, inscripcion.CursoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inscripción exitosa"})
}

// ObtenerCursosInscritos obtiene los cursos en los que un usuario está inscrito.
func (uc *UsuarioControlador) ObtenerCursosInscritos(c *gin.Context) {
	id := c.Param("id")

	cursos, err := uc.servicio.ObtenerCursosInscritos(id)
	if err != nil {
		c.JSON(
			500,
			gin.H{"error": err.Error()},
		)
		return
	}
	c.JSON(200, cursos)
}
