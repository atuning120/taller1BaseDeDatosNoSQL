package response

import (
	"go-API/models"
)

// CursoResponse define la estructura de la respuesta para un curso.
type CursoResponse struct {
	ID          string   `json:"id"`
	Nombre      string   `json:"nombre"`
	Descripcion string   `json:"descripcion"`
	Imagen      string   `json:"imagen_url"`
	Valoracion  float32  `json:"valoracion"`
	Unidades    []string `json:"unidades"` // IDs de las unidades
	Usuarios    int      `json:"cant_usuarios"`
	Comentarios []string `json:"comentarios"` // IDs de los comentarios
}

// NewCursoResponse convierte un modelo Curso en una respuesta CursoResponse.
func NewCursoResponse(curso models.Curso) CursoResponse {
	unidades := make([]string, len(curso.Unidades))
	for i, unidad := range curso.Unidades {
		unidades[i] = unidad.Hex()
	}

	comentarios := make([]string, len(curso.Comentarios))
	for i, comentario := range curso.Comentarios {
		comentarios[i] = comentario.Hex()
	}

	return CursoResponse{
		ID:          curso.ID.Hex(),
		Nombre:      curso.Nombre,
		Descripcion: curso.Descripcion,
		Imagen:      curso.Imagen,
		Valoracion:  curso.Valoracion,
		Unidades:    unidades,
		Usuarios:    curso.Usuarios,
		Comentarios: comentarios,
	}
}

// ErrorResponse define la estructura de las respuestas de error.
type ErrorResponse struct {
	Message string `json:"message"`
}

// CrearCurso define la respuesta cuando se crea un curso exitosamente.
type CrearCurso struct {
	InsertedID string `json:"inserted_id"`
}

// UpdateValoracionResponse define la estructura de la respuesta para la actualización de la valoración.
type UpdateValoracionResponse struct {
	Message               string  `json:"message"`
	ValoracionActualizada float32 `json:"valoracion_actualizada"`
}
