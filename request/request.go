package request

// CreateCursoRequest define el cuerpo de la solicitud para crear un curso.
type CreateCursoRequest struct {
	Nombre      string `json:"nombre" binding:"required"`
	Descripcion string `json:"descripcion"`
	Imagen      string `json:"imagen_url"`
}

// UpdateValoracionRequest define el cuerpo de la solicitud para actualizar la valoración.
type UpdateValoracionRequest struct {
	Valoracion float32 `json:"valoracion" binding:"required"`
}

// CreateUnidadRequest define el cuerpo de la solicitud para crear una unidad.
type CreateUnidadRequest struct {
	Nombre string `json:"nombre" binding:"required"`
}
