package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/courses", controllers.GetCourses)
	r.POST("/courses", controllers.CreateCourse)
}
