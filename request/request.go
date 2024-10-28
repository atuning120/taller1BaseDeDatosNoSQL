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

// CreateClaseRequest define los parámetros necesarios para crear una clase.
type CreateClaseRequest struct {
	Nombre      string `json:"nombre" binding:"required"`
	Descripcion string `json:"descripcion" binding:"required"`
	VideoURL    string `json:"video_url" binding:"required"`
}

// CreateComentarioRequest define los parámetros necesarios para crear un comentario.
type CreateComentarioRequest struct {
	Autor     string `json:"autor" binding:"required"`
	Titulo    string `json:"titulo" binding:"required"`
	Detalle   string `json:"detalle" binding:"required"`
	MeGusta   int    `json:"me_gusta" binding:"required"`
	NoMeGusta int    `json:"no_me_gusta" binding:"required"`
}
