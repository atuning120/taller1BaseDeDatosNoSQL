// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/cursos": {
            "get": {
                "description": "Devuelve todos los cursos disponibles",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Obtener cursos"
                ],
                "summary": "Devuelve todos los cursos",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.CursoResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "response.CursoResponse": {
            "type": "object",
            "properties": {
                "cant_usuarios": {
                    "type": "integer"
                },
                "comentarios": {
                    "description": "IDs de los comentarios",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "descripcion": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "imagen_url": {
                    "type": "string"
                },
                "nombre": {
                    "type": "string"
                },
                "unidades": {
                    "description": "IDs de las unidades",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "valoracion": {
                    "type": "number"
                }
            }
        },
        "response.ErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "API de Cursos y Usuarios",
	Description:      "Esta es una API para gestionar cursos y usuarios.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
